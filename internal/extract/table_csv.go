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
	for _, partPrefixCode := range e.RowPartsPositions {
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
