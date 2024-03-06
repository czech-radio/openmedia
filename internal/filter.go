package internal

import (
	"archive/zip"
	"fmt"
	"io/fs"
	"log/slog"
	"path/filepath"
	"strings"
)

type FilterOptions struct {
	SourceDirectory        string
	DestinationDirectory   string
	RecurseSourceDirectory bool
	InvalidFileContinue    bool

	FilterType WorkerTypeCode
	OutputType string // csv contacts,unique contact fields,
	CSVdelim   string
	CSVheader  bool

	DateFrom string
	DateTo   string
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
	Rows                  []Fields
	FieldsUniqueValues    map[int]UniqueValues // FiledID vs UniqueValues
	MaxUniqueCount        int                  // Field which has the highest unique values count - servers. Used when transforming rows to columns
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
	Years            []int
	IsoWeeks         []int
	PackageType      WorkerTypeCode
	// MatchPackageRegex regexp.Regexp
}

func MatchArchivePackage(packageFilename string, years []int, weeks []int, WorkerTypeCode WorkerTypeCode) bool {
	// fileName '2020_W05_MINIFIED.zip'
	packageType := WorkerTypeMap[WorkerTypeCode]
	fmt.Println(packageFilename, years, weeks, packageType)
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

func MatchArchivePackageFile(packageFilename string) {
}

func (af *ArchiveFolder) FolderListing(rootDir string, recursive bool) error {
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
		file_extension := filepath.Ext(filePath)
		switch file_extension {
		case ".zip":
			fileName := filepath.Base(filePath)
			ok := MatchArchivePackage(
				fileName, af.Years, af.IsoWeeks, af.PackageType,
			)
			if ok {
				af.PackageFilenames = append(af.PackageFilenames, filePath)
			}
		case ".xml":
			af.SimpleFilenames = append(af.SimpleFilenames, filePath)
		}
		return nil
	}
	return filepath.WalkDir(rootDir, dirWalker)
}

// NOTE: Do not forget to close all files
func (af *ArchiveFolder) FolderMap(
	rootDir string, recursive bool) error {
	err := af.FolderListing(rootDir, recursive)
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
