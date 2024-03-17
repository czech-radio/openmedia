package internal

import (
	"fmt"
	"strings"
)

// NEW STRUCTURE
// Table fields positions
type CSVrowPartFieldsPositions []string             // Field
type CSVrowPartsPositions []string                  // Part
type CSVrowPartsFieldsPositions map[string][]string // Row: partname vs partFieldsPositions

// Table fields values
type CSVrowPart map[string]CSVrowField // ObjectPrefix:CSVrowField
type CSVrow map[string]CSVrowPart      // Whole CSV line
type CSVtable []*CSVrow
type CSVtables map[string]*CSVtable

func (table *CSVtable) PrintToCSV(delim string) {
	// for row := range table {
	// }
}

func (row CSVrow) PrintToCSV(
	builder *strings.Builder,
	partsPos []string, partsFieldsPos CSVrowPartsFieldsPositions,
	delim string) string {
	for _, pos := range partsPos {
		// fmt.Println(pos)
		fieldsPos := partsFieldsPos[pos]
		// fmt.Println("fieldsPos", fieldsPos)
		part := row[pos]
		// fmt.Println("part", part)
		part.PrintToCSV(builder, fieldsPos, delim)
	}
	return builder.String()
}

func (part CSVrowPart) PrintToCSV(
	builder *strings.Builder, fieldsPosition CSVrowPartFieldsPositions, delim string,
) {
	var value string
	for _, pos := range fieldsPosition {
		field, ok := part[pos]
		if !ok {
			value = "NO_VALUE"
		} else {
			value = field.Value
		}
		fmt.Fprintf(builder, "%s%s", value, delim)
	}
}
