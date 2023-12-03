package internal

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"math/rand"
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

type Process struct {
	Options ProcessOptions
	Results ProcessResults
	Workers map[string]*ArchiveWorker
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
	ArchiveType            string
	InvalidFileRename      bool
	InvalidFileContinue    bool
}

type ProcessResults struct {
	Weeks          int
	FilesProcessed int
	FilesSuccess   int
	FilesFailure   int
	FilesCount     int
	SizeOriginal   int
	SizeBackup     int
	SizeMinified   int
	Errors         []error
}

type FileMeta struct {
	Year                 int
	Month                int
	Day                  int
	Week                 int
	Weekday              time.Weekday
	Archives             map[string]string
	RundownNameNew       string
	FilePathSource       string
	FilePathInArchive    string
	FileInfo             os.FileInfo
	FileReader           io.Reader
	DirectoryDestination string
	DirectorySource      string
}

type ArchiveWorker struct {
	Call          chan *FileMeta
	ArchivePath   string
	ArchiveFile   *os.File
	ArchiveWriter *zip.Writer
}

func (w *ArchiveWorker) Init(dstdir, archieName string) error {
	w.Call = make(chan *FileMeta)
	path := filepath.Join(dstdir, archieName)
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
	date time.Time, fileInfo os.FileInfo, archiveType, sourceDir, targetDir string, sourceFilePath string, reader io.Reader) error {
	year, week := date.ISOWeek()
	f.Year = year
	f.Month = int(date.Month())
	f.Day = date.Day()
	f.Week = week
	f.Weekday = date.Weekday()
	f.Archives = make(map[string]string)

	// archive names
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	randint := r1.Intn(100000)
	// name := fmt.Sprintf("MINIFIED_%03d", randint)
	name := "MINIFIED"

	f.Archives[name] = fmt.Sprintf("%d_W%02d_%s.%s", f.Year, f.Week, name, archiveType)
	name = "ORIGINAL"
	f.Archives[name] = fmt.Sprintf("%d_W%02d_%s.%s", f.Year, f.Week, name, archiveType)
	// f.RundownNameNew = fmt.Sprintf("RD_%s_W%02d_%04d_%02d_%02d", f.Weekday, f.Week, f.Year, f.Month, f.Day)
	f.RundownNameNew = fmt.Sprintf("RD_%s_W%02d_%04d_%02d_%02d_%03d", f.Weekday, f.Week, f.Year, f.Month, f.Day, randint)
	f.FileInfo = fileInfo
	f.DirectorySource = sourceDir
	f.DirectoryDestination = targetDir
	f.FilePathSource = sourceFilePath
	fmt.Println(f.FilePathSource)
	pathInArchive, err := filepath.Rel(sourceDir, sourceFilePath)
	if err != nil {
		return err
	}
	f.FilePathInArchive = filepath.Join(filepath.Dir(pathInArchive), f.RundownNameNew+filepath.Ext(sourceFilePath))
	f.FileReader = reader
	return nil
}

func (p *Process) InfoLog() {
	ds := p.Results
	msg := fmt.Sprintf("%d/%d, %d:%d",
		ds.FilesProcessed, ds.FilesCount,
		ds.FilesSuccess, ds.FilesFailure)
	slog.Debug(msg)
}

func (p *Process) ErrorHandle(errMain error, errorsPartial ...error) ControlFlowAction {
	if errMain == nil {
		return Continue
	}
	// p.Results.FilesFailure++
	p.Results.Errors = append(p.Results.Errors, errMain)
	p.Results.Errors = append(p.Results.Errors, errorsPartial...)

	if p.Options.InvalidFileContinue {
		return Skip
	} else {
		return Break
	}
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
	slog.Debug("waiting not")
	p.DestroyWorkers()
	return nil
}

func (p *Process) File(filePath string) error {
	// Open file
	p.Results.FilesProcessed++
	// p.InfoLog()
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

	// Infer rundown date from OM_HEADER field
	date, err := om.RundownDate()
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
	err = fileMetaOriginal.Parse(date, fileInfo, opts.ArchiveType, opts.SourceDirectory, opts.DestinationDirectory, filePath, reader)
	if err != nil {
		return err
	}

	// Archivate original
	err = p.CallArchivWorker(&fileMetaOriginal, "ORIGINAL")
	if err != nil {
		return err
	}
	p.WG.Add(1)

	// Transform output data
	pr = PipeRundownMarshal(om)
	pr = PipeRundownHeaderAdd(pr)

	// Archivate original
	fileMetaMinify := fileMetaOriginal
	fileMetaMinify.FileReader = pr
	err = p.CallArchivWorker(&fileMetaMinify, "MINIFIED")
	if err != nil {
		return err
	}
	p.WG.Add(1)
	return nil
}

func (p *Process) CallArchivWorker(fm *FileMeta, workerType string) error {
	workerName := fm.Archives[workerType]
	worker, ok := p.Workers[workerName]
	if !ok {
		worker = new(ArchiveWorker)
		err := worker.Init(fm.DirectoryDestination, workerName)
		if err != nil {
			return err
		}
		p.Workers[workerName] = worker
		go func() {
			for {
				f := <-worker.Call
				err := p.Archivate(worker, f)
				if err != nil {
					slog.Error(err.Error())
				}
				p.WG.Done()
			}
		}()
	}
	worker.Call <- fm
	return nil
}

func (p *Process) Archivate(worker *ArchiveWorker, f *FileMeta) error {
	// Create a new zip file header
	header, err := zip.FileInfoHeader(f.FileInfo)
	if err != nil {
		return err
	}
	header.Method = zip.Deflate
	// Set the name of the file within the zip archive
	header.Name = f.FilePathInArchive
	// Create a new entry in the zip archive
	entry, err := worker.ArchiveWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	fmt.Println(header)
	// Copy the file content to the zip entry
	n, err := io.Copy(entry, f.FileReader)
	if err != nil {
		return err
	}
	slog.Debug("written", "bytes", n)
	_, err = entry.Write([]byte("\n"))
	if err != nil {
		return err
	}

	return nil
}
