package internal

import (
	"log/slog"
	"regexp"
)

// var hoursRegex = regexp.MustCompile("^13:00-14:00")
func (e *Extractor) FilterByPartAndFieldID(partCode PartPrefixCode, fieldID string, fieldValuePatern string) []int {
	var res []int
	re := regexp.MustCompile(fieldValuePatern)
	partName := PartsPrefixMapProduction[partCode].Internal
	for i, row := range e.CSVtable {
		part, ok := row.CSVrow[partName]
		if !ok {
			slog.Debug(
				"filter not effective", "reason", "partname not found",
				"partName", partName,
			)
			return nil
		}
		field, ok := part[fieldID]
		if !ok {
			slog.Debug(
				"filter not effective", "reason", "fieldID not found",
				"partName", partName,
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

func (e *Extractor) TransformPartFieldID(partCode PartPrefixCode, fieldID string, fieldValuePatern string) {
}

func TransformEmptyString(input string) string {
	value := CSVspecialValues[CSVspecialValueEmptyString]
	if input == "" {
		return value
	}
	return input
}
