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

type Process struct {
	Options ProcessOptions
	Results ProcessResults
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
	Year            int
	Month           int
	Day             int
	Week            int
	Weekday         time.Weekday
	FileInfo        os.FileInfo
	FileReader      io.Reader
	ArchiveNameMini string
	ArchiveNameOrig string
	RundownNameNew  string
}

func (f *FileMeta) Parse(date time.Time, fileInfo os.FileInfo, archiveType string) {
	year, week := date.ISOWeek()
	f.Year = year
	f.Month = int(date.Month())
	f.Day = date.Day()
	f.Week = week
	f.Weekday = date.Weekday()
	f.ArchiveNameMini = fmt.Sprintf("%d_W%02d_MINIFIED.%s", f.Year, f.Week, archiveType)
	f.ArchiveNameOrig = fmt.Sprintf("%d_W%02d_ORIGINAL.%s", f.Year, f.Week, archiveType)
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
processFolder:
	for _, file := range validateResult.FilesValid {
		// Open file
		p.Results.FilesProcessed++
		p.InfoLog()
		fileHandle, err := os.Open(file)
		switch p.ErrorHandle(err, validateResult.Errors...) {
		case Break:
			break processFolder
		case Skip:
			continue processFolder
		}
		defer fileHandle.Close()

		// Read file data
		data, err := io.ReadAll(fileHandle)
		switch p.ErrorHandle(err, validateResult.Errors...) {
		case Break:
			break processFolder
		case Skip:
			continue processFolder
		}

		// Transform input data
		dataReader := bytes.NewReader(data)
		pr := PipeUTF16leToUTF8(dataReader)
		pr = PipeRundownHeaderAmmend(pr)

		// Unmarshal (Parse and validate)
		om, err := PipeRundownUnmarshal(pr) // treat EOF error
		switch p.ErrorHandle(err) {
		case Break:
			break processFolder
		case Skip:
			continue processFolder
		}

		// Infer rundown date from OM_HEADER field
		// _, err := om.RundownDate()
		date, err := om.RundownDate()
		switch p.ErrorHandle(err) {
		case Break:
			break processFolder
		case Skip:
			continue processFolder
		}

		// Send input data to backup archive
		// Get file info
		fileInfo, err := fileHandle.Stat()
		switch p.ErrorHandle(err, validateResult.Errors...) {
		case Break:
			break processFolder
		case Skip:
			continue processFolder
		}

		fm := new(FileMeta)
		fm.Parse(date, fileInfo, p.Options.ArchiveType)
		fmt.Printf("%+v\n", fm)

		// Transform output data
		pr = PipeRundownMarshal(om)
		pr = PipeRundownHeaderAdd(pr)
		// PipePrint(pr)
		PipeConsume(pr)

		// Send output data to minify archive
	}
	return nil
}

func (p *Process) GetMinifyChannel(fm *FileMeta) {
}

// func (p *Process) File() {
// }

// fileInfo, err := fileHandle.Stat()
// data []byte
// OM *OPENMEDIA
// date time.Time
// io.Reader
