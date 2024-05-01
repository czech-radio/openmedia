package cmd

import (
	"github/czech-radio/openmedia/internal/helper"
)

var rootCmdConfig = helper.CommandRoot

func RunRootAlt() {
	rootCmdConfig.Init()
	rootCmdConfig.AddSub("extAlt", ExtAltRun)
	rootCmdConfig.RunRoot()
}
