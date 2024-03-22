package internal

import (
	"fmt"
	"testing"
)

func TestExtractorInit(t *testing.T) {
	var ex Extractor
	ex.Init(nil, EXTproduction, CSVdelim)
	fmt.Println("prefixesInternal", ex.CSVrowPartsPositionsInternal)
	fmt.Println("prefixesExternal", ex.CSVrowPartsPositionsExternal)
	fmt.Println("partsFieldsPos", ex.CSVrowPartsFieldsPositions)
	fmt.Println("fieldsHeaderInternal", ex.CSVheaderInternal)
	fmt.Println("fieldsHeaderExternal", ex.CSVheaderExternal)
	fmt.Println("extractores", ex.OMextractors)
}
