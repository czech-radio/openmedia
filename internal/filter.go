package internal

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"os"
	"sort"
	"strconv"
	"strings"

	// "github.com/antchfx/xmlquery"
	"github.com/antchfx/xmlquery"
)

type FilterOptions struct {
	SourceDirectory        string
	DestinationDirectory   string
	RecurseSourceDirectory bool
	InvalidFileContinue    bool
	FileType               string
	DateFrom               string
	DateTo                 string
	OutputType             string
	CSVdelim               string
	CSVheader              bool
}

type ObjectAttributes = map[string]string
type Fields = map[int]string       // FieldID vs value
type UniqueValues = map[string]int // value vs count

type Filter struct {
	Options               FilterOptions
	Results               FilterResults
	Errors                []error
	ObjectHeader          []string
	ObjectsAttrValues     []ObjectAttributes
	HeaderFields          map[int]string
	HeaderFieldsIDsSorted []int
	HeaderFieldsIDsSubset map[int]bool
	Rows                  []Fields
	FieldsUniqueValues    map[int]UniqueValues // FiledID vs UniqueValues
	MaxUniqueCount        int
}

type FilterResults struct {
	FilesProcessed int
	FilesSuccess   int
	FilesFailure   int
	FilesCount     int
}

var FieldKeysToExtract []int = []int{
	7, 8, 304, 321,
	406, 407, 408, 421,
	422, 423, 424, 425,
	5015, 5016, 5087, 5088,
}

func (ft *Filter) KeysToExtract() {
	if ft.HeaderFieldsIDsSubset == nil {
		ft.HeaderFieldsIDsSubset = make(map[int]bool)
	}
	for _, i := range FieldKeysToExtract {
		ft.HeaderFieldsIDsSubset[i] = true
	}
}

func (ft *Filter) ErrorHandle(errMain error, errorsPartial ...error) ControlFlowAction {
	if errMain == nil {
		ft.Results.FilesSuccess++
		return Continue
	}

	ft.Results.FilesFailure++
	slog.Error(errMain.Error())
	ft.Errors = append(ft.Errors, errMain)
	if len(errorsPartial) > 0 {
		ft.Errors = append(ft.Errors, errorsPartial...)
	}

	if ft.Options.InvalidFileContinue {
		return Skip
	} else {
		return Break
	}
}

func (ft *Filter) LogResults(msg string) {
	slog.Info(msg, "results", fmt.Sprintf("%+v", ft.Results))
}
func (ft *Filter) Folder() error {
	validateResult, err := ValidateFilesInDirectory(ft.Options.SourceDirectory, ft.Options.RecurseSourceDirectory)
	if ft.ErrorHandle(err, validateResult.Errors...) == Break {
		return err
	}
	ft.Results.FilesCount = validateResult.FilesCount
	ft.Results.FilesFailure = validateResult.FilesFailure
	ft.LogResults("valid files")
	ft.KeysToExtract()
	ft.Rows = make([]Fields, 0, 10)
processFolder:
	for _, file := range validateResult.FilesValid {
		// Clear rows for next file
		// ft.Rows = make([]Fields, 0, 10)
		ft.ObjectsAttrValues = make([]ObjectAttributes, 0, 10)
		// Process file
		slog.Info("filter", "file", file)
		err = ft.File(file)
		if err != nil {
			err = fmt.Errorf("file: %s, err: %w", file, err)
		}
		flow := ft.ErrorHandle(err)
		switch flow {
		case Skip:
			continue processFolder
		case Break:
			break processFolder
		}
	}
	switch ft.Options.OutputType {
	case "csv":
		if ft.Options.CSVheader {
			ft.CSVheaderWrite()
		}
		err = ft.CSVwriteRows()
		if err != nil {
			slog.Error(err.Error())
		}
	case "unique":
		if ft.Options.CSVheader {
			ft.CSVheaderWriteB()
			ft.CSVwriteUniqueCountsB()
		}
		err = ft.CSVwriteUniqueFieldsValuesB()
		if err != nil {
			return err
		}
	}
	return nil
}

