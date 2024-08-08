package cmd

import (
	"fmt"
	ar "github/czech-radio/openmedia/internal/archive"
	"github/czech-radio/openmedia/internal/extract"
	"time"

	"github.com/triopium/go_utils/pkg/helper"
)

func commandExtractArchiveConfigure() {
	add := SubcommandConfig.AddOption
	OptionsCommonOutput()
	OptionsCommonExtractFilter()
	add("SourceDirectory", "sdir",
		"", "string", "",
		"Source rundown folder archive.", nil, helper.DirectoryExists)
	add("SourceDirectoryType", "sdirType",
		"MINIFIED.zip", "string", "",
		"type of source directory where the rundowns resides", nil, nil)
}

func (gc GlobalConfig) RunCommandExtractArchive() {
	af := extract.ArchiveFolderQuery{}
	commandExtractArchiveConfigure()
	SubcommandConfig.SubcommandOptionsParse(&af.ArchiveQueryCommon)
	SubcommandConfig.SubcommandOptionsParse(&af.ArchiveIO)
	SubcommandConfig.SubcommandOptionsParse(&af.FilterFile)

	_, offset := af.FilterDateFrom.Zone()
	offsetDuration := time.Duration(offset) * time.Second
	af.DateRange = [2]time.Time{
		af.FilterDateFrom.UTC().Add(offsetDuration + 1),
		af.FilterDateTo.UTC().Add(offsetDuration)}

	fmt.Printf("%+v\n", af.ArchiveQueryCommon)
	fmt.Printf("%+v\n", af.ArchiveIO)
	fmt.Printf("%+v\n", af.FilterFile)

	arf := extract.ArchiveFolder{
		PackageTypes: []ar.WorkerType{af.SourceDirectoryType}}

	extractors, ok := extract.ExtractorsCodeMap[af.ExtractorsCode]
	if !ok {
		panic(fmt.Errorf(
			"extractors name not defined: %s", af.ExtractorsCode))
	}
	af.Extractors = extractors

	// EXTRACT
	if err := arf.FolderMap(
		af.SourceDirectory, true, &af); err != nil {
		helper.Errors.ExitWithCode(err)
	}
	extracted := arf.FolderExtract(&af)
	extracted.ExportAll(&af.ArchiveQueryCommon, &af.ArchiveIO, &af.FilterFile)
	// extracted.ExportToDB(&af.ArchiveQueryCommon, &af.ArchiveIO, &af.FilterFile)
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
// }
