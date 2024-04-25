package extract

import (
	"os"
	"strings"

	"github.com/antchfx/xmlquery"
)

// Table fields positions
type FieldPosition struct {
	RowPartName string
	FieldID     string
	FieldName   string
}

type RowPartFieldsPositions []FieldPosition                         // Field
type RowPartsPositions []string                                     // Part
type RowPartsPositionsInternal []RowPartCode                        // Part
type RowPartsPositionsExternal []RowPartCode                        // Part
type RowPartsFieldsPositions map[RowPartCode]RowPartFieldsPositions // Row: partname vs partFieldsPositions

type RowField struct {
	FieldID   string
	FieldName string // Currently not needed here (will consume more memory). Alternative construct general list of fieldPrefix:fieldIDs vs FieldName
	Value     string
}

type RowPart map[string]RowField // FieldID:CSVrowField
type Row map[RowPartCode]RowPart // Whole CSV line PartPrefix:RowPart
type RowNode struct {
	Node *xmlquery.Node
	Row
}

// TableXMLheaders
type TableXMLheaders map[CSVheaderCodeName]Row

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
