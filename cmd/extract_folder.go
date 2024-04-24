package cmd

import (
	ar "github/czech-radio/openmedia/internal/archive"
	"github/czech-radio/openmedia/internal/extract"
	"github/czech-radio/openmedia/internal/helper"
	"log/slog"
)

type ConfigExtractFolder struct {
	SourceFile      string `cmd:"source_file; i; ; input file"`
	DestinationFile string `cmd:"destination_file; o; ; otput file"`
	CSVdelim        string `cmd:"csv_delim; csvd; \t; csv field delimiter"`
	WorkerType      int    `cmd:"worker_type; wt; 3; type of files to be extracted: 2-UTF8, 3-UTF16le"`
	OutputType      int    `cmd:"otput_type; ot; csv; type of otput format"`
}

func RunExtractFolder(
	rootCfg *ConfigRoot, opts *ConfigExtractFolder) {
	slog.Info("effective subcommand options", "options", opts)
	files, err := helper.ListDirFiles(opts.SourceFile)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	internalHeader := false
	externalHeader := true
	for i, filePath := range files {
		af := extract.ArchiveFile{}
		err := af.Init(ar.WorkerTypeCode(opts.WorkerType), filePath)
		if err != nil {
			helper.Errors.ExitWithCode(err)
		}
		err = af.ExtractByXMLquery(extract.EXTproduction)
		if err != nil {
			helper.Errors.ExitWithCode(err)
		}
		af.Extractor.TransformProduction()
		af.Extractor.PrintTableRowsToCSV(
			internalHeader, externalHeader, "\t")
		if i == 0 {
			internalHeader = false
			externalHeader = false
		}
	}
}
