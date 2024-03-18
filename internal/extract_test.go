package internal

import (
	"fmt"
	"path/filepath"
	"testing"
	"time"
)

func TestXmlQueryFields(t *testing.T) {
	ids := []string{"1", "2"}
	res := QueryFields("/OM_HEADER/OM_FIELD", ids)
	fmt.Println(res)
}

func TestXmlQuery(t *testing.T) {
	ids := []string{"1", "2"}
	res := QueryFields("/OM_HEADER/OM_FIELD", ids)
	fmt.Println(res)
}

func TestQueryObject(t *testing.T) {
	// ids := []string{"Rundown Name"}
	name := "Hourly Rundown"
	res, err := QueryObject(name)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res)
}

func TestGetLastPartOfObjectPath(t *testing.T) {
	// Define test cases
	testCases := []struct {
		input    string
		expected string
	}{
		{"/Radio Rundown", "Radio Rundown"},
		{"/Radio Rundown/Hourly Rundown", "Hourly Rundown"},
		{"", "."},
	}

	for _, tc := range testCases {
		// Call filepath.Base function
		result := filepath.Base(tc.input)

		// Check if result matches the expected output
		if result != tc.expected {
			t.Errorf("Expected Base(%q) to be %q, but got %q instead", tc.input, tc.expected, result)
		}
	}
}

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
