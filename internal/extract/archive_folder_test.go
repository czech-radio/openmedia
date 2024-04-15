package extract

// var srcFolder = "/home/jk/CRO/CRO_BASE/openmedia-archive_backup/Archive/"

// func TestArchiveFolderListing(t *testing.T) {
// 	// skipTest(t)
// 	workerTypes := []ar.WorkerTypeCode{ar.WorkerTypeZIPminified}
// 	arf := ArchiveFolder{
// 		PackageTypes: workerTypes,
// 	}
// 	dateFrom := time.Date(2020, 1, 1, 0, 0, 0, 0, ar.ArchiveTimeZone)
// 	dateTo := time.Date(2025, 2, 1, 0, 0, 0, 0, ar.ArchiveTimeZone)
// 	filterRange := [2]time.Time{dateFrom, dateTo}
// 	err := arf.FolderListing(srcFolder, true, filterRange)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	fmt.Printf("%+v\n", arf)
// }

// func TestArchiveFolderMap(t *testing.T) {
// 	workerTypes := []ar.WorkerTypeCode{ar.WorkerTypeZIPminified}
// 	arf := ArchiveFolder{
// 		PackageTypes: workerTypes,
// 	}

// 	dateFrom := time.Date(2020, 2, 1, 0, 0, 0, 0, ar.ArchiveTimeZone)
// 	dateTo := time.Date(2020, 2, 1, 10, 0, 0, 0, ar.ArchiveTimeZone)
// 	filterRange := [2]time.Time{dateFrom, dateTo}
// 	query := ArchiveFolderQuery{DateRange: filterRange}
// 	err := arf.FolderMap(srcFolder, true, &query)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	fmt.Println("packages", len(arf.Packages))
// 	for _, i := range arf.Packages {
// 		fmt.Println(len(i.PackageFiles))
// 	}
// }

// func TestArchiveFolderMap2(t *testing.T) {
// 	workerTypes := []ar.WorkerTypeCode{ar.WorkerTypeZIPminified}
// 	arf := ArchiveFolder{
// 		PackageTypes: workerTypes,
// 	}

// 	dateFrom := time.Date(2020, 2, 1, 0, 0, 0, 0, ar.ArchiveTimeZone)
// 	dateTo := time.Date(2020, 2, 1, 10, 0, 0, 0, ar.ArchiveTimeZone)
// 	filterRange := [2]time.Time{dateFrom, dateTo}

// 	query := ArchiveFolderQuery{
// 		RadioNames: map[string]bool{
// 			"Vltava":      true,
// 			"Radiožurnál": true,
// 		},
// 		DateRange: filterRange,
// 		IsoWeeks:  map[int]bool{},
// 		Months:    map[int]bool{},
// 		WeekDays:  map[time.Weekday]bool{},
// 	}

// 	err := arf.FolderMap(srcFolder, true, &query)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	fmt.Println("packages", len(arf.Packages))
// 	for _, i := range arf.Packages {
// 		fmt.Println(len(i.PackageFiles))
// 	}
// }
