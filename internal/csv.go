package internal

type FieldPosition struct {
	FieldPrefix string
	FieldID     string
	FieldName   string
}

// Table fields positions
// type CSVrowPartFieldsPositions []string // Field
type CSVrowPartFieldsPositions []FieldPosition // Field
// type CSVrowPartFieldsPositions map[string]FieldPosition // Field
type CSVrowPartsPositions []string // Part
// type CSVrowPartsFieldsPositions map[string][]string // Row: partname vs partFieldsPositions
type CSVrowPartsFieldsPositions map[string]CSVrowPartFieldsPositions // Row: partname vs partFieldsPositions
// type CSVrowsFieldIDheader []string                  // FieldID header
// type CSVrowsNameheader []string                     // FieldID header

// Table fields values
type CSVrowPart map[string]CSVrowField // ObjectPrefix:CSVrowField
type CSVrow map[string]CSVrowPart      // Whole CSV line
type CSVtable []*CSVrow
type CSVtables map[string]*CSVtable

func (table *CSVtable) PrintToCSV(delim string) {
	// for row := range table {
	// }
}

// func (fp CSVrowPartFieldsPositions) PrintHeader() {
// }

// func (row CSVrow) PrintToCSV(
// 	builder *strings.Builder,
// 	partsPos CSVrowPartFieldsPositions, partsFieldsPos CSVrowPartsFieldsPositions,
// 	delim string) string {
// 	for _, pos := range partsPos {
// 		// fmt.Println(pos)
// 		fieldsPos := partsFieldsPos[pos.FieldID]
// 		// fmt.Println("fieldsPos", fieldsPos)
// 		part := row[pos.FieldID]
// 		// fmt.Println("part", part)
// 		part.PrintToCSV(builder, fieldsPos, delim)
// 	}
// 	return builder.String()
// }

// func (part CSVrowPart) PrintToCSV(
// 	builder *strings.Builder, fieldsPosition CSVrowPartFieldsPositions, delim string,
// ) {
// 	var value string
// 	for _, pos := range fieldsPosition {
// 		field, ok := part[pos]
// 		if !ok {
// 			value = "NO_VALUE"
// 		} else {
// 			value = field.Value
// 		}
// 		fmt.Fprintf(builder, "%s%s", value, delim)
// 	}
// }
