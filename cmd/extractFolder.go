package cmd

import (
	"github/czech-radio/openmedia-archive/internal"
)

type ConfigExtractFolder struct {
	SourceFile      string `cmd:"source_file; i; ; input file"`
	DestinationFile string `cmd:"destination_file; o; ; otput file"`
	CSVdelim        string `cmd:"csv_delim; csvd; \t; csv field delimiter"`
}

func RunExtractFolder(rootCfg *ConfigRoot, filterCfg *ConfigExtractFolder) {
	files := []string{"/home/jk/CRO/CRO_BASE/openmedia-archive_backup/Archive/control/control_UTF16_RD_13-17_Plus_Tuesday_W01_2024_01_02.xml"}

	extractor := internal.Extractor{}
	delim := "\t"
	extractor.Init(nil, internal.EXTproduction, delim)
	for _, filePath := range files {
		af := internal.ArchiveFile{}
		err := af.Init(
			internal.WorkerTypeRundownXMLutf16le, filePath)
		if err != nil {
			internal.Errors.ExitWithCode(err)
		}
		err = af.ExtractByXMLqueryB(&extractor)
		if err != nil {
			internal.Errors.ExitWithCode(err)
		}
	}
}
