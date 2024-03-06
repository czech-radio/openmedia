package internal

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/snabb/isoweek"
)

// '2020_W05_MINIFIED.zip/RD_00-12_Vltava_Sunday_W05_2020_02_02.xml'
func MatchArchivePackage(packageFilename string, years []int, weeks []int, WorkerTypeCode WorkerTypeCode) bool {
	//DEPRECATED
	// fileName '2020_W05_MINIFIED.zip'
	packageType := WorkerTypeMap[WorkerTypeCode]
	// fmt.Println(packageFilename, years, weeks, packageType)
	var matchingYear bool
	var matchingWeek bool
	if !strings.Contains(packageFilename, packageType) {
		return false
	}
	for _, year := range years {
		idx := strings.Index(packageFilename, fmt.Sprintf("%04d", year))
		if idx == 0 {
			matchingYear = true
			break
		}
	}
	for _, week := range weeks {
		idx := strings.Index(packageFilename, fmt.Sprintf("%02d", week))
		if idx == 6 {
			matchingWeek = true
			break
		}
	}
	if len(years) > 0 && len(weeks) > 0 {
		return matchingYear && matchingWeek
	}
	if len(years) > 0 && len(weeks) == 0 {
		return matchingYear
	}
	if len(years) == 0 && len(weeks) > 0 {
		return matchingWeek
	}
	return true
}

// '2020_W05_MINIFIED.zip/RD_00-12_Vltava_Sunday_W05_2020_02_02.xml'
var packageNameRegex = regexp.MustCompile(`(\d\d\d\d)_W(\d\d)_(\s*.*)`)

func ArchivePackageNameParse(packageName string) (time.Time, time.Time, string, error) {
	//NOTE: What zone is the date given in rundowns? isoweek.Startime(1985, 1, time.UTC)
	// isoweek.StartDate
	res := packageNameRegex.FindStringSubmatch(packageName)
	if len(res) != 4 {
		return time.Time{}, time.Time{}, "", fmt.Errorf("unknown archive package name")
	}
	year, _ := strconv.Atoi(res[1])
	isoWeekNumber, _ := strconv.Atoi(res[2])
	packageType := res[3]
	_, m, d := isoweek.StartDate(year, isoWeekNumber)
	timeZone, _ := time.LoadLocation("")
	dateFrom := time.Date(year, m, d, 0, 0, 0, 0, timeZone)
	dateTo := dateFrom.AddDate(0, 0, 6)
	return dateFrom, dateTo, packageType, nil
}

func ArchivePackageMatch(packageName string, wtc WorkerTypeCode, dateFrom, dateTo time.Time) (bool, error) {
	wtcTypeName, ok := WorkerTypeMap[wtc]
	if !ok {
		panic("unknown workertype code")
	}
	packageStart, packageEnd, ptype, err := ArchivePackageNameParse(packageName)
	if err != nil {
		return false, err
	}
	if wtcTypeName != ptype {
		return false, nil
	}
	filterRange := [2]time.Time{dateFrom, dateTo}
	packageRange := [2]time.Time{packageStart, packageEnd}
	return DateIntervalsIntersec(filterRange, packageRange), nil
}
