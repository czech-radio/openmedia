package extcases

import (
	"github/czech-radio/openmedia-archive/internal"
	"testing"
)

func BenchmarkArchiveFileExtractByXMLquery(b *testing.B) {
	filePath := "/home/jk/CRO/CRO_BASE/openmedia-archive_backup/Archive/control/control_UTF16_RD_13-17_Plus_Tuesday_W01_2024_01_02.xml"
	af := internal.ArchiveFile{}
	err := af.Init(internal.WorkerTypeRundownXMLutf16le, filePath)
	if err != nil {
		b.Error(err.Error())
	}
	for i := 0; i < b.N; i++ {
		err = af.ExtractByXMLquery(EXTproduction)
		if err != nil {
			b.Error(err.Error())
		}
	}
}
