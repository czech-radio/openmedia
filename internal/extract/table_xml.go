package extract

import (
	"os"
	"strings"

	"github.com/antchfx/xmlquery"
)

// RowPartFieldPosition represents/describe general field in specific column
type RowPartFieldPosition struct {
	RowPartName string
	FieldID     string
	FieldName   string
}

// RowPartFieldsPositions
type RowPartFieldsPositions []RowPartFieldPosition // Field

// RowPartsPositions
type RowPartsPositions []RowPartCode

// type RowPartsPositionsExternal []RowPartCode                        // Part
type RowPartsFieldsPositions map[RowPartCode]RowPartFieldsPositions // Row: partname vs partFieldsPositions

// RowField
type RowField struct {
	FieldID   string
	FieldName string // Currently not needed here (will consume more memory). Alternative construct general list of fieldPrefix:fieldIDs vs FieldName
	Value     string
}

// RowPart
type RowPart map[string]RowField // FieldID:CSVrowField

// RowParts
type RowParts map[RowPartCode]RowPart // Whole CSV line PartPrefix:RowPart

// RowNode
type RowNode struct {
	Node *xmlquery.Node
	RowParts
}

// UniqueRow
type UniqueRow struct {
	Count         int
	TablePosition int
}

// TableXML
type TableXML struct {
	Rows              []*RowNode
	Headers           []string
	CSVrowsFiltered   []int
	RowPartsPositions []RowPartCode
	RowPartsFieldsPositions

	UniqueRowsOrder []int
	UniqueRows      map[string]int

	CSVwriterLocal *strings.Builder
	DstFilePath    string
	SrcFilePath    string
}

// TablesXML
type TablesXML struct {
	TablesPositions map[int]string       // pos:fileName
	Tables          map[string]*TableXML // fileName:CSVtable
	CSVwriterGlobal *strings.Builder
	DstFileGlobal   *os.File
}
