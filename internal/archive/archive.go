package internal

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/triopium/go_utils/pkg/helper"

	"io"
	"log/slog"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var ArchiveTimeZone, _ = time.LoadLocation("")

// WorkerTypeCode represents the selected line of processing.
type WorkerTypeCode int

const (
	WorkerTypeZIPoriginal WorkerTypeCode = iota
	WorkerTypeZIPminified
	WorkerTypeRundownXMLutf8
	WorkerTypeRundownXMLutf16le
	WorkerTypeCSVcontactsFields
	WorkerTypeCSVcontactsUniqueFields
	WorkerTypeCSVprodukce
)

var WorkerTypeMap = map[WorkerTypeCode]string{
	WorkerTypeZIPoriginal:             "ORIGINAL.zip",
	WorkerTypeZIPminified:             "MINIFIED.zip",
	WorkerTypeRundownXMLutf8:          "RD_utf8_xml",
	WorkerTypeRundownXMLutf16le:       "RD_utf16le_xml",
	WorkerTypeCSVcontactsFields:       "CONTACT_FIELDS.csv",
	WorkerTypeCSVcontactsUniqueFields: "CONTACT_FIELDS_UNIQUE.csv",
	WorkerTypeCSVprodukce:             "PRODUKCE.csv",
}

func WorkeTypeCodeGet(codeName string) WorkerTypeCode {
	for code, name := range WorkerTypeMap {
		if codeName == name {
			return code
		}
	}
	panic(fmt.Errorf("no worker code found for given name: %s", codeName))
}

func InferEncoding(wtc WorkerTypeCode) helper.FileEncodingCode {
	var enc helper.FileEncodingCode
	switch wtc {
	case WorkerTypeZIPminified, WorkerTypeRundownXMLutf8:
		enc = helper.UTF8
	case WorkerTypeZIPoriginal, WorkerTypeRundownXMLutf16le:
		enc = helper.UTF16le
	}
	return enc
}

type Archive struct {
	Options        ArchiveOptions
	Results        ArchiveResults
	ErrMap         helper.ErrMap
	Errors         []error
	ArchiveWorkers map[string]*ArchivePackageWorker // WorkerName->
	WG             *sync.WaitGroup
}

type ArchiveOptions struct {
	SourceDirectory          string
	OutputDirectory          string
	CompressionType          string
	InvalidFilenameContinue  bool
	InvalidFileContinue      bool
	InvalidFileRename        bool
	ProcessedFileRename      bool
	ProcessedFileDelete      bool
	PreserveFoldersInArchive bool
	RecurseSourceDirectory   bool

	// InputEncoding          string
	// OutputEncoding         string
	// ValidateWithDefaultXSD bool   // validate with bundled file
	// ValidateWithXSD        string // path to XSD file
	// ValidatePre            bool
	// ValidatePost           bool
}

type ArchiveResults struct {
	*ProcessStats
	SizeOriginal       uint64
	SizePackedBackup   uint64
	SizeMinified       uint64
	SizePackedMinified uint64
	Duplicates         map[string][]string // dstFile vs srcFile
	DuplicatesFound    int
}

type ArchiveItemFileMeta struct {
	Year              int
	Month             int
	Day               int
	Hour              int
	Minute            int
	Second            int
	Week              int
	Weekday           time.Weekday
	WorkerName        string
	CompressionType   string
	OpenMediaFileType *OpenMediaFileType

	RundownNameNew       string
	FileInfo             os.FileInfo
	FileReader           io.Reader
	FilePathSource       string
	FilePathInArchive    string
	DirectoryDestination string
	DirectorySource      string
}

type ArchivePackageWorker struct {
	Call           chan *ArchiveItemFileMeta
	WorkerName     string
	WorkerTypeName string
	WorkerTypeCode WorkerTypeCode
	ArchivePath    string
	ArchiveFile    *os.File
	ArchiveWriter  *zip.Writer
	ArchiveFiles   map[string]int // map filenames in archive
}

func (w *ArchivePackageWorker) MapOldArchive(archivePath string) (bool, error) {
	// Check if there is an old archive
	ok, err := helper.FileExists(archivePath)
	if err != nil {
		return false, err
	}
	if !ok {
		w.ArchiveFiles = make(map[string]int)
		return false, nil
	}
	// Read files already present in an archive
	if ok {
		r, err := zip.OpenReader(archivePath)
		if err != nil {
			return true, err
		}
		defer r.Close()
		w.ArchiveFiles = make(map[string]int, len(r.File))
		for _, file := range r.File {
			archiveNameAndFilePath := filepath.Join(w.WorkerName, file.Name)
			w.ArchiveFiles[archiveNameAndFilePath]++
		}
	}
	return true, nil
}

