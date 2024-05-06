package extract

import (
	"archive/zip"
	ar "github/czech-radio/openmedia/internal/archive"

	"io/fs"
	"log/slog"
	"path/filepath"
	"time"

	"github.com/triopium/go_utils/pkg/helper"
)

// ArchiveFolder
type ArchiveFolder struct {
	PackageTypes       []ar.WorkerTypeCode
	XMLencoding        helper.FileEncodingCode
	PackagesNamesOrder []PackageName
	Packages           map[PackageName]*ArchivePackage
	Files              []string
}

// PackageName
type PackageName string

// ArchivePackage
type ArchivePackage struct {
	PackageName       PackageName
	PackageReader     *zip.ReadCloser
	PackageFiles      map[string]*ArchivePackageFile
	PacakgeFilesOrder []string
}

// ArchiveFolderQuery
type ArchiveFolderQuery struct {
	RadioNames        map[string]bool
	DateRange         [2]time.Time
	IsoWeeks          map[int]bool
	Months            map[int]bool
	WeekDays          map[time.Weekday]bool
	Extractors        OMextractors
	ComputeUniqueRows bool
	PrintHeader       bool
	CSVdelim          string
	Transformer       TransformerCode
	FilterColumns     []FilterColumn

	SourceDirectory string
	OutputDirectory string
	OutputFileName  string

	ExtractorsName string

	FilterDateFrom   time.Time
	FilterDateTo     time.Time
	FilterRadioName  string
	FiltersDirectory string
}

// FolderListing
func (af *ArchiveFolder) FolderListing(
	rootDir string, recursive bool,
	filterRange [2]time.Time) error {
	dirWalker := func(filePath string, file fs.DirEntry, err error) error {
		slog.Debug(filePath)
		if err != nil {
			return err
		}
		if filePath == rootDir {
			return nil
		}
		if file.IsDir() && !recursive {
			return filepath.SkipDir
		}
		if file.IsDir() {
			return nil
		}
		for _, wtc := range af.PackageTypes {
			switch wtc {
			case ar.WorkerTypeZIPminified, ar.WorkerTypeZIPoriginal:
				enc := ar.InferEncoding(wtc)
				af.XMLencoding = enc
				ok, _ := ArchivePackageMatch(filePath, wtc, filterRange)
				if !ok {
					slog.Debug(
						"package omitted", "package", filePath)
					return nil
				}
				if ok {
					slog.Debug(
						"package matched", "package", filePath)
					packageName := PackageName(filePath)
					af.PackagesNamesOrder = append(
						af.PackagesNamesOrder, packageName)
				}
			}
		}
		return nil
	}
	return filepath.WalkDir(rootDir, dirWalker)
}

// NOTE: Do not forget to close all files
func (af *ArchiveFolder) FolderMap(
	rootDir string, recursive bool, q *ArchiveFolderQuery) error {
	err := af.FolderListing(rootDir, recursive, q.DateRange)
	if err != nil {
		return err
	}
	if af.Packages == nil {
		af.Packages = make(map[PackageName]*ArchivePackage)
	}
	filesCount := 0
	for _, packageName := range af.PackagesNamesOrder {
		archivePackage, count, err := PackageMap(packageName, q)
		if err != nil {
			return err
		}
		filesCount = filesCount + count
		if len(archivePackage.PackageFiles) > 0 {
			af.Packages[packageName] = archivePackage
		}
	}
	slog.Warn("packages matched", "count", len(af.Packages))
	slog.Warn("files inside packages matched", "count", filesCount)
	return nil
}

// FolderExtract
func (af *ArchiveFolder) FolderExtract(
	query *ArchiveFolderQuery) *Extractor {
	var extMain Extractor
	extMain.Init(nil, query.Extractors, query.CSVdelim)
	for _, packageName := range af.PackagesNamesOrder {
		slog.Warn("proccessing package", "package", packageName)
		archivePackage := af.Packages[packageName]
		for _, fileName := range archivePackage.PacakgeFilesOrder {
			file := archivePackage.PackageFiles[fileName]
			slog.Warn(
				"proccessing package", "package", packageName,
				"file", file.Reader.Name,
			)
			ext, err := file.ExtractByXMLquery(af.XMLencoding, query)
			if err != nil {
				slog.Error(err.Error())
			}
			ext.TableXML.NullXMLnode()
			slog.Warn("extracted lines", "count", len(ext.TableXML.Rows))
			extMain.TableXML.Rows = append(
				extMain.TableXML.Rows, ext.TableXML.Rows...)
		}
	}
	return &extMain
}
