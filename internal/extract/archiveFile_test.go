package extract

import (
	"fmt"
	ar "github/czech-radio/openmedia-archive/internal/archive"
	"testing"

	"github.com/antchfx/xmlquery"
)

func TestXMLqueryFile(t *testing.T) {
	filePath := "/home/jk/CRO/CRO_BASE/openmedia-archive_backup/Archive/control/control_UTF16_RD_13-17_Plus_Tuesday_W01_2024_01_02.xml"
	af := ArchiveFile{}
	err := af.Init(ar.WorkerTypeRundownXMLutf16le, filePath)
	if err != nil {
		t.Error(err.Error())
	}
	// query := "OM_OBJECT[@TemplateName='Radio Rundown']/OM_RECORD/OM_OBJECT[@TemplateName='Hourly Rundown']/OM_RECORD/OM_OBJECT[@TemplateName='Radio Story']"
	queries := map[string]string{
		"story":       "//OM_OBJECT[@TemplateName='Radio Story']",  // 130
		"audio":       "//OM_OBJECT[@TemplateName='Audioclip']",    // 84
		"contact":     "//OM_OBJECT[@TemplateName='Contact Item']", // 66
		"audo+contac": "//OM_OBJECT[@TemplateName='Audioclip' or @TemplateName='Contact Item']",
	}
	for i, q := range queries {
		subNodes := xmlquery.Find(af.BaseNode, q)
		fmt.Println(i, q, len(subNodes))
	}
}
