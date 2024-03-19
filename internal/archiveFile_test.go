package internal

import (
	"fmt"
	"testing"

	"github.com/antchfx/xmlquery"
)

func TestQueryXMLfile(t *testing.T) {
	filePath := "/home/jk/CRO/CRO_BASE/openmedia-archive_backup/Archive/control/control_UTF16_RD_13-17_Plus_Tuesday_W01_2024_01_02.xml"
	af := ArchiveFile{}
	err := af.Init(WorkerTypeRundownXMLutf16le, filePath)
	if err != nil {
		t.Error(err.Error())
	}
	query, err := QueryObject("*<OM_RECORD>")
	fmt.Println(query)
	if err != nil {
		t.Error(err.Error())
	}
	subNodes := xmlquery.Find(af.BaseNode, query)
	fmt.Println(len(subNodes))
}

func TestArchiveFileExtractByXMLquery(t *testing.T) {
	filePath := "/home/jk/CRO/CRO_BASE/openmedia-archive_backup/Archive/control/control_UTF16_RD_13-17_Plus_Tuesday_W01_2024_01_02.xml"
	af := ArchiveFile{}
	err := af.Init(WorkerTypeRundownXMLutf16le, filePath)
	if err != nil {
		t.Error(err.Error())
	}
	err = af.ExtractByXMLquery(EXTproduction)
	// err = af.ExtractByXMLquery(EXTproductionRECandHED)
	if err != nil {
		t.Error(err.Error())
	}
}
