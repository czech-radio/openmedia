package internal

import (
	"fmt"
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
	var builder strings.Builder
	for _, i := range e.CSVrowPartsPositions {
		pfp := e.CSVrowPartsFieldsPositions[i]
		for _, j := range pfp {
			fmt.Fprintf(&builder, "%s_%s%s", j.FieldPrefix, j.FieldID, delim)
		}
	}
	e.CSVrowHeader = builder.String()
}

func (e *Extractor) CSVheaderPrint() {
	fmt.Println(e.CSVrowHeader)
}

func (e *Extractor) PrintTableToCSV(header bool, delim string) {
	if header {
		fmt.Println(e.CSVrowHeader)
	}
	var sb strings.Builder
	for _, row := range e.CSVtable {
		row.PrintToCSV(
			&sb, e.CSVrowPartsPositions,
			e.CSVrowPartsFieldsPositions,
			delim,
		)
	}
	fmt.Println(sb.String())
}

func (row CSVrow) PrintToCSV(
	builder *strings.Builder,
	partsPos CSVrowPartsPositions,
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
	pos := "HourlyR-HED"
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

func (part CSVrowPart) PrintToCSV(
	builder *strings.Builder, fieldsPosition CSVrowPartFieldsPositions, delim string,
) {
	var value string
	for _, pos := range fieldsPosition {
		field, ok := part[pos.FieldID]
		if !ok {
			value = "FID NOT FOUND"
		} else {
			value = EscapeCSVdelim(field.Value)
		}
		fmt.Fprintf(builder, "%s%s", value, delim)
	}
}
