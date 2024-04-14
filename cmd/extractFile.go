package cmd

import (
	"github/czech-radio/openmedia-archive/internal"
	"github/czech-radio/openmedia-archive/internal/extcases"
	"github/czech-radio/openmedia-archive/internal/helper"
)

type ConfigExtractFile struct {
	SourceFile      string `cmd:"source_file; i; ; input file"`
	DestinationFile string `cmd:"destination_file; o; ; otput file"`
	CSVdelim        string `cmd:"csv_delim; csvd; \t; csv field delimiter"`
}

func RunExtractFile(rootCfg *ConfigRoot, filterCfg *ConfigExtractFile) {
	// filePath := "/home/jk/CRO/CRO_BASE/openmedia-archive_backup/Archive/control/control_UTF16_RD_13-17_Plus_Tuesday_W01_2024_01_02.xml"
	filePath := "/home/jk/CRO/CRO_BASE/openmedia-archive_backup/Archive/control2/RD_18-24_Radiožurnál_Friday_W09_2024_03_01_utf16le.xml"
	// filePath := "/home/jk/CRO/CRO_BASE/openmedia-archive_backup/Archive/control/control_UTF8_RD_13-17_Plus_Tuesday_W01_2024_01_02.xml"
	af := internal.ArchiveFile{}
	err := af.Init(
		internal.WorkerTypeRundownXMLutf16le, filePath)
	// internal.WorkerTypeRundownXMLutf8, filePath)
	if err != nil {
		helper.Errors.ExitWithCode(err)
	}
	// err = af.ExtractByXMLquery(internal.EXTtest)
	err = af.ExtractByXMLquery(extcases.EXTproduction)
	// err = af.ExtractByXMLquery(internal.EXTeuroVolby)
	if err != nil {
		helper.Errors.ExitWithCode(err)
	}
	af.Extractor.TransformProduction()
	// af.Extractor.TransformEurovolby()
	// af.Extractor.TransformTest()
	af.Extractor.PrintTableRowsToCSV(true, "\t")
}