func (w *ArchivePackageWorker) Init(dstdir, archiveName string) error {
	w.Call = make(chan *ArchiveItemFileMeta)
	archivePath := filepath.Join(dstdir, archiveName)
	w.ArchivePath = archivePath

	// 1. Open old zip file if present and read file list
	exists, err := w.MapOldArchive(archivePath)
	if err != nil {
		return err
	}

	var archive *os.File
	if !exists {
		archive, err = os.OpenFile(archivePath, os.O_RDWR|os.O_CREATE, 0600)
		if err != nil {
			return err
		}
		w.ArchiveFile = archive
		w.ArchiveWriter = zip.NewWriter(archive)
	}
	// 2. Skip if exists. TODO: Golang zip package does not support adding files to old archive.
	// archive must be reacreated such that the content from old archive is copied to the new archive. Alternatively use tar zip.
	if exists {
		return fmt.Errorf("archive already exists: %s", archiveName)
	}
	return nil
}

func (p *Archive) DestroyWorkers() {
	if p.ArchiveWorkers == nil {
		return
	}
	for _, worker := range p.ArchiveWorkers {
		worker.ArchiveWriter.Close()
		worker.ArchiveFile.Close()
	}
}

func (f *ArchiveItemFileMeta) Parse(
	metaInfo OMmetaInfo, fileInfo os.FileInfo, reader io.Reader,
	opts *ArchiveOptions, sourceFilePath string) error {
	date := metaInfo.Date
	if !helper.IsOlderThanOneISOweek(date, time.Now()) {
		return fmt.Errorf("file date not older than 1 ISOWeek: %s", sourceFilePath)
	}
	year, week := date.ISOWeek()
	f.Year = year
	f.Month = int(date.Month())
	f.Day = date.Day()
	f.Hour = date.Hour()
	f.Week = week
	f.Weekday = date.Weekday()
	f.Minute = date.Minute()
	f.Second = date.Second()
	omFileType, err := GetOMtypeByTemplateName(metaInfo.RundownType)
	if err != nil {
		return err
	}
	f.OpenMediaFileType = omFileType
	switch omFileType.Code {
	case OmFileTypeRundown:
		f.RundownNameNew = fmt.Sprintf(
			"%s_%s_%s_%s_W%02d_%04d_%02d_%02d",
			omFileType.ShortHand, metaInfo.HoursRange,
			metaInfo.Name,
			f.Weekday, f.Week, f.Year, f.Month, f.Day)
	case OmFileTypeContact:
		f.RundownNameNew = fmt.Sprintf(
			"%s_%s_%s_W%02d_%04d_%02d_%02dT%02d%02d%02d",
			omFileType.ShortHand, metaInfo.Name, f.Weekday, f.Week,
			f.Year, f.Month, f.Day, f.Hour, f.Minute, f.Second)
	}
	// f.DirectoryDestination = filepath.Join(opts.DestinationDirectory, omFileType.OutputDir)
	f.DirectoryDestination = opts.OutputDirectory
	f.FileInfo = fileInfo
	f.DirectorySource = opts.SourceDirectory
	f.FilePathSource = sourceFilePath
	pathInArchive, err := filepath.Rel(opts.SourceDirectory, sourceFilePath)
	if err != nil {
		return err
	}
	if opts.PreserveFoldersInArchive {
		f.FilePathInArchive = filepath.Join(
			filepath.Dir(pathInArchive), f.RundownNameNew+filepath.Ext(sourceFilePath))
	} else {
		f.FilePathInArchive = f.RundownNameNew + filepath.Ext(sourceFilePath)
	}
	f.FileReader = reader
	f.CompressionType = opts.CompressionType
	return nil
}

func (f *ArchiveItemFileMeta) SetWeekWorkerName(
	wtc WorkerTypeCode) string {
	workerTypeString := WorkerTypeMap[wtc]
	// f.WorkerName = fmt.Sprintf("%s/%d_W%02d_%s.%s", f.OpenMediaFileType.OutputDir, f.Year, f.Week, workerTypeString, f.CompressionType)
	f.WorkerName = fmt.Sprintf(
		"%s/%d_W%02d_%s", f.OpenMediaFileType.OutputDir,
		f.Year, f.Week, workerTypeString)
	return f.WorkerName
}

