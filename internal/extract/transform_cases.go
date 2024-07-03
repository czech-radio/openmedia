package extract

// TransformBase
func (e *Extractor) TransformBaseB() {
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

	// Stopaz
	stopazFields := []string{"38", "1002", "1005", "1010", "1036"}
	e.TransformColumnsFields(ValidateStopaz, false, stopazFields...)

	e.ComputeIndex()
	e.TransformHeaderExternal(RowPartCode_HourlyHead, "1000", "planovany_zacatek")

}

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

func (e *Extractor) TransformSpecialValues() {
	e.TransformColumnFields(RowPartCode_SubHead,
		"ObjectID", TransformEmptyToNoContain, true)
	e.TransformColumnFields(RowPartCode_StoryKategory,
		"TemplateName", TransformEmptyToNoContain, true)
	e.TransformColumnFields(RowPartCode_StoryKategory,
		"TemplateName", TransformEmptyToNoContain, true)
}

func (e *Extractor) TransformBase() {
	// e.AddColumn(RowPartCode_ComputedRID, "FileName")
	e.RowPartOmit(RowPartCode_StoryRec)
	indxs := e.FilterStoryPartRecordsDuds()
	e.DeleteNonMatchingRows(indxs)

	e.TreatStoryRecordsWithoutOMobject()
	e.TransformEmptyRowPart()
	e.TransformSpecialValues()
	e.ComputeIndex()

	indxs = e.FilterStoryPartsEmptyDupes()
	e.DeleteNonMatchingRows(indxs)
	indxs = e.FilterStoryPartsRedundant()
	e.DeleteNonMatchingRows(indxs)

	e.TransformHeaderExternal(RowPartCode_HourlyHead, "1000", "planovany_zacatek")
}

func (e *Extractor) TransformBeforeValidation() {
	e.TransformColumnsFields(TransformTema, false, "5016")
}

// TransformProduction
func (e *Extractor) TransformProduction() {
	// Convert dates
	e.TransformColumnsFields(TransformTimeDate, false, "1000", "1004", "1003")

	// // Convert stopaz
	// korekce 1029, 1035
	stopazFields := []string{"38", "1002", "1005", "1010", "1036"}
	e.TransformColumnsFields(TransformStopaz, false, stopazFields...)
	e.TransformColumnFields(RowPartCode_AudioClipHead,
		"1005", TransformStopaz, false)

	// Compute
	// e.ComputeRecordIDs(true)
	// e.SetFileNameColumn()
	e.TransformColumnsFields(TransformObjectID, false, "ObjectID")
	e.ComputeJoinNameAndSurname(RowPartCode_ComputedKON, "jmeno_spojene")
}
