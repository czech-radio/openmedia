package internal

import (
	"fmt"
	"strings"

	"github.com/antchfx/xmlquery"
)

// Table fields positions
type FieldPosition struct {
	FieldPrefix string
	FieldID     string
	FieldName   string
}

// type CSVrowPartFieldsPositions []string // Field
type CSVrowPartFieldsPositions []FieldPosition // Field
// type CSVrowPartFieldsPositions map[string]FieldPosition // Field
type CSVrowPartsPositions []string // Part
// type CSVrowPartsFieldsPositions map[string][]string // Row: partname vs partFieldsPositions
type CSVrowPartsFieldsPositions map[string]CSVrowPartFieldsPositions // Row: partname vs partFieldsPositions

type CSVrowField struct {
	// FieldPosition int
	FieldID   string
	FieldName string // Maybe not needed here. Must construct general list of fieldPrefix:fieldIDs vs FieldName
	Value     string
}

// Table fields values
type CSVrowPart map[string]CSVrowField // FieldID:CSVrowField

// type CSVrowPartNode struct {
// Node *xmlquery.Node
// CSVrowPart
// }

type CSVrow map[string]CSVrowPart // Whole CSV line PartPrefix:RowPart

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
	for _, pos := range partsPos {
		fieldsPos := partsFieldsPos[pos]
		part := row[pos]
		part.PrintToCSV(builder, fieldsPos, delim)
	}
	fmt.Fprintf(builder, "%s", "\n")
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