// func (p *Archive) ErrorHandleAdd(
// 	errMain error, errorsPartial ...error) helper.ControlFlowAction {
// 	if errMain == nil {
// 		return helper.Continue
// 	}
// 	if p.ErrMap == nil {
// 		p.ErrMap = make(map[string][]string)
// 	}
// 	_, ok := p.ErrMap[errMain.Error()]
// 	if !ok {
// 		p.ErrMap[errMain.Error()] = make([]string, 0)
// 	}
// 	if len(errorsPartial) == 0 {
// 		p.ErrMap[errMain.Error()] = append(
// 			p.ErrMap[errMain.Error()], "")
// 	}
// 	if len(errorsPartial) > 0 {
// 		for _, e := range errorsPartial {
// 			p.ErrMap[errMain.Error()] = append(
// 				p.ErrMap[errMain.Error()], e.Error())
// 		}
// 	}
// 	if p.Options.InvalidFileContinue {
// 		slog.Info("skipped processing file")
// 		return helper.Skip
// 	}
// 	slog.Info("halt processing files: bad file encountered")
// 	return helper.Break
// }

// func (p *Archive) ErrorHandle(
// 	errMain error, errorsPartial ...error) helper.ControlFlowAction {
// 	p.Results.FilesProcessed++
// 	if errMain == nil {
// 		p.Results.FilesSuccess++
// 		return helper.Continue
// 	}

// 	p.Results.FilesFailure++
// 	// Get info about function which called this hadnler
// 	fileName, funcName, line := helper.TraceFunction(2)
// 	slog.Error(errMain.Error(),
// 		"source", fileName, "function", funcName, "line", line)
// 	p.Errors = append(p.Errors, errMain)
// 	if len(errorsPartial) > 0 {
// 		p.Errors = append(p.Errors, errorsPartial...)
// 	}

// 	if p.Options.InvalidFileContinue {
// 		slog.Info("skipped processing file")
// 		return helper.Skip
// 	}
// 	slog.Info("halt processing files: bad file encountered")
// 	return helper.Break
// }

func (p *Archive) PrepareOutput() error {
	for _, t := range OpenMediaFileTypeMap {
		outputdir := filepath.Join(p.Options.OutputDirectory, t.OutputDir)
		if err := os.MkdirAll(outputdir, 0700); err != nil {
			return err
		}
		slog.Debug("created directory", "output_dir", outputdir)
	}
	return nil
}

func ValidRatio(valid, all int) string {
	return fmt.Sprintf("%d/%d", valid, all)
}

func (p *Archive) Folder() error {
	// 1. Folder files validation
	vr, err := ValidateFilenamesInDirectory(
		p.Options.SourceDirectory, p.Options.RecurseSourceDirectory)
	slog.Info("Archive filenames validation result",
		"valid_ratio", ValidRatio(vr.SuccessCount, vr.ProcessedCount))
	if err != nil {
		return err
	}
	err = vr.ErrMap.MarshalError("validation")
	if err != nil && !p.Options.InvalidFilenameContinue {
		return err
	}

	// 2. Prepare archivation
	p.Results.ProcessStats = new(ProcessStats)
	p.ErrMap = make(helper.ErrMap)
	p.Results.AllCount = vr.SuccessCount
	p.ArchiveWorkers = make(map[string]*ArchivePackageWorker)
	p.WG = new(sync.WaitGroup)
	err = p.PrepareOutput()
	if err != nil {
		return err
	}

	// 3. archive create
	p.WG.Add(1)
	for _, file := range vr.ValidFilenames {
		err := p.File(file)
		p.ErrMap.Aggregate(err, file)
		p.Results.ProcessedCount++
		if err != nil {
			p.Results.FailureCount++
		}
		if err != nil && !p.Options.InvalidFileContinue {
			return fmt.Errorf("%s %w", file, err)
		}
		p.Results.SuccessCount++
	}
	p.WG.Done()
	p.WG.Wait()

	// res := p.Results
	// p.WorkerLogInfo("GLOBAL_ORIGINAL",
	// 	res.SizeOriginal, res.SizePackedBackup, res.SizeOriginal,
	// 	p.Options.SourceDirectory, p.Options.OutputDirectory)
	// p.WorkerLogInfo("GLOBAL_MINIFY",
	// 	res.SizeOriginal, res.SizePackedMinified, res.SizeMinified,
	// 	p.Options.SourceDirectory, p.Options.OutputDirectory)

	// if p.Results.DuplicatesFound > 0 {
	// 	dupesErr := fmt.Errorf("duplicates found, cout: %d",
	// 		p.Results.DuplicatesFound)
	// 	p.ErrorHandle(dupesErr)
	// }

	// _ = ErrorsMarshalLog(p.Errors)
	p.DestroyWorkers()
	return nil
}

