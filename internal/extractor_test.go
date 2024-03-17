package internal

import (
	"fmt"
	"testing"
)

func TestExtractorInit(t *testing.T) {
	var ex Extractor
	ex.Init(EXTproduction)
	fmt.Println(ex.CSVrowPartsPositions)
	fmt.Println(ex.CSVrowPartsFieldsPositions)
	// fmt.Println(ex.CSVrowPartsFieldsHeader)
}

func TestPrintHeader(t *testing.T) {
	var ex Extractor
	ex.Init(EXTproduction)
	ex.CreateHeader(CSVdelim)
	fmt.Println(ex.CSVrowHeader)
	// fmt.Println(ex.CSVrowPartsPositions)
	// fmt.Println(ex.CSVrowPartsFieldsPositions)
	// fmt.Println(ex.CSVrowPartsFieldsHeader)
}
