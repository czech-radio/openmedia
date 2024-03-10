package internal

import (
	"fmt"
	"testing"
	"time"
)

func TestPackageNameParse(t *testing.T) {
	packName := "2024_W12_MINIFIED.zip"
	fmt.Println(ArchivePackageNameParse(packName))
}

func TestArchivePackageMatch(t *testing.T) {
	packageName := "2021_W09_MINIFIED.zip"
	dateFrom := time.Date(2021, 1, 1, 0, 0, 0, 0, ArchiveTimeZone)
	dateTo := time.Date(2022, 2, 1, 0, 0, 0, 0, ArchiveTimeZone)
	filterRange := [2]time.Time{dateFrom, dateTo}
	ok, err :=
		ArchivePackageMatch(packageName, WorkerTypeZIPminified, filterRange)
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Error(fmt.Errorf("should match"))
	}
}
func TestArchivePackageFileParse(t *testing.T) {
	files := []string{
		"RD_00-05_Radiožurnál_Saturday_W05_2020_02_01.xml",
		"RD_05-09_ČRo_Brno_Saturday_W05_2020_02_01.xml",
	}
	for i := range files {
		res, err := ArchivePackageFilenameParse(files[i])
		if err != nil {
			t.Error(res)
		}
		fmt.Println(res)
	}
}

func TestArchivePackageFileMatch(t *testing.T) {
	files := []string{
		"RD_00-05_Radiožurnál_Saturday_W05_2020_02_01.xml",
		"RD_05-09_ČRo_Brno_Saturday_W05_2020_02_01.xml",
	}
	dateFrom := time.Date(2021, 1, 1, 0, 0, 0, 0, ArchiveTimeZone)
	dateTo := time.Date(2022, 2, 1, 0, 0, 0, 0, ArchiveTimeZone)
	query := &ArchiveFolderQuery{
		DateRange: [2]time.Time{
			dateFrom,
			dateTo,
		},
		Months:     map[int]bool{},
		WeekDays:   map[time.Weekday]bool{},
		IsoWeeks:   map[int]bool{},
		RadioNames: map[string]bool{},
	}
	for i := range files {
		ok, err := ArchivePackageFileMatch(files[i], query)
		if err != nil {
			t.Error(err)
		}
		if ok != true {
			t.Error("not matched")
		}
	}

}

func TestArchivePackageFileMatchER(t *testing.T) {
	af := ArchiveFolderQuery{
		RadioNames: map[string]bool{
			"zurnal": false,
		},
	}
	fmt.Println(af)
}
