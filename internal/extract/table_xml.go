package extract

import (
	"os"
	"path/filepath"
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

// type RowPartsPositionsExternal []RowPartCode                        // Alternative
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
	Rows []*RowNode
	// RowsFiltered []int

	UniqueRowsOrder []int
	UniqueRows      map[string]int

	SrcFilePath          string
	DstFilePath          string
	CSVtableWriterLocal  *strings.Builder
	CSVheaderWriterLocal *strings.Builder
}

// TablesXML
type TablesXML struct {
	TablesPositions map[int]string       // pos:fileName
	Tables          map[string]*TableXML // fileName:CSVtable
	CSVwriterGlobal *strings.Builder
	DstFileGlobal   *os.File
}

// ConstructDestinationFilePath
func ConstructDestinationFilePath(srcPath string) string {
	srcDir, name := filepath.Split(srcPath)
	return filepath.Join(srcDir, "export"+name+".csv")
}

func (t *TableXML) NullXMLnode() {
	for i := range t.Rows {
		t.Rows[i].Node = nil
	}
}
