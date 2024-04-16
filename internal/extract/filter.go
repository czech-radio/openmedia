package extract

import "log/slog"

type FilterCode int

const (
	FilterCodeMatchPersonName FilterCode = iota
)

type FilterColumn struct {
	FilterName     FilterCode
	PartCodeCheck  PartPrefixCode
	FieldIDcheck   string
	PartCodeMark   PartPrefixCode
	FieldIDmark    string
	FileWithValues string
	Values         map[string]bool
}

var FilterCodeMap = map[FilterCode]FilterColumn{
	FilterCodeMatchPersonName: {
		FilterCodeMatchPersonName,
		FieldPrefix_ComputedID, "jmeno_spojene",
		FieldPrefix_ContactItemHead, "filtered", "", nil},
}

func (e *Extractor) FilterRun(f FilterColumn) {
	switch f.FilterName {
	case FilterCodeMatchPersonName:
		e.FilterMatchPersonName(f)
	}
}

func (e *Extractor) FiltersRun(filters []FilterColumn) {
	if filters == nil {
		slog.Debug("no filters to filter column specified")
		return
	}
	for f := range filters {
		e.FilterRun(filters[f])
	}
}

func (e *Extractor) FilterMatchPersonName(f FilterColumn) {
	for i, row := range e.CSVtable.Rows {
		matches, altValue := FieldValueMatchesValidValues(
			row.CSVrow, f.PartCodeCheck, f.FieldIDcheck, f.Values)
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
		// no need, but to be sure
		// e.CSVtable.Rows[i].CSVrow[f.PartCodeMark] = part
	}
}

func FieldValueMatchesValidValues(
	row CSVrow, partCode PartPrefixCode, fieldID string,
	validValues map[string]bool) (bool, string) {
	_, field, ok := GetPartAndField(row, partCode, fieldID)
	notFound := CSVspecialValues[CSVspecialValueChildNotFound]
	if !ok {
		return false, notFound
	}
	return validValues[field.Value], ""
}
