package internal

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var srcFolder = "/home/jk/CRO/CRO_BASE/openmedia-archive_backup/Archive/"

func skipTest(t *testing.T) {
	if os.Getenv("GO_TEST_TYPE") != "manual" {
		t.Skip("skipping test in CI environment")
	}
}

func TestArchiveFolderListing(t *testing.T) {
	skipTest(t)
	workerTypes := []WorkerTypeCode{WorkerTypeZIPminified}
	arf := ArchiveFolder{
		PackageTypes: workerTypes,
	}
	dateFrom := time.Date(2020, 1, 1, 0, 0, 0, 0, ArchiveTimeZone)
	dateTo := time.Date(2025, 2, 1, 0, 0, 0, 0, ArchiveTimeZone)
	filterRange := [2]time.Time{dateFrom, dateTo}
	err := arf.FolderListing(srcFolder, true, filterRange)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", arf)
}

func TestArchiveFolderMap(t *testing.T) {
	workerTypes := []WorkerTypeCode{WorkerTypeZIPminified}
	arf := ArchiveFolder{
		PackageTypes: workerTypes,
	}

	dateFrom := time.Date(2020, 2, 1, 0, 0, 0, 0, ArchiveTimeZone)
	dateTo := time.Date(2020, 2, 1, 23, 0, 0, 0, ArchiveTimeZone)
	filterRange := [2]time.Time{dateFrom, dateTo}
	query := ArchiveFolderQuery{DateRange: filterRange}
	err := arf.FolderMap(srcFolder, true, &query)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("packages", len(arf.Packages))
	// fmt.Println(len(arf.Packages))
	// for _, i := range arf.Packages {
	// fmt.Println(i.PackageName)
	// files := i.PackageFilenames
	// fmt.Println(len(files))
	// if i.PackageFilenames == nil {
	// fmt.Println(i, "is nil")
	// }
	// fmt.Println(i.PackageName, len(i.PackageFilenames))
	// }
	// fmt.Println(dateFrom.ISOWeek())
	// fmt.Println(dateFrom.Weekday())
	// fmt.Println(dateTo.ISOWeek())
	// fmt.Println(dateTo.Weekday())
	// dateC := time.Date(2020, 2, 3, 0, 0, 0, 0, ArchiveTimeZone)
	// fmt.Println(dateC.ISOWeek())
}
