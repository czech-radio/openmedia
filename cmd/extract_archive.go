package cmd

import (
	"fmt"
	"github/czech-radio/openmedia/internal/extract"
	"log/slog"
	"path/filepath"
	"time"

	ar "github/czech-radio/openmedia/internal/archive"

	c "github.com/triopium/go_utils/pkg/configure"
	"github.com/triopium/go_utils/pkg/helper"
)

var commandExtractArchiveConfig = c.CommanderConfig{}

func commandExtractArchiveConfigure() {
	add := commandExtractArchiveConfig.AddOption
	// Archive query
	add("SourceDirectory", "sdir",
		"/mnt/remote/cro/export-avo/Rundowns", "string", c.NotNill,
		"Source directory of rundown files.", nil, helper.DirectoryExists)
	add("SourceDirectoryType", "sdirType", "MINIFIED.zip", "string", "",
		"type of source directory where the rundowns resides", nil, nil)
	add("OutputDirectory", "odir", "/tmp/test/", "string", c.NotNill,
		"Destination directory or file", nil, helper.DirectoryExists)
	add("OutputFileName", "ofname", "", "string", c.NotNill,
		"Output file name.", nil, nil)

	// Filter query
	add("ExtractorsName", "exsn", "production2", "string", c.NotNill,
		"Name of extractor which specifies the parts of xml to be extracted", nil, nil)
	add("FilterDateFrom", "fdf", "", "date", c.NotNill,
		"Filter rundowns from date", nil, nil)
	add("FilterDateTo", "fdt", "", "date", c.NotNill,
		"Filter rundowns to date", nil, nil)
	add("FilterRadioName", "frn", "", "string", "",
		"Filter radio names", nil, nil)
	add("CSVdelim", "csvD", "\t", "string", "",
		"csv column field delimiter", []string{"\t", ";"}, nil)

	// Special filters
	// add("FiltersDirectory", "frdir", "", "string", "",
	// "Special filters directory", nil, nil)
	add("FilterFileName", "frfn", "", "string", "",
		"Special filters filename", nil, nil)
	add("FilterSheetName", "frsn", "data", "string", "",
		"Special filters filename", nil, nil)
	add("FilterTypeName", "frtn", "vysoka_politika", "string", "",
		"Special filters filename", nil, nil)
}

func PrepareConfig() *extract.ArchiveFolderQuery {
	q := extract.ArchiveFolderQuery{}
	commandExtractArchiveConfigure()
	commandExtractArchiveConfig.RunSub(&q)
	_, offset := q.FilterDateFrom.Zone()
	offsetDuration := time.Duration(offset) * time.Second
	q.DateRange = [2]time.Time{
		q.FilterDateFrom.UTC().Add(offsetDuration + 1),
		q.FilterDateTo.UTC().Add(offsetDuration - 1)}
	if q.FilterRadioName != "" {
		q.RadioNames = make(map[string]bool)
		q.RadioNames[q.FilterRadioName] = true
	}
	extCode := extract.ExtractorsPresetCode(q.ExtractorsName)
	extractors, ok := extract.ExtractorsCodeMap[extCode]
	if !ok {
		panic(fmt.Errorf("extractors name not defined: %s", q.ExtractorsName))
	}
	q.Extractors = extractors
	q.ExtractorsCode = extCode

	if q.FiltersFileName != "" {
		filterPath1 := filepath.Join(q.FiltersDirectory, q.FiltersFileName)
		ok, err1 := helper.FileExists(filterPath1)
		if !ok || err1 != nil {
			panic(fmt.Errorf("filter file: %s not readable: %s", filterPath1, err1))
		}
	}
	q.WorkerType = ar.WorkeTypeCodeGet(q.SourceDirectoryType)
	slog.Debug("effective subcommand config", "config", q)
	return &q
}

func PrepareFilter() *extract.NFilterColumn {
	filter := &extract.NFilterColumn{}
	err := commandExtractArchiveConfig.ParseFlags(filter)
	if err != nil {
		panic(err)
	}
	return filter
}

func RunCommandExtractArchive() {
	query := PrepareConfig()
	filter := PrepareFilter()
	arf := extract.ArchiveFolder{
		PackageTypes: []ar.WorkerTypeCode{query.WorkerType}}
	// EXTRACT
	if err := arf.FolderMap(query.SourceDirectory, true, query); err != nil {
		helper.Errors.ExitWithCode(err)
	}
	ext := arf.FolderExtract(query)

	// TRANSFORM
	// A) BASE
	var indxs []int
	ext.TransformBase()
	if query.ExtractorsCode == extract.ExtractorsProductionContacts {
		indxs = ext.FilterContacts()
	}
	ext.CSVtableBuild(false, false, query.CSVdelim, false, indxs)
	ext.TableOutputs(query.OutputDirectory, query.OutputFileName,
		query.ExtractorsName, "base", true)

	// B) PRODUCTION
	ext.TransformProduction()
	if filter.FilterFileName != "" {
		err := ext.FilterMatchPersonName(filter)
		// err := ext.FilterMatchPersonNameExact(filter)
		if err != nil {
			panic(err)
		}

		// err = ext.FilterMatchPersonAndParty(&filter1)
		err = ext.FilterMatchPersonIDandPolitics(filter)
		if err != nil {
			panic(err)
		}
	}
	ext.CSVtableBuild(false, false, query.CSVdelim, true, indxs)
	ext.TableOutputs(query.OutputDirectory, query.OutputFileName,
		query.ExtractorsName, "transformed", true)
}
