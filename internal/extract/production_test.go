package extract

// func TestExtractProduction(t *testing.T) {
// 	workerTypes := []ar.WorkerTypeCode{
// 		ar.WorkerTypeZIPoriginal}
// 	// internal.WorkerTypeZIPminified}
// 	arf := ArchiveFolder{
// 		PackageTypes: workerTypes,
// 	}
// 	dateFrom, _ := helper.CzechDateToUTC(2024, 3, 31, 0)
// 	dateTo, _ := helper.CzechDateToUTC(2024, 4, 1, 0)
// 	filterRange := [2]time.Time{dateFrom, dateTo}

// 	query := ArchiveFolderQuery{
// 		RadioNames: map[string]bool{
// 			// "Radiožurnál": true,
// 			// "Plus": true,
// 			// "Dvojka": true,
// 			// "ČRo_Vysočina": true,
// 			// "ČRo_Karlovy_Vary": true,
// 			// "ČRo_Brno": true,
// 		},
// 		DateRange:  filterRange,
// 		Extractors: EXTproduction,
// 		CSVdelim:   "\t",
// 	}

// 	srcFolder := "/mnt/remote/cro/export-avo/Rundowns"
// 	err := arf.FolderMap(
// 		srcFolder, true, &query)
// 	if err != nil {
// 		helper.Errors.ExitWithCode(err)
// 	}
// 	arf.FolderExtract(&query)

// }

// func BenchmarkArchiveFileExtractByXMLquery(b *testing.B) {
// 	filePath := "/home/jk/CRO/CRO_BASE/openmedia-archive_backup/Archive/control/control_UTF16_RD_13-17_Plus_Tuesday_W01_2024_01_02.xml"
// 	af := ArchiveFile{}
// 	err := af.Init(ar.WorkerTypeRundownXMLutf16le, filePath)
// 	if err != nil {
// 		b.Error(err.Error())
// 	}
// 	for i := 0; i < b.N; i++ {
// 		err = af.ExtractByXMLquery(EXTproduction)
// 		if err != nil {
// 			b.Error(err.Error())
// 		}
// 	}
// }
