package extract

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"fmt"
	ar "github/czech-radio/openmedia/internal/archive"

	"github.com/triopium/go_utils/pkg/helper"

	"github.com/antchfx/xmlquery"
)

// ArchiveFile
type ArchiveFile struct {
	SourceFilePath     string
	SourceFileEncoding string
	RundownType        string
	Reader             *bytes.Reader

	Tables   map[ar.WorkerTypeCode]TableXML
	BaseNode *xmlquery.Node
	Extractor
}

// Init
func (af *ArchiveFile) Init() error {
	// fileHandle, err := os.Open(af.SourceFilePath)
	// if err != nil {
	// 	return err
	// }

	// switch file encoding type
	// enc := ar.InferEncoding(wt)
	// data, err := helper.HandleFileEncoding(enc, fileHandle)
	// if err != nil {
	// 	return err
	// }
	// breader, err := ar.HandleXMLfileHeader(enc, data)
	// if err != nil {
	// 	return err
	// }
	// af.Reader = breader

	// baseNode, err := XMLgetOpenmediaBaseNode(af.Reader)
	// if err != nil {
	// 	return err
	// }
	// af.BaseNode = baseNode
	return nil
}

// InitOld
// func (af *ArchiveFile) InitOld(wt ar.WorkerTypeCode, filePath string) error {
// 	wr := ar.WorkerTypeMap[wt]
// 	instructions := strings.Split(wr, "_")
// 	af.RundownType = instructions[0]
// 	af.SourceFileEncoding = instructions[2]
// 	af.SourceFilePath = filePath
// 	fileHandle, err := os.Open(af.SourceFilePath)
// 	if err != nil {
// 		return err
// 	}

// 	// switch file encoding type
// 	enc := ar.InferEncoding(wt)
// 	data, err := helper.HandleFileEncoding(enc, fileHandle)
// 	if err != nil {
// 		return err
// 	}
// 	breader, err := ar.HandleXMLfileHeader(enc, data)
// 	if err != nil {
// 		return err
// 	}
// 	af.Reader = breader

// 	baseNode, err := XMLgetOpenmediaBaseNode(af.Reader)
// 	if err != nil {
// 		return err
// 	}
// 	af.BaseNode = baseNode
// 	return nil
// }

// ArchivePackageFile
type ArchivePackageFile struct {
	Reader *zip.File
	Tables map[ar.WorkerTypeCode]TableXML
}

// ExtractByXMLquery
func (af *ArchiveFile) ExtractByXMLquery(extrs OMextractors) error {
	// Extract specfied object fields
	var extractor Extractor
	extractor.Init(af.BaseNode, extrs, CSVdelim)
	err := extractor.ExtractTable(af.SourceFilePath)
	if err != nil {
		return err
	}
	af.Extractor = extractor
	return nil
}

// ExtractByXMLquery
func (apf *ArchivePackageFile) ExtractByXMLquery(
	enc helper.CharEncoding, q *ArchiveFolderQuery) (*Extractor, error) {
	var err error
	// Extract file from zip
	dataReader, err := ar.ZipXmlFileDecodeData(apf.Reader, enc)
	if err != nil {
		return nil, err
	}
	baseNode, err := XMLgetOpenmediaBaseNode(dataReader)
	if err != nil {
		return nil, err
	}
	// Extract specfied object fields
	var extractor Extractor
	extractor.Init(baseNode, q.Extractors, CSVdelim)

	err = extractor.ExtractTable(apf.Reader.Name)
	if err != nil {
		return nil, err
	}
	if q.ComputeUniqueRows {
		extractor.UniqueRows()
	}
	return &extractor, nil
}

// ExtractByParser
func (apf *ArchivePackageFile) ExtractByParser(
	enc helper.CharEncoding, q *ArchiveFolderQuery) error {
	dr, err := ar.ZipXmlFileDecodeData(apf.Reader, enc)
	if err != nil {
		return err
	}
	var OM ar.OPENMEDIA
	err = xml.NewDecoder(dr).Decode(&OM)
	if err != nil {
		return err
	}
	// var produkce CSVtable
	for _, i := range OM.OM_OBJECT.OM_RECORDS {
		// var row CSVrow
		// NOTE: REMAINING NOT IMPLEMENTED
		fmt.Println(i.OM_OBJECTS.OM_HEADER)
	}
	return nil
}
