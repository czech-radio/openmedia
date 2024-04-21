package extract

import (
	"fmt"
	ar "github/czech-radio/openmedia/internal/archive"
	"github/czech-radio/openmedia/internal/helper"
	"path/filepath"
	"testing"
)

var testerConfig = helper.TesterConfig{
	TestDataSource: "../../test/testdata",
}

func TestMain(m *testing.M) {
	testerConfig.TesterMain(m)
}

func TestSomething(t *testing.T) {
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t)
	testerConfig.PrintResult("fuck you")
}

func TestXmlQueryFields(t *testing.T) {
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t)
	ids := []string{"1", "2"}
	res := ar.XMLqueryFields("/OM_HEADER/OM_FIELD", ids)
	fmt.Println(res)
	// testerConfig.PrintResult(res)
}

func TestXmlQuery(t *testing.T) {
	ids := []string{"1", "2"}
	res := ar.XMLqueryFields("/OM_HEADER/OM_FIELD", ids)
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

// func TestArchiveFolderExtract(t *testing.T) {
// 	// workerTypes := []WorkerTypeCode{WorkerTypeZIPminified}
// 	workerTypes := []ar.WorkerTypeCode{ar.WorkerTypeZIPoriginal}
// 	arf := ArchiveFolder{
// 		PackageTypes: workerTypes,
// 	}
// 	dateFrom := time.Date(2020, 2, 1, 0, 0, 0, 0, ar.ArchiveTimeZone)
// 	dateTo := time.Date(2020, 2, 1, 3, 0, 0, 0, ar.ArchiveTimeZone)
// 	filterRange := [2]time.Time{dateFrom, dateTo}
// 	query := ArchiveFolderQuery{
// 		DateRange: filterRange,
// 		RadioNames: map[string]bool{
// 			// "Vltava": true,
// 			"Radiožurnál": true,
// 		},
// 	}
// 	err := arf.FolderMap(srcFolder, true, &query)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	arf.FolderExtract(&query)
// }

// func TestArchiveFolderExtractProdukce(t *testing.T) {
// 	// workerTypes := []WorkerTypeCode{WorkerTypeZIPminified}
// 	workerTypes := []ar.WorkerTypeCode{ar.WorkerTypeZIPoriginal}
// 	arf := ArchiveFolder{
// 		PackageTypes: workerTypes,
// 	}
// 	dateFrom := time.Date(2024, 2, 1, 13, 0, 0, 1, ar.ArchiveTimeZone)
// 	dateTo := time.Date(2024, 2, 1, 14, 0, 0, 0, ar.ArchiveTimeZone)
// 	filterRange := [2]time.Time{dateFrom, dateTo}
// 	query := ArchiveFolderQuery{
// 		DateRange: filterRange,
// 		RadioNames: map[string]bool{
// 			"Plus": true,
// 		},
// 	}
// 	err := arf.FolderMap(srcFolder, true, &query)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	arf.FolderExtract(&query)
// }
