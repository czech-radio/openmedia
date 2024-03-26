package internal

import (
	"fmt"
	"log/slog"
	"regexp"
	"time"
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

func GetRadioName(radioCode string) (string, error) {
	names, ok := RadioCodes[radioCode]
	if !ok {
		return "", fmt.Errorf("unknown radio code")
	}
	return names.Croapp_shortTitle, nil
}

func GetGenderName(genderCode string) (string, error) {
	gender, ok := GenderCodes[genderCode]
	if !ok {
		return "", fmt.Errorf("unknown gender code: %s", genderCode)
	}
	return gender, nil
}

func (e *Extractor) TransformField(
	partCode PartPrefixCode, fieldID string,
	fun func(string) (string, error)) {
	partName := PartsPrefixMapProduction[partCode].Internal

	for i, row := range e.CSVtable.Rows {
		part, ok := row.CSVrow[partName]
		if !ok {
			slog.Warn("row part name not found", "partName", partName)
			continue
		}
		field, ok := part[fieldID]
		if !ok {
			continue
		}
		name, err := fun(field.Value)
		if err != nil {
			slog.Warn(err.Error())
			continue
		}
		field.Value = name
		e.CSVtable.Rows[i].CSVrow[partName][fieldID] = field
	}
}

func GetPartAndField(
	row CSVrow, fieldPrefix PartPrefixCode,
	fieldID string) (
	CSVrowPart, CSVrowField, bool) {
	partName := PartsPrefixMapProduction[fieldPrefix].Internal
	part, ok := row[partName]
	if !ok {
		return part, CSVrowField{}, ok
	}
	field, ok := part[fieldID]
	if !ok {
		slog.Warn("fieldID not found", "fieldID", fieldID)
		return part, field, ok
	}
	return part, field, ok
}

func (e *Extractor) TransformDateToTime(
	prefixCode PartPrefixCode, fieldID string, addDate bool) {
	for _, row := range e.CSVtable.Rows {
		part, field, ok := GetPartAndField(
			row.CSVrow, prefixCode, fieldID)
		if !ok {
			continue
		}
		input := field.Value
		date, err := ParseXMLdate(input)
		if err != nil {
			continue
		}
		field.Value = fmt.Sprintf(
			"%02d:%02d:%02d", date.Hour(), date.Minute(), date.Second())
		part[fieldID] = field

		datestr := fmt.Sprintf("%02d.%02d.%04d", date.Day(), date.Month(), date.Year())
		if addDate {
			field.Value = datestr
			part["datum"] = field
		}
	}
}

func ParseXMLdate(input string) (time.Time, error) {
	layout := "20060102T150405.000"
	parsedTime, err := time.Parse(layout, input)
	if err != nil {
		slog.Error(err.Error())
		return parsedTime, err
	}
	return parsedTime, err
}

func (e *Extractor) ComputeID() {
	partName := PartsPrefixMapProduction[FieldPrefix_ComputedID].Internal
	targetFieldID := "ID"
	for i, row := range e.CSVtable.Rows {
		id := row.ConstructID()
		field := CSVrowField{
			FieldID:   targetFieldID,
			FieldName: "",
			Value:     id,
		}
		part := make(CSVrowPart)
		part[targetFieldID] = field
		e.CSVtable.Rows[i].CSVrow[partName] = part
	}
}

func (row CSVrow) ConstructID() string {
	partName := PartsPrefixMapProduction[FieldPrefix_StoryHead].Internal
	part, ok := row[partName]
	if !ok {
		return ""
	}
	// Z praktických důvodů bych poprosil o zavedení sloupce [ID] pro označení příspěvku. Půjde o kalkulované pole, které bude mít podobu
	// "[stanice] / [datum] / [cas_zacatku] - [cas konce] / [nazev]".
	// Story, 5081
	// Story, 1004 -> datum
	// Story, 1004 -> cas
	// Story, 1003 -> cas
	// Story, 8 Nazev
	stanice := part["5081"].Value
	datum := part["1004"].Value
	// cas := part["1004"]
	nazev := part["8"].Value
	zacatek := part["1004"].Value
	konec := part["1003"].Value
	out := fmt.Sprintf("%s/%s/%s - %s/%s",
		stanice, datum, zacatek, konec, nazev)
	return out
}

func TransformEmptyString(input string) string {
	value := CSVspecialValues[CSVspecialValueEmptyString]
	if input == "" {
		return value
	}
	return input
}