func ErrorsMarshalLog(errs []error) error {
	var results []string
	var marshalErrors []error
	if errs == nil {
		return nil
	}
	if len(errs) == 0 {
		return nil
	}
	for e := range errs {
		result, err := json.MarshalIndent(errs[e].Error(), "", "\t")
		if err != nil {
			slog.Warn("cannot marshal error", "error", err.Error())
			marshalErrors = append(marshalErrors, err)
			continue
		}
		results = append(results, string(result))
	}
	slog.Error("AggregatedErrors", "errors", results)
	if len(marshalErrors) > 0 {
		return fmt.Errorf(
			"error unmarshaling AggregatedErrors, count: %d",
			len(marshalErrors))
	}
	return nil
}

func (p *Archive) File(sourceFilePath string) error {
	// Open file
	fileHandle, err := os.Open(sourceFilePath)
	if err != nil {
		return err
	}
	defer fileHandle.Close()

	// Read file data
	data, err := io.ReadAll(fileHandle)
	if err != nil {
		return err
	}

	// Transform input data
	dataReader := bytes.NewReader(data)
	pr := PipeUTF16leToUTF8(dataReader)
	pr = PipeRundownHeaderAmmend(pr)

	// Unmarshal (Parse and validate)
	om, err := PipeRundownUnmarshal(pr) // treat EOF error
	if err != nil {
		return err
	}

	// Infer rundown meta info from OM_HEADER fields
	omMetaInfo, err := om.OMmetaInfoParse()
	if err != nil {
		return err
	}

	// Parse file meta information
	fileInfo, err := fileHandle.Stat()
	if err != nil {
		return err
	}
	fileMetaOriginal := ArchiveItemFileMeta{}
	reader := bytes.NewReader(data)
	err = fileMetaOriginal.Parse(omMetaInfo, fileInfo, reader,
		&p.Options, sourceFilePath)
	if err != nil {
		return err
	}

	// 1. Create archive from original files
	fileMetaOriginal.SetWeekWorkerName(WorkerTypeZIPoriginal)
	err = p.CallArchiveWorker(&fileMetaOriginal, WorkerTypeZIPoriginal)
	if err != nil {
		return err
	}
	p.WG.Add(1)

	// Transform output data
	pr = PipeRundownMarshal(om)
	pr = PipeRundownHeaderAdd(pr)

	// 2. Create archive from minified files
	fileMetaMinify := fileMetaOriginal
	fileMetaMinify.FileReader = pr
	fileMetaMinify.SetWeekWorkerName(WorkerTypeZIPminified)
	err = p.CallArchiveWorker(&fileMetaMinify, WorkerTypeZIPminified)
	if err != nil {
		return err
	}
	p.WG.Add(1)
	if p.Options.ProcessedFileRename {
		return helper.ProcessedFileRename(sourceFilePath)
	}
	if p.Options.ProcessedFileDelete {
		return os.Remove(sourceFilePath)
	}
	return nil
}

func (p *Archive) WorkerLogInfo(
	workerType string, sizeOrig, sizePacked, sizeMinified uint64,
	filePathSource, filePathDestination string) {
	archiveRatio := float64(sizePacked) / float64(sizeOrig)
	archiveRatioString := fmt.Sprintf("%.3f", archiveRatio)
	minifyRatio := float64(sizeMinified) / float64(sizeOrig)
	minifyRatioString := fmt.Sprintf("%.3f", minifyRatio)
	slog.Info(
		workerType, "ArhiveRatio", archiveRatioString,
		"MinifyRatio", minifyRatioString,
		"original", sizeOrig, "compressed", sizePacked,
		"minified", sizeMinified, "filePathSource", filePathSource, "filePathDestination", filePathDestination)
}

