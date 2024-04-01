package cmd

import "github/czech-radio/openmedia-archive/internal"

type ConfigExtractFile struct {
	SourceFile      string `cmd:"source_file; i; ; input file"`
	DestinationFile string `cmd:"destination_file; o; ; otput file"`
	CSVdelim        string `cmd:"csv_delim; csvd; \t; csv field delimiter"`
}

func RunExtractFile(rootCfg *ConfigRoot, filterCfg *ConfigExtractFile) {
	filePath := "/home/jk/CRO/CRO_BASE/openmedia-archive_backup/Archive/control/control_UTF16_RD_13-17_Plus_Tuesday_W01_2024_01_02.xml"
	af := internal.ArchiveFile{}
	err := af.Init(internal.WorkerTypeRundownXMLutf8, filePath)
	if err != nil {
		internal.Errors.ExitWithCode(err)
	}
	err = af.ExtractByXMLquery(internal.EXTtest)
	if err != nil {
		internal.Errors.ExitWithCode(err)
	}
	// af.Extractor.TransformEurovolby()
	af.Extractor.TransformTest()
	af.Extractor.PrintTableRowsToCSV(true, "\t")
}
