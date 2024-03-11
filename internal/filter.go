package internal

import (
	"fmt"
	"log/slog"
)

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

type ResultsCompounded map[string]*FilterResults

type FilterResults struct {
	FilesProcessed int
	FilesSuccess   int
	FilesFailure   int
	FilesCount     int
	ErrorsCount    int
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
