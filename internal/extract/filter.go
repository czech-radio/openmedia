package extract

import (
	"fmt"
	"log/slog"
	"regexp"

	"github.com/triopium/go_utils/pkg/files"
)

type FilterCode int

const (
	FilterCodeMatchPersonName FilterCode = iota
)

type NFilterColumn struct {
	FilterFileName  string
	FilterSheetName string
	ColumnHeaderRow int
	RowHeaderColumn int
	PartCodeMark    RowPartCode
	FieldIDmark     string
	FilterTable     files.Table
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

func (e *Extractor) FilterMatchPersonNameExact(f *NFilterColumn) error {
	newColumnName := "name_match"
	e.AddColumn(RowPartCode_ContactItemHead, newColumnName)
	sheetRows, err := files.ReadExcelFileSheetRows(f.FilterFileName, f.FilterSheetName)
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
		e.MarkField(r.RowParts, RowPartCode_ContactItemHead, newColumnName, mark)
	}
	return nil
}

func (e *Extractor) FilterMatchPersonName(f *NFilterColumn) error {
	newColumnName := "name_match"
	e.AddColumn(RowPartCode_ContactItemHead, newColumnName)
	sheetRows, err := files.ReadExcelFileSheetRows(f.FilterFileName, f.FilterSheetName)
	if err != nil {
		return err
	}
	sheetTableMapped := files.CreateTableTransformRowHeader(
		sheetRows, f.ColumnHeaderRow, f.RowHeaderColumn, TransformName)
	valueNP := RowFieldSpecialValueCodeMap[RowFieldValueNotPossible]

	rs := e.TableXML.Rows
	for _, r := range rs {
		_, field, ok := GetRowPartAndField(
			r.RowParts, RowPartCode_ComputedKON, "jmeno_spojene")
		if !ok {
			panic(ok)
		}
		valueTransformed := TransformName(field.Value)
		// _, ok = sheetTableMapped.RowHeaderToColumnMap[field.Value]
		_, ok = sheetTableMapped.RowHeaderToColumnMap[valueTransformed]
		// mark := MarkValue(ok, field.Value, valueNP)
		mark := MarkValue(ok, valueTransformed, valueNP)
		e.MarkField(r.RowParts, RowPartCode_ContactItemHead, newColumnName, mark)
	}
	return nil
}

func (e *Extractor) FilterMatchPersonAndParty(f *NFilterColumn) error {
	newColumnName := "name&party_match"
	e.AddColumn(RowPartCode_ContactItemHead, newColumnName)
	rows, err := files.ReadExcelFileSheetRows(f.FilterFileName, f.FilterSheetName)
	if err != nil {
		return err
	}
	table := files.CreateTable(rows, f.ColumnHeaderRow, f.RowHeaderColumn)
	rs := e.TableXML.Rows
	valNP := RowFieldSpecialValueCodeMap[RowFieldValueNotPossible]
	for _, r := range rs {
		_, nameField, ok := GetRowPartAndField(r.RowParts, RowPartCode_ComputedKON, "jmeno_spojene")
		if !ok {
			panic(fmt.Errorf("row part and fieldname not present in row"))
		}
		_, partyField, ok := GetRowPartAndField(r.RowParts, RowPartCode_ContactItemHead, "5015")
		if !ok {
			mark := MarkValue(ok, partyField.Value, valNP)
			e.MarkField(r.RowParts, RowPartCode_ContactItemHead, newColumnName, mark)
		}
		ok1 := table.MatchRow(nameField.Value, "navrhující strana", partyField.Value)
		ok2 := table.MatchRow(nameField.Value, "politická příslušnost", partyField.Value)
		res := ok1 || ok2
		mark := MarkValue(res, partyField.Value, valNP)
		e.MarkField(r.RowParts, RowPartCode_ContactItemHead, newColumnName, mark)
	}
	return nil
}

