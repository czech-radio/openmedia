package cmd

import (
	"github/czech-radio/openmedia-archive/internal"
	"log/slog"
)

type Config_extract struct {
	SourceDirectory        string `cmd:"source_directory; i; ; directory to be processed"`
	DestinationDirectory   string `cmd:"destination_directory; o; ; otput files"`
	RecurseSourceDirectory bool   `cmd:"recurse_source_directory; R; false; recurse source directory"`
	InvalidFileContinue    bool   `cmd:"invalid_file_continue; ifc; false; continue even though unprocessable file encountered"`
	FileType               string `cmd:"file_type; ft; rundown; files type to be processed"`
	DateFrom               string `cmd:"date_from; df; ; filter date from"`
	DateTo                 string `cmd:"date_to; dt; ; filter date to"`
	OutputType             string `cmd:"otput_type; ot; csv; type of otput format"`
	CSVdelim               string `cmd:"csv_delim; csvd; \t; csv field delimiter"`
	CSVheader              bool   `cmd:"csv_header; csvh; true; write csv column headers"`
}

func RunExtract(root_cfg *Config_root, filter_cfg *Config_extract) {
	options := internal.FilterOptions{}
	internal.CopyFields(filter_cfg, &options)
	slog.Info("effective subcommand options", "options", options)
	if root_cfg.DebugConfig {
		return
	}
	if root_cfg.DryRun {
		TEMP_DIR := internal.DirectoryCreateTemporaryOrPanic("openmedia_archive")
		options.DestinationDirectory = TEMP_DIR
	}
	internal.DirectoryIsReadableOrPanic(options.SourceDirectory)
	filter := internal.Filter{Options: options}
	err := filter.Folder()
	if err != nil {
		internal.Errors.ExitWithCode(err)
	}
}
