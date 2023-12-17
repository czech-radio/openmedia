package cmd

import (
	"flag"
	"fmt"
	"github/czech-radio/openmedia-archive/internal"
	"log/slog"
)

// Build tags set with -ldflags. Cannot set struct fields directly.
var (
	BuildGitTag    string
	BuildGitCommit string
	BuildBuildTime string
)

var VersionInfo = internal.VersionInfo{
	Version:   "0.9.0",
	GitTag:    BuildGitTag,
	GitCommit: BuildGitCommit,
	BuildTime: BuildBuildTime,
}

func VersionInfoPrint() {
	fmt.Printf("openmedia_archive:%+v\n", VersionInfo)
}

type Config_root struct {
	Version bool   `cmd:"version; V; false; version of the program"`
	Verbose string `cmd:"verbose; v; false; program verbosity level"`
}

func RunRoot() {
	rcfg := &Config_root{}
	internal.SetupRootFlags(rcfg)
	internal.SetLogLevel(rcfg.Verbose)
	if flag.NArg() < 1 {
		VersionInfoPrint()
		return
	}
	subcmd := flag.Arg(0)
	switch subcmd {
	case "archivate":
		slog.Info("subcommand called", "command", subcmd)
		acfg := &Config_archivate{}
		internal.SetupSubFlags(acfg)
		RunArchivate(acfg)
	default:
		slog.Error("unknown command", "command", subcmd)
	}
}
