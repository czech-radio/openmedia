package internal

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type ControlFlowAction int

const (
	Continue ControlFlowAction = iota
	Skip
	Break
)

type WorkerTypeCode int

const (
	WorkerTypeOriginal WorkerTypeCode = iota
	WorkerTypeMinified
	WorkerTypeCSV
)

var WorkerTypeMap map[WorkerTypeCode]string = map[WorkerTypeCode]string{
	WorkerTypeOriginal: "ORIGINAL",
	WorkerTypeMinified: "MINIFIED",
	WorkerTypeCSV:      "CSV",
}

type Process struct {
	Options ProcessOptions
	Results ProcessResults
	Errors  []error
	Workers map[string]*ArchiveWorker // WorkerName->
	WG      *sync.WaitGroup
}

type ProcessOptions struct {
	SourceDirectory        string
	DestinationDirectory   string
	InputEncoding          string
	OutputEncoding         string
	ValidateWithDefaultXSD bool   // validate with bundled file
	ValidateWithXSD        string // path to XSD file
	ValidatePre            bool
	ValidatePost           bool
	CompressionType        string
	InvalidFileRename      bool
	InvalidFileContinue    bool
}

type ProcessResults struct {
	Weeks              int
	FilesProcessed     int
	FilesSuccess       int
	FilesFailure       int
	FilesCount         int
	SizeOriginal       uint64
	SizePackedBackup   uint64
	SizeMinified       uint64
	SizePackedMinified uint64
	Duplicates         map[string][]string // dstFile vs srcFile
	DuplicatesFound    bool
}

type FileMeta struct {
	Year                 int
	Month                int
	Day                  int
	Hour                 int
	Minute               int
	Second               int
	Week                 int
	Weekday              time.Weekday
	WorkerName           string
	CompressionType      string
	OpenMediaFileType    *OpenMediaFileType
	RundownNameNew       string
	FileInfo             os.FileInfo
	FileReader           io.Reader
	FilePathSource       string
	FilePathInArchive    string
	DirectoryDestination string
	DirectorySource      string
}

type ArchiveWorker struct {
	Call           chan *FileMeta
	WorkrerName    string
	WorkerTypeName string
	WorkerTypeCode WorkerTypeCode
	ArchivePath    string
	ArchiveFile    *os.File
	ArchiveWriter  *zip.Writer
}

func (w *ArchiveWorker) Init(dstdir, archiveName string) error {
	w.Call = make(chan *FileMeta)
	path := filepath.Join(dstdir, archiveName)
	w.ArchivePath = path
	archive, err := os.Create(path)
	if err != nil {
		return err
	}
	w.ArchiveFile = archive
	w.ArchiveWriter = zip.NewWriter(archive)
	return nil
}

func (p *Process) DestroyWorkers() {
	if p.Workers == nil {
		return
	}
	for _, worker := range p.Workers {
		worker.ArchiveWriter.Close()
	}
}

func (f *FileMeta) Parse(
	metaInfo OMmetaInfo, fileInfo os.FileInfo, compressionType, sourceDir, targetDir string, sourceFilePath string, reader io.Reader) error {
	date := metaInfo.Date
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
	f.DirectoryDestination = filepath.Join(targetDir, omFileType.OutputDir)
	f.FileInfo = fileInfo
	f.DirectorySource = sourceDir
	f.FilePathSource = sourceFilePath
	pathInArchive, err := filepath.Rel(sourceDir, sourceFilePath)
	if err != nil {
		return err
	}
	f.FilePathInArchive = filepath.Join(filepath.Dir(pathInArchive), f.RundownNameNew+filepath.Ext(sourceFilePath))
	f.FileReader = reader
	f.CompressionType = compressionType
	return nil
}

func (f *FileMeta) SetWeekWorkerName(wtc WorkerTypeCode) string {
	workerName, _ := WorkerTypeMap[wtc]
	f.WorkerName = fmt.Sprintf("%d_W%02d_%s.%s", f.Year, f.Week, workerName, f.CompressionType)
	return f.WorkerName
}

func (p *Process) ErrorHandle(errMain error, errorsPartial ...error) ControlFlowAction {
	if errMain == nil {
		return Continue
	}
	// p.Results.FilesFailure++
	p.Errors = append(p.Errors, errMain)
	p.Errors = append(p.Errors, errorsPartial...)

	if p.Options.InvalidFileContinue {
		return Skip
	} else {
		return Break
	}
}

func (p *Process) CheckDuplicatesInArchive(fi *FileMeta) error {
	if p.Results.Duplicates == nil {
		p.Results.Duplicates = make(map[string][]string)
	}
	dupes, _ := p.Results.Duplicates[fi.FilePathInArchive]
	p.Results.Duplicates[fi.FilePathInArchive] = append(dupes, fi.FilePathSource)

	if len(dupes) > 0 {
		p.Results.DuplicatesFound = true
		return fmt.Errorf("file %s will result in duplicate %s file in archive", fi.FilePathSource, fi.FilePathInArchive)
	}
	return nil
}

func (p *Process) PrepareOutput() error {
	for _, t := range OpenMediaFileTypeMap {
		outputdir := filepath.Join(p.Options.DestinationDirectory, t.OutputDir)
		if err := os.MkdirAll(outputdir, 0700); err != nil {
			return err
		}
		slog.Debug("created directory", "output_dir", outputdir)
	}
	return nil
}

