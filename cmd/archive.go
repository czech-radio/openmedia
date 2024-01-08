package cmd

import (
	"github/czech-radio/openmedia-archive/internal"
	"log/slog"
)

type Config_archivate struct {
	SourceDirectory      string `cmd:"source_directory; i; ; directory to be processed"`
	DestinationDirectory string `cmd:"destination_directory; o; ; otput files"`
	InvalidFileContinue  bool   `cmd:"invalid_file_continue; f; 0; continue even though unprocesable file encountered"`
	InvalidFileRename    bool   `cmd:"invalid_file_rename; ifr; true; rename"`
	CompressionType      string `cmd:"compression_type; ct; zip; type of file compression [zip]."`
}

func RunArchivate(root_cfg *Config_root, archivate_cfg *Config_archivate) {
	slog.Info("running command", "command", "archivate", "root_cfg", root_cfg)
	options := internal.ProcessOptions{}
	internal.CopyFields(archivate_cfg, &options)
	if root_cfg.DryRun {
		TEMP_DIR := internal.DirectoryCreateTemporaryOrPanic("openmedia_reduce")
		options.DestinationDirectory = TEMP_DIR
	}
	slog.Info("running command", "command", "archivate", "options", options)
	internal.DirectoryIsReadableOrPanic(options.SourceDirectory)
	internal.DirectoryIsReadableOrPanic(options.DestinationDirectory)
	process := internal.Process{Options: options}
	err := process.Folder()
	if err != nil {
		internal.Errors.ExitWithCode(err)
	}
}
