package cmd

import (
	"github/czech-radio/openmedia/internal/helper"
)

var ExtAlt = helper.CommandConfig{}

func ExtAltOpts() {
	add := ExtAlt.AddOption
	add("sourceDirectory", "i", "", "string", "Source rundown file.",
		nil, nil)
	add("outputFile", "o", "", "string", "Destination directory or file",
		nil, nil)
}

type ExtAltCfg struct {
	SourceFile string
}

var my = ExtAltCfg{}

func ExtAltRun() {
	ExtAltOpts()
	ExtAlt.RunSub(&my)
}
