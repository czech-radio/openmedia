package extract

// TransformProductionCSV
func (e *Extractor) TransformProductionCSV() {
	e.TransformColumnFields(RowPartCode_StoryHead,
		"5016", TransformTema, false)
	e.TransformColumnFields(RowPartCode_ContactItemHead,
		"5016", TransformTema, false)
	e.TransformColumnFields(RowPartCode_SubHead,
		"ObjectID", TransformEmptyToNoContain, true)
	e.TransformColumnFields(RowPartCode_StoryKategory,
		"TemplateName", TransformEmptyToNoContain, true)
	e.TransformColumnFields(RowPartCode_StoryKategory,
		"TemplateName", TransformEmptyToNoContain, true)

	e.ComputeName()
	e.ComputeIndex()
}

// TransformProduction
func (e *Extractor) TransformProduction() {
	// Convert dates
	e.TransformDateToTime(RowPartCode_SubHead, "1004", false)
	e.TransformDateToTime(RowPartCode_SubHead, "1003", false)
	e.TransformDateToTime(RowPartCode_StoryHead, "1004", true)
	e.TransformDateToTime(RowPartCode_StoryHead, "1003", false)

	e.TransformColumnFields(RowPartCode_HourlyHead,
		"1000", TransformTimeDate, false)

	// Convert stopaz
	stopazFields := []string{"38", "1002", "1005", "1010", "1036"}
	e.TransformColumnsFields(TransformStopaz, false, stopazFields...)

	// Korekce
	e.TransformColumnFields(RowPartCode_StoryHead,
		"1029", TransformStopaz, false)
	e.TransformColumnFields(RowPartCode_StoryHead,
		"1035", TransformStopaz, false)

	// Audio
	e.TransformColumnFields(RowPartCode_AudioClipHead,
		"1005", TransformStopaz, false)

	// Compute
	e.ComputeID()
	e.TransformColumnFields(RowPartCode_ComputedKON,
		"ID", TransformShortenField, false,
	)
	e.ComputeRecordIDs(true)
	e.SetFileNameColumn()

	// Transform Special Values
	e.TransformColumnsFields(TransformEmptyToNoContain, true, "TemplateName")
	e.ComputeIndex()

	e.TransformColumnsFields(TransformObjectID, false, "TemplateName")
	e.ComputeName()
}
