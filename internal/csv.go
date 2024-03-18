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

type CSVrow2 struct {
	CurrentNode  *xmlquery.Node
	RowNodePath  string
	FieldsPrefix string
	RowParts     map[string]CSVrowPart
}

type CSVtableOrig []*CSVrow
type CSVtable []*CSVrowNode
type CSVtables map[string]*CSVtable

func (e *Extractor) PrintRows() error {
	for i, row := range e.Rows {
		fmt.Println(i, row)
	}
	return nil
}

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
	// var builder strings.Builder
	for i, row := range e.CSVtable {
		fmt.Println(i, row)
	}
	// for i, row := range e.Rows {
	// fmt.Println(i, row.NodePath)
	// row.PrintToCSV()
	// }
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
}

func (part CSVrowPart) PrintToCSV(
	builder *strings.Builder, fieldsPosition CSVrowPartFieldsPositions, delim string,
) {
	var value string
	for _, pos := range fieldsPosition {
		field, ok := part[pos.FieldID]
		if !ok {
			value = "NO_VALUE"
		} else {
			value = field.Value
		}
		fmt.Fprintf(builder, "%s%s", value, delim)
	}
}
