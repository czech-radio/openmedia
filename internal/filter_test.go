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
		IsoWeeks:    []int{6},
		PackageType: "MINIFIED",
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
		{"2020_W09_MINIFIED.zip", []int{2020}, []int{}, "MINIFIED", true},
		{"2021_W09_MINIFIED", []int{2021}, []int{9}, "MINIFIED", true},
		{"2021_W09_MINIFIED", []int{}, []int{9}, "MINIFIED", true},
		{"2021_W09_MINIFIED", []int{2020, 2021}, []int{9}, "MINIFIED", true},
		{"2021_W09_MINIFIED", []int{2020, 2021}, []int{9}, "ORIGINAL", true},
		// should not match
		{"2021_W09_MINIFIED", []int{2020}, []int{}, "MINIFIED", false},
		{"2021_W09_MINIFIED", []int{2021}, []int{8, 7}, "MINIFIED", false},
		{"2021_W09_MINIFIED", []int{2021, 2022}, []int{8, 7}, "MINIFIED", false},
		{"2021_W09_MINIFIED", []int{2020, 2022}, []int{9}, "FUCK", false},
	}
	for _, p := range pairs {
		fileName := p[0].(string)
		years := p[1].([]int)
		weeks := p[2].([]int)
		packageType := p[3].(string)
		ok := MatchArchivePackage(fileName, years, weeks, packageType)
		if ok != p[4].(bool) {
			t.Error(fmt.Errorf("not matching: %+v, result: %v", p, ok))
		}
	}
}
