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
	for i, row := range e.CSVtable.Rows {
		matches, altValue := FieldValueMatchesValidValues(
			row.CSVrow, f.PartCodeCheck, f.FieldIDcheck, f.Values,
		)
		part, ok := e.CSVtable.Rows[i].CSVrow[f.PartCodeMark]
		if !ok {
			part = make(CSVrowPart)
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
		field := CSVrowField{
			FieldID:   f.FieldIDmark,
			FieldName: "",
			Value:     mark,
		}
		part[f.FieldIDmark] = field
		e.CSVtable.Rows[i].CSVrow[f.PartCodeMark] = part
	}
}

// FieldValueMatchesValidValues
func FieldValueMatchesValidValues(
	row CSVrow, partCode RowPartCode, fieldID string,
	validValues map[string]bool) (bool, string) {
	_, field, ok := GetPartAndField(row, partCode, fieldID)
	notFound := CSVspecialValues[CSVspecialValueChildNotFound]
	if !ok {
		return false, notFound
	}
	return validValues[field.Value], ""
}
