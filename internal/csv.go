package internal

import (
	"fmt"
	"log/slog"
	"regexp"
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

func (e *Extractor) PrintTableToCSV(header bool, delim string) {
	if header {
		e.CSVheaderPrint()
	}
	var sb strings.Builder
	for _, row := range e.CSVtable {
		row.PrintToCSV(
			&sb, e.CSVrowPartsPositionsInternal,
			e.CSVrowPartsFieldsPositions,
			delim,
		)
	}
	fmt.Println(sb.String())
}

func (row CSVrow) PrintToCSV(
	builder *strings.Builder,
	partsPos []string,
	partsFieldsPos CSVrowPartsFieldsPositions,
	delim string,
) {
	ok := FilterByHours(row)
	if !ok {
		return
	}
	for _, pos := range partsPos {
		fieldsPos := partsFieldsPos[pos]
		part := row[pos]
		part.PrintToCSV(builder, fieldsPos, delim)
	}
	fmt.Fprintf(builder, "%s", "\n")
}

var hoursRegex = regexp.MustCompile("^13:00-14:00")

func FilterByHours(row CSVrow) bool {
	pos := PartsPrefixMapProduction[FieldPrefix_HourlyHead].Internal
	part, ok := row[pos]
	if !ok {
		return true
	}
	hours, ok := part["8"]
	if !ok {
		return true
	}
	ok = hoursRegex.MatchString(hours.Value)
	return ok
}

func TransformEmptyString(input string) string {
	if input == "" {
		return "(NEUVEDENO)"
	}
	return input
}

func (part CSVrowPart) PrintToCSV(
	builder *strings.Builder, fieldsPosition CSVrowPartFieldsPositions, delim string,
) {
	for _, pos := range fieldsPosition {
		field, ok := part[pos.FieldID]
		if !ok {
			value := "(NELZE)"
			fmt.Fprintf(builder, "%s%s", value, delim)
			continue
		}

		value := strings.TrimSpace(field.Value)
		value = TransformEmptyString(value)
		value = EscapeCSVdelim(value)
		fmt.Fprintf(builder, "%s%s", value, delim)
	}
}
