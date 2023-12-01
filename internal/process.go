package internal

type ControlFlowAction int

const (
	Continue ControlFlowAction = iota
	Skip     ControlFlowAction = iota
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
	if errMain != nil {
		p.Results.Errors = append(p.Results.Errors, errMain)
		p.Results.Errors = append(p.Results.Errors, errorsPartial...)
		if p.Options.InvalidFileContinue {
			return Skip
		} else {
			return Break
		}
	}
	return Continue
}

func (p *Process) Folder() error {
	validateResult, err := ValidateFilenamesInDirectory(p.Options.SourceDirectory)
	if p.ErrorHandle(err, validateResult.Errors...) == Break {
		return err
	}
	return nil
}
