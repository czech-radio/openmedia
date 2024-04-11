package internal

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

func (e *Extractor) TransformCodedFields() {
	// af.TransformField(
	// internal.FieldPrefix_StoryHead,
	// "5081", internal.GetRadioName)

	// af.TransformField(
	// internal.FieldPrefix_ContactItemHead,
	// "5088", internal.GetGenderName)
}

func (e *Extractor) TransformProduction() {
	// Convert dates
	e.TransformDateToTime(FieldPrefix_SubHead, "1004", false)
	e.TransformDateToTime(FieldPrefix_SubHead, "1003", false)
	e.TransformDateToTime(FieldPrefix_StoryHead, "1004", true)
	e.TransformDateToTime(FieldPrefix_StoryHead, "1003", false)

	// Convert stopaz
	e.TransformField(FieldPrefix_SubHead,
		"1005", TransformStopaz, false)
	e.TransformField(FieldPrefix_StoryHead,
		"1036", TransformStopaz, false)
	e.TransformField(FieldPrefix_StoryHead,
		"1005", TransformStopaz, false)
	e.TransformField(FieldPrefix_StoryHead,
		"1010", TransformStopaz, false)
	e.TransformField(FieldPrefix_StoryHead,
		"1002", TransformStopaz, false)
	e.TransformField(FieldPrefix_AudioClipHead,
		"38", TransformStopaz, false)

	// Korekce
	e.TransformField(FieldPrefix_StoryHead,
		"1029", TransformStopaz, false)
	e.TransformField(FieldPrefix_StoryHead,
		"1035", TransformStopaz, false)

	// Audio
	e.TransformField(FieldPrefix_AudioClipHead,
		"1005", TransformStopaz, false)

	// Compute
	e.ComputeID()
	e.TransformField(FieldPrefix_ComputedID,
		"ID", TransformShortenField, false,
	)
	e.ComputeRecordIDs(true)
	e.SetFileNameColumn()

	// Transform Special Values
	e.TransformField(FieldPrefix_SubHead,
		"TemplateName", TransformEmptyToNoContain, true)
	e.TransformField(FieldPrefix_StoryKategory,
		"TemplateName", TransformEmptyToNoContain, true)
	e.ComputeIndex()

	e.TransformField(FieldPrefix_StoryHead, "ObjectID", TransformObjectID, false)
	e.TransformField(FieldPrefix_StoryKategory, "ObjectID", TransformObjectID, false)
	e.TransformField(FieldPrefix_SubHead, "ObjectID", TransformObjectID, false)
}

func (e *Extractor) TransformTest() {
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
