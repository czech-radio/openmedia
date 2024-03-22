package internal

import (
	"fmt"
	"testing"

	"github.com/antchfx/xmlquery"
)

func TestXMLqueryFile(t *testing.T) {
	filePath := "/home/jk/CRO/CRO_BASE/openmedia-archive_backup/Archive/control/control_UTF16_RD_13-17_Plus_Tuesday_W01_2024_01_02.xml"
	af := ArchiveFile{}
	err := af.Init(WorkerTypeRundownXMLutf16le, filePath)
	if err != nil {
		t.Error(err.Error())
	}
	query := "OM_OBJECT[@TemplateName='Radio Rundown']/OM_RECORD/OM_OBJECT[@TemplateName='Hourly Rundown']/OM_RECORD/OM_OBJECT[@TemplateName='Radio Story']"
	// pat := "*<OM_RECORD>"
	// pat := "Radio Rundown"
	// query := "OM_OBJECT[@TemplateName='Radio Rundown']/OM_RECORD/OM_OBJECT[@TemplateName='Hourly Rundown']/OM_RECORD/OM_OBJECT[@TemplateName='Radio Story']"
	// 47
	// query := "OM_OBJECT[@TemplateName='Radio Rundown']/OM_RECORD/OM_OBJECT[@TemplateName='Hourly Rundown']/OM_RECORD//OM_OBJECT[@TemplateName='Contact Item']"
	// query := "OM_OBJECT[@TemplateName='Radio Rundown']/OM_RECORD/OM_OBJECT[@TemplateName='Hourly Rundown']/OM_RECORD/OM_OBJECT[@TemplateName='Sub Rundown']/OM_RECORD/OM_OBJECT[@TemplateName='Radio Story']"
	// 83
	// query := "OM_OBJECT[@TemplateName='Radio Rundown']/OM_RECORD/OM_OBJECT[@TemplateName='Hourly Rundown']/OM_RECORD/OM_OBJECT[@TemplateName='Sub Rundown']/OM_RECORD//OM_OBJECT[@TemplateName='Radio Story']"
	// pat := "*Radio Story"
	// query, err := QueryObject(pat)
	// fmt.Println(query)
	// if err != nil {
	// t.Error(err.Error())
	// }
	subNodes := xmlquery.Find(af.BaseNode, query)
	fmt.Println(len(subNodes))
	subNodes[0].SelectAttr("TemplateName")
}

func TestArchiveFileExtractByXMLquery(t *testing.T) {
	filePath := "/home/jk/CRO/CRO_BASE/openmedia-archive_backup/Archive/control/control_UTF16_RD_13-17_Plus_Tuesday_W01_2024_01_02.xml"
	af := ArchiveFile{}
	err := af.Init(WorkerTypeRundownXMLutf16le, filePath)
	if err != nil {
		t.Error(err.Error())
	}
	err = af.ExtractByXMLquery(EXTproduction)
	if err != nil {
		t.Error(err.Error())
	}
	patern := "13:00-14:00"
	rowIdx := af.Extractor.FilterByPartAndFieldID(
		FieldPrefix_HourlyHead, "8", patern,
	)
	af.Extractor.PrintTableRowsToCSV(true, "\t", rowIdx)
}
