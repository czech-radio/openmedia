package extract

import (
	"log/slog"
	"path/filepath"
	"strings"

	"github.com/triopium/go_utils/pkg/helper"
)

func (e *Extractor) OutputRowsSpecific(
	processName string, qc *ArchiveQueryCommon, qio *ArchiveIO, indexes []int) {
	e.CSVtableBuild(false, false, qio.CSVdelim, true, indexes)
	e.TableOutputs(qio.OutputDirectory, qio.OutputFileName,
		string(qc.ExtractorsCode), processName, true)
}

func (e *Extractor) OutputRowsAll(
	processName string, qc *ArchiveQueryCommon, qio *ArchiveIO) {
	e.CSVtableBuild(false, false, qio.CSVdelim, true)
	e.TableOutputs(qio.OutputDirectory, qio.OutputFileName,
		string(qc.ExtractorsCode), processName, true)
}

func (e *Extractor) ExportToDB(qc *ArchiveQueryCommon, qio *ArchiveIO, qf *FilterFile) {
	e.AddColumn(RowPartCode_ComputedRID, "FileName")
	if qc.AddRecordNumbers {
		e.ColumnCompute_RecordIDs(true)
		e.ColumnCompute_FilenameWithRecordIDs(true)
	}
	indxs := e.FilterStoryPartRecordsDuds()
	e.DeleteNonMatchingRows(indxs)
	e.RowPartOmit(RowPartCode_StoryRec)
	processName := "db_export"
	e.ExportBase(processName, qc, qio)
}

func (e *Extractor) ExportAll(
	qc *ArchiveQueryCommon, qio *ArchiveIO, qf *FilterFile,
) {
	// Pretreat: Add/Remove columns
	e.AddColumn(RowPartCode_ComputedRID, "FileName")
	if qc.AddRecordNumbers {
		e.ColumnCompute_RecordIDs(false)
		e.ColumnCompute_FilenameWithRecordIDs(false)
	}

	processName := "RIDodd_DIFF"
	indxs := e.FilterStoryPartRecordsDuds()
	e.OmitRecordIDsColumn()
	e.OutputRowsSpecific(processName, qc, qio, ReverseIndexes(indxs))
	e.DeleteNonMatchingRows(indxs)
	e.RowPartOmit(RowPartCode_StoryRec)

	e.TransformBase()
	// e.RowPartOmit(RowPartCode_StoryRec)
	processName = "base_NCparts"
	e.OutputRowsAll(processName, qc, qio)

	// Removes duplicate (NC) story parts empty
	processName = "base_NCparts_DIFF"
	indxs = e.FilterStoryPartsEmptyDupes()
	e.OutputRowsSpecific(processName, qc, qio, ReverseIndexes(indxs))
	e.DeleteNonMatchingRows(indxs)

	processName = "base"
	e.OutputRowsAll(processName, qc, qio)
	// e.ExportBase(processName, qc, qio)

	processName += "_validated"
	err := e.ExportValidated(
		processName, qc, qio)
	if err != nil {
		helper.Errors.ExitWithCode(err)
	}

	processName += "_filtered"
	err = e.ExportFiltered(
		processName, qc, qio, qf)
	if err != nil {
		helper.Errors.ExitWithCode(err)
	}
}

func (e *Extractor) ExportBase(

	processName string, qc *ArchiveQueryCommon, qio *ArchiveIO) {
	e.TransformBase()
	if qc.ExtractorsCode == ExtractorsProductionContacts {
		indxs := e.FilterPeculiarContacts()
		e.DeleteNonMatchingRows(indxs)
	}
	e.CSVtableBuild(false, false, qio.CSVdelim, true)
	e.TableOutputs(qio.OutputDirectory, qio.OutputFileName,
		string(qc.ExtractorsCode), processName, true)
}

func (e *Extractor) ExportValidated(
	processName string, qc *ArchiveQueryCommon, qio *ArchiveIO) error {
	if qc.ValidatorFileName == "" {
		slog.Info("validation_warning", "msg", "validation receipe file not specified")
		return nil
	}
	e.TransformBeforeValidation()
	// e.ValidateAllColumns(qc.ValidatorFileName)
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

func (e *Extractor) ExportFiltered(
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
