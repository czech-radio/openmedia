package extract

import (
	"log/slog"
	"path/filepath"
	"strings"
)

func (e *Extractor) OutputBaseDataset(
	processName string, queryOpts *ArchiveFolderQuery) {
	e.TransformBase()
	if queryOpts.ExtractorsCode == ExtractorsProductionContacts {
		indxs := e.FilterPeculiarContacts()
		e.DeleteNonMatchingRows(indxs)
	}
	e.CSVtableBuild(false, false, queryOpts.CSVdelim, true)
	e.TableOutputs(queryOpts.OutputDirectory, queryOpts.OutputFileName,
		queryOpts.ExtractorsName, processName, true)
}

func (e *Extractor) OutputValidatedDataset(
	processName string, queryOpts *ArchiveFolderQuery) error {
	e.TransformBeforeValidation()
	e.ValidateAllColumns(queryOpts.ValidatorFileName)
	e.CSVtableBuild(false, false, queryOpts.CSVdelim, true)
	e.TableOutputs(queryOpts.OutputDirectory, queryOpts.OutputFileName,
		queryOpts.ExtractorsName, processName, true)

	// Write validation to validation log file
	validationLogfile := strings.Join(
		[]string{queryOpts.OutputFileName, processName, "log"}, "_")
	logFilePath := filepath.Join(
		queryOpts.OutputDirectory, validationLogfile+".csv")
	return e.ValidationLogWrite(logFilePath, queryOpts.CSVdelim, true)
}

func (e *Extractor) OutputFilteredDataset(
	processName string, queryOpts *ArchiveFolderQuery,
	filterOpts *FilterFile,
) error {
	if filterOpts.FilterFileName == "" {
		slog.Info("filter file not specified")
		return nil
	}
	e.TransformProduction()
	filters := make(FilterFileCodes)
	filters.AddFilters(
		FilterFileOposition,
		// e.FilterMatchPersonNameJoinedNoDiacritics,
		e.FilterMatchPersonNameSurnameNormalized,
		e.FilterMatchPersonIDandHighPolitics,
	)
	filters.AddFilters(
		FilterFileEuroElection,
		// e.FilterMatchPersonNameSurnameNormalized,
		e.FilterMatchPersonName,
		e.FilterMatchPersonAndParty,
	)
	err := filters.FiltersApply(filterOpts)
	if err != nil {
		return err
	}

	e.CSVtableBuild(false, false, queryOpts.CSVdelim, true)
	e.TableOutputs(queryOpts.OutputDirectory, queryOpts.OutputFileName,
		queryOpts.ExtractorsName, processName, true)
	return nil
}