func (ft *Filter) File(filePath string) error {
	fileHandle, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer fileHandle.Close()

	// Read file data
	data, err := io.ReadAll(fileHandle)
	if err != nil {
		return err
	}

	// Transform input data
	dataReader := bytes.NewReader(data)
	pr := PipeUTF16leToUTF8(dataReader)
	pr = PipeRundownHeaderAmmend(pr)

	doc, err := xmlquery.Parse(pr)
	if err != nil {
		return err
	}
	err = ft.FilterObjectByTemplateName(doc, "Contact Item")
	if err != nil {
		return err
	}
	// switch writePerFile {
	// case "csv":
	// 	if ft.Options.CSVheader {
	// 		ft.CSVheaderWrite()
	// 	}
	// 	err = ft.CSVwriteRows()
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

func (ft *Filter) FilterObjectByTemplateName(
	doc *xmlquery.Node, templateName string) error {
	templateQuery := fmt.Sprintf("//OM_OBJECT[@TemplateName='%s']", templateName)
	templates := xmlquery.Find(doc, templateQuery)
	for i := range templates {
		if i == 0 {
			ft.ObjectHeaderNamesAdd(templates[i])
		}
		ft.ObjectHeaderFields(templates[i])
		fieldQuery := "//OM_HEADER/OM_FIELD"
		row := xmlquery.Find(templates[i], fieldQuery)
		err := ft.Fields(row)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ft *Filter) ObjectHeaderFields(node *xmlquery.Node) {
	values := make(map[string]string)
	for _, a := range node.Attr {
		values[a.Name.Local] = a.Value
	}
	ft.ObjectsAttrValues = append(ft.ObjectsAttrValues, values)
}

func (ft *Filter) ObjectHeaderNamesAdd(node *xmlquery.Node) {
	if ft.ObjectHeader == nil {
		ft.ObjectHeader = make([]string, 0, 6)
	} else {
		return
	}
	for _, a := range node.Attr {
		ft.ObjectHeader = append(ft.ObjectHeader, a.Name.Local)
	}
}

func (ft *Filter) UniqueFieldValues(fieldID int, fieldValue string) {
	if ft.FieldsUniqueValues == nil {
		ft.FieldsUniqueValues = make(map[int]UniqueValues)
	}
	_, ok := ft.FieldsUniqueValues[fieldID]
	if !ok {
		ft.FieldsUniqueValues[fieldID] = make(map[string]int)
	}
	uv := ft.FieldsUniqueValues[fieldID]
	uv[fieldValue]++
}

func (ft *Filter) Fields(nodes []*xmlquery.Node) error {
	if ft.HeaderFields == nil {
		ft.HeaderFields = make(map[int]string)
	}
	var row Fields = make(Fields)

	for _, n := range nodes {
		fieldID, err := strconv.Atoi(n.SelectAttr("FieldID"))
		if err != nil {
			return err
		}
		_, ok := ft.HeaderFieldsIDsSubset[fieldID]
		if !ok {
			continue
		}

		fieldName := n.SelectAttr("FieldName")
		fieldValue := n.InnerText()
		ft.HeaderFields[fieldID] = fieldName
		// row[fieldID] = fieldValue
		row[fieldID] = strings.TrimSpace(fieldValue)
		ft.UniqueFieldValues(fieldID, fieldValue)
	}
	ft.Rows = append(ft.Rows, row)

	// Sort all FieldIDs
	keys := make([]int, 0, len(ft.HeaderFields))
	for k := range ft.HeaderFields {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	ft.HeaderFieldsIDsSorted = keys
	return nil
}

func (ft *Filter) CSVheaderWrite() {
	var headerNamed strings.Builder
	// var headerIDs strings.Builder

	// OJECT ATTRS
	// 1.) Wrtite object header IDs
	// for _, h := range f.ObjectHeader {
	// fmt.Fprintf(&headerNamed, "%s%s", h, f.Options.CSVdelim)
	// }
	// 2.) Wrtir object header names
	// for _, h := range f.ObjectHeader {
	// fmt.Fprintf(&headerIDs, "%s%s", h, f.Options.CSVdelim)
	// }

	// FIELDS
	// keys := ft.HeaderFieldsIDsSorted
	keys := FieldKeysToExtract
	// 1.) Write fields IDs
	// for ki, _ := range keys {
	// fmt.Fprintf(&headerIDs, "%d%s", keys[ki], f.Options.CSVdelim)
	// }
	// fmt.Fprintf(os.Stdout, "%s\n", headerNamed.String())
	// 2.) Write fields names
	for ki, _ := range keys {
		key := keys[ki]
		fmt.Fprintf(&headerNamed, "%s%s", ft.HeaderFields[key], ft.Options.CSVdelim)
	}
	fmt.Fprintf(os.Stdout, "%s\n", headerNamed.String())

	// Do not write header for subsequent files
	ft.Options.CSVheader = false
}

func (ft *Filter) CSVwriteRows() error {
	if len(ft.Rows) == 0 {
		return fmt.Errorf("no rows to export")
	}
	// Write rows
	for i := range ft.Rows {
		var csvRow strings.Builder
		//Write object attrib values
		// for _, h := range f.ObjectHeader {
		// val, ok := objValues[h]
		// if !ok {
		// return fmt.Errorf("missing object attribute: %s", h)
		// }
		// fmt.Fprintf(&csvRow, "%s%s", val, f.Options.CSVdelim)
		// }
		//Write fields values
		for _, id := range FieldKeysToExtract {
			rowField, ok := ft.Rows[i][id]
			if !ok {
				slog.Warn("missing field value", "fieldID", id)
				rowField = ""
			}
			fmt.Fprintf(&csvRow, "%s%s", EscapeDelim(rowField), ft.Options.CSVdelim)
		}
		//Row to stdout
		fmt.Fprintf(os.Stdout, "%s\n", csvRow.String())
	}
	return nil
}

func (ft *Filter) GetMaxUniqueFieldsValues() int {
	var max int
	for _, uv := range ft.FieldsUniqueValues {
		luv := len(uv)
		if max < luv {
			max = luv
		}
	}
	return max
}

func (ft *Filter) CSVwriteUniqueCounts() {
	var csvRow strings.Builder
	for _, key := range FieldKeysToExtract {
		count := len(ft.FieldsUniqueValues[key])
		fmt.Fprintf(&csvRow, "%d%s", count, ft.Options.CSVdelim)
	}
	fmt.Fprintf(os.Stdout, "%s\n", csvRow.String())
}

func (ft *Filter) CSVwriteUniqueCountsB() {
	var csvRow strings.Builder
	for _, key := range FieldKeysToExtract {
		count := len(ft.FieldsUniqueValues[key])
		fmt.Fprintf(&csvRow, "%d%s%s", count, ft.Options.CSVdelim, ft.Options.CSVdelim)
	}
	fmt.Fprintf(os.Stdout, "%s\n", csvRow.String())
}

func EscapeDelim(value string) string {
	out := strings.TrimSpace(value)
	out = strings.ReplaceAll(out, "\t", "\\t")
	out = strings.ReplaceAll(out, "\n", "\\n")
	return out
}

func (ft *Filter) CSVwriteUniqueFieldsValues() error {
	rowsCount := ft.GetMaxUniqueFieldsValues()
	rows := make([][]string, rowsCount, rowsCount)
	rowLen := len(ft.FieldsUniqueValues)
	for i, key := range FieldKeysToExtract {
		fvals, _ := ft.FieldsUniqueValues[key]
		rown := 0
		for j, k := range fvals {
			if len(rows[rown]) == 0 {
				rows[rown] = make([]string, rowLen, rowLen)
			}
			// rows[rown][i] = j
			rows[rown][i] = fmt.Sprintf("%d: %s", k, j)
			rown++
		}
	}
	for i := range rows {
		var csvRow strings.Builder
		for j := range rows[i] {
			row := rows[i]
			fmt.Fprintf(&csvRow, "%s%s", row[j], ft.Options.CSVdelim)
		}
		fmt.Fprintf(os.Stdout, "%s\n", csvRow.String())
	}
	return nil
}

func (ft *Filter) CSVheaderWriteB() {
	var headerNamed strings.Builder
	keys := FieldKeysToExtract
	// Write fields names
	for ki, _ := range keys {
		key := keys[ki]
		// fmt.Fprintf(&headerNamed, "%s%s", ft.HeaderFields[key], ft.Options.CSVdelim)
		fmt.Fprintf(&headerNamed, "c(%s)%s", ft.HeaderFields[key], ft.Options.CSVdelim)
		fmt.Fprintf(&headerNamed, "%s%s", ft.HeaderFields[key], ft.Options.CSVdelim)
	}
	fmt.Fprintf(os.Stdout, "%s\n", headerNamed.String())

	// Do not write header for subsequent files
	ft.Options.CSVheader = false
}

func (ft *Filter) CSVwriteUniqueFieldsValuesB() error {
	// counts in separate column
	rowsCount := ft.GetMaxUniqueFieldsValues()
	rows := make([][]string, rowsCount, rowsCount)
	rowLen := len(ft.FieldsUniqueValues)
	for i, key := range FieldKeysToExtract {
		// i:		index (column)
		// key: FiledID
		fvals, _ := ft.FieldsUniqueValues[key]
		rown := 0
		for j, k := range fvals {
			if len(rows[rown]) == 0 {
				rows[rown] = make([]string, 2*rowLen, 2*rowLen)
			}
			// unique value frequency (četnost)
			rows[rown][i*2] = fmt.Sprint(k)
			// unique value
			rows[rown][i*2+1] = EscapeDelim(j)
			rown++
		}
	}
	for i := range rows {
		var csvRow strings.Builder
		for j := range rows[i] {
			row := rows[i]
			fmt.Fprintf(&csvRow, "%s%s", row[j], ft.Options.CSVdelim)
		}
		fmt.Fprintf(os.Stdout, "%s\n", csvRow.String())
	}
	return nil
}

func (ft *Filter) CSVwriteUniqueFieldsValuesSorted() error {
	// TODO
	return nil
}
