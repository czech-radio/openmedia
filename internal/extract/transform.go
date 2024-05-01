package extract

import (
	"fmt"
	ar "github/czech-radio/openmedia/internal/archive"
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
	TransformerProductionCSV
	TransformerEurovolby
)

// Transform
func (e *Extractor) Transform(code TransformerCode) {
	switch code {
	case TransformerMock:
		e.TransformMock()
	case TransformerCodedFields:
		e.TransformCodedFields()
	case TransformerProduction:
		e.TransformProduction()
	case TransformerProductionCSV:
		e.TransformBase()
	case TransformerEurovolby:
		e.TransformEurovolby()
	}
}

// UniqueRows
func (e *Extractor) UniqueRows() {
	if e.TableXML.UniqueRows == nil {
		e.TableXML.UniqueRows = make(map[string]int)
	}
	for i, row := range e.TableXML.Rows {
		rowBuilder := strings.Builder{}
		row.RowParts.CSVrowBuild(
			&rowBuilder,
			e.RowPartsPositions,
			e.RowPartsFieldsPositions, CSVdelim)
		rowStr := rowBuilder.String()
		_, ok := e.TableXML.UniqueRows[rowStr]
		if !ok {
			e.TableXML.UniqueRowsOrder = append(e.TableXML.UniqueRowsOrder, i)
		}
		e.TableXML.UniqueRows[rowStr]++
	}

	slog.Warn("unique rows count",
		"uniqueCount", len(e.TableXML.UniqueRowsOrder),
		"allCount", len(e.TableXML.Rows),
	)
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

// RemoveColumn
func (e *Extractor) RemoveColumn(
	fieldPrefix RowPartCode, fieldID string) {
	var newPos RowPartFieldsPositions
	positions := e.RowPartsFieldsPositions[fieldPrefix]
	for _, pos := range positions {
		if pos.FieldID == fieldID {
			continue
		}
		newPos = append(newPos, pos)
	}
	e.RowPartsFieldsPositions[fieldPrefix] = newPos
	e.HeaderBuild()
}

// AddColumn
func (e *Extractor) AddColumn(fieldPrefix RowPartCode, fieldID string) {
	// NOTE: new part with new fieldPrefix will not be added. Currently must be specified in []OMextractors
	positions := e.RowPartsFieldsPositions[fieldPrefix]
	var newPos RowPartFieldsPositions
	alreadyContains := false
	for _, pos := range positions {
		if pos.FieldID == fieldID {
			alreadyContains = true
		}
		newPos = append(newPos, pos)
	}
	if !alreadyContains {
		fieldPos := RowPartFieldPosition{
			RowPartName: "",
			FieldID:     fieldID,
			FieldName:   "",
		}
		newPos = append(newPos, fieldPos)
	}
	e.RowPartsFieldsPositions[fieldPrefix] = newPos
	e.HeaderBuild()
}

// GETERS
func GetRadioName(radioCode string) (string, error) {
	names, ok := RadioCodes[radioCode]
	if !ok {
		return "", fmt.Errorf("unknown radio code")
	}
	return names.Croapp_shortTitle, nil
}

// GetGenderName
func GetGenderName(genderCode string) (string, error) {
	gender, ok := GenderCodes[genderCode]
	if !ok {
		return "", fmt.Errorf("unknown gender code: %s", genderCode)
	}
	return gender, nil
}

// GetRowPartAndField
func GetRowPartAndField(
	row RowParts, partCode RowPartCode,
	fieldID string) (
	RowPart, RowField, bool) {
	part, ok := row[partCode]
	if !ok {
		return part, RowField{}, ok
	}
	field, ok := part[fieldID]
	if !ok {
		slog.Debug("fieldID not found", "fieldID", fieldID)
		return part, field, ok
	}
	return part, field, ok
}

// TRANSFORMERS

// TransformColumnFields transform specified column fields with tranformer function. Force will transform field even if it is missing.
func (e *Extractor) TransformColumnFields(
	partCode RowPartCode, fieldID string,
	fun func(string) (string, error), force bool) {
	partName := RowPartsCodeMapProduction[partCode].Internal

	for i, row := range e.TableXML.Rows {
		part, ok := row.RowParts[partCode]
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
		if e.TableXML.Rows[i].RowParts[partCode] == nil {
			e.TableXML.Rows[i].RowParts[partCode] = make(RowPart)
		}
		field.Value = name
		e.TableXML.Rows[i].RowParts[partCode][fieldID] = field
	}
}

// TransformColumnsFields
func (e *Extractor) TransformColumnsFields(
	fun func(string) (string, error),
	force bool, fieldIDs ...string) {
	for partCode := range e.RowPartsFieldsPositions {
		for _, field := range fieldIDs {
			e.TransformColumnFields(partCode, field, fun, force)
		}
	}
}

func (e *Extractor) TransformHeaderExternal(
	rowPartCode RowPartCode, fieldID, newName string) {
	part, ok := e.RowPartsFieldsPositions[rowPartCode]
	if !ok {
		return
	}
	for i, f := range part {
		if f.FieldID == fieldID {
			part[i].FieldName = newName
		}
	}
	e.HeaderBuild()
	// slog.Warn("fucked", "external", e.HeaderExternal)
	// helper.Sleeper(10, "s")
}

// TransformObjectID
func TransformObjectID(objectID string) (string, error) {
	specVal := CheckIfMapContainsKeyValue(RowFieldSpecialValueCodeMap, objectID)
	if objectID == "" || specVal {
		return objectID, nil
	}
	return fmt.Sprintf("ID_%s", objectID), nil
}

// TransformTimeDate
func TransformTimeDate(dateString string) (string, error) {
	date, err := ParseXMLdate(dateString)
	if err != nil {
		return "", err
	}
	return date.Format("2006-01-02 15:04:05"), nil
}

// RD_00-05_Radiožurnál_Sunday_W13_2024_03_31.xml
var temaregex = regexp.MustCompile(`(\d\d)-`)

// TransformTema
func TransformTema(tema string) (string, error) {
	var sb strings.Builder
	res := temaregex.FindAllStringSubmatch(tema, -1)
	if len(res) == 0 {
		return tema, nil
	}
	for _, r := range res {
		fmt.Fprintf(&sb, "%s;", r[1])
	}
	return sb.String(), nil
}

// TransformStopaz
func TransformStopaz(stopaz string) (string, error) {
	var sign string
	milliSeconds, err := strconv.ParseInt(stopaz, 10, 64)
	if err != nil {
		return stopaz, err
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

// TransformDateToTime
func (e *Extractor) TransformDateToTime(
	prefixCode RowPartCode, fieldID string, addDate bool) {
	for _, row := range e.TableXML.Rows {
		part, field, ok := GetRowPartAndField(
			row.RowParts, prefixCode, fieldID)
		if !ok {
			continue
		}
		input := field.Value
		date, err := ParseXMLdate(input)
		if err != nil {
			continue
		}
		field.Value = fmt.Sprintf(
			// "%02d:%02d:%02d,%03d",
			// date.Hour(), date.Minute(), date.Second(), date.Nanosecond()/1000000)
			"%02d:%02d:%02d",
			date.Hour(), date.Minute(), date.Second(),
		)
		part[fieldID] = field

		datestr := fmt.Sprintf("%02d.%02d.%04d", date.Day(), date.Month(), date.Year())
		if addDate {
			field.Value = datestr
			part["datum"] = field
			e.AddColumn(prefixCode, "datum")
		}
	}
}

// ParseXMLdate
func ParseXMLdate(input string) (time.Time, error) {
	layout := ar.DateLayout_RundownDate
	parsedTime, err := time.Parse(layout, input)
	if err != nil {
		return parsedTime, err
	}
	return parsedTime, err
}

// TransformEmptyString
func TransformEmptyString(input string) string {
	value := RowFieldSpecialValueCodeMap[RowFieldValueEmptyString]
	if input == "" {
		return value
	}
	return input
}

// TransformEmptyRowPart transform whole row part fields to special value if
// the part was not extracted from xml.
func (e *Extractor) TransformEmptyRowPart() {
	specValNotPossible := RowFieldSpecialValueCodeMap[RowFieldValueNotPossible]
	for _, row := range e.TableXML.Rows {
		for _, partCode := range e.RowPartsPositions {
			_, ok := row.RowParts[partCode]
			if ok {
				continue
			}
			part := make(RowPart)
			fields := e.RowPartsFieldsPositions[partCode]
			for _, field := range fields {
				part[field.FieldID] = RowField{
					FieldID:   field.FieldID,
					FieldName: "",
					Value:     specValNotPossible,
				}
			}
			row.RowParts[partCode] = part
		}
	}
}

// TransformEmptyToNoContain
func TransformEmptyToNoContain(input string) (string, error) {
	childVal := RowFieldSpecialValueCodeMap[RowFieldValueChildNotFound]
	parentVal := RowFieldSpecialValueCodeMap[RowFieldValueParentNotFound]
	out := input
	switch input {
	case childVal, parentVal, "":
		out = RowFieldSpecialValueCodeMap[RowFieldValueParentNotFound]
	}
	return out, nil
}

// TransformShortenField
func TransformShortenField(input string) (string, error) {
	targetLength := 248
	if len(input) <= targetLength {
		return input, nil
	} else {
		// Shorten the string to specified characters
		return input[:targetLength], nil
	}
}

// FIELD CONSTRUCTERS/COMPUTERS

// ConsructRecordIDs
func (row RowParts) ConsructRecordIDs() string {
	prefixes := []RowPartCode{
		RowPartCode_RadioRec,
		RowPartCode_HourlyRec,
		RowPartCode_SubRec,
		RowPartCode_StoryRec,
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

func (e *Extractor) ComputeJoinNameAndSurname(
	newColumnRowPart RowPartCode, newColumnName string) {
	// targetFieldID := "jmeno_spojene"
	targetFieldID := newColumnName
	e.AddColumn(newColumnRowPart, newColumnName)
	for i := range e.TableXML.Rows {
		part, ok := e.TableXML.Rows[i].RowParts[RowPartCode_ContactItemHead]
		if !ok {
			continue
		}
		jmeno := part["421"].Value
		prijmeni := part["422"].Value
		// dstPart, ok := e.TableXML.Rows[i].RowParts[RowPartCode_ComputedKON]
		dstPart, ok := e.TableXML.Rows[i].RowParts[newColumnRowPart]
		if !ok {
			dstPart = make(RowPart)
		}
		field := RowField{
			FieldID:   targetFieldID,
			FieldName: "",
			Value:     fmt.Sprintf("%s %s", prijmeni, jmeno),
		}
		dstPart[targetFieldID] = field
		e.TableXML.Rows[i].RowParts[newColumnRowPart] = dstPart
		// e.TableXML.Rows[i].RowParts[RowPartCode_ComputedKON] = dstPart
	}
}

// ComputeIndex
func (e *Extractor) ComputeIndex() {
	targetFieldID := "C-index"
	var value string
	indexComponents := IndexComponents{}
	for i, row := range e.TableXML.Rows {
		dstPart, ok := row.RowParts[RowPartCode_ComputedRID]
		if !ok {
			dstPart = make(RowPart)
		}
		value, indexComponents = ComputeIndexCreate(
			row.RowParts, indexComponents)
		field := RowField{
			FieldID:   targetFieldID,
			FieldName: "",
			Value:     value,
		}
		dstPart[targetFieldID] = field
		e.TableXML.Rows[i].RowParts[RowPartCode_ComputedRID] = dstPart
	}
}

// IndexComponents
type IndexComponents struct {
	RundownIndex, BlockIndex, StoryIndex int
	RundownPrev, BlockPrev, StoryPrev    string
}

// ComputeIndexCreate
func ComputeIndexCreate(
	row RowParts, comps IndexComponents) (
	string, IndexComponents) {
	// var indexBlock int
	_, nazev, _ := GetRowPartAndField(
		row, RowPartCode_RadioHead, "8")
	_, stanice, _ := GetRowPartAndField(
		row, RowPartCode_RadioHead, "5081")
	_, dateTime, _ := GetRowPartAndField(
		row, RowPartCode_RadioHead, "1000")
	_, story, _ := GetRowPartAndField(
		row, RowPartCode_StoryHead, "ObjectID")
	date, err := ParseXMLdate(dateTime.Value)
	var dateStr string
	if err == nil {
		dateStr = date.Format("2006-01-02")
	}
	if err != nil {
		dateStr = dateTime.Value
	}
	_, blok, _ := GetRowPartAndField(
		row, RowPartCode_HourlyHead, "8")

	if comps.RundownPrev != nazev.Value {
		comps.BlockIndex = 0
		comps.StoryIndex = 0
		comps.RundownPrev = nazev.Value
	}
	if comps.BlockPrev != blok.Value {
		comps.BlockIndex++
		comps.StoryIndex = 0
		comps.BlockPrev = blok.Value
	}

	if comps.StoryPrev != story.Value {
		comps.StoryIndex++
		comps.StoryPrev = story.Value
	}

	res := fmt.Sprintf(
		"%s/%s/%s/%02d/%02d",
		stanice.Value, dateStr, nazev.Value[0:5],
		// stanice.Value, dateStr, blok.Value,
		comps.BlockIndex, comps.StoryIndex)
	return res, comps
}

// ComputeIndexOld
func (e *Extractor) ComputeIndexOld() {
	targetFieldID := "C-index"
	prevIDs := []string{"", "", "", ""}
	prevPos := []int{0, 0, 0}
	var index string
	for i, row := range e.TableXML.Rows {
		dstPart, ok := e.TableXML.Rows[i].RowParts[RowPartCode_ComputedRID]
		if !ok {
			dstPart = make(RowPart)
		}
		if i == 0 {
			index, prevIDs, _ = ComputeIndexCreateOld(row.RowParts, prevIDs, prevPos)
		}
		if i > 0 {
			index, prevIDs, prevPos = ComputeIndexCreateOld(row.RowParts, prevIDs, prevPos)
		}
		field := RowField{
			FieldID:   targetFieldID,
			FieldName: "",
			Value:     index,
		}
		dstPart[targetFieldID] = field
		e.TableXML.Rows[i].RowParts[RowPartCode_ComputedRID] = dstPart
	}
}

// ComputeIndexCreateOld
func ComputeIndexCreateOld(
	row RowParts, prevIDs []string, prevPos []int) (string, []string, []int) {
	_, blok, _ := GetRowPartAndField(
		row, RowPartCode_HourlyHead, "8")
	_, sub, _ := GetRowPartAndField(
		row, RowPartCode_SubHead, "ObjectID")
	_, story, _ := GetRowPartAndField(
		row, RowPartCode_StoryHead, "ObjectID")
	_, storyPart, _ := GetRowPartAndField(
		row, RowPartCode_StoryKategory, "ObjectID")
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

// ComputeIndexCreateB
func ComputeIndexCreateB(
	// Story order
	row RowParts, prevIDs []string, prevPos []int) (string, []string, []int) {
	_, blok, _ := GetRowPartAndField(
		row, RowPartCode_HourlyHead, "8")
	_, sub, _ := GetRowPartAndField(
		row, RowPartCode_SubHead, "ObjectID")
	_, story, _ := GetRowPartAndField(
		row, RowPartCode_StoryHead, "ObjectID")
	_, storyPart, _ := GetRowPartAndField(
		row, RowPartCode_StoryKategory, "ObjectID")
	// slog.Warn("compare", "cur", blok.Value, "prev", prevIDs[0])
	if story.Value != prevIDs[2] {
		prevPos[0]++
	}
	prev := []string{blok.Value, sub.Value, story.Value, storyPart.Value}
	index := ComputeIndexCast(blok.Value, prevPos)
	return index, prev, prevPos
}

// ComputeIndexCast
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

// ComputeRecordIDs
func (e *Extractor) ComputeRecordIDs(removeSrcColumns bool) {
	targetFieldID := "C-RID"
	for i, row := range e.TableXML.Rows {
		id := row.ConsructRecordIDs()
		field := RowField{
			FieldID:   targetFieldID,
			FieldName: "",
			Value:     id,
		}
		part, ok := e.TableXML.Rows[i].RowParts[RowPartCode_ComputedRID]
		if !ok {
			part = make(RowPart)
		}
		part[targetFieldID] = field
		e.TableXML.Rows[i].RowParts[RowPartCode_ComputedRID] = part
	}
	if removeSrcColumns {
		e.RemoveColumn(
			RowPartCode_RadioRec, "RecordID")
		e.RemoveColumn(
			RowPartCode_HourlyRec, "RecordID")
		e.RemoveColumn(
			RowPartCode_SubRec, "RecordID")
		e.RemoveColumn(
			RowPartCode_StoryRec, "RecordID")
	}
}

// SetFileNameColumn
func (e *Extractor) SetFileNameColumn() {
	targetFieldID := "FileName"
	for i := range e.TableXML.Rows {
		field := RowField{
			FieldID:   targetFieldID,
			FieldName: "",
			Value:     e.TableXML.SrcFilePath,
		}
		part, ok := e.TableXML.Rows[i].RowParts[RowPartCode_ComputedRID]
		if !ok {
			part = make(RowPart)
		}
		part[targetFieldID] = field
		e.TableXML.Rows[i].RowParts[RowPartCode_ComputedRID] = part
	}
}

// ComputeID
func (e *Extractor) ComputeID() {
	targetFieldID := "ID"
	for i, row := range e.TableXML.Rows {
		id := row.ConstructID()
		field := RowField{
			FieldID:   targetFieldID,
			FieldName: "",
			Value:     id,
		}
		part := make(RowPart)
		part[targetFieldID] = field
		e.TableXML.Rows[i].RowParts[RowPartCode_ComputedKON] = part
	}
}

// ConstructID
func (row RowParts) ConstructID() string {
	partStoryHead, ok := row[RowPartCode_StoryHead]
	if !ok {
		return ""
	}
	partHourlyHead, ok := row[RowPartCode_HourlyHead]
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
