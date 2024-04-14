package cmd

import (
	// "github/czech-radio/openmedia-archive/internal"
	"github/czech-radio/openmedia-archive/internal"
	"github/czech-radio/openmedia-archive/internal/helper"
	"time"
)

type ConfigExtractArchive struct {
	SourceDirectory        string `cmd:"source_directory; i; ; directory to be processed"`
	DestinationDirectory   string `cmd:"destination_directory; o; ; otput files"`
	RecurseSourceDirectory bool   `cmd:"recurse_source_directory; R; false; recurse source directory"`
	InvalidFileContinue    bool   `cmd:"invalid_file_continue; ifc; false; continue even though unprocessable file encountered"`

	OutputType string `cmd:"otput_type; ot; csv; type of otput format"`
	CSVdelim   string `cmd:"csv_delim; csvd; \t; csv field delimiter"`
	CSVheader  bool   `cmd:"csv_header; csvh; true; write csv column headers"`

	DateFrom string `cmd:"date_from; df; ; filter date from"`
	DateTo   string `cmd:"date_to; dt; ; filter date to"`

	ComputeUniqueRows      string `cmd:"compute_unique_rows; cur; false; compute unique rows for all tables"`
	ProccessPerArchiveFile string `cmd:"process_per_archive_file; ppaf; true; run process for each file alone do not group tables"`
	ProccessPerPackage     string `cmd:"process_per_archive_package; ppap; false; run process for each file alone do not group tables"`
}

func RunExtractArchive(rootCfg *ConfigRoot, cfg *ConfigExtractArchive) {
	workerTypes := []internal.WorkerTypeCode{
		internal.WorkerTypeZIPoriginal}
	// internal.WorkerTypeZIPminified}
	arf := internal.ArchiveFolder{
		PackageTypes: workerTypes,
	}

	// brezen
	// dateFrom, _ := helper.CzechDateToUTC(2024, 2, 1, 0)
	// dateTo, _ := helper.CzechDateToUTC(2024, 4, 1, 0)
	// week13
	// dateFrom, _ := helper.CzechDateToUTC(2024, 3, 1, 0)
	// dateFrom, _ := helper.CzechDateToUTC(2023, 12, 1, 0)
	// dateFrom, _ := helper.CzechDateToUTC(2024, 3, 25, 0)
	dateFrom, _ := helper.CzechDateToUTC(2024, 3, 31, 0)
	// dateTo, _ := helper.CzechDateToUTC(2024, 3, 1, 0)
	dateTo, _ := helper.CzechDateToUTC(2024, 4, 1, 0)

	filterRange := [2]time.Time{dateFrom, dateTo}

	// extractor := helper.Extractor{}
	// extractor.Init(nil, helper.EXTproduction, helper.CSVdelim)

	query := internal.ArchiveFolderQuery{
		RadioNames: map[string]bool{
			// "Radiožurnál": true,
			// "Plus": true,
			// "Dvojka": true,
			// "ČRo_Vysočina": true,
			// "ČRo_Karlovy_Vary": true,
			// "ČRo_Brno": true,
		},
		DateRange:  filterRange,
		Extractors: internal.EXTproduction,
		// Extractors: internal.EXTeuroVolby,
		CSVdelim: cfg.CSVdelim,
	}

	// var extractor Extractor
	// extractor.Init(openMedia, q.Extractors, CSVdelim)
	srcFolder := "/mnt/remote/cro/export-avo/Rundowns"
	// srcFolder := "/home/jk/CRO/CRO_BASE/openmedia-archive_backup/Archive/"

	err := arf.FolderMap(
		srcFolder, true, &query)
	if err != nil {
		helper.Errors.ExitWithCode(err)
	}
	arf.FolderExtract(&query)
}
