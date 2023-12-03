package internal

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"os"
	"time"
)

type ControlFlowAction int

const (
	Continue ControlFlowAction = iota
	Skip
	Break
)

type Worker chan *FileMeta

type Process struct {
	Options ProcessOptions
	Results ProcessResults
	Workers map[string]Worker
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
	Year           int
	Month          int
	Day            int
	Week           int
	Weekday        time.Weekday
	Archives       map[string]string
	RundownNameNew string
	FileInfo       os.FileInfo
	FileReader     io.Reader
}

func (f *FileMeta) Parse(date time.Time, fileInfo os.FileInfo, archiveType string) {
	year, week := date.ISOWeek()
	f.Year = year
	f.Month = int(date.Month())
	f.Day = date.Day()
	f.Week = week
	f.Weekday = date.Weekday()
	f.Archives = make(map[string]string)
	// archive names
	name := "MINIFIED"
	f.Archives[name] = fmt.Sprintf("%d_W%02d_%s.%s", f.Year, f.Week, name, archiveType)
	name = "ORIGINAL"
	f.Archives[name] = fmt.Sprintf("%d_W%02d_%s.%s", f.Year, f.Week, name, archiveType)
	f.RundownNameNew = fmt.Sprintf("RD_%s_W%02d_%04d_%02d_%02d", f.Weekday, f.Week, f.Year, f.Month, f.Day)
	f.FileInfo = fileInfo
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
	p.Workers = make(map[string]Worker)
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
	return nil
}

func (p *Process) File(file string) error {
	// Open file
	p.Results.FilesProcessed++
	// p.InfoLog()
	fileHandle, err := os.Open(file)
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
	fm := new(FileMeta)
	fm.Parse(date, fileInfo, p.Options.ArchiveType)
	fmt.Printf("%+v\n", fm)
	p.CallWorker(fm, "MINIFIED")

	// Transform output data
	pr = PipeRundownMarshal(om)
	pr = PipeRundownHeaderAdd(pr)
	// PipePrint(pr)
	PipeConsume(pr)
	p.CallWorker(fm, "ORIGINAL")
	return nil
	// Send output data to minify archive
}

func (p *Process) CallWorker(fm *FileMeta, workerType string) {
	// var workerMini Worker
	workerName := fm.Archives[workerType]
	worker, ok := p.Workers[workerName]
	fmt.Printf("%+v,%v\n", worker, ok)
	if !ok {
		worker = make(Worker)
		p.Workers[workerName] = worker
		go func() {
			for {
				f := <-worker
				fmt.Println(f)
			}
		}()
	}
	worker <- fm
}
