package extract

import (
	"archive/zip"
	"fmt"
	ar "github/czech-radio/openmedia/internal/archive"

	"log/slog"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/triopium/go_utils/pkg/helper"

	"github.com/ncruces/go-strftime"
	"github.com/snabb/isoweek"
)

// '2020_W05_MINIFIED.zip/RD_00-12_Vltava_Sunday_W05_2020_02_02.xml'
var packageNameRegex = regexp.MustCompile(`(\d\d\d\d)_W(\d\d)_(\s*.*)`)

// ArchivePackageNameParse
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

// ArchivePackageMatch
func ArchivePackageMatch(
	packageName string, wtc ar.WorkerTypeCode, filterRange [2]time.Time) (bool, error) {
	wtcTypeName, ok := ar.WorkerTypeMap[wtc]
	if !ok {
		panic("unknown workertype code")
	}
	slog.Debug(packageName, "range", filterRange)
	packageStart, packageEnd, ptype, err := ArchivePackageNameParse(packageName)
	if err != nil {
		return false, err
	}
	if wtcTypeName != ptype {
		return false, nil
	}
	packageRange := [2]time.Time{packageStart, packageEnd}
	_, ok = helper.DateRangesIntersection(filterRange, packageRange)
	return ok, nil
}

// var packageFileNameRegex = regexp.MustCompile(`^(RD)_(\d\d)_(\d\d)_(\s*.*)$`)
// 'RD_00-05_Radiožurnál_Saturday_W05_2020_02_01.xml'
// 'RD_05-09_ČRo_Brno_Saturday_W05_2020_02_01.xml'
// `^(RD)_(\d\d)-(\d\d)_(\p{L}+_)*W(\d\d)_(\d\d\d\d)_(\d\d)_(\d\d).xml$`)
var packageFileNameRegex = regexp.MustCompile(
	`^(RD)_(\d\d)-(\d\d)_(.*)_W(\d\d)_(\d\d\d\d)_(\d\d)_(\d\d).xml$`)

// RundownName
type RundownName struct {
	Type          string
	DateRange     [2]time.Time
	RadioName     string
	IsoWeekNumber int
	WeekDay       time.Weekday
}

// ArchivePackageFilenameParse
func ArchivePackageFilenameParse(fileName string) (RundownName, error) {
	var out RundownName
	res := packageFileNameRegex.FindStringSubmatch(fileName)
	if len(res) != 9 {
		return out, fmt.Errorf("unknown archive package file")
	}
	rundownType := res[1]
	hourFrom := res[2]
	hourTo := res[3]
	isoWeekStr := res[5]
	isoWeek, err := strconv.Atoi(isoWeekStr)
	if err != nil {
		return RundownName{}, err
	}

	year := res[6]
	month := res[7]
	day := res[8]
	strDateFrom := fmt.Sprintf("%s%s%s%s", year, month, day, hourFrom)
	timeFormat := "%Y%m%d%H"
	dateFrom, err := strftime.Parse(timeFormat, strDateFrom)
	if err != nil {
		return out, err
	}
	//Max hour number is 23
	if hourTo == "24" {
		hourTo = "00"
	}
	strDateTo := fmt.Sprintf("%s%s%s%s", year, month, day, hourTo)
	dateTo, err := strftime.Parse(timeFormat, strDateTo)
	if err != nil {
		return out, err
	}
	if hourTo == "24" {
		dateTo = dateTo.AddDate(0, 0, 1)
	}

	splited := strings.Split(res[4], "_")
	RadionName := strings.Join(splited[0:len(splited)-1], "_")

	out = RundownName{
		Type: rundownType,
		DateRange: [2]time.Time{
			dateFrom,
			dateTo,
		},
		RadioName:     RadionName,
		IsoWeekNumber: isoWeek,
		WeekDay:       dateFrom.Weekday(),
	}
	return out, nil
}

// ArchivePackageFileMatch
func ArchivePackageFileMatch(
	nestedFileName string, q *ArchiveFolderQuery) (
	bool, error) {
	if q == nil {
		return false, nil
	}
	meta, err := ArchivePackageFilenameParse(nestedFileName)
	if err != nil {
		slog.Debug(
			"filename match filename", "filename", nestedFileName,
			"matched", false)
		return false, err
	}
	if len(q.RadioNames) > 0 && !q.RadioNames[meta.RadioName] {
		slog.Debug(
			"filename match radioname", "filename", nestedFileName,
			"matched", false)
		return false, nil
	}
	if len(q.WeekDays) > 0 && !q.WeekDays[meta.WeekDay] {
		slog.Debug(
			"filename match weekdays", "filename", nestedFileName,
			"matched", false)
		return false, nil
	}
	_, ok := helper.DateRangesIntersection(q.DateRange, meta.DateRange)
	if !ok {
		slog.Warn(
			"filename match daterange", "filename", nestedFileName,
			"matched", false)
		return false, nil
	}

	return true, nil
}

// PackageMap
func PackageMap(
	packageName PackageName, q *ArchiveFolderQuery) (
	*ArchivePackage, int, error) {
	zipr, err := zip.OpenReader(string(packageName))
	var count int
	if err != nil {
		return nil, count, err
	}
	var ap ArchivePackage
	ap.PackageFiles = make(map[string]*ArchivePackageFile)
	for _, fr := range zipr.File {
		ok, err := ArchivePackageFileMatch(fr.Name, q)
		if err != nil {
			return nil, count, err
		}
		if !ok {
			slog.Debug(
				"package no_match", "package", packageName,
				"file", fr.Name, "query", q.DateRange,
				"matched", false)
			continue
		}
		slog.Debug(
			"package matched", "package", packageName,
			"file", fr.Name, "query", q.DateRange,
			"matched", true)
		ap.PackageName = packageName
		ap.PackageReader = zipr
		apf := ArchivePackageFile{}
		apf.Reader = fr
		ap.PackageFiles[fr.Name] = &apf
		ap.PacakgeFilesOrder = append(ap.PacakgeFilesOrder, fr.Name)
	}
	count += len(ap.PackageFiles)
	slog.Warn(
		"filenames in all packages", "count", count,
		"matched", true)
	return &ap, count, nil
}
