package cmd

import (
	"github/czech-radio/openmedia-archive/internal"
)

type ConfigExtractFile struct {
	SourceFile      string `cmd:"source_file; i; ; input file"`
	DestinationFile string `cmd:"destination_file; o; ; otput file"`
	CSVdelim        string `cmd:"csv_delim; csvd; \t; csv field delimiter"`
}

// 4.  března 2024
func RunExtractFile(rootCfg *ConfigRoot, filterCfg *ConfigExtractFile) {
	filePath := "/home/jk/CRO/CRO_BASE/openmedia-archive_backup/Archive/control/control_UTF16_RD_13-17_Plus_Tuesday_W01_2024_01_02.xml"
	internal.ExtractProductionVer1(filePath, true)
}

// func RunExtractFileB(rootCfg *ConfigRoot, filterCfg *ConfigExtractFile) {
// 	filePath := "/home/jk/CRO/CRO_BASE/openmedia-archive_backup/Archive/control/control_UTF16_RD_13-17_Plus_Tuesday_W01_2024_01_02.xml"
// 	af := internal.ArchiveFile{}
// 	err := af.Init(internal.WorkerTypeRundownXMLutf16le, filePath)
// 	if err != nil {
// 		internal.Errors.ExitWithCode(err)
// 	}
// 	err = af.ExtractByXMLquery(internal.EXTproduction)
// 	if err != nil {
// 		internal.Errors.ExitWithCode(err)
// 	}
// 	// af.Ex
// 	// af.Extractor.CastTablesToCSV()
// 	internal.SetLogLevel("-4")

// 	// af.TransformField(
// 	// internal.FieldPrefix_StoryHead,
// 	// "5081", internal.GetRadioName)

// 	// af.TransformField(
// 	// internal.FieldPrefix_ContactItemHead,
// 	// "5088", internal.GetGenderName)

// 	// Convert dates
// 	af.TransformDateToTime(internal.FieldPrefix_SubHead, "1004", false)
// 	af.TransformDateToTime(internal.FieldPrefix_SubHead, "1003", false)
// 	af.TransformDateToTime(internal.FieldPrefix_StoryHead, "1004", true)
// 	af.TransformDateToTime(internal.FieldPrefix_StoryHead, "1003", false)

// 	// Convert stopaz
// 	af.TransformField(
// 		internal.FieldPrefix_SubHead,
// 		"1005", internal.TransformStopaz)

// 	af.TransformField(
// 		internal.FieldPrefix_StoryHead,
// 		"1005", internal.TransformStopaz)

// 	af.TransformField(
// 		internal.FieldPrefix_StoryHead,
// 		"1036", internal.TransformStopaz)

// 	af.TransformField(
// 		internal.FieldPrefix_StoryHead,
// 		"1010", internal.TransformStopaz)

// 	af.TransformField(
// 		internal.FieldPrefix_StoryHead,
// 		"1002", internal.TransformStopaz)

// 	af.TransformField(
// 		internal.FieldPrefix_AudioClipHead,
// 		"1005", internal.TransformStopaz)

// 	af.ComputeID()

// 	// rowsIDx := af.Extractor.FilterByPartAndFieldID(internal.FieldPrefix_HourlyHead, "8", "13:00-14:00")
// 	// af.Extractor.PrintTableRowsToCSV(true, "\t", rowsIDx)
// 	af.Extractor.PrintTableRowsToCSV(true, "\t")
// 	// workerTypes := []internal.WorkerTypeCode{
// 	// internal.WorkerTypeRundownXMLutf16le}
// 	// archiveFile := internal.ArchiveFile{}
// 	// archiveFile.Init(internal.WorkerTypeRundownXMLutf16le)
// }
