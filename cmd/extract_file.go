package cmd

import (
	"fmt"
	"github/czech-radio/openmedia/internal/extract"

	"github.com/triopium/go_utils/pkg/helper"
)

func commandExtractFileConfigure() {
	add := SubcommandConfig.AddOption
	OptionsCommonOutput()
	OptionsCommonExtractFilter()
	add("SourceFilePath", "sfp",
		"", "string", "",
		"Source rundown file name.", nil, helper.FileExists)
	//NOTE: char encoding is infered by reading first two bytes. Some files may violate the convention. Alternatively option for hard specifing charencoding can be added.
	// add("SourceFileEncoding", "sfe", "UTF8", "string", "",
	//	"Source file encoding. Original files has UTF16le encoding. Minified files has UTF8 encoding.",
	//	[]string{"UTF8", "UTF16le"}, nil)
}

func (gc GlobalConfig) RunCommandExtractFile() {
	af := extract.ArchiveFile{}
	commandExtractFileConfigure()
	SubcommandConfig.SubcommandOptionsParse(&af.ArchiveQueryCommon)
	SubcommandConfig.SubcommandOptionsParse(&af.ArchiveIO)
	SubcommandConfig.SubcommandOptionsParse(&af.FilterFile)
	fmt.Printf("%+v\n", af.ArchiveQueryCommon)
	fmt.Printf("%+v\n", af.ArchiveIO)
	fmt.Printf("%+v\n", af.FilterFile)
	err := af.Init()
	if err != nil {
		panic(err)
	}
	exCode := af.ArchiveQueryCommon.ExtractorsCode
	extractors, ok := extract.ExtractorsCodeMap[exCode]
	if !ok {
		panic(fmt.Errorf("extractors name not defined: %s", exCode))
	}
	err = af.ExtractByXMLquery(extractors)
	if err != nil {
		panic(err)
	}
	af.ExportAll(
		&af.ArchiveQueryCommon, &af.ArchiveIO, &af.FilterFile)
}
