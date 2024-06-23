package extract

import (
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/triopium/go_utils/pkg/files"
)

type ValidateColumnAddressParams struct {
	SheetName string
	RowPartCode
	FieldID     string
	FormatRegex string
}

type ValidationColumnAccount map[string]int // Invalid Value VS count

type ValidationTableAccount struct {
	InvalidValues    map[string]ValidationColumnAccount // Column name VS ColumnValidation
	ColumnsPositions []string
}

func LogValidationError(err error) {
	slog.Error("validation_error", "msg", err.Error())
}

// VALIDATE UNIT FUNCTIONS
func (e *Extractor) ColumnAccountPrepare(
	pc RowPartCode, fieldID string) ValidationColumnAccount {
	if e.ValidationTableAccount.InvalidValues == nil {
		e.ValidationTableAccount.InvalidValues = make(map[string]ValidationColumnAccount)
	}
	if e.ValidationTableAccount.ColumnsPositions == nil {
		e.ValidationTableAccount.ColumnsPositions = make([]string, 0)
	}
	columnName := GetColumnHeaderExternal(pc, fieldID)
	columnAccount, ok := e.ValidationTableAccount.InvalidValues[columnName]
	if !ok {
		columnAccount = make(map[string]int)
		e.ValidationTableAccount.InvalidValues[columnName] = columnAccount
	}
	return columnAccount
}

func FieldPrevalidate(rp RowParts, rpc RowPartCode, fieldID string) bool {
	_, field, ok := GetRowPartAndField(
		rp, rpc, fieldID)

	errFuncLog := func(cause string) {
		slog.Debug("validation_warning", "filed", field, "cause", cause)
	}

	if !ok {
		errFuncLog("field missing")
		return false
	}
	if field.Value == "" {
		errFuncLog("field empty")
		return false
	}
	if CheckIfFieldValueIsSpecialValue(field.Value) {
		errFuncLog("special value")
		return false
	}
	return true
}

func ValueInvalidAccount(
	vca ValidationColumnAccount,
	part RowPart, field RowField,
	valid bool,
) {
	valueNC := RowFieldSpecialValueCodeMap[RowFieldValueNotValid]
	if !valid {
		part[field.FieldID] = RowField{
			FieldID:   field.FieldID,
			FieldName: "",
			Value:     valueNC,
		}
		vca[field.Value]++
	}
}

func FormatFieldBeforeValidation(input string) string {
	out := strings.ToLower(input)                // all lower letters
	out = strings.ReplaceAll(out, "\u00a0", " ") // non-breaking space (NBSP) not needed if strings.Join(strings.Fields(out)," ") is used as it removes NBSP also
	// out = strings.Replace(out, " ", "", -1)      // replace space
	out = strings.Replace(out, "\t", "", -1)
	out = strings.Replace(out, "\n", "", -1)
	out = strings.Join(strings.Fields(out), " ")
	return out
}

func FormatFieldValuesListBeforeValidation(input string) []string {
	// split list by various delimiters to slice
	out := strings.ToLower(input)                // all lower letters
	out = strings.ReplaceAll(out, "\u00a0", " ") // non-breaking space (NBSP) not needed if strings.Join(strings.Fields(out)," ") is used as it removes NBSP also
	out = strings.Replace(out, "\t", " ", -1)
	out = strings.Replace(out, "\n", " ", -1)
	out = strings.Replace(out, ";", " ", -1)
	out = strings.Replace(out, ",", " ", -1)
	return strings.Fields(out)
}

// ValidateStopaz
func ValidateStopaz(stopaz string) (string, error) {
	milliSeconds, err := strconv.ParseInt(stopaz, 10, 64)
	specVal := RowFieldSpecialValueCodeMap[RowFieldValueNotValid]
	if err != nil {
		return stopaz, err
	}
	if milliSeconds < 0 {
		// negative stopaz is invalid
		return specVal, err
	}
	return stopaz, nil
}

func (e *Extractor) ValidateAllColumns(
	xlsxValidationReceipeFile string) {
	e.ValidateColumnsValues(xlsxValidationReceipeFile)
	e.ValidateColumnsValuesList(xlsxValidationReceipeFile)
	e.ValidateColumnsFormat(xlsxValidationReceipeFile)
}

// VALIDATE COLUMNS VALUES LIST
func ValidateValuesAgainstList(
	allovedVals map[string][]string, valuesList string) bool {
	// valueList is string containing values delimited by delimiter
	values := FormatFieldValuesListBeforeValidation(valuesList)
	for _, val := range values {
		// _, ok := allovedVals[val+";"]
		_, ok := allovedVals[val]
		// NOTE: in allovedVals the values are "01;" from source xlsx file
		if !ok {
			return false
		}
	}
	return true
}

