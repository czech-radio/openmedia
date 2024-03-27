package internal

import (
	"archive/zip"
	"io/fs"
	"log/slog"
	"path/filepath"
	"time"
)

type ArchiveFolder struct {
	PackageTypes       []WorkerTypeCode
	XMLencoding        FileEncodingNumber
	PackagesNamesOrder []PackageName
	Packages           map[PackageName]*ArchivePackage
	Files              []string
}

type FileName string
type PackageName string

type ArchivePackage struct {
	PackageName       PackageName
	PackageReader     *zip.ReadCloser
	PackageFiles      map[string]*ArchivePackageFile
	PacakgeFilesOrder []string
}

type ArchiveFolderQuery struct {
	RadioNames  map[string]bool
	DateRange   [2]time.Time
	IsoWeeks    map[int]bool
	Months      map[int]bool
	WeekDays    map[time.Weekday]bool
	Extractors  OMextractors
	PrintHeader bool
}

func (af *ArchiveFolder) FolderListing(
	rootDir string, recursive bool,
	filterRange [2]time.Time) error {
	dirWalker := func(filePath string, file fs.DirEntry, err error) error {
		slog.Info(filePath)
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
			case WorkerTypeZIPminified, WorkerTypeZIPoriginal:
				af.InferEncoding(wtc)
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
					af.PackagesNamesOrder = append(af.PackagesNamesOrder, packageName)
				}
			}
		}
		return nil
	}
	return filepath.WalkDir(rootDir, dirWalker)
}

func (af *ArchiveFolder) InferEncoding(wtc WorkerTypeCode) FileEncodingNumber {
	var enc FileEncodingNumber
	switch wtc {
	case WorkerTypeZIPminified:
		enc = UTF8
	case WorkerTypeZIPoriginal:
		enc = UTF16le
	}
	af.XMLencoding = enc
	return enc
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

func (af *ArchiveFolder) FolderExtract(query *ArchiveFolderQuery) {
	query.PrintHeader = true
	for _, packageName := range af.PackagesNamesOrder {
		slog.Warn("proccessing package", "package", packageName)
		// for _, pf := range p.PackageFiles {
		archivePackage := af.Packages[packageName]
		for _, fileName := range archivePackage.PacakgeFilesOrder {
			file := archivePackage.PackageFiles[fileName]
			slog.Warn(
				"proccessing package", "package", packageName,
				"file", file.Reader.Name,
			)
			err := file.ExtractByXMLquery(af.XMLencoding, query)
			if err != nil {
				slog.Error(err.Error())
			}
			query.PrintHeader = false
			break
		}
	}
}
