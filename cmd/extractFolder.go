package cmd

import (
	"github/czech-radio/openmedia-archive/internal"
	"log/slog"
)

type ConfigExtractFolder struct {
	SourceFile      string `cmd:"source_file; i; ; input file"`
	DestinationFile string `cmd:"destination_file; o; ; otput file"`
	CSVdelim        string `cmd:"csv_delim; csvd; \t; csv field delimiter"`
}

func RunExtractFolder(rootCfg *ConfigRoot, filterCfg *ConfigExtractFolder) {
	// folder := "/home/jk/CRO/CRO_BASE/openmedia-archive_backup/Archive/conrol_brezen/Vysocina"
	// folder := "/home/jk/CRO/CRO_BASE/openmedia-archive_backup/Archive/conrol_brezen/Radiozur"
	// folder := "/home/jk/CRO/CRO_BASE/openmedia-archive_backup/Archive/conrol_brezen/Dvojka"
	folder := "/home/jk/CRO/CRO_BASE/openmedia-archive_backup/Archive/conrol_brezen/Plus"
	files, err := internal.ListDirFiles(folder)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	header := true
	for i, f := range files {
		internal.ExtractProductionVer1(f, header)
		if i == 0 {
			header = false
		}
	}
}
