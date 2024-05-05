package extract

import (
	"log/slog"

	"github.com/triopium/go_utils/pkg/files"
)

type FilterCode int

const (
	FilterCodeMatchPersonName FilterCode = iota
)

type NFilterColumn struct {
	FilterFileName  string
	SheetName       string
	ColumnHeaderRow int
	RowHeaderColumn int
	PartCodeMark    RowPartCode
	FieldIDmark     string
	FilterTable     files.Table
}

func (e *Extractor) MarkMatched(
	matches bool, rowParts RowParts, partCode RowPartCode, fieldID string) {
	part, ok := rowParts[partCode]
	if !ok {
		part = make(RowPart)
	}
	var mark string
	if matches {
		mark = "1"
	} else {
		mark = "0"
	}
	field := RowField{
		FieldID:   fieldID,
		FieldName: "",
		Value:     mark,
	}
	part[fieldID] = field
	rowParts[partCode] = part
}

func (e *Extractor) FilterMatchPersonName(f *NFilterColumn) error {
	newColumnName := "name_match"
	e.AddColumn(RowPartCode_ContactItemHead, newColumnName)
	rows, err := files.ReadExcelFileSheetRows(f.FilterFileName, f.SheetName)
	if err != nil {
		return err
	}
	table := files.CreateTable(rows, f.ColumnHeaderRow, f.RowHeaderColumn)
	rs := e.TableXML.Rows
	for _, r := range rs {
		_, field, ok := GetRowPartAndField(r.RowParts, RowPartCode_ComputedKON, "jmeno_spojene")
		if !ok {
			panic(ok)
		}
		_, ok = table.RowHeaderToColumnMap[field.Value]
		e.MarkMatched(ok, r.RowParts, RowPartCode_ContactItemHead, newColumnName)
	}
	return nil
}

func (e *Extractor) FilterMatchPersonAndParty(f *NFilterColumn) error {
	newColumnName := "name&party_match"
	e.AddColumn(RowPartCode_ContactItemHead, newColumnName)
	rows, err := files.ReadExcelFileSheetRows(f.FilterFileName, f.SheetName)
	if err != nil {
		return err
	}
	table := files.CreateTable(rows, f.ColumnHeaderRow, f.RowHeaderColumn)
	rs := e.TableXML.Rows
	for _, r := range rs {
		_, nameField, ok := GetRowPartAndField(r.RowParts, RowPartCode_ComputedKON, "jmeno_spojene")
		if !ok {
			panic(ok)
		}
		_, partyField, ok := GetRowPartAndField(r.RowParts, RowPartCode_ContactItemHead, "5015")
		if !ok {
			panic(ok)
		}
		ok1 := table.MatchRow(nameField.Value, "navrhující strana", partyField.Value)
		ok2 := table.MatchRow(nameField.Value, "politická příslušnost", partyField.Value)
		res := ok1 || ok2
		e.MarkMatched(res, r.RowParts, RowPartCode_ContactItemHead, newColumnName)
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
