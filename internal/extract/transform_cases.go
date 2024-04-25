package extract

// NOTE: docasne nepouzite bude dospecifikovano
func (e *Extractor) TransformCodedFields() {
	e.TransformColumnFields(
		RowPartCode_StoryHead,
		"5081", GetRadioName, false)

	e.TransformColumnFields(
		RowPartCode_ContactItemHead,
		"5088", GetGenderName, false)
}

func (e *Extractor) TransformEurovolby() {
	// Convert dates
	e.TransformDateToTime(RowPartCode_SubHead, "1004", false)
	e.TransformDateToTime(RowPartCode_SubHead, "1003", false)
	e.TransformDateToTime(RowPartCode_StoryHead, "1004", true)
	e.TransformDateToTime(RowPartCode_StoryHead, "1003", false)

	// Convert stopaz
	e.TransformColumnFields(RowPartCode_SubHead,
		"1005", TransformStopaz, false)
	e.TransformColumnFields(RowPartCode_StoryHead,
		"1005", TransformStopaz, false)
	e.TransformColumnFields(RowPartCode_StoryHead,
		"1036", TransformStopaz, false)
	e.TransformColumnFields(RowPartCode_StoryHead,
		"1010", TransformStopaz, false)
	e.TransformColumnFields(RowPartCode_StoryHead,
		"1002", TransformStopaz, false)

	// korekce
	e.TransformColumnFields(RowPartCode_StoryHead,
		"1029", TransformStopaz, false)

	e.TransformColumnFields(RowPartCode_StoryHead,
		"1035", TransformStopaz, false)

	// Audio
	// e.TransformField(
	// RowPartCode_AudioClipHead,
	// "1005", TransformStopaz,false)

	e.ComputeID()
	// RecordIDs
	// e.ComputeRecordIDs(false)

	// FILTER ROWS
	// rowsIDx := e.Extractor.FilterByPartAndFieldID(internal.RowPartCode_HourlyHead, "8", "13:00-14:00")
	// e.Extractor.PrintTableRowsToCSV(true, "\t", rowsIDx)
	// e.PrintTableRowsToCSV(printHeader, "\t")
}

func (e *Extractor) TransformMock() {
	// Convert dates
	e.TransformDateToTime(RowPartCode_SubHead,
		"1004", false)
	e.TransformDateToTime(RowPartCode_SubHead,
		"1003", false)
	e.TransformDateToTime(RowPartCode_StoryHead,
		"1004", true)
	e.TransformDateToTime(RowPartCode_StoryHead,
		"1003", false)

	// Convert stopaz
	e.TransformColumnFields(RowPartCode_SubHead,
		"1005", TransformStopaz, false)

	e.TransformColumnFields(RowPartCode_SubHead,
		"38", TransformStopaz, false)

	e.TransformColumnFields(RowPartCode_StoryHead,
		"1005", TransformStopaz, false)

	e.TransformColumnFields(RowPartCode_StoryHead,
		"1036", TransformStopaz, false)

	e.TransformColumnFields(RowPartCode_StoryHead,
		"1010", TransformStopaz, false)

	e.TransformColumnFields(RowPartCode_StoryHead,
		"1002", TransformStopaz, false)

	// Korekce
	e.TransformColumnFields(RowPartCode_StoryHead,
		"1029", TransformStopaz, false)

	e.TransformColumnFields(RowPartCode_StoryHead,
		"1035", TransformStopaz, false)

	// Audio
	e.TransformColumnFields(RowPartCode_AudioClipHead,
		"1005", TransformStopaz, false)

	// COMPUTE
	e.ComputeID()
	e.ComputeRecordIDs(true)
}
