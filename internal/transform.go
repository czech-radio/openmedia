package internal

import (
	"log/slog"
	"regexp"
)

func (e *Extractor) FilterByPartAndFieldID(
	partCode PartPrefixCode, fieldID string, fieldValuePatern string) []int {
	var res []int
	re := regexp.MustCompile(fieldValuePatern)
	partName := PartsPrefixMapProduction[partCode].Internal
	for i, row := range e.CSVtable.Rows {
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

func (e *Extractor) ComputeID(
	// Z praktických důvodů bych poprosil o zavedení sloupce [ID] pro označení příspěvku. Půjde o kalkulované pole, které bude mít podobu
	// "[stanice] / [datum] / [cas_zacatku] - [cas konce] / [nazev]".
	// Story, 5081
	// Story, 1004 -> datum
	// Story, 1004 -> cas
	// Story, 1003 -> cas
	// Story, 8 Nazev

	partCode PartPrefixCode, fieldID string, fieldValuePatern string) {
}

func TransformEmptyString(input string) string {
	value := CSVspecialValues[CSVspecialValueEmptyString]
	if input == "" {
		return value
	}
	return input
}
