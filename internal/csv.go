package internal

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/antchfx/xmlquery"
)

// Table fields positions
type FieldPosition struct {
	FieldPrefix string
	FieldID     string
	FieldName   string
}
type CSVrowPartFieldsPositions []FieldPosition                       // Field
type CSVrowPartsPositions []string                                   // Part
type CSVrowPartsPositionsInternal []string                           // Part
type CSVrowPartsPositionsExternal []string                           // Part
type CSVrowPartsFieldsPositions map[string]CSVrowPartFieldsPositions // Row: partname vs partFieldsPositions

// Table fields values
type CSVrowField struct {
	FieldID   string
	FieldName string // Maybe not needed here. Must construct general list of fieldPrefix:fieldIDs vs FieldName
	Value     string
}
type CSVrowPart map[string]CSVrowField // FieldID:CSVrowField
type CSVrow map[string]CSVrowPart      // Whole CSV line PartPrefix:RowPart
type CSVrowNode struct {
	Node *xmlquery.Node
	CSVrow
}
type CSVtable []*CSVrowNode
type CSVtables map[string]*CSVtable

func (e *Extractor) CSVheaderCreate(delim string) {
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
				slog.Warn("fieldname for given attribute not defined", "attribute", attr)
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

	// Print specified rows
	if len(rowsIndexes) == 1 {
		for _, index := range rowsIndexes[0] {
			e.CSVtable[index].PrintRowToCSV(
				&sb, e.CSVrowPartsPositionsInternal,
				e.CSVrowPartsFieldsPositions,
				delim,
			)
		}
		fmt.Print(sb.String())
		return
	}

	// Print whole table
	for _, row := range e.CSVtable {
		row.PrintRowToCSV(
			&sb, e.CSVrowPartsPositionsInternal,
			e.CSVrowPartsFieldsPositions,
			delim,
		)
	}
	fmt.Print(sb.String())
}

func (row CSVrow) PrintRowToCSV(
	builder *strings.Builder,
	partsPos []string,
	partsFieldsPos CSVrowPartsFieldsPositions,
	delim string,
) {
	for _, pos := range partsPos {
		fieldsPos := partsFieldsPos[pos]
		part := row[pos]
		part.PrintPartToCSV(builder, fieldsPos, delim)
	}
	fmt.Fprintf(builder, "%s", "\n")
}

func (part CSVrowPart) PrintPartToCSV(
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
