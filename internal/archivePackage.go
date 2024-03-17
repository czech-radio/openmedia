package internal

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"log/slog"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/antchfx/xmlquery"
	"github.com/ncruces/go-strftime"
	"github.com/snabb/isoweek"
)

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

func ArchivePackageMatch(
	packageName string, wtc WorkerTypeCode, filterRange [2]time.Time) (bool, error) {
	wtcTypeName, ok := WorkerTypeMap[wtc]
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
	_, ok = DateRangesIntersection(filterRange, packageRange)
	return ok, nil
}

// var packageFileNameRegex = regexp.MustCompile(`^(RD)_(\d\d)_(\d\d)_(\s*.*)$`)
// 'RD_00-05_Radiožurnál_Saturday_W05_2020_02_01.xml'
// 'RD_05-09_ČRo_Brno_Saturday_W05_2020_02_01.xml'
// `^(RD)_(\d\d)-(\d\d)_(\p{L}+_)*W(\d\d)_(\d\d\d\d)_(\d\d)_(\d\d).xml$`)
var packageFileNameRegex = regexp.MustCompile(
	`^(RD)_(\d\d)-(\d\d)_(.*)_W(\d\d)_(\d\d\d\d)_(\d\d)_(\d\d).xml$`)

type RundownName struct {
	Type          string
	DateRange     [2]time.Time
	RadioName     string
	IsoWeekNumber int
	WeekDay       time.Weekday
}

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

func ArchivePackageFileMatch(nestedFileName string, q *ArchiveFolderQuery) (bool, error) {
	if q == nil {
		return false, nil
	}
	meta, err := ArchivePackageFilenameParse(nestedFileName)
	if err != nil {
		return false, err
	}
	if len(q.RadioNames) > 0 && !q.RadioNames[meta.RadioName] {
		slog.Debug("not matched radioname", "filename", nestedFileName)
		return false, nil
	}
	if len(q.WeekDays) > 0 && !q.WeekDays[meta.WeekDay] {
		slog.Debug("no matched weekdays", "filename", nestedFileName)
		return false, nil
	}
	_, ok := DateRangesIntersection(q.DateRange, meta.DateRange)
	if !ok {
		slog.Debug("not matched daterange", "filename", nestedFileName)
		return false, nil
	}

	return true, nil
}

func PackageMap(packageName PackageName, q *ArchiveFolderQuery) (*ArchivePackage, error) {
	zipr, err := zip.OpenReader(string(packageName))
	if err != nil {
		return nil, err
	}
	var ap ArchivePackage
	// ap.PackageFiles = make(map[string]*zip.File)
	ap.PackageFiles = make(map[string]*ArchivePackageFile)
	for _, fr := range zipr.File {
		ok, err := ArchivePackageFileMatch(fr.Name, q)
		if err != nil {
			return nil, err
		}
		if !ok {
			slog.Debug("package file does not match", "package", packageName, "file", fr.Name, "query", q.DateRange)
			continue
		}
		slog.Debug("package file matches", "package", packageName, "file", fr.Name, "query", q.DateRange)
		ap.PackageName = packageName
		ap.PackageReader = zipr
		apf := ArchivePackageFile{}
		apf.Reader = fr
		ap.PackageFiles[fr.Name] = &apf
	}
	return &ap, nil
}

type ArchivePackageFile struct {
	Reader *zip.File
	Tables map[WorkerTypeCode]CSVtable
}

func (apf *ArchivePackageFile) ExtractByParser(
	enc FileEncodingNumber, q *ArchiveFolderQuery) error {
	dr, err := ZipXmlFileDecodeData(apf.Reader, enc)
	if err != nil {
		return err
	}
	var OM OPENMEDIA
	err = xml.NewDecoder(dr).Decode(&OM)
	if err != nil {
		return err
	}
	// var produkce CSVtable
	for _, i := range OM.OM_OBJECT.OM_RECORDS {
		// var row CSVrow
		fmt.Println(i.OM_OBJECTS.OM_HEADER)
	}
	return nil
}

func (apf *ArchivePackageFile) ExtractByXMLquery(
	enc FileEncodingNumber, q *ArchiveFolderQuery) error {
	var err error
	// Extract file from zip
	dataReader, err := ZipXmlFileDecodeData(apf.Reader, enc)
	if err != nil {
		return err
	}
	// Parse base xml node
	baseNode, err := xmlquery.Parse(dataReader)
	if err != nil {
		return err
	}
	openMedia := xmlquery.Find(baseNode, "/OPENMEDIA")
	if len(openMedia) != 1 {
		return fmt.Errorf(
			"unknown opendmedia file, nodes found count: %d, should be 1", len(openMedia))
	}

	// Extract specfied object fields
	var extractor Extractor
	csvDelim := "\t"
	extractor.Init(openMedia[0], EXTproduction, csvDelim)
	err = extractor.ExtractRows()
	if err != nil {
		return err
	}
	extractor.PrintRowsToCSV(true, csvDelim)
	// result, err := ExtractBaseObjectRows(openMedia[0], EXTproduction)
	// PrintRowPayloads("RESULT", result)
	PrintRowPayloads("RESULT", extractor.Rows)
	return nil
}
