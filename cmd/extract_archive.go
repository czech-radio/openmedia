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
		"Source rundown file.", nil, helper.DirectoryExists)
	add("SourceDirectoryType", "sdirType", "MINIFIED.zip", "string", "",
		"type of source folder where the rundowns resides", nil, nil)
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
	add("FiltersDirectory", "frdir", "", "string", "",
		"Special filters directory", nil, nil)
	add("FiltersFileName", "frfn", "", "string", "",
		"Special filters filename", nil, nil)
	add("FiltersLoad", "frload", "", "[]string", "",
		"filter files to load", nil, nil)
	add("FilterRecords", "frrec", "false", "bool", "",
		"filtere records", nil, nil)
}

func PrepareConfig() *extract.ArchiveFolderQuery {
	q := extract.ArchiveFolderQuery{}
	commandExtractArchiveConfigure()
	commandExtractArchiveConfig.RunSub(&q)
	q.DateRange = [2]time.Time{q.FilterDateFrom, q.FilterDateTo}
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

func RunCommandExtractArchive() {
	q := PrepareConfig()
	arf := extract.ArchiveFolder{PackageTypes: []ar.WorkerTypeCode{q.WorkerType}}
	// EXTRACT
	if err := arf.FolderMap(q.SourceDirectory, true, q); err != nil {
		helper.Errors.ExitWithCode(err)
	}
	ext := arf.FolderExtract(q)

	var indxs []int
	// indxs = ext.FilterContacts()

	// A) BASE
	ext.TransformBase()
	ext.CSVtableBuild(false, false, q.CSVdelim, false, indxs)
	ext.TableOutputs(q.OutputDirectory, q.OutputFileName,
		q.ExtractorsName, "base", true)

	// // B) Transformed
	// ext.TransformProduction()
	// // filter1 := extract.NFilterColumn{
	// // 	FilterFileName: filterPath1,
	// // 	SheetName:      "data",
	// // }
	// // err = ext.FilterMatchPersonName(&filter1)
	// // if err != nil {
	// // 	panic(err)
	// // }
	// // err = ext.FilterMatchPersonIDandPolitics(&filter1)
	// // // err = ext.FilterMatchPersonAndParty(&filter1)
	// // if err != nil {
	// // 	panic(err)
	// // }

	// // B) Opozice
	// ext.CSVtableBuild(false, false, q.CSVdelim, true, indxs)
	// ext.TableOutputs(q.OutputDirectory, q.OutputFileName,
	// 	q.ExtractorsName, "transformed", true)
}
