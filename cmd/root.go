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
	DryRun  bool   `cmd:"dry_run; n; false; run program in dry run mode which does not make any pernament or dangerous action. Useful for testing purposes."`
	LogType string `cmd:"log_type; lt; txt; use logger type [txt, json]"`
}

func RunRoot() {
	rcfg := &Config_root{}
	internal.SetupRootFlags(rcfg)
	internal.SetLogLevel(rcfg.Verbose, rcfg.LogType)
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
		RunArchivate(rcfg, acfg)
	default:
		slog.Error("unknown command", "command", subcmd)
	}
}
