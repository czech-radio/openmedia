package extract

import (
	"fmt"
	ar "github/czech-radio/openmedia/internal/archive"
	"testing"

	"github.com/antchfx/xmlquery"
)

func TestXMLqueryFile(t *testing.T) {
	testSubdir := "rundowns_mock"
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t, testSubdir)
	tp := testerConfig.TempSourcePathGeter(testSubdir)
	filePath := tp("RD_05-09_Dvojka_2_1605299_20200304234745.xml")
	af := ArchiveFile{}
	err := af.Init(ar.WorkerTypeRundownXMLutf16le, filePath)
	if err != nil {
		t.Error(err.Error())
	}

	type pathCount struct {
		Path  string
		Count int
	}

	queries := map[string]pathCount{
		"story":       {"//OM_OBJECT[@TemplateName='Radio Story']", 120}, // 130
		"audio":       {"//OM_OBJECT[@TemplateName='Audioclip']", 44},    // 84
		"contact":     {"//OM_OBJECT[@TemplateName='Contact Item']", 7},  // 66
		"audo+contac": {"//OM_OBJECT[@TemplateName='Audioclip' or @TemplateName='Contact Item']", 51},
	}

	for i, q := range queries {
		subNodes := xmlquery.Find(af.BaseNode, q.Path)
		fmt.Println(i, q, len(subNodes))
		if len(subNodes) != q.Count {
			t.Errorf(
				"conut of found subnodes is different than expected: found %d expected %d, path %s", len(subNodes), q.Count, q.Path,
			)
		}
	}
}
