package internal

import (
	"testing"
	"time"
)

func TestReplaceParentRowTrue(t *testing.T) {
	EXTproduction.ReplaceParentRowTrueChecker()
	PrintObjJson("FEK", EXTproduction)
}

func TestArchiveFolderExtract(t *testing.T) {
	// workerTypes := []WorkerTypeCode{WorkerTypeZIPminified}
	workerTypes := []WorkerTypeCode{WorkerTypeZIPoriginal}
	arf := ArchiveFolder{
		PackageTypes: workerTypes,
	}
	dateFrom := time.Date(2020, 2, 1, 0, 0, 0, 0, ArchiveTimeZone)
	dateTo := time.Date(2020, 2, 1, 3, 0, 0, 0, ArchiveTimeZone)
	filterRange := [2]time.Time{dateFrom, dateTo}
	query := ArchiveFolderQuery{
		DateRange: filterRange,
		RadioNames: map[string]bool{
			// "Vltava": true,
			"Radiožurnál": true,
		},
	}
	err := arf.FolderMap(srcFolder, true, &query)
	if err != nil {
		t.Error(err)
	}
	arf.FolderExtract(&query)
}
