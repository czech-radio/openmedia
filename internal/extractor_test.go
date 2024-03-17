package internal

import (
	"fmt"
	"testing"
)

var CSVdelim = "\t"

func TestExtractorInit(t *testing.T) {
	var ex Extractor
	ex.Init(nil, EXTproduction, CSVdelim)
	fmt.Println("partsPos", ex.CSVrowPartsPositions)
	fmt.Println("partsFieldsPos", ex.CSVrowPartsFieldsPositions)
	fmt.Println("fieldsHeader", ex.CSVrowHeader)
	fmt.Println("extractores", ex.OMobjExtractors)
}

func TestPrintTable(t *testing.T) {
}
