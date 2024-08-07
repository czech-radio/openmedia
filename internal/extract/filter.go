package extract

import (
	"fmt"
	"log/slog"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/triopium/go_utils/pkg/files"
)

type FilterFileCode string

const (
	FilterFileOposition    FilterFileCode = "filtr_opozice"
	FilterFileEuroElection FilterFileCode = "filtr_eurovolby"
)

type FilterFunction func(*FilterFile) error
type FilterFileCodes map[FilterFileCode][]FilterFunction

var FilterFileCodeMap = FilterFileCodes{
	FilterFileOposition:    {},
	FilterFileEuroElection: {},
}

func (ffc FilterFileCodes) AddFilters(
	filterCode FilterFileCode, filterFuncs ...FilterFunction) {
	ffc[filterCode] = filterFuncs
}

func (ffc FilterFileCodes) FiltersApply(ff *FilterFile) error {
	filterCode, err := ffc.GetFilterFileCode(ff.FilterFileName)
	if err != nil {
		return err
	}
	filterFuncs, ok := ffc[filterCode]
	if !ok {
		return fmt.Errorf("No filter functions defined for filter file: %v", filterCode)
	}
	for _, ffunc := range filterFuncs {
		err := ffunc(ff)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ffc *FilterFileCodes) GetFilterFileCode(filePath string) (FilterFileCode, error) {
	fileName := filepath.Base(filePath)
	for f := range *ffc {
		rg := regexp.MustCompile("^" + string(f))
		if rg.MatchString(fileName) {
			return f, nil
		}
	}
	return FilterFileCode(""), fmt.Errorf("Unknown filter filename: %s", filePath)
}

type FilterFile struct {
	FilterFileName  string
	FilterSheetName string
	ColumnHeaderRow int
	RowHeaderColumn int
}

func (e *Extractor) MarkField(rowParts RowParts,
	partCode RowPartCode, fieldID, value string) {
	part, ok := rowParts[partCode]
	if !ok {
		part = make(RowPart)
	}
	field := RowField{
		FieldID:   fieldID,
		FieldName: "",
		Value:     value,
	}
	part[fieldID] = field
	rowParts[partCode] = part

}

func MarkValue(matches bool, fieldValue, nulValue string) string {
	if matches {
		return "1"
	}
	if fieldValue == nulValue {
		return nulValue
	}
	return "0"
}

// FilterByPartAndFieldIDregexp
func (e *Extractor) FilterByPartAndFieldIDregexp(
	partCode RowPartCode, fieldID string,
	fieldValuePatern string) []int {
	var res []int
	re := regexp.MustCompile(fieldValuePatern)
	for i, row := range e.TableXML.Rows {
		part, ok := row.RowParts[partCode]
		if !ok {
			slog.Debug(
				"filter not effective", "reason", "partname not found",
				"partName", partCode,
			)
			return nil
		}
		field, ok := part[fieldID]
		if !ok {
			slog.Debug(
				"filter not effective", "reason", "fieldID not found",
				"partName", partCode,
			)
			return nil
		}
		ok = re.MatchString(field.Value)
		if ok {
			res = append(res, i)
		}
	}
	return res
}

func (e *Extractor) FilterMatchPersonNameExact(f *FilterFile) error {
	newColumnName := "name_match"
	e.AddColumn(RowPartCode_ContactItemHead, newColumnName)
	sheetRows, err := files.ReadExcelFileSheetRows(
		f.FilterFileName, f.FilterSheetName)
	if err != nil {
		return err
	}
	sheetTableMapped := files.CreateTable(
		sheetRows, f.ColumnHeaderRow, f.RowHeaderColumn)
	valueNP := RowFieldSpecialValueCodeMap[RowFieldValueNotPossible]

	rs := e.TableXML.Rows
	for _, r := range rs {
		_, field, ok := GetRowPartAndField(
			r.RowParts, RowPartCode_ComputedKON, "jmeno_spojene")
		if !ok {
			panic(ok)
		}
		_, ok = sheetTableMapped.RowHeaderToColumnMap[field.Value]
		mark := MarkValue(ok, field.Value, valueNP)
		e.MarkField(r.RowParts, RowPartCode_ContactItemHead,
			newColumnName, mark)
	}
	return nil
}

func (e *Extractor) FilterMatchPersonName(f *FilterFile) error {
	newColumnName := "name_match"
	e.AddColumn(RowPartCode_ContactItemHead, newColumnName)
	sheetRows, err := files.ReadExcelFileSheetRows(
		f.FilterFileName, f.FilterSheetName)
	if err != nil {
		return err
	}
	sheetTableMapped := files.CreateTableTransformRowHeader(
		sheetRows, f.ColumnHeaderRow, f.RowHeaderColumn, TransformPersonName)
	valueNP := RowFieldSpecialValueCodeMap[RowFieldValueNotPossible]

	rs := e.TableXML.Rows
	for _, r := range rs {
		_, field, ok := GetRowPartAndField(
			r.RowParts, RowPartCode_ComputedKON, "jmeno_spojene")
		if !ok {
			panic(ok)
		}
		valueTransformed := TransformPersonName(field.Value)
		_, ok = sheetTableMapped.RowHeaderToColumnMap[valueTransformed]
		mark := MarkValue(ok, valueTransformed, valueNP)
		e.MarkField(r.RowParts, RowPartCode_ContactItemHead, newColumnName, mark)
	}
	return nil
}

// MatchStringElement check if element is present in slice
func MatchStringElement(slice []string, element string) bool {
	for _, s1 := range slice {
		if s1 == element {
			return true
		}
	}
	return false
}

// MatchStringElements check if at least count elements are present in both slices
func MatchStringElements(str1, str2 []string, count int) bool {
	countMatched := 0
	curIdx := 0
	for _, s1 := range str1 {
		if MatchStringElement(str2[curIdx:], s1) {
			curIdx++
			countMatched++
		}
	}
	return countMatched >= count
}

func MatchPersonNameSurnameNormalized(
	table *files.Table,
	rowParts RowParts,
) (string, bool) {
	xColumnSurnameName := "Příjmení a jméno"
	xColumnSurname := "Příjmení"
	xColumnName := "Jméno"
	colsMap := table.ColumnHeaderMap
	xColumnSurnameNameIdx, ok0 := colsMap[xColumnSurnameName]
	xColumnSurnameIdx, ok1 := colsMap[xColumnSurname]
	xColumnNameIdx, ok2 := colsMap[xColumnName]
	if !(ok0 && ok1 && ok2) {
		panic("At least one of column header not found")
		// slog.Error("At least one of column header not found")
	}
	_, spojene, _ := GetRowPartAndField(
		rowParts, RowPartCode_ComputedKON, "jmeno_spojene")

	tableMap := table.RowHeaderToColumnMap
	xRow, ok := tableMap[TransformPersonNameNoDiacritcs(spojene.Value)]
	if ok {
		return xRow[xColumnSurnameNameIdx], ok
	}
	_, name, _ := GetRowPartAndField(
		rowParts, RowPartCode_ContactItemHead, "421")
	namest := strings.Fields(TransformPersonNameNoDiacritcs(name.Value))
	_, surname, _ := GetRowPartAndField(
		rowParts, RowPartCode_ContactItemHead, "422")
	surnamest := strings.Fields(TransformPersonNameNoDiacritcs(surname.Value))

	xtable := table.RowHeaderToColumnMap
	for _, trow := range xtable {
		xnames := strings.Fields(
			TransformPersonNameNoDiacritcs(trow[xColumnNameIdx]))
		xsurnames := strings.Fields(
			TransformPersonNameNoDiacritcs(trow[xColumnSurnameIdx]))
		ok1 := MatchStringElements(xnames, namest, 1)
		ok2 := MatchStringElements(xsurnames, surnamest, 1)
		if ok1 && ok2 {
			nameRef := trow[xColumnSurnameNameIdx]
			return nameRef, true
		}
	}

	return "", false
}

// FilterMatchPersonNameSurnameNormalized match at least one Name and one Surname from Name, Surname table. Name and surnames are normalized to all lowercase
func (e *Extractor) FilterMatchPersonNameSurnameNormalized(f *FilterFile) error {
	columnNameMatch := "name_match"
	e.AddColumn(RowPartCode_ContactItemHead, columnNameMatch)
	ColumnNameRefered := "referencni_jmeno"
	e.AddColumn(RowPartCode_ContactItemHead, ColumnNameRefered)

	sheetRows, err := files.ReadExcelFileSheetRows(f.FilterFileName, f.FilterSheetName)
	if err != nil {
		return err
	}
	sheetTableMapped := files.CreateTableTransformRowHeader(
		sheetRows, f.ColumnHeaderRow, f.RowHeaderColumn, TransformPersonNameNoDiacritcs)
	sheetTableMapped.MapTableRowKeyTransform(sheetRows, f.ColumnHeaderRow, f.RowHeaderColumn, TransformPersonNameNoDiacritcs)
	valueNP := RowFieldSpecialValueCodeMap[RowFieldValueNotPossible]

	rs := e.TableXML.Rows
	for _, r := range rs {
		nameRef, ok := MatchPersonNameSurnameNormalized(
			sheetTableMapped, r.RowParts)
		var refnameMark string
		var nameMatchMark string
		if ok {
			refnameMark = nameRef
			nameMatchMark = "1"
		}
		if !ok {
			refnameMark = valueNP
			nameMatchMark = "0"
		}
		e.MarkField(
			r.RowParts, RowPartCode_ContactItemHead,
			columnNameMatch, nameMatchMark)
		e.MarkField(
			r.RowParts, RowPartCode_ContactItemHead,
			ColumnNameRefered, refnameMark)
	}
	return nil
}

func (e *Extractor) FilterMatchPersonAndParty(f *FilterFile) error {
	newColumnName := "name&party_match"
	e.AddColumn(RowPartCode_ContactItemHead, newColumnName)
	rows, err := files.ReadExcelFileSheetRows(
		f.FilterFileName, f.FilterSheetName)
	if err != nil {
		return err
	}
	table := files.CreateTable(
		rows, f.ColumnHeaderRow, f.RowHeaderColumn)
	rs := e.TableXML.Rows
	valNP := RowFieldSpecialValueCodeMap[RowFieldValueNotPossible]
	for _, r := range rs {
		_, nameField, ok := GetRowPartAndField(r.RowParts, RowPartCode_ComputedKON, "jmeno_spojene")
		if !ok {
			panic(fmt.Errorf("row part and fieldname not present in row"))
		}
		_, partyField, ok := GetRowPartAndField(
			r.RowParts, RowPartCode_ContactItemHead, "5015")
		if !ok {
			mark := MarkValue(ok, partyField.Value, valNP)
			e.MarkField(r.RowParts, RowPartCode_ContactItemHead, newColumnName, mark)
		}
		ok1 := table.MatchRow(
			nameField.Value, "navrhující strana", partyField.Value)
		ok2 := table.MatchRow(
			nameField.Value, "politická příslušnost", partyField.Value)
		res := ok1 || ok2
		mark := MarkValue(res, partyField.Value, valNP)
		e.MarkField(r.RowParts, RowPartCode_ContactItemHead, newColumnName, mark)
	}
	return nil
}

func (e *Extractor) FilterMatchPersonIDandHighPolitics(f *FilterFile) error {
	newColumnName := "vysoka_politika"
	e.AddColumn(RowPartCode_ContactItemHead, newColumnName)
	rows, err := files.ReadExcelFileSheetRows(f.FilterFileName, f.FilterSheetName)
	if err != nil {
		return err
	}
	table := files.CreateTable(rows, f.ColumnHeaderRow, 3)
	valNP := RowFieldSpecialValueCodeMap[RowFieldValueNotPossible]
	rs := e.TableXML.Rows
	mark := ""
	for _, r := range rs {
		_, contactIDfield, ok := GetRowPartAndField(
			r.RowParts, RowPartCode_ContactItemHead, "5068")
		if !ok {
			slog.Error("row part and fieldname not present in row")
			continue
		}
		if contactIDfield.Value == valNP {
			e.MarkField(r.RowParts, RowPartCode_ContactItemHead, newColumnName, valNP)
			continue
		}
		row, ok := table.RowHeaderToColumnMap[contactIDfield.Value]
		if !ok {
			mark = "99"
		} else {
			mark = row[0]
		}
		e.MarkField(r.RowParts, RowPartCode_ContactItemHead, newColumnName, mark)
	}
	return nil
}

func (e *Extractor) FilterPeculiarContacts() []int {
	indxs := make([]int, 0, len(e.Rows))
	rows := e.TableXML.Rows
	for i, row := range rows {
		_, objid, _ := GetRowPartAndField(
			row.RowParts, RowPartCode_StoryKategory, "TemplateName")
		switch objid.Value {
		case "Contact Item", "Contact Bin",
			"UNKNOWN-Contact Item", "UNKNOWN-Contact Bin":
			indxs = append(indxs, i)
		default:
			continue
		}
	}
	return indxs
}

var rgxRecordDuds = regexp.MustCompile(`\d\d\d\d\d`)

// FilterStoryPartRecordsDuds get row indexes of records which has ObjectID or rows which corresponding RecordID is smaller than rgxRecordDuds
func (e *Extractor) FilterStoryPartRecordsDuds() []int {
	indxs := make([]int, 0, len(e.Rows))
	rows := e.TableXML.Rows
	for i, row := range rows {
		_, objid, _ := GetRowPartAndField(
			row.RowParts, RowPartCode_StoryKategory, "ObjectID")
		_, recid, _ := GetRowPartAndField(
			row.RowParts, RowPartCode_StoryRec, "RecordID")
		// _, storytype, _ := GetRowPartAndField(row.RowParts, RowPartCode_StoryRec, "5001")
		if objid.Value != "" {
			indxs = append(indxs, i)
			continue
		}
		if recid.Value != "" {
			ok := rgxRecordDuds.MatchString(recid.Value)
			if ok {
				continue
			}
		}
		indxs = append(indxs, i)
	}
	return indxs
}

func FieldValueIsEmpty(value string) bool {
	if value == "" {
		return true
	}
	return CheckIfFieldValueIsSpecialValue(value)
}

func (e *Extractor) FilterStoryPartsEmptyDupes() []int {
	indxs := make([]int, 0, len(e.Rows)/100)
	rows := e.TableXML.Rows
	var storyIDprevious string
	var addEmptyRow int = -1

	for i, row := range rows {
		_, storyID, _ := GetRowPartAndField(
			row.RowParts, RowPartCode_StoryHead, "ObjectID")
		_, storyPartTemplate, _ := GetRowPartAndField(
			row.RowParts, RowPartCode_StoryKategory, "TemplateName")

		if storyID.Value != storyIDprevious {
			storyIDprevious = storyID.Value
			if addEmptyRow > -1 {
				indxs = append(indxs, addEmptyRow)
				addEmptyRow = -1
			}
			if FieldValueIsEmpty(storyPartTemplate.Value) {
				addEmptyRow = i
				continue
			}
			indxs = append(indxs, i)
			continue
		}

		if !FieldValueIsEmpty(storyPartTemplate.Value) {
			indxs = append(indxs, i)
		}
	}
	if addEmptyRow > -1 {
		indxs = append(indxs, addEmptyRow)
	}
	return indxs
}

// DeleteNonMatchingRows delete row indexes not specified in rowIndxsFiltered
func (e *Extractor) DeleteNonMatchingRows(rowIdxsFiltered []int) {
	slog.Info("rows count before deletion",
		"parsed", len(e.Rows), "filtered", len(rowIdxsFiltered))
	out := make([]*RowNode, len(rowIdxsFiltered))
	rowIdxFilteredCurrent := 0
	for ri := range e.Rows {
		if ri == rowIdxsFiltered[rowIdxFilteredCurrent] {
			out[rowIdxFilteredCurrent] = e.Rows[ri]
			rowIdxFilteredCurrent++
		}
		if rowIdxFilteredCurrent+1 > len(rowIdxsFiltered) {
			break
		}
	}
	e.Rows = out
	slog.Info("rows count after deletion", "count_after", len(out))
}

func ReverseIndexes(rowsIndxs []int) []int {
	nonMatching := make([]int, 0)
	curIndex := 0
	for i := range rowsIndxs {
		for j := curIndex; j < rowsIndxs[i]; j++ {
			nonMatching = append(nonMatching, j)
		}
		curIndex = rowsIndxs[i] + 1
	}
	return nonMatching
}