func (e *Extractor) ValidateColumnsValuesList(xlsxValidationReceipeFile string) {
	columnsParams := []ValidateColumnAddressParams{
		{"druh", RowPartCode_StoryHead, "16", ""},
		{"tema", RowPartCode_StoryHead, "5016", ""},
		{"tema", RowPartCode_ContactItemHead, "5016", ""},
	}
	for _, c := range columnsParams {
		err := e.ValidateColumnValuesList(xlsxValidationReceipeFile, c)
		if err != nil {
			LogValidationError(err)
		}
	}
}

func (e *Extractor) ValidateColumnValuesList(
	xlsxValidationReceipeFile string,
	vp ValidateColumnAddressParams,
) error {
	sheetRows, err := files.ReadExcelFileSheetRows(
		xlsxValidationReceipeFile, vp.SheetName)
	if err != nil {
		return err
	}
	sheetTableMapped := files.CreateTableTransformRowHeader(
		sheetRows, 0, 0, FormatFieldBeforeValidation)
	allovedVals := sheetTableMapped.RowHeaderToColumnMap
	columnAccount := e.ColumnAccountPrepare(vp.RowPartCode, vp.FieldID)
	rs := e.TableXML.Rows

	for _, r := range rs {
		if !FieldPrevalidate(r.RowParts, vp.RowPartCode, vp.FieldID) {
			continue
		}
		part, field, _ := GetRowPartAndField(
			r.RowParts, vp.RowPartCode, vp.FieldID)
		validOK := ValidateValuesAgainstList(allovedVals, field.Value)
		ValueInvalidAccount(columnAccount, part, field, validOK)
	}
	return nil
}

// VALIDATE COLUMNS VALUES
func (e *Extractor) ValidateColumnsValues(
	xlsxValidationReceipeFile string) {
	columnsParams := []ValidateColumnAddressParams{
		{"stanice_RR", RowPartCode_RadioHead, "5081", ""},
		{"format_SR", RowPartCode_SubHead, "321", ""},
		{"format", RowPartCode_StoryHead, "321", ""},
		{"redakce", RowPartCode_StoryHead, "12", ""},
		{"cil_vyroby", RowPartCode_StoryHead, "5079", ""},
		{"pohlavi_KON", RowPartCode_ContactItemHead, "5088", ""},
	}
	for _, c := range columnsParams {
		err := e.ValidateColumnValues(xlsxValidationReceipeFile, c)
		if err != nil {
			LogValidationError(err)
		}
	}
}

func (e *Extractor) ValidateColumnValues(
	xlsxValidationReceipeFile string,
	vp ValidateColumnAddressParams,
) error {
	sheetRows, err := files.ReadExcelFileSheetRows(
		xlsxValidationReceipeFile, vp.SheetName)
	if err != nil {
		fmt.Println("FUCK2")
		return err
	}

	sheetTableMapped := files.CreateTableTransformRowHeader(
		sheetRows, 0, 0, FormatFieldBeforeValidation)
	columnAccount := e.ColumnAccountPrepare(vp.RowPartCode, vp.FieldID)

	rs := e.TableXML.Rows
	for _, r := range rs {
		if !FieldPrevalidate(r.RowParts, vp.RowPartCode, vp.FieldID) {
			continue
		}
		part, field, _ := GetRowPartAndField(
			r.RowParts, vp.RowPartCode, vp.FieldID)
		valueToValidate := FormatFieldBeforeValidation(field.Value)
		_, validOK := sheetTableMapped.RowHeaderToColumnMap[valueToValidate]

		ValueInvalidAccount(columnAccount, part, field, validOK)
	}
	return nil
}

