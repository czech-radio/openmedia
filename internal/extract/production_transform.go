package extract

func (e *Extractor) TransformProductionCSV() {
	e.TransformField(FieldPrefix_StoryHead,
		"5016", TransformTema, false)
	e.TransformField(FieldPrefix_ContactItemHead,
		"5016", TransformTema, false)
	e.TransformField(FieldPrefix_SubHead,
		"ObjectID", TransformEmptyToNoContain, true)
	e.TransformField(FieldPrefix_StoryKategory,
		"TemplateName", TransformEmptyToNoContain, true)
	e.ComputeName()
	e.ComputeIndex()
}

func (e *Extractor) TransformProduction() {
	// Convert dates
	e.TransformDateToTime(FieldPrefix_SubHead, "1004", false)
	e.TransformDateToTime(FieldPrefix_SubHead, "1003", false)
	e.TransformDateToTime(FieldPrefix_StoryHead, "1004", true)
	e.TransformDateToTime(FieldPrefix_StoryHead, "1003", false)

	e.TransformField(FieldPrefix_HourlyHead,
		"1000", TransformTimeDate, false)

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
	e.TransformField(FieldPrfix_ComputedKON,
		"ID", TransformShortenField, false,
	)
	e.ComputeRecordIDs(true)
	e.SetFileNameColumn()

	// Transform Special Values
	e.TransformField(FieldPrefix_SubHead,
		"TemplateName", TransformEmptyToNoContain, true)
	e.TransformField(FieldPrefix_StoryKategory,
		"TemplateName", TransformEmptyToNoContain, true)
	e.ComputeIndexOld()

	e.TransformField(FieldPrefix_HourlyHead,
		"ObjectID", TransformObjectID, false)
	e.TransformField(FieldPrefix_StoryHead,
		"ObjectID", TransformObjectID, false)
	e.TransformField(FieldPrefix_StoryKategory,
		"ObjectID", TransformObjectID, false)
	e.TransformField(FieldPrefix_SubHead,
		"ObjectID", TransformObjectID, false)
	e.ComputeName()
}
