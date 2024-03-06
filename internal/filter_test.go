package internal

import (
	"fmt"
	"os"
	"testing"
)

var srcFolder = "/home/jk/CRO/CRO_BASE/openmedia-archive_backup/Archive/"

func skipTest(t *testing.T) {
	if os.Getenv("GO_TEST_TYPE") != "manual" {
		t.Skip("skipping test in CI environment")
	}
}

func TestArchiveFolderListing(t *testing.T) {
	skipTest(t)
	arf := new(ArchiveFolder)
	err := arf.FolderListing(srcFolder, true)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", arf)
}

func TestArchiveFolderMap(t *testing.T) {
	arf := ArchiveFolder{
		Years:       []int{2020},
		IsoWeeks:    []int{},
		PackageType: WorkerTypeZIPminified,
	}
	err := arf.FolderMap(srcFolder, true)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(arf)
}

func TestMatchArchivePackage(t *testing.T) {
	pairs := [][]any{
		// should match
		{"2020_W09_MINIFIED.zip", []int{2020}, []int{}, WorkerTypeZIPminified, true},
		{"2021_W09_MINIFIED.zip", []int{2021}, []int{9}, WorkerTypeZIPminified, true},
		{"2021_W09_MINIFIED.zip", []int{}, []int{9}, WorkerTypeZIPminified, true},
		{"2021_W09_MINIFIED.zip", []int{2020, 2021}, []int{9}, WorkerTypeZIPminified, true},
		{"2021_W09_MINIFIED.zip", []int{2020, 2021}, []int{9}, WorkerTypeZIPminified, true},
		// should not match
		{"2021_W09_MINIFIED.zip", []int{2020}, []int{}, WorkerTypeZIPminified, false},
		{"2021_W09_MINIFIED.zip", []int{2021}, []int{8, 7}, WorkerTypeZIPminified, false},
		{"2021_W09_MINIFIED.zip", []int{2021, 2022}, []int{8, 7}, WorkerTypeZIPminified, false},
		{"2021_W09_MINIFIED.zip", []int{2020, 2022}, []int{9}, WorkerTypeCode(10), false},
		{"2021_W09_MINIFIED.zip", []int{2020, 2021}, []int{9}, WorkerTypeZIPoriginal, false},
	}
	for _, p := range pairs {
		fileName := p[0].(string)
		years := p[1].([]int)
		weeks := p[2].([]int)
		packageType := p[3].(WorkerTypeCode)
		ok := MatchArchivePackage(fileName, years, weeks, packageType)
		if ok != p[4].(bool) {
			t.Error(fmt.Errorf("not matching: %+v, result: %v", p, ok))
		}
	}
}