// VALIDATE COLUMNS FORMAT
func (e *Extractor) ValidateColumnsFormat(xlsxValidationReceipeFile string) {
	rgx_ID := `^[a-fA-F\d]{8}(-[a-fA-F\d]{4}){3}-[a-fA-F\d]{8}$`
	// UUID variant eg.: F763E219-4759-4698-B2D6-5E6DFDBC
	rgx_Stopaz := `^\d{1,16}$`         // only numbers
	rgx_Datum := `^\d{8}T\d{6},\d{3}$` // 20240102T130010,333
	rgx_Incode := `^OM\d{3,}$`         // OM + 3 or more digits
	rgx_Itemcode := `^.{4,}$`          // 4 or more digits
	rgx_Nazev_HR := `^\d\d:\d\d-\d\d:\d\d$`

	columnsParams := []ValidateColumnAddressParams{
		// ID
		{"ID_KON", RowPartCode_ContactItemHead, "5087", rgx_ID},
		{"ID_KON", RowPartCode_ContactItemHead, "5068", rgx_ID},

		// STOPAZ
		{"stopaz", RowPartCode_SubHead, "1005", rgx_Stopaz},
		{"stopaz", RowPartCode_SubHead, "1002", rgx_Stopaz},
		{"stopaz", RowPartCode_StoryHead, "1005", rgx_Stopaz},
		{"stopaz", RowPartCode_StoryHead, "1002", rgx_Stopaz},
		{"stopaz", RowPartCode_AudioClipHead, "38", rgx_Stopaz},

		// DATUM 4x
		{"datum", RowPartCode_RadioHead, "1000", rgx_Datum},
		{"datum", RowPartCode_HourlyHead, "1000", rgx_Datum},
		{"datum", RowPartCode_SubHead, "1004", rgx_Datum},
		{"datum", RowPartCode_StoryHead, "1004", rgx_Datum},

		{"incode", RowPartCode_StoryHead, "5072", rgx_Incode},
		{"itemcode", RowPartCode_AudioClipHead, "5082", rgx_Itemcode},
		{"nazev_HR", RowPartCode_HourlyHead, "8", rgx_Nazev_HR},
	}
	for _, c := range columnsParams {
		err := e.ValidateColumnFormat(xlsxValidationReceipeFile, c)
		if err != nil {
			LogValidationError(err)
		}
	}
}

func (e *Extractor) ValidateColumnFormat(
	xlsxValidationReceipeFile string,
	vp ValidateColumnAddressParams,
) error {
	formatRegex := regexp.MustCompile(vp.FormatRegex)
	columnAccount := e.ColumnAccountPrepare(vp.RowPartCode, vp.FieldID)

	rs := e.TableXML.Rows
	for i, r := range rs {
		if !FieldPrevalidate(r.RowParts, vp.RowPartCode, vp.FieldID) {
			continue
		}
		part, field, _ := GetRowPartAndField(
			r.RowParts, vp.RowPartCode, vp.FieldID)
		// fieldValueWithoutSpace := TrimAllWhiteSpace(field.Value)
		validOK := formatRegex.MatchString(field.Value)
		if !validOK {
			slog.Debug("validation_warning", "msg", "not valid", "field", field, "rowi", i, "row", r)
		}
		ValueInvalidAccount(columnAccount, part, field, validOK)
	}
	return nil
}

func (e *Extractor) ValidationLogWrite(
	logFilePath, delim string, overWrite bool) error {
	header := e.ValidationLogHeader()
	sb := new(strings.Builder)
	fmt.Fprintf(sb, "%s\n", strings.Join(
		[]string{"column", "value", "count"}, delim))
	for _, h := range header {
		columnAccount := e.ValidationTableAccount.InvalidValues[h]
		// Sorted per column name
		keys := SortColumnAccount(columnAccount)
		for _, k := range keys {
			// fmt.Fprintf(sb, "%s%s%s%s%d\n",
			fmt.Fprintf(sb, "%s%s%q%s%d\n",
				h, delim, k, delim, columnAccount[k])
		}
		// Not sorted
		// for value, count := range columnAccount {
		// fmt.Fprintf(sb, "%s%s%s%s%d\n", h, delim, value, delim, count)
		// }
	}
	perms := FileOverwritePermissions(overWrite)
	outputFile, err := os.OpenFile(logFilePath, perms, 0600)
	if err != nil {
		return err
	}
	defer outputFile.Close()
	n, err := outputFile.WriteString(sb.String())
	if err != nil {
		return err
	}
	slog.Info("written bytes to file", "fileName", logFilePath, "bytesCount", n)
	return nil
}

func SortColumnAccount(vca ValidationColumnAccount) []string {
	type KeyValue struct {
		Key   string
		Value int
	}
	var kvs []KeyValue
	for k, v := range vca {
		kvs = append(kvs, KeyValue{k, v})
	}

	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].Value > kvs[j].Value
	})
	out := make([]string, len(kvs))
	for i, kv := range kvs {
		out[i] = kv.Key
	}
	return out
}

func (e *Extractor) ValidationLogHeader() []string {
	header := make([]string, 0)
	for _, partPos := range e.HeaderExternal {
		columnAccount, ok := e.ValidationTableAccount.InvalidValues[partPos]
		if !ok {
			continue
		}
		if len(columnAccount) == 0 {
			continue
		}
		header = append(header, partPos)
	}
	return header
}
