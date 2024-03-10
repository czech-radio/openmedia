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

	dateFrom := time.Date(2020, 1, 1, 0, 0, 0, 0, ArchiveTimeZone)
	dateTo := time.Date(2022, 2, 1, 0, 0, 0, 0, ArchiveTimeZone)
	filterRange := [2]time.Time{dateFrom, dateTo}
	err := arf.FolderMap(srcFolder, true, filterRange)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(arf)
}
