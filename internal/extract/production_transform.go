package extract

// TransformBase
func (e *Extractor) TransformBase() {
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

	// e.ComputeName()
	e.ComputeIndex()
	e.TransformHeaderExternal(RowPartCode_HourlyHead, "1000", "planovany_zacatek")

}

// TransformProduction
func (e *Extractor) TransformProduction() {
	// add field column
	// e.AddColumn(RowPartCode_HourlyHead, "kek")

	// Convert dates
	// e.TransformDateToTime(RowPartCode_RadioHead, "1000", false)
	// e.TransformDateToTime(RowPartCode_HourlyHead, "1000", false)
	// e.TransformDateToTime(RowPartCode_SubHead, "1004", false)
	// e.TransformDateToTime(RowPartCode_SubHead, "1003", false)
	// e.TransformDateToTime(RowPartCode_StoryHead, "1004", false)
	// e.TransformDateToTime(RowPartCode_StoryHead, "1004", true)
	// e.TransformDateToTime(RowPartCode_StoryHead, "1003", false)

	// Convert dates
	e.TransformColumnsFields(TransformTimeDate, false, "1000", "1004", "1003")

	// // Convert stopaz
	stopazFields := []string{"38", "1002", "1005", "1010", "1036"}
	e.TransformColumnsFields(TransformStopaz, false, stopazFields...)
	e.TransformColumnFields(RowPartCode_AudioClipHead,
		"1005", TransformStopaz, false)

	// Korekce
	// e.TransformColumnFields(RowPartCode_StoryHead,
	// 	"1029", TransformStopaz, false)
	// e.TransformColumnFields(RowPartCode_StoryHead,
	// 	"1035", TransformStopaz, false)

	// Compute
	// e.ComputeRecordIDs(true)
	// e.SetFileNameColumn()

	// e.TransformColumnsFields(TransformObjectID, false, "TemplateName")
	e.TransformColumnsFields(TransformObjectID, false, "ObjectID")
	e.ComputeJoinNameAndSurname(RowPartCode_ComputedKON, "jmeno_spojene")
}
