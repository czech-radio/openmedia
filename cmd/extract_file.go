package cmd

import (
	"fmt"
	"github/czech-radio/openmedia/internal/extract"

	c "github.com/triopium/go_utils/pkg/configure"
	"github.com/triopium/go_utils/pkg/helper"
)

func commandExtractFileConfigure() {
	add := SubcommandConfig.AddOption
	add("SourceFilePath", "sfp",
		"", "string", c.NotNil,
		"Source rundown file.", nil, helper.FileExists)
	add("OutputDirectory", "odir",
		"", "string", c.NotNil,
		"Output file path for extracted data.", nil,
		helper.DirectoryExists)
	add("OutputFileName", "ofn",
		"", "string", c.NotNil,
		"Output file path for extracted data.", nil,
		nil)
	// helper.File)
	// add("SourceFileEncoding", "sfe", "UTF8", "string", "",
	// 	"Source file encoding. Original files has UTF16le encoding. Minified files has UTF8 encoding.",
	// 	[]string{"UTF8", "UTF16le"}, nil)
	CommonExtractOptions()
}

func (gc GlobalConfig) RunCommandExtractFile() {
	af := extract.ArchiveFile{}
	commandExtractFileConfigure()
	SubcommandConfig.SubcommandOptionsParse(&af)
	fmt.Printf("%+v\n", af)
	err := af.Init()
	if err != nil {
		panic(err)
	}

	extCode := extract.ExtractorsPresetCode(af.ExtractorsName)
	extractors, ok := extract.ExtractorsCodeMap[extCode]
	if !ok {
		panic(fmt.Errorf("extractors name not defined: %s", extCode))
	}
	err = af.ExtractByXMLquery(extractors)
	if err != nil {
		panic(err)
	}
	// queryOpts := ParseConfigOptions()
	// processName := "base"
	// af.OutputBaseDataset(processName, queryOpts)
}
