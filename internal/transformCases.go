package internal

func (e *Extractor) TransformEurovolby() {
	// Convert dates
	e.TransformDateToTime(FieldPrefix_SubHead, "1004", false)
	e.TransformDateToTime(FieldPrefix_SubHead, "1003", false)
	e.TransformDateToTime(FieldPrefix_StoryHead, "1004", true)
	e.TransformDateToTime(FieldPrefix_StoryHead, "1003", false)

	// Convert stopaz
	e.TransformField(
		FieldPrefix_SubHead,
		"1005", TransformStopaz)

	e.TransformField(
		FieldPrefix_StoryHead,
		"1005", TransformStopaz)

	e.TransformField(
		FieldPrefix_StoryHead,
		"1036", TransformStopaz)

	e.TransformField(
		FieldPrefix_StoryHead,
		"1010", TransformStopaz)

	e.TransformField(
		FieldPrefix_StoryHead,
		"1002", TransformStopaz)

	// korekce
	e.TransformField(
		FieldPrefix_StoryHead,
		"1029", TransformStopaz)

	e.TransformField(
		FieldPrefix_StoryHead,
		"1035", TransformStopaz)

	// Audio
	// e.TransformField(
	// FieldPrefix_AudioClipHead,
	// "1005", TransformStopaz)

	e.ComputeID()
	// RecordIDs
	// e.ComputeRecordIDs(false)

	// FILTER ROWS
	// rowsIDx := e.Extractor.FilterByPartAndFieldID(internal.FieldPrefix_HourlyHead, "8", "13:00-14:00")
	// e.Extractor.PrintTableRowsToCSV(true, "\t", rowsIDx)
	// e.PrintTableRowsToCSV(printHeader, "\t")
}

func (e *Extractor) TransformProduction() {
	// af.TransformField(
	// internal.FieldPrefix_StoryHead,
	// "5081", internal.GetRadioName)

	// af.TransformField(
	// internal.FieldPrefix_ContactItemHead,
	// "5088", internal.GetGenderName)

	// Convert dates
	e.TransformDateToTime(FieldPrefix_SubHead, "1004", false)
	e.TransformDateToTime(FieldPrefix_SubHead, "1003", false)
	e.TransformDateToTime(FieldPrefix_StoryHead, "1004", true)
	e.TransformDateToTime(FieldPrefix_StoryHead, "1003", false)

	// Convert stopaz
	e.TransformField(
		FieldPrefix_SubHead,
		"1005", TransformStopaz)

	e.TransformField(
		FieldPrefix_StoryHead,
		"1005", TransformStopaz)

	e.TransformField(
		FieldPrefix_StoryHead,
		"1036", TransformStopaz)

	e.TransformField(
		FieldPrefix_StoryHead,
		"1010", TransformStopaz)

	e.TransformField(
		FieldPrefix_StoryHead,
		"1002", TransformStopaz)

	// korekce
	e.TransformField(
		FieldPrefix_StoryHead,
		"1029", TransformStopaz)

	e.TransformField(
		FieldPrefix_StoryHead,
		"1035", TransformStopaz)

	// Audio
	e.TransformField(
		FieldPrefix_AudioClipHead,
		"1005", TransformStopaz)

	e.ComputeID()
	e.ComputeRecordIDs(true)
}

func (e *Extractor) TransformTest() {
	// Convert dates
	e.TransformDateToTime(FieldPrefix_SubHead, "1004", false)
	e.TransformDateToTime(FieldPrefix_SubHead, "1003", false)
	e.TransformDateToTime(FieldPrefix_StoryHead, "1004", true)
	e.TransformDateToTime(FieldPrefix_StoryHead, "1003", false)

	// Convert stopaz
	e.TransformField(
		FieldPrefix_SubHead,
		"1005", TransformStopaz)

	e.TransformField(
		FieldPrefix_StoryHead,
		"1005", TransformStopaz)

	e.TransformField(
		FieldPrefix_StoryHead,
		"1036", TransformStopaz)

	e.TransformField(
		FieldPrefix_StoryHead,
		"1010", TransformStopaz)

	e.TransformField(
		FieldPrefix_StoryHead,
		"1002", TransformStopaz)

	// korekce
	e.TransformField(
		FieldPrefix_StoryHead,
		"1029", TransformStopaz)

	e.TransformField(
		FieldPrefix_StoryHead,
		"1035", TransformStopaz)

	// Audio
	e.TransformField(
		FieldPrefix_AudioClipHead,
		"1005", TransformStopaz)

	// COMPUTE
	e.ComputeID()
	// e.ComputeRecordIDs(false)
	e.ComputeRecordIDs(true)
}
