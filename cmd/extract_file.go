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
	add("OutputFilePath", "ofp",
		".", "string", c.NotNil,
		"Output file path for extracted data.", nil, nil)
	add("SourceFileEncoding", "sfe", "UTF8", "string", "",
		"Source file encoding. Original files has UTF16le encoding. Minified files has UTF8 encoding.",
		[]string{"UTF8", "UTF16le"}, nil)
	CommonExtractOptions()
}

func (gc GlobalConfig) RunCommandExtractFile() {
	af := &extract.ArchiveFile{}
	commandExtractFileConfigure()
	commandArchiveConfig.SubcommandOptionsParse(af)
	fmt.Printf("%+v\n", af)
}
