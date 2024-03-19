package cmd

import (
	"github/czech-radio/openmedia-archive/internal"
	"time"
)

type ConfigExtract struct {
	SourceDirectory        string `cmd:"source_directory; i; ; directory to be processed"`
	DestinationDirectory   string `cmd:"destination_directory; o; ; otput files"`
	RecurseSourceDirectory bool   `cmd:"recurse_source_directory; R; false; recurse source directory"`
	InvalidFileContinue    bool   `cmd:"invalid_file_continue; ifc; false; continue even though unprocessable file encountered"`

	OutputType string `cmd:"otput_type; ot; csv; type of otput format"`
	CSVdelim   string `cmd:"csv_delim; csvd; \t; csv field delimiter"`
	CSVheader  bool   `cmd:"csv_header; csvh; true; write csv column headers"`

	DateFrom string `cmd:"date_from; df; ; filter date from"`
	DateTo   string `cmd:"date_to; dt; ; filter date to"`

	FilterTypeNumber int `cmd:"filter_type; ft; 0; files type to be processed"`
}

func RunExtract(rootCfg *ConfigRoot, filterCfg *ConfigExtract) {
	// options := internal.FilterOptions{}
	// internal.CopyFields(filterCfg, &options)
	// options.FilterType = internal.WorkerTypeCode(filterCfg.FilterTypeNumber)
	// slog.Info("effective subcommand options", "options", options)
	// if rootCfg.DebugConfig {
	// return
	// }
	// internal.DirectoryIsReadableOrPanic(options.SourceDirectory)
	// ORIG:
	// filter := internal.Filter{Options: options}
	// err := filter.Folder()
	// if err != nil {
	// internal.Errors.ExitWithCode(err)
	// }

	workerTypes := []internal.WorkerTypeCode{internal.WorkerTypeZIPoriginal}
	arf := internal.ArchiveFolder{
		PackageTypes: workerTypes,
	}
	dateFrom := time.Date(2020, 2, 1, 0, 0, 0, 0, internal.ArchiveTimeZone)
	dateTo := time.Date(2020, 2, 1, 3, 0, 0, 0, internal.ArchiveTimeZone)
	filterRange := [2]time.Time{dateFrom, dateTo}
	query := internal.ArchiveFolderQuery{
		DateRange: filterRange,
		RadioNames: map[string]bool{
			// "Vltava": true,
			"Radiožurnál": true,
		},
	}
	srcFolder := "/home/jk/CRO/CRO_BASE/openmedia-archive_backup/Archive/"
	err := arf.FolderMap(srcFolder, true, &query)
	if err != nil {
		internal.Errors.ExitWithCode(err)
	}
	arf.FolderExtract(&query)
}
