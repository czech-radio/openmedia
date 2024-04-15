package extract

import (
	"fmt"
	"log/slog"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type TransformerCode int

const (
	TransformerMock TransformerCode = iota
	TransformerCodedFields
	TransformerProduction
	TransformerEurovolby
)

// var ExtractorTransformerCodeMap map[ExtractorTransformerCode]

func (e *Extractor) Transform(code TransformerCode) {
	switch code {
	case TransformerMock:
		e.TransformMock()
	case TransformerCodedFields:
		e.TransformCodedFields()
	case TransformerProduction:
		e.TransformProduction()
	case TransformerEurovolby:
		e.TransformEurovolby()
	}
}

func (e *Extractor) UniqueRows() {
	if e.CSVtable.UniqueRows == nil {
		e.CSVtable.UniqueRows = make(map[string]int)
	}
	for i, row := range e.CSVtable.Rows {
		rowBuilder := strings.Builder{}
		row.CSVrow.CastToCSV(
			&rowBuilder,
			e.CSVrowPartsPositionsInternal,
			e.CSVrowPartsFieldsPositions, CSVdelim)
		rowStr := rowBuilder.String()
		_, ok := e.CSVtable.UniqueRows[rowStr]
		if !ok {
			e.CSVtable.UniqueRowsOrder = append(e.CSVtable.UniqueRowsOrder, i)
		}
		e.CSVtable.UniqueRows[rowStr]++
	}

	slog.Warn("unique rows count",
		"uniqueCount", len(e.CSVtable.UniqueRowsOrder),
		"allCount", len(e.CSVtable.Rows),
	)
}

func (e *Extractor) FilterByPartAndFieldID(
	partCode PartPrefixCode, fieldID string, fieldValuePatern string) []int {
	var res []int
	re := regexp.MustCompile(fieldValuePatern)
	for i, row := range e.CSVtable.Rows {
		part, ok := row.CSVrow[partCode]
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

func (e *Extractor) RemoveColumn(
	fieldPrefix PartPrefixCode, fieldID string) {
	var newPos CSVrowPartFieldsPositions
	positions := e.CSVrowPartsFieldsPositions[fieldPrefix]
	for _, pos := range positions {
		if pos.FieldID == fieldID {
			continue
		}
		newPos = append(newPos, pos)
	}
	e.CSVrowPartsFieldsPositions[fieldPrefix] = newPos
	e.CreateTablesHeader(CSVdelim)
}

// Geters
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

func GetPartAndField(
	row CSVrow, partCode PartPrefixCode,
	fieldID string) (
	CSVrowPart, CSVrowField, bool) {
	part, ok := row[partCode]
	if !ok {
		return part, CSVrowField{}, ok
	}
	field, ok := part[fieldID]
	if !ok {
		slog.Debug("fieldID not found", "fieldID", fieldID)
		return part, field, ok
	}
	return part, field, ok
}

// Transformers
func (e *Extractor) TransformField(
	partCode PartPrefixCode, fieldID string,
	fun func(string) (string, error), force bool) {
	partName := PartsPrefixMapProduction[partCode].Internal

	for i, row := range e.CSVtable.Rows {
		part, ok := row.CSVrow[partCode]
		if !ok && !force {
			slog.Debug("row part name not found",
				"partName", partName)
			continue
		}
		field, ok := part[fieldID]
		if !ok && !force {
			continue
		}
		name, err := fun(field.Value)
		if err != nil {
			slog.Debug(err.Error())
			continue
		}
		if e.CSVtable.Rows[i].CSVrow[partCode] == nil {
			e.CSVtable.Rows[i].CSVrow[partCode] = make(CSVrowPart)
		}
		field.Value = name
		e.CSVtable.Rows[i].CSVrow[partCode][fieldID] = field
	}
}

func TransformObjectID(objectID string) (string, error) {
	if objectID == "" || objectID == CSVspecialValues[CSVspecialValueChildNotFound] {
		return objectID, nil
	}
	return fmt.Sprintf("ID_%s", objectID), nil
}

func TransformStopaz(stopaz string) (string, error) {
	var sign string
	milliSeconds, err := strconv.ParseInt(stopaz, 10, 64)
	if err != nil {
		return "", err
	}
	if milliSeconds < 0 {
		milliSeconds = -milliSeconds
		sign = "-"
	}
	duration := time.Duration(milliSeconds) * time.Millisecond
	// hoursF := int64(60 * 60 * 1000)
	// hours := milliSeconds / hoursF

	format := "%s%02d:%02d:%02d,%03d"
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
	milis := int(duration.Milliseconds()) % 1000

	value := fmt.Sprintf(
		format, sign, hours, minutes, seconds, milis)
	return value, nil
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
			// "%02d:%02d:%02d,%03d", date.Hour(), date.Minute(), date.Second(), date.Nanosecond()/1000000)
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

func TransformEmptyString(input string) string {
	value := CSVspecialValues[CSVspecialValueEmptyString]
	if input == "" {
		return value
	}
	return input
}

func TransformEmptyToNoContain(input string) (string, error) {
	childVal := CSVspecialValues[CSVspecialValueChildNotFound]
	if input == childVal {
		return CSVspecialValues[CSVspecialValueParentNotFound], nil
	}
	if input == "" {
		return CSVspecialValues[CSVspecialValueParentNotFound], nil
	}
	return input, nil
}

func TransformShortenField(input string) (string, error) {
	targetLength := 248
	if len(input) <= targetLength {
		return input, nil
	} else {
		// Shorten the string to specified characters
		return input[:targetLength], nil
	}
}

// Field constructers/computers

func (row CSVrow) ConsructRecordIDs() string {
	prefixes := []PartPrefixCode{
		FieldPrefix_RadioRec,
		FieldPrefix_HourlyRec,
		FieldPrefix_SubRec,
		FieldPrefix_StoryRec,
	}
	var sb strings.Builder
	var value string
	for _, prefix := range prefixes {
		value = ""
		part, ok := row[prefix]
		if ok {
			value = part["RecordID"].Value
		}
		if value == "" {
			value = "-"
		}
		fmt.Fprintf(&sb, "/%s", value)
	}
	return sb.String()
}

func (e *Extractor) ComputeName() {
	targetFieldID := "jmeno_spojene"
	for i := range e.CSVtable.Rows {
		part, ok := e.CSVtable.Rows[i].CSVrow[FieldPrefix_ContactItemHead]
		if !ok {
			continue
		}
		jmeno := part["421"].Value
		prijmeni := part["422"].Value
		slog.Warn("FUCK", jmeno, prijmeni)
		dstPart, ok := e.CSVtable.Rows[i].CSVrow[FieldPrefix_ComputedID]
		if !ok {
			dstPart = make(CSVrowPart)
		}
		field := CSVrowField{
			FieldID:   targetFieldID,
			FieldName: "",
			Value:     fmt.Sprintf("%s %s", prijmeni, jmeno),
		}
		dstPart[targetFieldID] = field
		e.CSVtable.Rows[i].CSVrow[FieldPrefix_ComputedID] = dstPart
	}
}

func (e *Extractor) ComputeIndex() {
	targetFieldID := "C-index"
	prevIDs := []string{"", "", "", ""}
	prevPos := []int{0, 0, 0}
	var index string
	for i, row := range e.CSVtable.Rows {
		dstPart, ok := e.CSVtable.Rows[i].CSVrow[FieldPrefix_ComputedRID]
		if !ok {
			dstPart = make(CSVrowPart)
		}
		if i == 0 {
			index, prevIDs, _ = ComputeIndexCreate(row.CSVrow, prevIDs, prevPos)
		}
		if i > 0 {
			index, prevIDs, prevPos = ComputeIndexCreate(row.CSVrow, prevIDs, prevPos)
		}
		field := CSVrowField{
			FieldID:   targetFieldID,
			FieldName: "",
			Value:     index,
		}
		dstPart[targetFieldID] = field
		e.CSVtable.Rows[i].CSVrow[FieldPrefix_ComputedRID] = dstPart
	}
}

func ComputeIndexCreate(
	row CSVrow, prevIDs []string, prevPos []int) (string, []string, []int) {
	_, blok, _ := GetPartAndField(
		row, FieldPrefix_HourlyHead, "8")
	_, sub, _ := GetPartAndField(
		row, FieldPrefix_SubHead, "ObjectID")
	_, story, _ := GetPartAndField(
		row, FieldPrefix_StoryHead, "ObjectID")
	_, storyPart, _ := GetPartAndField(
		row, FieldPrefix_StoryKategory, "ObjectID")
	// slog.Warn("compare", "cur", blok.Value, "prev", prevIDs[0])
	if sub.Value == prevIDs[0] {
		goto skip_sub
	}
	if story.Value != prevIDs[2] {
		prevPos[0]++
		if sub.Value == "" {
			prevPos[1] = 0
		} else {
			prevPos[1]++
		}
	}
skip_sub:
	if story.Value == prevIDs[2] {
		prevPos[2]++
	}
	if story.Value != prevIDs[2] {
		prevPos[2] = 1
	}

	if blok.Value != prevIDs[0] {
		// reset for new blok
		prevPos[0] = 1
		prevPos[1] = 0
		prevPos[2] = 0
	}
	prev := []string{blok.Value, sub.Value, story.Value, storyPart.Value}
	index := ComputeIndexCast(blok.Value, prevPos)
	return index, prev, prevPos
}

func ComputeIndexCreateB(
	// Story order
	row CSVrow, prevIDs []string, prevPos []int) (string, []string, []int) {
	_, blok, _ := GetPartAndField(
		row, FieldPrefix_HourlyHead, "8")
	_, sub, _ := GetPartAndField(
		row, FieldPrefix_SubHead, "ObjectID")
	_, story, _ := GetPartAndField(
		row, FieldPrefix_StoryHead, "ObjectID")
	_, storyPart, _ := GetPartAndField(
		row, FieldPrefix_StoryKategory, "ObjectID")
	// slog.Warn("compare", "cur", blok.Value, "prev", prevIDs[0])
	if story.Value != prevIDs[2] {
		prevPos[0]++
	}
	prev := []string{blok.Value, sub.Value, story.Value, storyPart.Value}
	index := ComputeIndexCast(blok.Value, prevPos)
	return index, prev, prevPos
}

func ComputeIndexCast(blok string, pos []int) string {
	var index strings.Builder
	if len(blok) > 5 {
		fmt.Fprintf(&index, "%s", blok[0:5])
	} else {
		fmt.Fprintf(&index, "%s", blok)
	}
	// for i := range pos[0:2] {
	for i := range pos {
		fmt.Fprintf(&index, "/%02d", pos[i])
	}
	return index.String()
}

func (e *Extractor) ComputeRecordIDs(removeSrcColumns bool) {
	targetFieldID := "C-RID"
	for i, row := range e.CSVtable.Rows {
		id := row.ConsructRecordIDs()
		field := CSVrowField{
			FieldID:   targetFieldID,
			FieldName: "",
			Value:     id,
		}
		part, ok := e.CSVtable.Rows[i].CSVrow[FieldPrefix_ComputedRID]
		if !ok {
			part = make(CSVrowPart)
		}
		part[targetFieldID] = field
		e.CSVtable.Rows[i].CSVrow[FieldPrefix_ComputedRID] = part
	}
	if removeSrcColumns {
		e.RemoveColumn(
			FieldPrefix_RadioRec, "RecordID")
		e.RemoveColumn(
			FieldPrefix_HourlyRec, "RecordID")
		e.RemoveColumn(
			FieldPrefix_SubRec, "RecordID")
		e.RemoveColumn(
			FieldPrefix_StoryRec, "RecordID")
	}
}

func (e *Extractor) SetFileNameColumn() {
	targetFieldID := "FileName"
	for i := range e.CSVtable.Rows {
		field := CSVrowField{
			FieldID:   targetFieldID,
			FieldName: "",
			Value:     e.CSVtable.SrcFilePath,
		}
		part, ok := e.CSVtable.Rows[i].CSVrow[FieldPrefix_ComputedRID]
		if !ok {
			part = make(CSVrowPart)
		}
		part[targetFieldID] = field
		e.CSVtable.Rows[i].CSVrow[FieldPrefix_ComputedRID] = part
	}
}

func (e *Extractor) ComputeID() {
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
		e.CSVtable.Rows[i].CSVrow[FieldPrefix_ComputedID] = part
	}
}

func (row CSVrow) ConstructID() string {
	partStoryHead, ok := row[FieldPrefix_StoryHead]
	if !ok {
		return ""
	}
	partHourlyHead, ok := row[FieldPrefix_HourlyHead]
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
	blok := partHourlyHead["8"].Value
	stanice := partStoryHead["5081"].Value
	datum := partStoryHead["datum"].Value
	nazev := partStoryHead["8"].Value
	zacatek := partStoryHead["1004"].Value
	konec := partStoryHead["1003"].Value
	out := fmt.Sprintf("%s/%s/%s/%s - %s/%s",
		stanice, blok, datum, zacatek, konec, nazev)
	return out
}
