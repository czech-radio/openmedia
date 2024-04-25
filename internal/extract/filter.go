package extract

import "log/slog"

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

// FilterRun
func (e *Extractor) FilterRun(f FilterColumn) {
	switch f.FilterName {
	case FilterCodeMatchPersonName:
		e.FilterMatchPersonName(f)
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
func (e *Extractor) FilterMatchPersonName(f FilterColumn) {
	for i, row := range e.TableXML.Rows {
		matches, altValue := FieldValueMatchesValidValues(
			row.Row, f.PartCodeCheck, f.FieldIDcheck, f.Values,
		)
		part, ok := e.TableXML.Rows[i].Row[f.PartCodeMark]
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
		e.TableXML.Rows[i].Row[f.PartCodeMark] = part
	}
}

// FieldValueMatchesValidValues
func FieldValueMatchesValidValues(
	row Row, partCode RowPartCode, fieldID string,
	validValues map[string]bool) (bool, string) {
	_, field, ok := GetRowPartAndField(row, partCode, fieldID)
	notFound := RowFieldValueCodeMap[RowFieldValueChildNotFound]
	if !ok {
		return false, notFound
	}
	return validValues[field.Value], ""
}
