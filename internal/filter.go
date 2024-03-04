package internal

import (
	"fmt"
	"log/slog"
)

type FilterOptions struct {
	SourceDirectory        string
	DestinationDirectory   string
	RecurseSourceDirectory bool
	InvalidFileContinue    bool
	FileType               string
	DateFrom               string
	DateTo                 string
	OutputType             string // csv contacts,unique contact fields,
	CSVdelim               string
	CSVheader              bool
}

type Filter struct {
	Options               FilterOptions
	Results               FilterResults
	Errors                []error
	ObjectHeader          []string
	ObjectsAttrValues     []ObjectAttributes
	HeaderFields          map[int]string
	HeaderFieldsIDsSorted []int
	HeaderFieldsIDsSubset map[int]bool
	Rows                  []Fields
	FieldsUniqueValues    map[int]UniqueValues // FiledID vs UniqueValues
	MaxUniqueCount        int
}

type FilterResults struct {
	FilesProcessed int
	FilesSuccess   int
	FilesFailure   int
	FilesCount     int
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
