package cmd

import (
	ar "github/czech-radio/openmedia/internal/archive"
	"github/czech-radio/openmedia/internal/extract"
	"github/czech-radio/openmedia/internal/helper"
	"time"
)

type ConfigExtractArchive struct {
	SourceDirectory        string `cmd:"source_directory; i; ; directory to be processed"`
	DestinationDirectory   string `cmd:"destination_directory; o; ; otput files"`
	RecurseSourceDirectory bool   `cmd:"recurse_source_directory; R; false; recurse source directory"`
	InvalidFileContinue    bool   `cmd:"invalid_file_continue; ifc; false; continue even though unprocessable file encountered"`
	WorkerType             string `cmd:"worker_type; wt; ; type of files to be extracted"`

	// CSV config
	OutputType string `cmd:"otput_type; ot; csv; type of otput format"`
	CSVdelim   string `cmd:"csv_delim; csvd; \t; csv field delimiter"`
	CSVheader  bool   `cmd:"csv_header; csvh; true; write csv column headers"`

	// Query config
	DateFrom    string `cmd:"date_from; df; ; filter date from"`
	DateTo      string `cmd:"date_to; dt; ; filter date to"`
	RadioNames  string `cmd:"radio_names; rn; ; list of radio names"`
	Transformer string `cmd:"transformer; tr; ; csv fields transformer name"`

	// Processing specification
	ComputeUniqueRows      string `cmd:"compute_unique_rows; cur; false; compute unique rows for all tables"`
	ProccessPerArchiveFile string `cmd:"process_per_archive_file; ppaf; true; run process for each file alone do not group tables"`
	ProccessPerPackage     string `cmd:"process_per_archive_package; ppap; false; run process for each file alone do not group tables"`
}

func RunExtractArchive(rootCfg *ConfigRoot, cfg *ConfigExtractArchive) {

	workerTypes := []ar.WorkerTypeCode{
		ar.WorkerTypeZIPoriginal}
	// internal.WorkerTypeZIPminified}
	arf := extract.ArchiveFolder{
		PackageTypes: workerTypes,
	}

	// DateFrom
	// dateFrom, _ := helper.CzechDateToUTC(2024, 2, 1, 0)
	// dateFrom, _ := helper.CzechDateToUTC(2024, 3, 1, 0)
	// dateFrom, _ := helper.CzechDateToUTC(2023, 12, 1, 0)
	// dateFrom, _ := helper.CzechDateToUTC(2024, 3, 25, 0)
	// dateFrom, _ := helper.CzechDateToUTC(2024, 3, 31, 0)

	// DateTo
	// dateTo, _ := helper.CzechDateToUTC(2024, 2, 1, 0)
	// dateTo, _ := helper.CzechDateToUTC(2024, 3, 1, 0)
	// dateTo, _ := helper.CzechDateToUTC(2024, 4, 1, 0)

	// TEST WEEK 13
	// dateFrom, _ := helper.CzechDateToUTC(2024, 3, 25, 0)
	// dateTo, _ := helper.CzechDateToUTC(2024, 4, 1, 0)

	// TEST VZOR
	dateFrom, _ := helper.CzechDateToUTC(2024, 1, 2, 15)
	dateTo, _ := helper.CzechDateToUTC(2024, 1, 2, 17)
	// dateTo, _ := helper.CzechDateToUTC(2024, 1, 3, 5)

	filterRange := [2]time.Time{dateFrom, dateTo}

	radioNames := map[string]bool{
		// "Radiožurnál": true,
		"Plus": true,
		// "Dvojka": true,
		// "ČRo_Vysočina": true,
		// "ČRo_Karlovy_Vary": true,
		// "ČRo_Brno": true,
	}

	// Filter columns
	var filterColumns []extract.FilterColumn
	// filterFile := "/home/jk/CRO/CRO_BASE/openmedia_backup/filters/filtrace - zadání.xlsx"
	filterFile := "/home/jk/CRO/CRO_BASE/openmedia_backup/filters/filtrace - zadání.xlsx"
	values, err := helper.MapExcelSheetColumn(filterFile, "seznam", 0)
	if err != nil {
		helper.Errors.ExitWithCode(err)
	}
	filter := extract.FilterCodeMap[extract.FilterCodeMatchPersonName]
	filter.FileWithValues = filterFile
	filter.Values = values
	filterColumns = append(filterColumns, filter)

	// Build query
	query := extract.ArchiveFolderQuery{
		RadioNames: radioNames,
		DateRange:  filterRange,
		Extractors: extract.EXTproduction,
		// Transformer: extract.TransformerProduction,
		Transformer: extract.TransformerProductionCSV,
		// Extractors: internal.EXTeuroVolby,
		FilterColumns: filterColumns,
		CSVdelim:      cfg.CSVdelim,
	}

	// Query run on folder
	srcFolder := "/mnt/remote/cro/export-avo/Rundowns"
	// srcFolder := "/home/jk/CRO/CRO_BASE/openmedia_backup/Archive/"

	err = arf.FolderMap(
		srcFolder, true, &query)
	if err != nil {
		helper.Errors.ExitWithCode(err)
	}
	arf.FolderExtract(&query)
}
