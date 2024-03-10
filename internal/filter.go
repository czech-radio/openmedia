package internal

import (
	"archive/zip"
	"fmt"
	"io/fs"
	"log/slog"
	"path/filepath"
	"time"
)

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

type ArchiveFolderPackageQuery struct {
	RadioNames map[string]bool
	DateRange  [2]time.Time
	IsoWeeks   map[int]bool
	Months     map[int]bool
	WeekDays   map[time.Weekday]bool
}

type ResultsCompounded map[string]*FilterResults

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

type FilterResults struct {
	FilesProcessed int
	FilesSuccess   int
	FilesFailure   int
	FilesCount     int
	ErrorsCount    int
}

type ArchivePackage struct {
	FilePath    string
	Type        string
	FilesNested []string
}

type ArchiveFolder struct {
	PackageFilenames []string
	SimpleFilenames  []string
	PackageReaders   []*zip.ReadCloser // zipr.File //[]*File contains all zip files
	PackageTypes     []WorkerTypeCode
	// MatchPackageRegex regexp.Regexp
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
				if ok {
					af.PackageFilenames = append(af.PackageFilenames, filePath)
				}
			}
		}
		return nil
	}
	return filepath.WalkDir(rootDir, dirWalker)
}

// NOTE: Do not forget to close all files
func (af *ArchiveFolder) FolderMap(
	rootDir string, recursive bool, filterRange [2]time.Time) error {
	err := af.FolderListing(rootDir, recursive, filterRange)
	if err != nil {
		return err
	}
	for _, folder := range af.PackageFilenames {
		zipr, err := zip.OpenReader(folder)
		if err != nil {
			return err
		}
		af.PackageReaders = append(af.PackageReaders, zipr)
	}
	return nil
}

func (af *ArchiveFolder) FolderProcess() {
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
