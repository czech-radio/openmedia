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
}