func (p *Archive) CheckArchiveWorkerDupes(
	worker *ArchivePackageWorker, fm *ArchiveItemFileMeta) error {
	if p.Results.Duplicates == nil {
		p.Results.Duplicates = make(map[string][]string)
	}
	archiveNameAndFileName := filepath.Join(
		worker.WorkerName, fm.FilePathInArchive)
	_, filePresent := worker.ArchiveFiles[archiveNameAndFileName]
	if !filePresent {
		worker.ArchiveFiles[archiveNameAndFileName]++
		return nil
	}
	p.Results.DuplicatesFound++
	dupes := p.Results.Duplicates[fm.FilePathInArchive]
	p.Results.Duplicates[fm.FilePathInArchive] = append(
		dupes, fm.FilePathSource)
	return fmt.Errorf(
		"file %s will result in duplicate in %s",
		fm.FilePathSource, archiveNameAndFileName)
}

func (p *Archive) CallArchiveWorker(
	fm *ArchiveItemFileMeta, workerTypeCode WorkerTypeCode) error {
	worker, ok := p.ArchiveWorkers[fm.WorkerName]
	if !ok {
		worker = new(ArchivePackageWorker)
		worker.WorkerName = fm.WorkerName
		err := worker.Init(fm.DirectoryDestination, fm.WorkerName)
		if err != nil {
			return fmt.Errorf("file not proccessed: %s, %w",
				fm.FilePathSource, err)
		}
		p.ArchiveWorkers[fm.WorkerName] = worker
		go func(w *ArchivePackageWorker, workerTypeCode WorkerTypeCode) {
			for {
				workerParams := <-w.Call
				origSize, compressedSize, bytesWritten, err := p.AddFileToArchive(
					worker, workerParams)
				if err != nil {
					slog.Error(err.Error())
					break
				}
				fileDestinationInArchive := filepath.Join(
					worker.WorkerName, workerParams.FilePathInArchive)
				p.WorkerLogInfo(
					fm.WorkerName, origSize, compressedSize,
					bytesWritten, workerParams.FilePathSource,
					fileDestinationInArchive,
				)
				// Update results stats
				switch workerTypeCode {
				case WorkerTypeZIPminified:
					p.Results.SizePackedMinified += compressedSize
					p.Results.SizeMinified += bytesWritten
				case WorkerTypeZIPoriginal:
					p.Results.SizePackedBackup += compressedSize
					p.Results.SizeOriginal += origSize
				}
				p.WG.Done()
			}
		}(worker, workerTypeCode)
	}
	err := p.CheckArchiveWorkerDupes(worker, fm)
	if err != nil {
		return err
	}
	worker.Call <- fm
	return nil
}

func (p *Archive) AddFileToArchive(
	worker *ArchivePackageWorker, f *ArchiveItemFileMeta) (
	uint64, uint64, uint64, error) {
	var fileSize, compressedSize, minifiedSize uint64
	// Create a new zip file header
	header, err := zip.FileInfoHeader(f.FileInfo)
	if err != nil {
		return fileSize, compressedSize, minifiedSize, err
	}
	finfo, err := worker.ArchiveFile.Stat()
	if err != nil {
		return fileSize, compressedSize, minifiedSize, err
	}
	before := finfo.Size()
	// fileSize := header.UncompressedSize64 // Not working -> 0
	fileSize = uint64(f.FileInfo.Size())
	header.Method = zip.Deflate
	// Set the name of the file within the zip archive
	header.Name = f.FilePathInArchive
	// Create a new entry in the zip archive
	entry, err := worker.ArchiveWriter.CreateHeader(header)
	if err != nil {
		return fileSize, compressedSize, minifiedSize, err
	}
	// Copy the file content to the zip entry
	bytesCount, err := io.Copy(entry, f.FileReader)
	minifiedSize = uint64(bytesCount)
	if err != nil {
		return fileSize, compressedSize, minifiedSize, err
	}
	// compressedSize := header.CompressedSize64 // Not working ->0
	finfo, err = worker.ArchiveFile.Stat()
	if err != nil {
		return fileSize, compressedSize, minifiedSize, err
	}
	after := finfo.Size()
	compressedSize = uint64(after - before)
	return fileSize, compressedSize, minifiedSize, nil
}
