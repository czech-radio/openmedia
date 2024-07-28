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

func CommonExtractOptions() {
	add := SubcommandConfig.AddOption
	// Output spcification
	add("CSVdelim", "csvD", "\t", "string", "",
		"csv column field delimiter", []string{"\t", ";"}, nil)

	// Filter query
	add("ExtractorsName", "exsn", "production_all", "string", c.NotNil,
		"Name of extractor which specifies the parts of xml to be extracted", nil, nil)
	add("FilterDateFrom", "fdf",
		helper.ISOweekStartLocal(-1).String(), "date", c.NotNil,
		"Filter rundowns from date. Format of the date is given in form 'YYYY-mm-ddTHH:mm:ss' e.g. 2024, 2024-02-01 or 2024-02-01T10. The precission of date given is arbitrary.", nil, nil)
	add("FilterDateTo", "fdt",
		helper.ISOweekStartLocal(0).String(), "date", c.NotNil,
		"Filter rundowns to date", nil, nil)

	add("FilterRadioNames", "frns", "", "[]string", "",
		"Filter radio names", nil, nil)

	// Special columns
	add("AddRecordsNumbers", "arn", "false", "bool", "",
		"Add record numbers columns and dependent columns", "", nil)

	// Special filters
	add("FilterFileName", "frfn", "", "string", "",
		"Special filters filename. The filter filename specifies how the file is parsed and how it is used", nil, CheckFileExistsIfNotNull)
	add("FilterSheetName", "frsn", "data", "string", "",
		"Special filters sheetname", nil, nil)
	add("ValidatorFileName", "valfn", "", "string", "",
		"xlsx file containing validation receipe", nil, CheckFileExistsIfNotNull)
}

func commandExtractArchiveConfigure() {
	add := SubcommandConfig.AddOption
	// Archive query
	add("SourceDirectory", "sdir",
		// "/mnt/remote/cro/export-avo/Rundowns", "string", c.NotNil,
		"/tmp/test2", "string", c.NotNil,
		"Source directory of rundown files.", nil, helper.DirectoryExists)
	add("SourceDirectoryType", "sdirType",
		"MINIFIED.zip", "string", "",
		"type of source directory where the rundowns resides", nil, nil)
	add("OutputDirectory", "odir",
		"/tmp/test/", "string", c.NotNil,
		"Destination directory or file", nil, helper.DirectoryExists)
	add("OutputFileName", "ofname",
		"default", "string", c.NotNil,
		"Output file name.", nil, nil)
	CommonExtractOptions()
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
	SubcommandConfig.SubcommandOptionsParse(&q)
	_, offset := q.FilterDateFrom.Zone()
	offsetDuration := time.Duration(offset) * time.Second
	q.DateRange = [2]time.Time{
		q.FilterDateFrom.UTC().Add(offsetDuration + 1),
		q.FilterDateTo.UTC().Add(offsetDuration)}
	// if q.FilterRadioName != "" {
	// q.RadioNames = make(map[string]bool)
	// q.RadioNames[q.FilterRadioName] = true
	// }
	extCode := extract.ExtractorsPresetCode(q.ExtractorsName)
	extractors, ok := extract.ExtractorsCodeMap[extCode]
	if !ok {
		panic(fmt.Errorf("extractors name not defined: %s", q.ExtractorsName))
	}

	q.AddRecordsNumbers = true
	if q.AddRecordsNumbers {
		extractors.AddRecordsColumns()
	}

	q.Extractors = extractors
	q.ExtractorsCode = extCode
	wtc, ok := ar.WorkerTypeGetCode(q.SourceDirectoryType)
	if !ok {
		panic(fmt.Errorf("unknown directoy type: %s", q.SourceDirectoryType))
	}
	q.WorkerType = wtc

	slog.Info("effective subcommand config", "config", q)
	return &q
}

func ParseFilterOptions() *extract.FilterFile {
	filter := &extract.FilterFile{}
	err := SubcommandConfig.ParseFlags(filter)
	if err != nil {
		panic(err)
	}
	return filter
}

func (gc GlobalConfig) RunCommandExtractArchive() {
	queryOpts := ParseConfigOptions()
	fmt.Printf("%+v\n", queryOpts)
	filterOpts := ParseFilterOptions()
	fmt.Printf("%+v\n", filterOpts)
	if true {
		fmt.Println("hello")
		return
	}
	// queryOpts := ParseConfigOptions()
	arf := extract.ArchiveFolder{
		PackageTypes: []ar.WorkerType{queryOpts.WorkerType}}

	// EXTRACT
	if err := arf.FolderMap(
		queryOpts.SourceDirectory, true, queryOpts); err != nil {
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

	// E) EXPORT CSV FILES IN DIR TO XLSX
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
