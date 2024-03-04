package cmd

import (
	"github/czech-radio/openmedia-archive/internal"
	"log/slog"
)

// NOTE: Consider define command options in map[string][]string
// cmdmap[command_name]=["source_dir","i","","directory to be processed]
// TODO: Add test for commands. (dont know how to avoid circular dependency)

type ConfigArchive struct {
	SourceDirectory          string `cmd:"source_directory; i; ; directory to be processed"`
	DestinationDirectory     string `cmd:"destination_directory; o; ; otput files"`
	CompressionType          string `cmd:"compression_type; ct; zip; type of file compression [zip]."`
	InvalidFileContinue      bool   `cmd:"invalid_file_continue; ifc; false; continue even though unprocessable file encountered"`
	InvalidFileRename        bool   `cmd:"invalid_file_rename; ifr; false; rename invalid files"`
	ProcessedFileRename      bool   `cmd:"processed_file_rename; pfr; false; rename processed files"`
	ProcessedFileDelete      bool   `cmd:"processed_file_delete; pfd; false; delete processed files"`
	PreserveFoldersInArchive bool   `cmd:"PreserveFoldersInArchive; pfia; false; preserve source folder structure in archive"`
	RecurseSourceDirectory   bool   `cmd:"recurse_source_directory; R; false; recurse source directory"`
}

func RunArchive(
	rootCfg *ConfigRoot, createCfg *ConfigArchive) {
	options := internal.ArchiveOptions{}
	internal.CopyFields(createCfg, &options)
	slog.Info("effective subcommand options", "options", options)
	if rootCfg.DebugConfig {
		return
	}
	if rootCfg.DryRun {
		TempDir := internal.DirectoryCreateTemporaryOrPanic("openmedia_archive")
		options.DestinationDirectory = TempDir
	}
	internal.DirectoryIsReadableOrPanic(options.SourceDirectory)
	internal.DirectoryIsReadableOrPanic(options.DestinationDirectory)
	process := internal.Archive{Options: options}
	err := process.Folder()
	if err != nil {
		internal.Errors.ExitWithCode(err)
	}
}
