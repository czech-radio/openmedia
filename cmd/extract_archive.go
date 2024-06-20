package cmd

import (
	"fmt"
	"github/czech-radio/openmedia/internal/extract"
	"log/slog"
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
		"/mnt/remote/cro/export-avo/Rundowns", "string", c.NotNil,
		"Source directory of rundown files.", nil, helper.DirectoryExists)
	add("SourceDirectoryType", "sdirType", "MINIFIED.zip", "string", "",
		"type of source directory where the rundowns resides", nil, nil)
	add("OutputDirectory", "odir", "/tmp/test/", "string", c.NotNil,
		"Destination directory or file", nil, helper.DirectoryExists)
	add("OutputFileName", "ofname", "", "string", c.NotNil,
		"Output file name.", nil, nil)

	// Filter query
	add("ExtractorsName", "exsn", "production_all", "string", c.NotNil,
		"Name of extractor which specifies the parts of xml to be extracted", nil, nil)
	add("FilterDateFrom", "fdf", "", "date", c.NotNil,
		"Filter rundowns from date", nil, nil)
	add("FilterDateTo", "fdt", "", "date", c.NotNil,
		"Filter rundowns to date", nil, nil)
	add("FilterRadioName", "frn", "", "string", "",
		"Filter radio names", nil, nil)
	add("CSVdelim", "csvD", "\t", "string", "",
		"csv column field delimiter", []string{"\t", ";"}, nil)

	// Special filters
	add("FilterFileName", "frfn", "", "string", "",
		"Special filters filename", nil, CheckFileExistsIfNotNull)
	add("FilterSheetName", "frsn", "data", "string", "",
		"Special filters sheetname", nil, nil)
	add("ValidatorFileName", "valfn", "", "string", "",
		"xlsx file containing validation receipe", nil, CheckFileExistsIfNotNull)
}

func CheckFileExistsIfNotNull(fileName string) (bool, error) {
	if fileName != "" {
		return helper.FileExists(fileName)
	}
	return true, nil
}

func ParseConfigOptions() *extract.ArchiveFolderQuery {
	q := extract.ArchiveFolderQuery{}
	commandExtractArchiveConfigure()
	commandExtractArchiveConfig.RunSub(&q)
	_, offset := q.FilterDateFrom.Zone()
	offsetDuration := time.Duration(offset) * time.Second
	q.DateRange = [2]time.Time{
		q.FilterDateFrom.UTC().Add(offsetDuration + 1),
		q.FilterDateTo.UTC().Add(offsetDuration)}
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
	q.WorkerType = ar.WorkeTypeCodeGet(q.SourceDirectoryType)
	slog.Debug("effective subcommand config", "config", q)
	return &q
}

func ParseFilterOptions() *extract.FilterFile {
	filter := &extract.FilterFile{}
	err := commandExtractArchiveConfig.ParseFlags(filter)
	if err != nil {
		panic(err)
	}
	return filter
}

// func RunCommandExtractArchive(rcfg *c.RootConfig) {
func (gc GlobalConfig) RunCommandExtractArchive() {
	queryOpts := ParseConfigOptions()
	filterOpts := ParseFilterOptions()
	arf := extract.ArchiveFolder{
		PackageTypes: []ar.WorkerTypeCode{queryOpts.WorkerType}}

	// EXTRACT
	if err := arf.FolderMap(queryOpts.SourceDirectory, true, queryOpts); err != nil {
		helper.Errors.ExitWithCode(err)
	}
	extracted := arf.FolderExtract(queryOpts)

	// OUTPUT
	processName := "base"
	extracted.OutputBaseDataset(processName, queryOpts)

	processName += "_validated"
	err := extracted.OutputValidatedDataset(processName, queryOpts)
	if err != nil {
		helper.Errors.ExitWithCode(err)
	}

	processName += "_filtered"
	err = extracted.OutputFilteredDataset(processName, queryOpts, filterOpts)
	if err != nil {
		helper.Errors.ExitWithCode(err)
	}

	// D) PREVOD KÓDŮ"
	// 	// E) EXPORT CSV FILES IN DIR TO XLSX
	// 	if false {
	// 		delimRunes := []rune(queryOpts.CSVdelim)
	// 		if len(delimRunes) != 1 {
	// 			slog.Error("cannot use delim")
	// 			return
	// 		}
	// 		// errex := files.CSVdirToXLSX(query.OutputDirectory, delimRunes[0])
	// 		errex := extract.CSVdirToXLSX(queryOpts.OutputDirectory, delimRunes[0])
	// 		if errex != nil {
	// 			slog.Error(errex.Error())
	// 		}
	// 	}
	// }
}
