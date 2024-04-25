package extract

import (
	"fmt"
	"github/czech-radio/openmedia/internal/helper"
	"log/slog"
	"os"
	"strings"
)

// CSVtableBuildHeader
func (e *Extractor) CSVtableBuildHeader(delim string) {
	var builderInternal strings.Builder
	var builderExternal strings.Builder
	for _, partPrefixCode := range e.RowPartsPositions {
		rowPart := e.RowPartsFieldsPositions[partPrefixCode]
		for _, field := range rowPart {
			internal := HeaderColumnInternalCreate(
				partPrefixCode, field.FieldID, e.CSVdelim,
			)
			fmt.Fprint(&builderInternal, internal)
			external := HeaderColumnExternalCreate(
				partPrefixCode, field.FieldID, e.CSVdelim,
			)
			fmt.Fprint(&builderExternal, external)
		}
	}
	e.CSVheaderInternal = builderInternal.String()
	e.CSVheaderExternal = builderExternal.String()
}

// CSVtableSaveToFile
func (table *TableXML) CSVtableSaveToFile(dstFilePath string) (int, error) {
	outputFile, err := os.OpenFile(dstFilePath, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return 0, err
	}
	defer outputFile.Close()
	return outputFile.WriteString(table.CSVwriterLocal.String())
}

// CSVtableBuild
func (table *TableXML) CSVtableBuild(
	header bool, delim string, rowsIndexes ...[]int) {
	if len(rowsIndexes) > 1 {
		slog.Error("not implemented multiple indexes' slices")
	}
	var count int
	// TODO: simplify for one loop
	// Print specified rows
	if len(rowsIndexes) == 1 {
		for _, index := range rowsIndexes[0] {
			table.Rows[index].CSVrowBuild(
				table.CSVwriterLocal, table.RowPartsPositions,
				table.RowPartsFieldsPositions,
				delim,
			)
			count++
		}
		slog.Debug("lines casted to CSV", "count", count)
		return
	}

	// Print whole table
	for _, row := range table.Rows {
		row.CSVrowBuild(
			table.CSVwriterLocal, table.RowPartsPositions,
			table.RowPartsFieldsPositions,
			delim,
		)
		count++
	}
	slog.Debug("lines casted to CSV", "count", count)
}

// CSVrowBuild
func (row RowParts) CSVrowBuild(
	builder *strings.Builder,
	partsPos RowPartsPositions,
	partsFieldsPos RowPartsFieldsPositions,
	delim string,
) {
	for _, pos := range partsPos {
		fieldsPos := partsFieldsPos[pos]
		part := row[pos]
		part.CSVrowPartBuild(builder, fieldsPos, delim)
	}
	fmt.Fprintf(builder, "%s", "\n")
}

// CSVrowPartBuild
func (part RowPart) CSVrowPartBuild(
	builder *strings.Builder,
	fieldsPosition RowPartFieldsPositions,
	delim string,
) {
	specValEmpty := RowFieldValueCodeMap[RowFieldValueEmptyString]
	for _, pos := range fieldsPosition {
		field, ok := part[pos.FieldID]
		if !ok {
			value := specValEmpty
			fmt.Fprintf(builder, "%s%s", value, delim)
			continue
		}
		value := strings.TrimSpace(field.Value)
		value = TransformEmptyString(value)
		value = helper.EscapeCSVdelim(value)
		fmt.Fprintf(builder, "%s%s", value, delim)
	}
}

// CSVheaderPrintDirect
func (e *Extractor) CSVheaderPrintDirect(internal, external bool) {
	if internal {
		fmt.Println(strings.Join(e.HeaderInternal, e.CSVdelim))
	}
	if external {
		fmt.Println(strings.Join(e.HeaderExternal, e.CSVdelim))
	}
}

// CSVtablePrintDirect
func (e *Extractor) CSVtablePrintDirect(
	internalHeader, externalHeader bool,
	delim string, rowsIndexes ...[]int) {
	var sb strings.Builder
	e.CSVheaderPrintDirect(internalHeader, externalHeader)

	if len(rowsIndexes) > 1 {
		slog.Error("not implemented for multiple indexes' slices")
	}

	var count int
	// Print specified rows
	if len(rowsIndexes) == 1 {
		for _, index := range rowsIndexes[0] {
			e.TableXML.Rows[index].CSVrowBuild(
				&sb,
				e.RowPartsPositions,
				e.RowPartsFieldsPositions,
				delim,
			)
			count++
		}
		fmt.Print(sb.String())
		slog.Debug("lines printed", "count", count)
		return
	}

	// Print whole table
	for _, row := range e.TableXML.Rows {
		row.CSVrowBuild(
			&sb,
			e.RowPartsPositions,
			e.RowPartsFieldsPositions,
			delim,
		)
		count++
	}
	fmt.Print(sb.String())
	slog.Warn("lines printed", "count", count)
}
