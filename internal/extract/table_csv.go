package extract

import (
	"fmt"
	"github/czech-radio/openmedia/internal/helper"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

// CSVtableBuildHeader
func (e *Extractor) CSVtableBuildHeader(delim string) {
	var builderInternal strings.Builder
	var builderExternal strings.Builder
	for _, partPrefixCode := range e.RowPartsPositionsInternal {
		prefix := RowPartsCodeMapProduction[partPrefixCode]

		rowPart := e.RowPartsFieldsPositions[partPrefixCode]
		for _, field := range rowPart {
			fmt.Fprintf(
				&builderInternal, "%s_%s%s",
				prefix.Internal, field.FieldID, delim,
			)
			fieldName := FieldsIDsNamesProduction[field.FieldID]
			externalName := BuildHeaderNameExternal(
				partPrefixCode, fieldName)
			fmt.Fprintf(
				&builderExternal, "%s%s", externalName, delim,
			)
		}
	}
	e.CSVheaderInternal = builderInternal.String()
	e.CSVheaderExternal = builderExternal.String()
}

// CSVtablesBuild
func (e *Extractor) CSVtablesBuild(
	header bool, delim string, separateTables bool) {
	if !separateTables {
		e.CSVwriterGlobal = new(strings.Builder)
	}
	// Header write global
	if !separateTables && header {
		e.CSVwriterGlobal.WriteString(e.CSVheaderInternal)
		e.CSVwriterGlobal.WriteString(e.CSVheaderExternal)
	}

	for i, table := range e.TablesXML.Tables {
		if separateTables && header {
			table.CSVtableBuild(header, delim)
		}
		table.CSVtableBuild(header, delim)
		slog.Debug(
			"casting table to CSV", "current", i, "count", len(e.Tables))
	}
}

// SaveTablesToFile
func (e *Extractor) SaveTablesToFile(
	separateTables bool, dstFilePath string) error {
	if !separateTables {
		outputFile, err := os.OpenFile(dstFilePath, os.O_RDWR|os.O_CREATE, 0600)
		if err != nil {
			return err
		}
		defer outputFile.Close()
		n, err := outputFile.WriteString(e.CSVwriterGlobal.String())
		if err != nil {
			return err
		}
		slog.Debug("writen bytes to one file", "filename", dstFilePath, "bytesCount", n)
		return nil
	}
	current := 0
	bytesCountCumulative := 0
	for i, table := range e.Tables {
		current++
		dstFilePath := ConstructDstFilePath(table.SrcFilePath)
		outputFile, err := os.OpenFile(dstFilePath, os.O_RDWR|os.O_CREATE, 0600)
		if err != nil {
			return err
		}
		n, err := outputFile.WriteString(table.CSVwriterLocal.String())
		if err != nil {
			return err
		}
		sequnece := fmt.Sprintf("%d/%d", current, len(e.Tables))
		slog.Debug(
			"writen bytes to file in sequence", "sequence", sequnece,
			"filename", dstFilePath,
			"srcFile", i, "bytesCount", n,
		)

	}
	slog.Debug("finished writing files in sequence",
		"bytesCount", bytesCountCumulative,
		"filesCount", len(e.Tables))
	return nil
}

// BuildHeaderNameExternal
func BuildHeaderNameExternal(
	rowPartCode RowPartCode, fieldName string) string {
	prefix := RowPartsCodeMapProduction[rowPartCode]
	if prefix.External == "" {
		return fieldName
	}
	return fmt.Sprintf("%s_%s", fieldName, prefix.External)
}

// ConstructDstFilePath
func ConstructDstFilePath(srcPath string) string {
	srcDir, name := filepath.Split(srcPath)
	return filepath.Join(srcDir, "export"+name+".csv")
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
	// Print header
	if header {
		table.CSVheadersBuild(table.Headers...)
	}

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

// CSVheadersBuild
func (table *TableXML) CSVheadersBuild(headers ...string) {
	for _, header := range headers {
		if header != "" {
			table.CSVwriterLocal.WriteString(header)
		}
	}
}

// CSVrowBuild
func (row Row) CSVrowBuild(
	builder *strings.Builder,
	partsPos RowPartsPositionsInternal,
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
	// specValNotPossible := CSVspecialValues[CSVspecialValueChildNotFound]
	specValEmpty := RowFieldValueCodeMap[RowFieldValueEmptyString]
	for _, pos := range fieldsPosition {
		// if part == nil {
		// fmt.Fprintf(builder, "%s%s", specValNotPossible, delim)
		// continue
		// }
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

// CSVheaderPrint
func (e *Extractor) CSVheaderPrint(internal, external bool) {
	if internal {
		fmt.Println(e.CSVheaderInternal)
	}
	if external {
		fmt.Println(e.CSVheaderExternal)
	}
}

// CSVtablePrint
func (e *Extractor) CSVtablePrint(
	internalHeader, externalHeader bool,
	delim string, rowsIndexes ...[]int) {
	var sb strings.Builder
	e.CSVheaderPrint(internalHeader, externalHeader)

	if len(rowsIndexes) > 1 {
		slog.Error("not implemented for multiple indexes' slices")
	}

	var count int
	// Print specified rows
	if len(rowsIndexes) == 1 {
		for _, index := range rowsIndexes[0] {
			e.TableXML.Rows[index].CSVrowBuild(
				&sb,
				e.RowPartsPositionsInternal,
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
			e.RowPartsPositionsInternal,
			e.RowPartsFieldsPositions,
			delim,
		)
		count++
	}
	fmt.Print(sb.String())
	slog.Warn("lines printed", "count", count)
}
