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
	Version              bool   `cmd:"version; V; false; version of the program"`
	Verbose              string `cmd:"verbose; v; 0; program verbosity level: DEBUG (-4), INFO (0), WARN (4), and ERROR (8)"`
	DryRun               bool   `cmd:"dry_run; n; false; run program in dry run mode which does not make any pernament or dangerous action. Useful for testing purposes."`
	LogType              string `cmd:"log_type; lt; json; use logger type [json,plain]"`
	SourceDirectory      string `cmd:"source_directory; i; ; directory to be processed"`
	DestinationDirectory string `cmd:"destination_directory; o; ; otput files"`
	CompressionType      string `cmd:"compression_type; ct; zip; type of file compression [zip]."`
	InvalidFileContinue  bool   `cmd:"invalid_file_continue; ifc; 0; continue even though unprocesable file encountered"`
	InvalidFileRename    bool   `cmd:"invalid_file_rename; ifr; false; rename invalid files"`
	ProcessedFileRename  bool   `cmd:"processed_file_rename; pfr; false; rename processesd files"`
	ProcessedFileDelete  bool   `cmd:"processed_file_delete; pfd; false; delete processed files"`
}

func RunRoot() {
	rcfg := &Config_root{}
	internal.SetupRootFlags(rcfg)
	internal.SetLogLevel(rcfg.Verbose, rcfg.LogType)
	if flag.NFlag() < 1 || rcfg.Version {
		VersionInfoPrint()
		return
	}
	slog.Info("running command", "config", rcfg)
	options := internal.ProcessOptions{}
	internal.CopyFields(rcfg, &options)
	if rcfg.DryRun {
		TEMP_DIR := internal.DirectoryCreateTemporaryOrPanic("openmedia_archive")
		options.DestinationDirectory = TEMP_DIR
	}
	slog.Info("running process", "options", options)
	internal.DirectoryIsReadableOrPanic(options.SourceDirectory)
	internal.DirectoryIsReadableOrPanic(options.DestinationDirectory)
	process := internal.Process{Options: options}
	err := process.Folder()
	if err != nil {
		internal.Errors.ExitWithCode(err)
	}
}
