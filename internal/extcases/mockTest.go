package extcases

import (
	"fmt"
	"github/czech-radio/openmedia-archive/internal"
	"testing"
)

func TestExtractorInit(t *testing.T) {
	var ex internal.Extractor
	// ex.Init(nil, EXTproduction, CSVdelim)
	ex.Init(nil, EXTmock, "\t")
	fmt.Println("prefixesInternal", ex.CSVrowPartsPositionsInternal)
	fmt.Println("prefixesExternal", ex.CSVrowPartsPositionsExternal)
	fmt.Println("partsFieldsPos", ex.CSVrowPartsFieldsPositions)
	fmt.Println("fieldsHeaderInternal", ex.CSVheaderInternal)
	fmt.Println("fieldsHeaderExternal", ex.CSVheaderExternal)
	// fmt.Printf("extractores: %+v\n", ex.OMextractors)
	fmt.Printf("part codes: %+v\n", ex.CSVrowPartsPositionsExternal)
	// out, _ := json.MarshalIndent(ex, "", "\t")
	// fmt.Println("extractor:", string(out))
}

func TestArchiveFileExtractByXMLqueryFilter(t *testing.T) {
	filePath := "/home/jk/CRO/CRO_BASE/openmedia-archive_backup/Archive/control/control_UTF16_RD_13-17_Plus_Tuesday_W01_2024_01_02.xml"
	af := internal.ArchiveFile{}
	err := af.Init(internal.WorkerTypeRundownXMLutf16le, filePath)
	if err != nil {
		t.Error(err.Error())
	}
	// err = af.ExtractByXMLquery(EXTproduction)
	err = af.ExtractByXMLquery(EXTmock)
	if err != nil {
		t.Error(err.Error())
	}
	patern := "13:00-14:00"
	rowIdx := af.Extractor.FilterByPartAndFieldID(
		internal.FieldPrefix_HourlyHead, "8", patern,
	)
	af.Extractor.PrintTableRowsToCSV(true, "\t", rowIdx)
}

func TestArchiveFileExtractByXMLquery(t *testing.T) {
	filePath := "/home/jk/CRO/CRO_BASE/openmedia-archive_backup/Archive/control/control_UTF16_RD_13-17_Plus_Tuesday_W01_2024_01_02.xml"
	af := internal.ArchiveFile{}
	err := af.Init(internal.WorkerTypeRundownXMLutf16le, filePath)
	if err != nil {
		t.Error(err.Error())
	}
	// err = af.ExtractByXMLquery(EXTproduction)
	err = af.ExtractByXMLquery(EXTmock)
	if err != nil {
		t.Error(err.Error())
	}
	af.Extractor.PrintTableRowsToCSV(true, "\t")
}
