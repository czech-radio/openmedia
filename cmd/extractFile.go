package cmd

import "github/czech-radio/openmedia-archive/internal"

type ConfigExtractFile struct {
	SourceFile      string `cmd:"source_file; i; ; input file"`
	DestinationFile string `cmd:"destination_file; o; ; otput file"`
	CSVdelim        string `cmd:"csv_delim; csvd; \t; csv field delimiter"`
}

func RunExtractFile(rootCfg *ConfigRoot, filterCfg *ConfigExtract) {
	filePath := "/home/jk/CRO/CRO_BASE/openmedia-archive_backup/Archive/control/control_UTF16_RD_13-17_Plus_Tuesday_W01_2024_01_02.xml"
	af := internal.ArchiveFile{}
	err := af.Init(internal.WorkerTypeRundownXMLutf16le, filePath)
	if err != nil {
		internal.Errors.ExitWithCode(err)
	}
	err = af.ExtractByXMLquery()
	if err != nil {
		internal.Errors.ExitWithCode(err)
	}
	// workerTypes := []internal.WorkerTypeCode{
	// internal.WorkerTypeRundownXMLutf16le}
	// archiveFile := internal.ArchiveFile{}
	// archiveFile.Init(internal.WorkerTypeRundownXMLutf16le)
}
