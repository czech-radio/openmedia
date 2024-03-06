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
		{"2021_w09_minified.zip", []int{2020, 2021}, []int{9}, WorkerTypeZIPoriginal, false},
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
func TestArchivePackageMatch(t *testing.T) {
	packageName := "2021_W09_MINIFIED.zip"
	timeZone, _ := time.LoadLocation("")
	dateFrom := time.Date(2021, 1, 1, 0, 0, 0, 0, timeZone)
	dateTo := time.Date(2022, 2, 1, 0, 0, 0, 0, timeZone)
	ok, err := ArchivePackageMatch(packageName, WorkerTypeZIPminified, dateFrom, dateTo)
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Error(fmt.Errorf("should match"))
	}
}
