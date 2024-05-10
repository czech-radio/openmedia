package cmd

import (
	"fmt"
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
	// Archive query
	add("SourceDirectory", "sdir", "", "string",
		"Source rundown file.", nil, nil)
	add("OutputDirectory", "odir", "", "string",
		"Destination directory or file", nil, nil)
	add("OutputFileName", "ofname", "", "string",
		"Output file name.", nil, nil)

	// Filter query
	add("ExtractorsName", "exsn", "production1", "string",
		"Name of extractor which specifies the parts of xml to be extracted", nil, nil)
	add("FilterDateFrom", "fdf", "", "date",
		"Filter rundowns from date", nil, nil)
	add("FilterDateTo", "fdt", "", "date",
		"Filter rundowns to date", nil, nil)
	add("FilterRadioName", "frn", "", "string",
		"Filter radio names", nil, nil)
	add("CSVdelim", "csvD", "\t", "string",
		"csv column field delimiter", []any{"\t"}, nil)

	// Special filters
	add("FiltersDirectory", "frdir", "", "string",
		"Special filters directory", nil, nil)
	// add("FiltersLoad", "frload", "", "[]string",
	// "filter files to load", nil, nil)
	add("FilterRecords", "frrec", "false", "bool",
		"filtere records", nil, nil)
}

func RunCommandExtractArchive() {
	// PREPARE CONFIG
	q := extract.ArchiveFolderQuery{}
	commandExtractArchiveConfigure()
	commandExtractArchiveConfig.RunSub(&q)
	q.DateRange = [2]time.Time{q.FilterDateFrom, q.FilterDateTo}
	if q.FilterRadioName != "" {
		q.RadioNames = make(map[string]bool)
		q.RadioNames[q.FilterRadioName] = true
	}
	// TODO: use internal configure check instead
	extCode := extract.ExtractorsPresetCode(q.ExtractorsName)
	extractors, ok := extract.ExtractorsCodeMap[extCode]
	if !ok {
		panic(fmt.Errorf("extractors name not defined: %s", q.ExtractorsName))
	}
	// q.Extractors = extract.EXTproduction
	q.Extractors = extractors
	slog.Debug("effective subcommand config", "config", q)

	filterPath1 := filepath.Join(q.FiltersDirectory, "eurovolby - zadání.xlsx")
	ok, err1 := helper.FileExists(filterPath1)
	if !ok || err1 != nil {
		panic(fmt.Errorf("filter file not readable: %s", err1))
	}

	workerTypes := []ar.WorkerTypeCode{
		// ar.WorkerTypeZIPoriginal}
		ar.WorkerTypeZIPminified}
	arf := extract.ArchiveFolder{
		PackageTypes: workerTypes,
	}

	// EXTRACT
	err := arf.FolderMap(
		q.SourceDirectory, true, &q)
	if err != nil {
		helper.Errors.ExitWithCode(err)
	}
	if q.OutputDirectory == "" {
		panic("no output directory specified")
	}
	ext := arf.FolderExtract(&q)

	// TREAT ODD RECORD
	ext.RowPartOmit(extract.RowPartCode_StoryRec)
	var indxs []int
	if q.FilterRecords {
		indxs = ext.FilterStoryPartRecordsDuds()
	}
	// indxs = []int.(nil)

	ext.TreatStoryRecordsWithoutOMobject()

	// A) BASE
	ext.TransformEmptyRowPart()
	ext.TransformBase()
	ext.CSVtableBuild(false, false, q.CSVdelim, false, indxs)

	ext.CSVtableOutputs(q.OutputDirectory, q.OutputFileName,
		q.ExtractorsName, "base", true)

	// B) EUROVOLBY
	ext.TransformProduction()
	filter1 := extract.NFilterColumn{
		FilterFileName: filterPath1,
		SheetName:      "data",
	}
	err = ext.FilterMatchPersonName(&filter1)
	if err != nil {
		panic(err)
	}
	err = ext.FilterMatchPersonAndParty(&filter1)
	if err != nil {
		panic(err)
	}

	ext.CSVtableBuild(false, false, q.CSVdelim, true, indxs)
	ext.CSVtableOutputs(q.OutputDirectory, q.OutputFileName,
		q.ExtractorsName, "eurovolby", true)
}
