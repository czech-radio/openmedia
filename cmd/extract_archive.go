package cmd

import (
	ar "github/czech-radio/openmedia/internal/archive"
	"github/czech-radio/openmedia/internal/extract"
	"github/czech-radio/openmedia/internal/helper"
	"path/filepath"
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

	// LEDEN-BREZEN
	dateFrom, _ := helper.CzechDateToUTC(2024, 1, 1, 0)
	dateTo, _ := helper.CzechDateToUTC(2024, 3, 1, 0)

	// TEST WEEK 13
	// dateFrom, _ := helper.CzechDateToUTC(2024, 3, 25, 0)
	// dateTo, _ := helper.CzechDateToUTC(2024, 4, 1, 0)

	// TEST VZOR
	// dateFrom, _ := helper.CzechDateToUTC(2024, 1, 2, 15)
	// dateTo, _ := helper.CzechDateToUTC(2024, 1, 2, 17)

	filterRange := [2]time.Time{dateFrom, dateTo}

	radioNames := map[string]bool{
		// "Radiožurnál": true,
		// "Plus": true,
		// "Dvojka": true,
		// "ČRo_Vysočina": true,
		// "ČRo_Karlovy_Vary": true,
		// "ČRo_Brno": true,
	}
	// expRange := "test"
	// expRange := "W13"
	expRange := "2024_leden_unor"

	// Build query
	query := extract.ArchiveFolderQuery{
		RadioNames: radioNames,
		DateRange:  filterRange,
		Extractors: extract.EXTproduction,
		// Extractors: internal.EXTeuroVolby,
		CSVdelim: cfg.CSVdelim,
	}

	// Query run on folder
	srcFolder := "/mnt/remote/cro/export-avo/Rundowns"
	// srcFolder := "/home/jk/CRO/CRO_BASE/openmedia_backup/Archive/"
	dstDir := "/mnt/remote/cro/R/GŘ/Strategický rozvoj/Analytická sekce/Analýzy/Produkce/Tests/2024_04_30_produkce/"

	dstFile1 := filepath.Join(dstDir, expRange+"_base_wheader.csv")
	dstFile2 := filepath.Join(dstDir, expRange+"_base_woheader.csv")

	err := arf.FolderMap(
		srcFolder, true, &query)
	if err != nil {
		helper.Errors.ExitWithCode(err)
	}
	ext := arf.FolderExtract(&query)

	// A) BASE
	// Transform
	ext.TransformEmptyRowPart()
	ext.TransformBase()
	ext.CSVtableBuild(false, false, query.CSVdelim, false)

	// with internal header
	ext.CSVheaderBuild(true, true)
	ext.CSVheaderWrite(dstFile1)
	ext.CSVtableWrite(dstFile1)

	// without header
	ext.CSVheaderBuild(false, true)
	ext.CSVheaderWrite(dstFile2)
	ext.CSVtableWrite(dstFile2)

	// B) PRODUCTION
	dstFile3 := filepath.Join(dstDir, expRange+"_eurovolby_wheader.csv")
	dstFile4 := filepath.Join(dstDir, expRange+"_eurovolby_woheader.csv")
	ext.TransformProduction()

	// filter
	filter := extract.NFilterColumn{
		FilterFileName:  "/home/jk/CRO/CRO_BASE/openmedia_backup/filters/eurovolby - zadání.xlsx",
		SheetName:       "data",
		ColumnHeaderRow: 0,
		RowHeaderColumn: 0,
		PartCodeMark:    0,
		FieldIDmark:     "",
	}

	err = ext.FilterMatchPersonName(&filter)
	if err != nil {
		panic(err)
	}
	err = ext.FilterMatchPersonAndParty(&filter)
	if err != nil {
		panic(err)
	}

	// build
	ext.CSVtableBuild(false, false, query.CSVdelim, true)

	// write file
	ext.CSVheaderBuild(true, true)
	ext.CSVheaderWrite(dstFile3)
	ext.CSVtableWrite(dstFile3)

	ext.CSVheaderBuild(false, true)
	ext.CSVheaderWrite(dstFile4)
	ext.CSVtableWrite(dstFile4)
}
