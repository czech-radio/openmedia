package cmd

import (
	"github/czech-radio/openmedia-archive/internal"
	"log/slog"
)

type Config_create struct {
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

func RunCreate(root_cfg *Config_root, create_cfg *Config_create) {
	options := internal.ProcessOptions{}
	internal.CopyFields(create_cfg, &options)
	slog.Info("effective subcommand options", "options", options)
	if root_cfg.DebugConfig {
		return
	}
	if root_cfg.DryRun {
		TEMP_DIR := internal.DirectoryCreateTemporaryOrPanic("openmedia_archive")
		options.DestinationDirectory = TEMP_DIR
	}
	internal.DirectoryIsReadableOrPanic(options.SourceDirectory)
	internal.DirectoryIsReadableOrPanic(options.DestinationDirectory)
	process := internal.Process{Options: options}
	err := process.Folder()
	if err != nil {
		internal.Errors.ExitWithCode(err)
	}
}
