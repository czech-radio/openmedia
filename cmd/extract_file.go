package cmd

import (
	ar "github/czech-radio/openmedia/internal/archive"
	"github/czech-radio/openmedia/internal/extract"
	"github/czech-radio/openmedia/internal/helper"
)

type ConfigExtractFile struct {
	SourceFile      string `cmd:"source_file; i; ; input file"`
	DestinationFile string `cmd:"destination_file; o; ; otput file"`
	CSVdelim        string `cmd:"csv_delim; csvd; \t; csv field delimiter"`
}

func RunExtractFile(rootCfg *ConfigRoot, filterCfg *ConfigExtractFile) {
	// filePath := "/home/jk/CRO/CRO_BASE/openmedia_backup/Archive/control/control_UTF16_RD_13-17_Plus_Tuesday_W01_2024_01_02.xml"
	filePath := "/home/jk/CRO/CRO_BASE/openmedia_backup/Archive/control2/RD_18-24_Radiožurnál_Friday_W09_2024_03_01_utf16le.xml"
	// filePath := "/home/jk/CRO/CRO_BASE/openmedia_backup/Archive/control/control_UTF8_RD_13-17_Plus_Tuesday_W01_2024_01_02.xml"
	af := extract.ArchiveFile{}
	err := af.Init(
		ar.WorkerTypeRundownXMLutf16le, filePath)
	// internal.WorkerTypeRundownXMLutf8, filePath)
	if err != nil {
		helper.Errors.ExitWithCode(err)
	}
	// err = af.ExtractByXMLquery(internal.EXTtest)
	err = af.ExtractByXMLquery(extract.EXTproduction)
	// err = af.ExtractByXMLquery(internal.EXTeuroVolby)
	if err != nil {
		helper.Errors.ExitWithCode(err)
	}
	af.Extractor.TransformProduction()
	af.Extractor.CSVtableBuild(false, true, "\t", true)
	// af.Extractor.TransformEurovolby()
	// af.Extractor.TransformTest()
	// af.Extractor.CSVtablePrintDirect()
}
