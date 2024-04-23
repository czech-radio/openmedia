package extract

// NOTE: docasne nepouzite bude dospecifikovano
func (e *Extractor) TransformCodedFields() {
	e.TransformField(
		FieldPrefix_StoryHead,
		"5081", GetRadioName, false)

	e.TransformField(
		FieldPrefix_ContactItemHead,
		"5088", GetGenderName, false)
}

func (e *Extractor) TransformEurovolby() {
	// Convert dates
	e.TransformDateToTime(FieldPrefix_SubHead, "1004", false)
	e.TransformDateToTime(FieldPrefix_SubHead, "1003", false)
	e.TransformDateToTime(FieldPrefix_StoryHead, "1004", true)
	e.TransformDateToTime(FieldPrefix_StoryHead, "1003", false)

	// Convert stopaz
	e.TransformField(FieldPrefix_SubHead,
		"1005", TransformStopaz, false)
	e.TransformField(FieldPrefix_StoryHead,
		"1005", TransformStopaz, false)
	e.TransformField(FieldPrefix_StoryHead,
		"1036", TransformStopaz, false)
	e.TransformField(FieldPrefix_StoryHead,
		"1010", TransformStopaz, false)
	e.TransformField(FieldPrefix_StoryHead,
		"1002", TransformStopaz, false)

	// korekce
	e.TransformField(FieldPrefix_StoryHead,
		"1029", TransformStopaz, false)

	e.TransformField(FieldPrefix_StoryHead,
		"1035", TransformStopaz, false)

	// Audio
	// e.TransformField(
	// FieldPrefix_AudioClipHead,
	// "1005", TransformStopaz,false)

	e.ComputeID()
	// RecordIDs
	// e.ComputeRecordIDs(false)

	// FILTER ROWS
	// rowsIDx := e.Extractor.FilterByPartAndFieldID(internal.FieldPrefix_HourlyHead, "8", "13:00-14:00")
	// e.Extractor.PrintTableRowsToCSV(true, "\t", rowsIDx)
	// e.PrintTableRowsToCSV(printHeader, "\t")
}

func (e *Extractor) TransformMock() {
	// Convert dates
	e.TransformDateToTime(FieldPrefix_SubHead,
		"1004", false)
	e.TransformDateToTime(FieldPrefix_SubHead,
		"1003", false)
	e.TransformDateToTime(FieldPrefix_StoryHead,
		"1004", true)
	e.TransformDateToTime(FieldPrefix_StoryHead,
		"1003", false)

	// Convert stopaz
	e.TransformField(FieldPrefix_SubHead,
		"1005", TransformStopaz, false)

	e.TransformField(FieldPrefix_SubHead,
		"38", TransformStopaz, false)

	e.TransformField(FieldPrefix_StoryHead,
		"1005", TransformStopaz, false)

	e.TransformField(FieldPrefix_StoryHead,
		"1036", TransformStopaz, false)

	e.TransformField(FieldPrefix_StoryHead,
		"1010", TransformStopaz, false)

	e.TransformField(FieldPrefix_StoryHead,
		"1002", TransformStopaz, false)

	// Korekce
	e.TransformField(FieldPrefix_StoryHead,
		"1029", TransformStopaz, false)

	e.TransformField(FieldPrefix_StoryHead,
		"1035", TransformStopaz, false)

	// Audio
	e.TransformField(FieldPrefix_AudioClipHead,
		"1005", TransformStopaz, false)

	// COMPUTE
	e.ComputeID()
	e.ComputeRecordIDs(true)
}
