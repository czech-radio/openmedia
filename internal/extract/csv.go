package extract

import (
	"fmt"
	"github/czech-radio/openmedia/internal/helper"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/antchfx/xmlquery"
)

// Table fields positions
type FieldPosition struct {
	RowPart   string
	FieldID   string
	FieldName string
}
type CSVrowPartFieldsPositions []FieldPosition                            // Field
type CSVrowPartsPositions []string                                        // Part
type CSVrowPartsPositionsInternal []RowPartCode                           // Part
type CSVrowPartsPositionsExternal []RowPartCode                           // Part
type CSVrowPartsFieldsPositions map[RowPartCode]CSVrowPartFieldsPositions // Row: partname vs partFieldsPositions

// Table fields values
type CSVrowField struct {
	FieldID   string
	FieldName string // Currently not needed here (will consume more memory). Alternative construct general list of fieldPrefix:fieldIDs vs FieldName
	Value     string
}
type CSVrowPart map[string]CSVrowField // FieldID:CSVrowField
type CSVrow map[RowPartCode]CSVrowPart // Whole CSV line PartPrefix:RowPart
type CSVrowNode struct {
	Node *xmlquery.Node
	CSVrow
}

// CSVheaders
type CSVheaders map[CSVheaderCodeName]CSVrow

// UniqueRow
type UniqueRow struct {
	Count         int
	TablePosition int
}

// CSVtable
type CSVtable struct {
	Rows              []*CSVrowNode
	Headers           []string
	CSVrowsFiltered   []int
	RowPartsPositions []RowPartCode
	CSVrowPartsFieldsPositions

	UniqueRowsOrder []int
	UniqueRows      map[string]int

	CSVwriterLocal *strings.Builder
	DstFilePath    string
	SrcFilePath    string
}

// CSVtables
type CSVtables struct {
	TablesPositions map[int]string       // pos:fileName
	Tables          map[string]*CSVtable // fileName:CSVtable
	CSVwriterGlobal *strings.Builder
	DstFileGlobal   *os.File
}

// BuildHeaderNameExternal
func BuildHeaderNameExternal(
	prefixCode RowPartCode, fieldName string) string {
	prefix := PartsPrefixMapProduction[prefixCode]
	if prefix.External == "" {
		return fieldName
	}
	return fmt.Sprintf("%s_%s", fieldName, prefix.External)
}

// CreateTablesHeader
func (e *Extractor) CreateTablesHeader(delim string) {
	var builderInternal strings.Builder
	var builderExternal strings.Builder
	for _, partPrefixCode := range e.CSVrowPartsPositionsInternal {
		prefix := PartsPrefixMapProduction[partPrefixCode]

		rowPart := e.CSVrowPartsFieldsPositions[partPrefixCode]
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

// CastTablesToCSV
func (e *Extractor) CastTablesToCSV(
	header bool, delim string, separateTables bool) {
	if !separateTables {
		e.CSVwriterGlobal = new(strings.Builder)
	}
	// Header write global
	if !separateTables && header {
		e.CSVwriterGlobal.WriteString(e.CSVheaderInternal)
		e.CSVwriterGlobal.WriteString(e.CSVheaderExternal)
	}

	for i, table := range e.Tables {
		if separateTables && header {
			table.CastTableToCSV(header, delim)
		}
		table.CastTableToCSV(header, delim)
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

// ConstructDstFilePath
func ConstructDstFilePath(srcPath string) string {
	srcDir, name := filepath.Split(srcPath)
	return filepath.Join(srcDir, "export"+name+".csv")
}

// SaveTableToFile
func (table *CSVtable) SaveTableToFile(dstFilePath string) (int, error) {
	outputFile, err := os.OpenFile(dstFilePath, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return 0, err
	}
	defer outputFile.Close()
	return outputFile.WriteString(table.CSVwriterLocal.String())
}

// CastTableToCSV
func (table *CSVtable) CastTableToCSV(
	header bool, delim string, rowsIndexes ...[]int) {
	// Print header
	if header {
		table.CastHeaderToCSV(table.Headers...)
	}

	if len(rowsIndexes) > 1 {
		slog.Error("not implemented multiple indexes' slices")
	}
	var count int
	// TODO: simplify for one loop
	// Print specified rows
	if len(rowsIndexes) == 1 {
		for _, index := range rowsIndexes[0] {
			table.Rows[index].CastToCSV(
				table.CSVwriterLocal, table.RowPartsPositions,
				table.CSVrowPartsFieldsPositions,
				delim,
			)
			count++
		}
		slog.Debug("lines casted to CSV", "count", count)
		return
	}

	// Print whole table
	for _, row := range table.Rows {
		row.CastToCSV(
			table.CSVwriterLocal, table.RowPartsPositions,
			table.CSVrowPartsFieldsPositions,
			delim,
		)
		count++
	}
	slog.Debug("lines casted to CSV", "count", count)
}

// CastHeaderToCSV
func (table *CSVtable) CastHeaderToCSV(headers ...string) {
	for _, header := range headers {
		if header != "" {
			table.CSVwriterLocal.WriteString(header)
		}
	}
}

// CastToCSV
func (row CSVrow) CastToCSV(
	builder *strings.Builder,
	partsPos CSVrowPartsPositionsInternal,
	partsFieldsPos CSVrowPartsFieldsPositions,
	delim string,
) {
	for _, pos := range partsPos {
		fieldsPos := partsFieldsPos[pos]
		part := row[pos]
		part.CastToCSV(builder, fieldsPos, delim)
	}
	fmt.Fprintf(builder, "%s", "\n")
}

// CastToCSV
func (part CSVrowPart) CastToCSV(
	builder *strings.Builder,
	fieldsPosition CSVrowPartFieldsPositions,
	delim string,
) {
	// specValNotPossible := CSVspecialValues[CSVspecialValueChildNotFound]
	specValEmpty := CSVspecialValues[CSVspecialValueEmptyString]
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

// PrintTableRowsToCSV
func (e *Extractor) PrintTableRowsToCSV(
	internalHeader, externalHeader bool,
	delim string, rowsIndexes ...[]int) {
	var sb strings.Builder
	e.CSVheaderPrint(internalHeader, externalHeader)
	if len(rowsIndexes) > 1 {
		slog.Error("not implemented multiple indexes' slices")
	}
	var count int
	// Print specified rows
	if len(rowsIndexes) == 1 {
		for _, index := range rowsIndexes[0] {
			e.CSVtable.Rows[index].CastToCSV(
				&sb, e.CSVrowPartsPositionsInternal,
				e.CSVrowPartsFieldsPositions,
				delim,
			)
			count++
		}
		fmt.Print(sb.String())
		slog.Debug("lines printed", "count", count)
		return
	}

	// Print whole table
	for _, row := range e.CSVtable.Rows {
		row.CastToCSV(
			&sb, e.CSVrowPartsPositionsInternal,
			e.CSVrowPartsFieldsPositions,
			delim,
		)
		count++
	}
	fmt.Print(sb.String())
	slog.Warn("lines printed", "count", count)
}