func (p *Process) Folder() error {
	validateResult, err := ValidateFilenamesInDirectory(p.Options.SourceDirectory)
	if p.ErrorHandle(err, validateResult.Errors...) == Break {
		return err
	}
	p.Results.FilesCount = validateResult.FilesCount
	p.Results.FilesFailure = validateResult.FilesFailure
	p.Workers = make(map[string]*ArchiveWorker)
	p.WG = new(sync.WaitGroup)
	err = p.PrepareOutput()
	if err != nil {
		return err
	}
processFolder:
	for _, file := range validateResult.FilesValid {
		err := p.File(file)
		flow := p.ErrorHandle(err)
		switch flow {
		case Skip:
			continue processFolder
		case Break:
			break processFolder
		}
	}
	p.WG.Wait()
	res := p.Results
	p.WorkerLogInfo("GLOBAL_ORIGINAL", res.SizeOriginal, res.SizePackedBackup, res.SizeOriginal, p.Options.SourceDirectory)
	p.WorkerLogInfo("GLOBAL_MINIFY", res.SizeOriginal, res.SizePackedMinified, res.SizeMinified, p.Options.SourceDirectory)
	if p.Results.DuplicatesFound {
		slog.Error("duplicates found")
		ms, err := json.MarshalIndent(p.Results.Duplicates, "", "  ")
		if err != nil {
			slog.Error("cannot marshal dupes")
		}
		fmt.Printf("%s\n", ms)
	}
	p.DestroyWorkers()
	return nil
}

func (p *Process) File(filePath string) error {
	// Open file
	fileHandle, err := os.Open(filePath)
	if err != nil {
		return err
	}
	// defer fileHandle.Close()

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
	fileMetaOriginal := FileMeta{}
	reader := bytes.NewReader(data)
	opts := p.Options
	err = fileMetaOriginal.Parse(omMetaInfo, fileInfo, opts.CompressionType, opts.SourceDirectory, opts.DestinationDirectory, filePath, reader)
	if err != nil {
		return err
	}

	// Check resulting dupes
	err = p.CheckDuplicatesInArchive(&fileMetaOriginal)
	if err != nil {
		return err
	}

	// 1. Archivate original
	fileMetaOriginal.SetWeekWorkerName(WorkerTypeOriginal)
	err = p.CallArchivWorker(&fileMetaOriginal, WorkerTypeOriginal)
	if err != nil {
		return err
	}
	p.WG.Add(1)

	// Transform output data
	pr = PipeRundownMarshal(om)
	pr = PipeRundownHeaderAdd(pr)

	// 2. Archivate original
	fileMetaMinify := fileMetaOriginal
	fileMetaMinify.FileReader = pr
	fileMetaMinify.SetWeekWorkerName(WorkerTypeMinified)
	err = p.CallArchivWorker(&fileMetaMinify, WorkerTypeMinified)
	if err != nil {
		return err
	}
	p.WG.Add(1)
	return nil
}

func (p *Process) WorkerLogInfo(workerType string, sizeOrig, sizePacked, sizeMinified uint64, filePath string) {
	archiveRatio := float64(sizePacked) / float64(sizeOrig)
	archiveRatioString := fmt.Sprintf("%.3f", archiveRatio)
	minifyRatio := float64(sizeMinified) / float64(sizeOrig)
	minifyRatioString := fmt.Sprintf("%.3f", minifyRatio)
	slog.Info(
		workerType, "ArhiveRatio", archiveRatioString,
		"MinifyRatio", minifyRatioString,
		"original", sizeOrig, "compressed", sizePacked,
		"minified", sizeMinified, "file", filePath)
}

func (p *Process) CallArchivWorker(fm *FileMeta, workerTypeCode WorkerTypeCode) error {
	worker, ok := p.Workers[fm.WorkerName]
	if !ok {
		worker = new(ArchiveWorker)
		err := worker.Init(fm.DirectoryDestination, fm.WorkerName)
		if err != nil {
			return err
		}
		p.Workers[fm.WorkerName] = worker
		go func(w *ArchiveWorker, workerTypeCode WorkerTypeCode) {
			for {
				workerParams := <-w.Call
				origSize, compressedSize, bytesWritten, err := p.Archivate(worker, workerParams)
				if err != nil {
					slog.Error(err.Error())
					break
				}
				p.WorkerLogInfo(fm.WorkerName, origSize, compressedSize,
					bytesWritten, workerParams.FilePathSource)
				// Update results stats
				switch workerTypeCode {
				case WorkerTypeMinified:
					p.Results.SizePackedMinified += compressedSize
					p.Results.SizeMinified += bytesWritten
					p.Results.FilesProcessed++
				case WorkerTypeOriginal:
					p.Results.SizePackedBackup += compressedSize
					p.Results.SizeOriginal += origSize
				}
				p.WG.Done()
			}
		}(worker, workerTypeCode)
	}
	worker.Call <- fm
	return nil
}

func (p *Process) Archivate(worker *ArchiveWorker, f *FileMeta) (
	uint64, uint64, uint64, error) {
	var fileSize, compressedSize, minifiedSize uint64
	// Create a new zip file header
	header, err := zip.FileInfoHeader(f.FileInfo)
	if err != nil {
		return fileSize, compressedSize, minifiedSize, nil
	}
	finfo, err := worker.ArchiveFile.Stat()
	if err != nil {
		return fileSize, compressedSize, minifiedSize, nil
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
		return fileSize, compressedSize, minifiedSize, nil
	}
	// Copy the file content to the zip entry
	bytesCount, err := io.Copy(entry, f.FileReader)
	minifiedSize = uint64(bytesCount)
	if err != nil {
		return fileSize, compressedSize, minifiedSize, nil
	}
	// compressedSize := header.CompressedSize64 // Not working ->0
	finfo, err = worker.ArchiveFile.Stat()
	if err != nil {
		return fileSize, compressedSize, minifiedSize, nil
	}
	after := finfo.Size()
	compressedSize = uint64(after - before)
	return fileSize, compressedSize, minifiedSize, nil
}
