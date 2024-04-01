// Package cmd defines subcommands for binary
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

// VersionInfo Binary version info
var VersionInfo = internal.VersionInfo{
	Version:   "0.9.9",
	GitTag:    BuildGitTag,
	GitCommit: BuildGitCommit,
}

// VersionInfoPrint print binary version info
func VersionInfoPrint() {
	fmt.Printf("openmedia_archive:%+v\n", VersionInfo)
}

// ConfigRoot Config for root command
type ConfigRoot struct {
	// "long flag; short flag; default value; description"
	Version     bool   `cmd:"version; V; false; version of the program"`
	Verbose     string `cmd:"verbose; v; 0; program verbosity level: DEBUG (-4), INFO (0), WARN (4), and ERROR (8)"`
	DebugConfig bool   `cmd:"debug_config; dc; false; print effective config variables"`
	DryRun      bool   `cmd:"dry_run; n; false; run program in dry run mode which does not make any pernament or dangerous action. Useful for testing purposes."`
	LogType     string `cmd:"log_type; lt; json; use logger type [json,plain]"`
}

func RunRoot() {
	rcfg := &ConfigRoot{}
	internal.SetupRootFlags(rcfg)
	internal.SetLogLevel(rcfg.Verbose, rcfg.LogType)
	if flag.NArg() < 1 {
		VersionInfoPrint()
		return
	}
	subcmd := flag.Arg(0)
	slog.Info("root config", "config", rcfg)
	slog.Info("subcommand called", "subcommand", subcmd)

	switch subcmd {
	case "archive":
		cmdCfg := &ConfigArchive{}
		internal.SetupSubFlags(cmdCfg)
		RunArchive(rcfg, cmdCfg)
	case "extractArchive":
		cmdCfg := &ConfigExtractArchive{}
		internal.SetupSubFlags(cmdCfg)
		RunExtractArchive(rcfg, cmdCfg)
	case "extractFile":
		cmdCfg := &ConfigExtractFile{}
		internal.SetupSubFlags(cmdCfg)
		RunExtractFile(rcfg, cmdCfg)
	case "extractFolder":
		cmdCfg := &ConfigExtractFolder{}
		internal.SetupSubFlags(cmdCfg)
		RunExtractFolder(rcfg, cmdCfg)
	default:
		slog.Error("unknown command", "command", subcmd)
	}
}
