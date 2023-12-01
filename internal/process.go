package internal

import (
	"bytes"
	"fmt"
	"io"
	"os"
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
	Weeks        int
	Files        int
	SizeOriginal int
	SizeBackup   int
	SizeMinified int
	Errors       []error
}

func (p *Process) ErrorHandle(errMain error, errorsPartial ...error) ControlFlowAction {
	if errMain == nil {
		return Continue
	}
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
processFolder:
	for i, file := range validateResult.FilesValid {
		fmt.Printf("%d: %+v\n", i, file)

		// Open file
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
		date, err := om.RundownDate()
		switch p.ErrorHandle(err) {
		case Break:
			break processFolder
		case Skip:
			continue processFolder
		}

		// Send input data to backup archive

		// Transform output data
		pr = PipeRundownMarshal(om)
		pr = PipeRundownHeaderAdd(pr)
		PipePrint(pr)

		// Send output data to minify archive
	}
	return nil
}

func (p *Process) File() {
}
