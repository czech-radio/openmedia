package internal

import "testing"

func TestArchiveFileExtract(t *testing.T) {
	filePath := "/home/jk/CRO/CRO_BASE/openmedia-archive_backup/Archive/control/control_UTF16_RD_13-17_Plus_Tuesday_W01_2024_01_02.xml"
	af := ArchiveFile{}
	err := af.Init(WorkerTypeRundownXMLutf16le, filePath)
	if err != nil {
		t.Error(err.Error())
	}
	err = af.ExtractByXMLquery()
	if err != nil {
		t.Error(err.Error())
	}
}
