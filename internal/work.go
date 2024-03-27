package internal

func (e *Extractor) TransformEurovolby(printHeader bool) {
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
	e.ComputeKategory()
	e.RemoveColumn(
		FieldPrefix_ContactItemHead, "TemplateName")
	e.RemoveColumn(
		FieldPrefix_AudioClipHead, "TemplateName")

	// FILTER ROWS
	// rowsIDx := e.Extractor.FilterByPartAndFieldID(internal.FieldPrefix_HourlyHead, "8", "13:00-14:00")
	// e.Extractor.PrintTableRowsToCSV(true, "\t", rowsIDx)
	e.PrintTableRowsToCSV(printHeader, "\t")
}

func ExtractProductionVer1(filePath string, header bool) {
	af := ArchiveFile{}
	err := af.Init(WorkerTypeRundownXMLutf16le, filePath)
	if err != nil {
		Errors.ExitWithCode(err)
	}
	err = af.ExtractByXMLquery(EXTproduction)
	if err != nil {
		Errors.ExitWithCode(err)
	}
	SetLogLevel("-4")

	// af.TransformField(
	// internal.FieldPrefix_StoryHead,
	// "5081", internal.GetRadioName)

	// af.TransformField(
	// internal.FieldPrefix_ContactItemHead,
	// "5088", internal.GetGenderName)

	// Convert dates
	af.TransformDateToTime(FieldPrefix_SubHead, "1004", false)
	af.TransformDateToTime(FieldPrefix_SubHead, "1003", false)
	af.TransformDateToTime(FieldPrefix_StoryHead, "1004", true)
	af.TransformDateToTime(FieldPrefix_StoryHead, "1003", false)

	// Convert stopaz
	af.TransformField(
		FieldPrefix_SubHead,
		"1005", TransformStopaz)

	af.TransformField(
		FieldPrefix_StoryHead,
		"1005", TransformStopaz)

	af.TransformField(
		FieldPrefix_StoryHead,
		"1036", TransformStopaz)

	af.TransformField(
		FieldPrefix_StoryHead,
		"1010", TransformStopaz)

	af.TransformField(
		FieldPrefix_StoryHead,
		"1002", TransformStopaz)

	// korekce
	af.TransformField(
		FieldPrefix_StoryHead,
		"1029", TransformStopaz)

	af.TransformField(
		FieldPrefix_StoryHead,
		"1035", TransformStopaz)

	// Audio
	af.TransformField(
		FieldPrefix_AudioClipHead,
		"1005", TransformStopaz)

	af.ComputeID()
	af.ComputeKategory()
	af.RemoveColumn(
		FieldPrefix_ContactItemHead, "TemplateName")
	af.RemoveColumn(
		FieldPrefix_AudioClipHead, "TemplateName")

	// FILTER ROWS
	// rowsIDx := af.Extractor.FilterByPartAndFieldID(internal.FieldPrefix_HourlyHead, "8", "13:00-14:00")
	// af.Extractor.PrintTableRowsToCSV(true, "\t", rowsIDx)
	af.Extractor.PrintTableRowsToCSV(header, "\t")
}
