package cmd

import (
	"github/czech-radio/openmedia/internal/extract"
	"log/slog"
	"path/filepath"
	"time"

	ar "github/czech-radio/openmedia/internal/archive"

	"github.com/triopium/go_utils/pkg/configure"
	"github.com/triopium/go_utils/pkg/helper"
)

var commandExtractArchiveConfig = configure.CommandConfig{}

func commandExtractArchiveConfigure() {
	add := commandExtractArchiveConfig.AddOption
	add("SourceDirectory", "sdir", "", "string",
		"Source rundown file.", nil, nil)
	add("OutputDirectory", "odir", "", "string",
		"Destination directory or file", nil, nil)
	add("OutputFileName", "ofname", "", "string",
		"Output file name.", nil, nil)
	add("FilterDateFrom", "fdf", "", "date",
		"Filter rundowns from date", nil, nil)
	add("FilterDateTo", "fdt", "", "date",
		"Filter rundowns to date", nil, nil)
	add("FilterRadioName", "frn", "", "string",
		"Filter radio names", nil, nil)
	add("CSVdelim", "csvD", "\t", "string",
		"csv column field delimiter", []any{"\t"}, nil)
}

func RunCommandExtractArchive() {
	q := extract.ArchiveFolderQuery{}
	commandExtractArchiveConfigure()
	commandExtractArchiveConfig.RunSub(&q)
	q.DateRange = [2]time.Time{q.FilterDateFrom, q.FilterDateTo}
	if q.FilterRadioName != "" {
		q.RadioNames = make(map[string]bool)
		q.RadioNames[q.FilterRadioName] = true
	}
	q.Extractors = extract.EXTproduction
	slog.Debug("effective subcommand config", "config", q)
	workerTypes := []ar.WorkerTypeCode{
		ar.WorkerTypeZIPoriginal}
	arf := extract.ArchiveFolder{
		PackageTypes: workerTypes,
	}

	err := arf.FolderMap(
		q.SourceDirectory, true, &q)
	if err != nil {
		helper.Errors.ExitWithCode(err)
	}
	if q.OutputDirectory == "" {
		panic("no output directory specified")
	}
	ext := arf.FolderExtract(&q)
	// A) BASE
	dstFile1 := filepath.Join(q.OutputDirectory, q.OutputFileName+"_base_woheader.csv")
	ext.TransformEmptyRowPart()
	ext.CSVtableBuild(false, false, q.CSVdelim, false)
	ext.CSVheaderBuild(true, true)
	ext.CSVheaderWrite(dstFile1, true)
	ext.CSVtableWrite(dstFile1, false)
}
