package cmd

import (
	"fmt"
	"github/czech-radio/openmedia/internal/helper"
)

var CommandExtAlt = helper.CommandConfig{}

func CommandExtAltOpts() {
	add := CommandExtAlt.AddOption
	add("sourceFile", "sf", "", "string", "Source rundown file.",
		nil, nil)
}

func RunExtAlt() {
	fmt.Println("running sub")
	// CommandExtAltOpts()
	// CommandExtAlt.DeclareFlags()
}
