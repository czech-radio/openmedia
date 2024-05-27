package extract

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
