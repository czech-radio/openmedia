package extract

import (
	"log/slog"

	"github.com/triopium/go_utils/pkg/files"
)

type FilterCode int

const (
	FilterCodeMatchPersonName FilterCode = iota
)

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

type FilterColumnNew struct {
	FilterFileName  string
	FilterSheetName string
	TransformKey    func(string) string
	CheckingFunc    func(RowParts) bool

	CheckColumnPart  RowPartCode
	CheckColumnField string
	MarkColumnPart   RowPartCode
	MarkColumnField  string
	MarkSpecial      string
}

func (e *Extractor) Filter(f *FilterColumnNew) error {
	e.AddColumn(f.MarkColumnPart, f.MarkColumnField)
	sheetRows, err := files.ReadExcelFileSheetRows(f.FilterFileName, f.FilterSheetName)
	if err != nil {
		return err
	}
	sheetTableMapped := files.CreateTableTransformRowHeader(
		sheetRows, 0, 0, f.TransformKey)

	rs := e.TableXML.Rows
	for _, r := range rs {
		_, field, ok := GetRowPartAndField(
			r.RowParts, f.CheckColumnPart, f.CheckColumnField)
		if !ok {
			panic(ok)
		}
		valueTransformed := f.TransformKey(field.Value)
		_, ok = sheetTableMapped.RowHeaderToColumnMap[valueTransformed]
		mark := MarkValue(ok, valueTransformed, f.MarkSpecial)
		e.MarkField(r.RowParts, f.MarkColumnPart, f.MarkSpecial, mark)
	}
	return nil
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

// func (e *Extractor) ApplyFilter(filterFile string) {
// 	if filterFile == "" {
// 		return
// 	}
// 	filterCode, err := GetFilterFileCode(filterFile)
// 	if err != nil {
// 		panic(err)
// 	}
// 	switch filterCode {
// 	}
// }
