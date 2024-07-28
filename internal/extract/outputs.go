package extract

import (
	"log/slog"
	"path/filepath"
	"strings"

	"github.com/triopium/go_utils/pkg/helper"
)

func (e *Extractor) OutputAll(
	qc *ArchiveQueryCommon, qio *ArchiveIO, qf *FilterFile,
) {
	processName := "base"
	e.OutputBaseDatasetNew(
		processName, qc, qio)

	processName += "_validated"
	err := e.OutputValidatedDatasetNew(
		processName, qc, qio)
	if err != nil {
		helper.Errors.ExitWithCode(err)
	}

	processName += "_filtered"
	err = e.OutputFilteredDatasetNew(
		processName, qc, qio, qf)
	if err != nil {
		helper.Errors.ExitWithCode(err)
	}
}

func (e *Extractor) OutputBaseDatasetNew(
	processName string, qc *ArchiveQueryCommon, qio *ArchiveIO) {
	e.TransformBase()
	if qc.AddRecordNumbers {
		e.ComputeRecordIDs(true)
	}
	if qc.ExtractorsCode == ExtractorsProductionContacts {
		indxs := e.FilterPeculiarContacts()
		e.DeleteNonMatchingRows(indxs)
	}
	e.CSVtableBuild(false, false, qio.CSVdelim, true)
	e.TableOutputs(qio.OutputDirectory, qio.OutputFileName,
		string(qc.ExtractorsCode), processName, true)
}

func (e *Extractor) OutputValidatedDatasetNew(
	processName string, qc *ArchiveQueryCommon, qio *ArchiveIO) error {
	if qc.ValidatorFileName == "" {
		slog.Info("validation_warning", "msg", "validation receipe file not specified")
		return nil
	}
	e.TransformBeforeValidation()
	e.ValidateAllColumns(qc.ValidatorFileName)
	e.CSVtableBuild(false, false, qio.CSVdelim, true)
	e.TableOutputs(
		qio.OutputDirectory, qio.OutputFileName,
		string(qc.ExtractorsCode), processName, true)

	// Write validation to validation log file
	validationLogfile := strings.Join(
		[]string{qio.OutputFileName, processName, "log"}, "_")
	logFilePath := filepath.Join(
		qio.OutputDirectory, validationLogfile+".csv")
	return e.ValidationLogWrite(logFilePath, qio.CSVdelim, true)
}

func (e *Extractor) OutputFilteredDatasetNew(
	processName string, qc *ArchiveQueryCommon, qio *ArchiveIO,
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

	e.CSVtableBuild(false, false, qio.CSVdelim, true)
	e.TableOutputs(qio.OutputDirectory, qio.OutputFileName,
		string(qc.ExtractorsCode), processName, true)
	return nil
}
