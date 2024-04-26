package cmd

import (
	"fmt"
	"github/czech-radio/openmedia/internal/helper"
	"log/slog"
)

var rootCmdConfig = helper.CommandRoot

func RunRootAlt() {
	slog.Warn("running")
	rootCmdConfig.DeclareFlags()
	rcfg := &helper.RootCfg{}
	err := rootCmdConfig.ParseFlags(rcfg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", rcfg)
}
