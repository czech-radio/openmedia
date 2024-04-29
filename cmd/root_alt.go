package cmd

import (
	"github/czech-radio/openmedia/internal/helper"
)

var rootCmdConfig = helper.CommandRoot

func RunRootAlt() {
	rootCmdConfig.Init()
	rootCmdConfig.AddSub("extAlt", RunExtAlt)
	rootCmdConfig.RunRoot()
}

// func RunRootAltB() {
// 	rootCmdConfig.DeclareFlags()
// 	rcfg := &helper.RootCfg{}
// 	err := rootCmdConfig.ParseFlags(rcfg)
// 	if err != nil {
// 		panic(err)
// 	}
// 	helper.SetLogLevel(strconv.Itoa(rcfg.Verbose), rcfg.LogType)
// 	if flag.NArg() < 1 {
// 		VersionInfoPrint()
// 		return
// 	}
// 	rootCmdConfig.AddSub("extAlt", RunExtAlt)

// 	subcmd := flag.Arg(0)
// 	slog.Info("root config", "config", rcfg)
// 	slog.Info("subcommand called", "subcommand", subcmd)
// 	fmt.Printf("config %+v\n", rcfg)
// }
