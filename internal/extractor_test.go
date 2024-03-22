package internal

import (
	"fmt"
	"testing"
)

var CSVdelim = "\t"

func TestExtractorInit(t *testing.T) {
	var ex Extractor
	ex.Init(nil, EXTproduction, CSVdelim)
	// fmt.Println("prefixesInternal", ex.CSVrowPartsPositionsInternal)
	// fmt.Println("prefixesExternal", ex.CSVrowPartsPositionsExternal)
	// fmt.Println("partsFieldsPos", ex.CSVrowPartsFieldsPositions)
	fmt.Println("fieldsHeader", ex.CSVheaderInternal)
	// fmt.Println("extractores", ex.OMextractors)
}

func TestPrintTable(t *testing.T) {
}
