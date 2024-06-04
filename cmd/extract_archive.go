package cmd

import (
	"fmt"
	"github/czech-radio/openmedia/internal/extract"
	"log/slog"
	"path/filepath"
	"strings"
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

func PrepareConfig() *extract.ArchiveFolderQuery {
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

	if true {
		// EXTRACT
		if err := arf.FolderMap(query.SourceDirectory, true, query); err != nil {
			helper.Errors.ExitWithCode(err)
		}

		ext := arf.FolderExtract(query)
		// TRANSFORM
		// A) BASE
		process_name := "base"
		var indxs []int
		ext.TransformBase()
		if query.ExtractorsCode == extract.ExtractorsProductionContacts {
			indxs = ext.FilterContacts()
		}
		ext.CSVtableBuild(false, false, query.CSVdelim, false, indxs)
		ext.TableOutputs(query.OutputDirectory, query.OutputFileName,
			query.ExtractorsName, process_name, true)

		// B) VALIDATE
		process_name += "_validated"
		ext.TransformBeforeValidation()
		ext.ValidateAllColumns(query.ValidatorFileName)
		ext.CSVtableBuild(false, false, query.CSVdelim, true, indxs)
		ext.TableOutputs(query.OutputDirectory, query.OutputFileName,
			query.ExtractorsName, process_name, true)

		logFileName := strings.Join(
			[]string{query.OutputFileName, process_name, "log"}, "_")
		logFilePath := filepath.Join(
			query.OutputDirectory, logFileName+".csv")
		err := ext.ValidationLogWrite(logFilePath, query.CSVdelim, true)
		if err != nil {
			panic(err)
		}

		// C) FILTER
		process_name += "_filtered"
		ext.TransformProduction()

		if filter.FilterFileName != "" {
			// Match filename: // eurovolby,oposition
			err := ext.FilterMatchPersonName(filter)
			if err != nil {
				panic(err)
			}

			// Match filename: // eurovolby
			// err = ext.FilterMatchPersonAndParty(filter)
			// if err != nil {
			// panic(err)
			// }

			// Match filename: // oposition
			err = ext.FilterMatchPersonIDandPolitics(filter)
			if err != nil {
				panic(err)
			}
		}

		ext.CSVtableBuild(false, false, query.CSVdelim, true, indxs)
		ext.TableOutputs(query.OutputDirectory, query.OutputFileName,
			query.ExtractorsName, process_name, true)
	}
	// D) PREVOD KÓDŮ"
	// E) EXPORT CSV FILES IN DIR TO XLSX
	if false {
		delimRunes := []rune(query.CSVdelim)
		if len(delimRunes) != 1 {
			slog.Error("cannot use delim")
			return
		}
		// errex := files.CSVdirToXLSX(query.OutputDirectory, delimRunes[0])
		errex := extract.CSVdirToXLSX(query.OutputDirectory, delimRunes[0])
		if errex != nil {
			slog.Error(errex.Error())
		}
	}
}