func (e *Extractor) FilterMatchPersonIDandPolitics(f *NFilterColumn) error {
	newColumnName := "vysoka_politika"
	e.AddColumn(RowPartCode_ContactItemHead, newColumnName)
	rows, err := files.ReadExcelFileSheetRows(f.FilterFileName, f.FilterSheetName)
	if err != nil {
		return err
	}
	table := files.CreateTable(rows, f.ColumnHeaderRow, 1)
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

// FilterColumn
type FilterColumn struct {
	FilterName     FilterCode
	PartCodeCheck  RowPartCode
	FieldIDcheck   string
	PartCodeMark   RowPartCode
	FieldIDmark    string
	FileWithValues string
	Values         map[string]bool
}

var FilterCodeMap = map[FilterCode]FilterColumn{
	FilterCodeMatchPersonName: {
		FilterCodeMatchPersonName,
		RowPartCode_ComputedKON, "jmeno_spojene",
		RowPartCode_ContactItemHead, "filtered", "", nil},
}

func (e *Extractor) FilterContacts() []int {
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

func (e *Extractor) FilterStoryPartRecordsDuds() []int {
	indxs := make([]int, 0, len(e.Rows))
	rows := e.TableXML.Rows
	for i, row := range rows {
		_, objid, _ := GetRowPartAndField(row.RowParts, RowPartCode_StoryKategory, "ObjectID")
		_, recid, _ := GetRowPartAndField(row.RowParts, RowPartCode_StoryRec, "RecordID")
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

func (e *Extractor) DeleteNonMatchingRows(rowIdxsFiltered []int) {
	slog.Warn("rows count before deletion", "parsed", len(e.Rows), "filtered", len(rowIdxsFiltered))
	out := make([]*RowNode, len(rowIdxsFiltered))
	rowIdxFilteredCurrent := 0
	for ri := range e.Rows {
		if ri == rowIdxsFiltered[rowIdxFilteredCurrent] {
			out[rowIdxFilteredCurrent] = e.Rows[ri]
			rowIdxFilteredCurrent++
		}
	}
	e.Rows = out
	slog.Warn("rows count after deletion", "count_after", len(out))
}

// FilterRun
func (e *Extractor) FilterRun(f FilterColumn) {
	switch f.FilterName {
	case FilterCodeMatchPersonName:
		e.FilterMatchPersonNameB(f)
	}
}

// FiltersRun
func (e *Extractor) FiltersRun(filters []FilterColumn) {
	if filters == nil {
		slog.Debug("no filters to filter column specified")
		return
	}
	for f := range filters {
		e.FilterRun(filters[f])
	}
}

// FilterMatchPersonName
func (e *Extractor) FilterMatchPersonNameB(f FilterColumn) {
	for i, row := range e.TableXML.Rows {
		matches, altValue := FieldValueMatchesValidValues(
			row.RowParts, f.PartCodeCheck, f.FieldIDcheck, f.Values,
		)
		part, ok := e.TableXML.Rows[i].RowParts[f.PartCodeMark]
		if !ok {
			part = make(RowPart)
		}
		var mark string
		if matches {
			mark = "1"
		} else {
			mark = "0"
		}
		if altValue != "" {
			mark = altValue
		}
		field := RowField{
			FieldID:   f.FieldIDmark,
			FieldName: "",
			Value:     mark,
		}
		part[f.FieldIDmark] = field
		e.TableXML.Rows[i].RowParts[f.PartCodeMark] = part
	}
}

// FieldValueMatchesValidValues
func FieldValueMatchesValidValues(
	row RowParts, partCode RowPartCode, fieldID string,
	validValues map[string]bool) (bool, string) {
	_, field, ok := GetRowPartAndField(row, partCode, fieldID)
	notFound := RowFieldSpecialValueCodeMap[RowFieldValueChildNotFound]
	if !ok {
		return false, notFound
	}
	return validValues[field.Value], ""
}

// FilterByPartAndFieldID
func (e *Extractor) FilterByPartAndFieldID(
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

// func UniqSliceInt(slices ...[]int) []int {
// 	outMap := make(map[int]bool)
// 	outSlice []int
// 	for s := range slices {
// 		for _, r := range slices[s] {
// 			if outMap[r] {
// 			}
// 			// outMap[r] = true
// 		}
// 	}
// }
