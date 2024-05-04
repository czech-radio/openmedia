package cmd

import (
	"fmt"
	"github/czech-radio/openmedia/internal/extract"
	"github/czech-radio/openmedia/internal/helper"
)

var extract_Arch_Config = helper.CommandConfig{}

func ExtractArchConfigure() {
	add := extract_Arch_Config.AddOption
	add("SourceDirectory", "sdir", "", "string", "Source rundown file.",
		nil, nil)
	add("OutputDirectory", "odir", "", "string", "Destination directory or file",
		nil, nil)
	add("OutputFileName", "ofname", "", "string", "Output file name.",
		nil, nil)
	add("FilterDateFrom", "fdf", "", "date", "Filter rundowns from date",
		nil, nil)
	add("FilterDateTo", "fdt", "", "date", "Filter rundowns to date",
		nil, nil)
	add("FilterRadioNames", "frn", "", "string", "Filter radio names",
		nil, nil)
}

type My struct {
	OutputFileName string
}

func RunExtractArch() {
	query := extract.ArchiveFolderQuery{}
	// query := My{}
	ExtractArchConfigure()
	fmt.Println("trek", extract_Arch_Config.Opts)
	extract_Arch_Config.RunSub(&query)
	fmt.Printf("kek %+v\n", query)
}
