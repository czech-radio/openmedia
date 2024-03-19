package cmd

type ConfigExtractFile struct {
	SourceFile      string `cmd:"source_file; i; ; input file"`
	DestinationFile string `cmd:"destination_file; o; ; otput file"`
	CSVdelim        string `cmd:"csv_delim; csvd; \t; csv field delimiter"`
}

func RunExtractFile(rootCfg *ConfigRoot, filterCfg *ConfigExtract) {
	// workerTypes := []internal.WorkerTypeCode{
	// internal.WorkerTypeRundownXMLutf16le}
	// archiveFile := internal.ArchiveFile{}
	// archiveFile.Init(internal.WorkerTypeRundownXMLutf16le)
}
