package extract

import (
	"fmt"
	"github/czech-radio/openmedia-archive/internal/helper"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/antchfx/xmlquery"
)

// Table fields positions
type FieldPosition struct {
	FieldPrefix string
	FieldID     string
	FieldName   string
}
type CSVrowPartFieldsPositions []FieldPosition                               // Field
type CSVrowPartsPositions []string                                           // Part
type CSVrowPartsPositionsInternal []PartPrefixCode                           // Part
type CSVrowPartsPositionsExternal []PartPrefixCode                           // Part
type CSVrowPartsFieldsPositions map[PartPrefixCode]CSVrowPartFieldsPositions // Row: partname vs partFieldsPositions

// Table fields values
type CSVrowField struct {
	FieldID   string
	FieldName string // Currently not needed here (will consume more memory). Alternative construct general list of fieldPrefix:fieldIDs vs FieldName
	Value     string
}
type CSVrowPart map[string]CSVrowField    // FieldID:CSVrowField
type CSVrow map[PartPrefixCode]CSVrowPart // Whole CSV line PartPrefix:RowPart
type CSVrowNode struct {
	Node *xmlquery.Node
	CSVrow
}

type CSVheaders map[CSVheaderCodeName]CSVrow

type UniqueRow struct {
	Count         int
	TablePosition int
}

type CSVtable struct {
	Rows              []*CSVrowNode
	Headers           []string
	CSVrowsFiltered   []int
	RowPartsPositions []PartPrefixCode
	CSVrowPartsFieldsPositions

	UniqueRowsOrder []int
	UniqueRows      map[string]int

	CSVwriterLocal *strings.Builder
	DstFilePath    string
	SrcFilePath    string
}

type CSVtables struct {
	TablesPositions map[int]string       // pos:fileName
	Tables          map[string]*CSVtable // fileName:CSVtable
	CSVwriterGlobal *strings.Builder
	DstFileGlobal   *os.File
}

func BuildHeaderNameExternal(
	prefixCode PartPrefixCode, fieldName string) string {
	prefix := PartsPrefixMapProduction[prefixCode]
	if prefix.External == "" {
		return fieldName
	}
	return fmt.Sprintf("%s_%s", fieldName, prefix.External)
}

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

func ConstructDstFilePath(srcPath string) string {
	srcDir, name := filepath.Split(srcPath)
	return filepath.Join(srcDir, "export"+name+".csv")
}

func (table *CSVtable) SaveTableToFile(dstFilePath string) (int, error) {
	outputFile, err := os.OpenFile(dstFilePath, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return 0, err
	}
	defer outputFile.Close()
	return outputFile.WriteString(table.CSVwriterLocal.String())
}

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

func (table *CSVtable) CastHeaderToCSV(headers ...string) {
	for _, header := range headers {
		if header != "" {
			table.CSVwriterLocal.WriteString(header)
		}
	}
}

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

func (part CSVrowPart) CastToCSV(
	builder *strings.Builder,
	fieldsPosition CSVrowPartFieldsPositions,
	delim string,
) {
	specVal := CSVspecialValues[CSVspecialValueChildNotFound]
	for _, pos := range fieldsPosition {
		field, ok := part[pos.FieldID]
		if !ok {
			value := specVal
			fmt.Fprintf(builder, "%s%s", value, delim)
			continue
		}

		value := strings.TrimSpace(field.Value)
		value = TransformEmptyString(value)
		value = helper.EscapeCSVdelim(value)
		fmt.Fprintf(builder, "%s%s", value, delim)
	}
}

func (e *Extractor) CSVheaderPrint() {
	fmt.Println(e.CSVheaderInternal)
	fmt.Println(e.CSVheaderExternal)
}

func (e *Extractor) PrintTableRowsToCSV(
	header bool, delim string, rowsIndexes ...[]int) {
	var sb strings.Builder
	// Print header
	if header {
		e.CSVheaderPrint()
	}

	if len(rowsIndexes) > 1 {
		slog.Error("not implemented multiple indexes' slices")
	}
	var count int
	// Print specified rows
	if len(rowsIndexes) == 1 {
		for _, index := range rowsIndexes[0] {
			// e.CSVtable[index].PrintRowToCSV(
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
