package internal

import (
	"fmt"
	"testing"
)

func TestExtractorInit(t *testing.T) {
	var ex Extractor
	ex.Init(EXTproduction, CSVdelim)
	fmt.Println("partsPos", ex.CSVrowPartsPositions)
	fmt.Println("partsFieldsPos", ex.CSVrowPartsFieldsPositions)
	fmt.Println("fieldsHeader", ex.CSVrowHeader)
}
