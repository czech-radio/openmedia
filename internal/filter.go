package internal

import (
	"archive/zip"
	"fmt"
	"io/fs"
	"log/slog"
	"path/filepath"
	"time"
)

type Filter struct {
	Options               FilterOptions
	Results               FilterResults
	MainResults           ResultsCompounded
	Errors                []error
	ObjectHeader          []string
	ObjectsAttrValues     []ObjectAttributes
	HeaderFields          map[int]string
	HeaderFieldsIDsSorted []int
	HeaderFieldsIDsSubset map[int]bool
	FieldsUniqueValues    map[int]UniqueValues // FiledID vs UniqueValues
	MaxUniqueCount        int                  // Field which has the highest unique values count - servers. Used when transforming rows to columns
	Rows                  []Fields
}

type FilterOptions struct {
	SourceDirectory        string
	DestinationDirectory   string
	RecurseSourceDirectory bool
	InvalidFileContinue    bool

	FilterType WorkerTypeCode
	CSVdelim   string
	CSVheader  bool

	DateFrom   string
	DateTo     string
	WeekDays   string
	RadioNames string
}

type ResultsCompounded map[string]*FilterResults

type FilterResults struct {
	FilesProcessed int
	FilesSuccess   int
	FilesFailure   int
	FilesCount     int
	ErrorsCount    int
}

type PackageName string

type ArchiveFolder struct {
	PackageTypes  []WorkerTypeCode
	PackagesNames []PackageName
	Packages      map[PackageName]*ArchivePackage
	// PackageReaders []*zip.ReadCloser // zipr.File //[]*File contains all zip files
}

type ArchivePackage struct {
	PackageName      PackageName
	PackageReader    *zip.ReadCloser
	PackageFilenames []string
}

type ArchiveFolderQuery struct {
	RadioNames map[string]bool
	DateRange  [2]time.Time
	IsoWeeks   map[int]bool
	Months     map[int]bool
	WeekDays   map[time.Weekday]bool
}

func (af *ArchiveFolder) FolderListing(
	rootDir string, recursive bool, filterRange [2]time.Time) error {
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
				ok, _ := ArchivePackageMatch(filePath, wtc, filterRange)
				if !ok {
					slog.Debug("package omitted", "package", filePath)
					return nil
				}
				if ok {
					slog.Debug("package matched", "package", filePath)
					packageName := PackageName(filePath)
					af.PackagesNames = append(af.PackagesNames, packageName)
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
	for _, packageName := range af.PackagesNames {
		archivePackage, err := PackageMap(packageName, q)
		if err != nil {
			return err
		}
		af.Packages[packageName] = archivePackage
	}
	return nil
}

func PackageMap(packageName PackageName, q *ArchiveFolderQuery) (*ArchivePackage, error) {
	zipr, err := zip.OpenReader(string(packageName))
	if err != nil {
		return nil, err
	}
	var ap ArchivePackage
	for _, f := range zipr.File {
		ok, err := ArchivePackageFileMatch(f.Name, q)
		if err != nil || !ok {
			slog.Debug("no match", f.Name, q.DateRange)
			return nil, err
		}
		slog.Debug("matches", f.Name, q.DateRange)
		ap.PackageName = packageName
		ap.PackageReader = zipr
		ap.PackageFilenames = append(ap.PackageFilenames, f.Name)
	}
	return &ap, nil
}

func (ft *Filter) ErrorHandle(errMain error, errorsPartial ...error) ControlFlowAction {
	if errMain == nil {
		ft.Results.FilesSuccess++
		return Continue
	}

	ft.Results.FilesFailure++
	slog.Error(errMain.Error())
	ft.Errors = append(ft.Errors, errMain)
	if len(errorsPartial) > 0 {
		ft.Errors = append(ft.Errors, errorsPartial...)
	}

	if ft.Options.InvalidFileContinue {
		return Skip
	}
	return Break
}

func (ft *Filter) LogResults(msg string) {
	slog.Info(msg, "results", fmt.Sprintf("%+v", ft.Results))
}
