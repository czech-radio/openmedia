package cmd

import (
	"github/czech-radio/openmedia-archive/internal"
	"github/czech-radio/openmedia-archive/internal/helper"
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
	// folder := "/home/jk/CRO/CRO_BASE/openmedia-archive_backup/Archive/conrol_brezen/Plus"
	folder := "/home/jk/CRO/CRO_BASE/openmedia-archive_backup/Archive/landa_control"
	files, err := helper.ListDirFiles(folder)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	header := true
	for i, filePath := range files {
		af := internal.ArchiveFile{}
		err := af.Init(internal.WorkerTypeRundownXMLutf8, filePath)
		if err != nil {
			helper.Errors.ExitWithCode(err)
		}
		err = af.ExtractByXMLquery(internal.EXTtest)
		if err != nil {
			helper.Errors.ExitWithCode(err)
		}
		af.Extractor.TransformEurovolby()
		af.Extractor.PrintTableRowsToCSV(header, "\t")
		if i == 0 {
			header = false
		}
	}
}
