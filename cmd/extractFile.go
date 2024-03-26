package cmd

import (
	"github/czech-radio/openmedia-archive/internal"
)

type ConfigExtractFile struct {
	SourceFile      string `cmd:"source_file; i; ; input file"`
	DestinationFile string `cmd:"destination_file; o; ; otput file"`
	CSVdelim        string `cmd:"csv_delim; csvd; \t; csv field delimiter"`
}

func RunExtractFile(rootCfg *ConfigRoot, filterCfg *ConfigExtractFile) {
	filePath := "/home/jk/CRO/CRO_BASE/openmedia-archive_backup/Archive/control/control_UTF16_RD_13-17_Plus_Tuesday_W01_2024_01_02.xml"
	af := internal.ArchiveFile{}
	err := af.Init(internal.WorkerTypeRundownXMLutf16le, filePath)
	if err != nil {
		internal.Errors.ExitWithCode(err)
	}
	err = af.ExtractByXMLquery(internal.EXTproduction)
	if err != nil {
		internal.Errors.ExitWithCode(err)
	}
	// af.Ex
	// af.Extractor.CastTablesToCSV()
	internal.SetLogLevel("-4")
	rowsIDx := af.Extractor.FilterByPartAndFieldID(internal.FieldPrefix_HourlyHead, "8", "13:00-14:00")

	af.TransformField(
		internal.FieldPrefix_StoryHead,
		"5081", internal.GetRadioName)
	af.TransformField(
		internal.FieldPrefix_ContactItemHead,
		"5088", internal.GetGenderName)

	af.TransformDateToTime(internal.FieldPrefix_SubHead, "1004")
	af.TransformDateToTime(internal.FieldPrefix_SubHead, "1003")
	af.TransformDateToTime(internal.FieldPrefix_StoryHead, "1004")
	af.TransformDateToTime(internal.FieldPrefix_StoryHead, "1003")
	af.ComputeID()

	af.Extractor.PrintTableRowsToCSV(true, "\t", rowsIDx)
	// fmt.Println("FICK")
	// workerTypes := []internal.WorkerTypeCode{
	// internal.WorkerTypeRundownXMLutf16le}
	// archiveFile := internal.ArchiveFile{}
	// archiveFile.Init(internal.WorkerTypeRundownXMLutf16le)
}
