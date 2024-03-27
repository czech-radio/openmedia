package internal

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestExtractorInit(t *testing.T) {
	var ex Extractor
	ex.Init(nil, EXTproductionAudioContacts, CSVdelim)
	fmt.Println("prefixesInternal", ex.CSVrowPartsPositionsInternal)
	fmt.Println("prefixesExternal", ex.CSVrowPartsPositionsExternal)
	fmt.Println("partsFieldsPos", ex.CSVrowPartsFieldsPositions)
	fmt.Println("fieldsHeaderInternal", ex.CSVheaderInternal)
	fmt.Println("fieldsHeaderExternal", ex.CSVheaderExternal)
	fmt.Printf("extractores: %+v\n", ex.OMextractors)
	out, _ := json.MarshalIndent(ex, "", "\t")
	fmt.Println("extractor:", string(out))
}
