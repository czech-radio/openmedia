package internal

import (
	"fmt"
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
type CSVrowPartFieldsPositions []FieldPosition     // Field
type CSVrowPartsPositions []string                 // Part
type CSVrowPartsPositionsInternal []PartPrefixCode // Part
type CSVrowPartsPositionsExternal []PartPrefixCode // Part
// type CSVrowPartsFieldsPositions map[string]CSVrowPartFieldsPositions         // Row: partname vs partFieldsPositions
type CSVrowPartsFieldsPositions map[PartPrefixCode]CSVrowPartFieldsPositions // Row: partname vs partFieldsPositions

// Table fields values
type CSVrowField struct {
	FieldID   string
	FieldName string // Maybe not needed here. Must construct general list of fieldPrefix:fieldIDs vs FieldName
	Value     string
}
type CSVrowPart map[string]CSVrowField    // FieldID:CSVrowField
type CSVrow map[PartPrefixCode]CSVrowPart // Whole CSV line PartPrefix:RowPart
type CSVrowNode struct {
	Node *xmlquery.Node
	CSVrow
}

type CSVheaders map[CSVheaderCodeName]CSVrow

type CSVtable struct {
	Rows              []*CSVrowNode
	Headers           []string
	CSVrowsFiltered   []int
	RowPartsPositions []PartPrefixCode
	CSVrowPartsFieldsPositions

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

func (e *Extractor) CreateTablesHeader(delim string) {
	var internalBuilder strings.Builder
	var externalBuilder strings.Builder
	for _, extr := range e.OMextractors {
		prefix := PartsPrefixMapProduction[extr.PartPrefixCode]
		for _, attr := range extr.ObjectAttrsNames {
			fmt.Fprintf(
				&internalBuilder, "%s_%s%s",
				prefix.Internal, attr, delim,
			)
			attrName, ok := FieldsIDsNamesProduction[attr]
			if !ok {
				slog.Warn(
					"fieldname for given attribute not defined",
					"attribute", attr,
				)
			}
			fmt.Fprintf(
				&externalBuilder, "%s_%s%s",
				prefix.External, attrName, delim,
			)
		}
		for _, fieldID := range extr.FieldIDs {
			fmt.Fprintf(
				&internalBuilder, "%s_%s%s",
				prefix.Internal, fieldID, delim,
			)
			fieldName, ok := FieldsIDsNamesProduction[fieldID]
			if !ok {
				slog.Warn("fieldname for given fieldID not defined", "filedID", fieldID)
			}
			fmt.Fprintf(
				&externalBuilder, "%s_%s%s",
				prefix.External, fieldName, delim,
			)
		}
	}
	e.CSVheaderInternal = internalBuilder.String()
	e.CSVheaderExternal = externalBuilder.String()
}

func (e *Extractor) CreateTablesHeaderB(delim string) {
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
			fmt.Fprintf(
				&builderExternal, "%s_%s%s",
				prefix.External, field.FieldID, delim,
			)
		}
	}
	fmt.Println(builderInternal.String())
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
		value = EscapeCSVdelim(value)
		fmt.Fprintf(builder, "%s%s", value, delim)
	}
}
