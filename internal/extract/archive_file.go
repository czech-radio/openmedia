package extract

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"fmt"
	ar "github/czech-radio/openmedia/internal/archive"
	"github/czech-radio/openmedia/internal/helper"
	"os"
	"strings"

	"github.com/antchfx/xmlquery"
)

// ArchiveFile
type ArchiveFile struct {
	Reader      *bytes.Reader
	Tables      map[ar.WorkerTypeCode]TableXML
	Encoding    string
	RundownType string
	FilePath    string
	BaseNode    *xmlquery.Node
	Extractor
}

// Init
func (af *ArchiveFile) Init(wt ar.WorkerTypeCode, filePath string) error {
	wr := ar.WorkerTypeMap[wt]
	instructions := strings.Split(wr, "_")
	af.RundownType = instructions[0]
	af.Encoding = instructions[2]
	af.FilePath = filePath
	fileHandle, err := os.Open(af.FilePath)
	if err != nil {
		return err
	}

	// switch file encoding type
	enc := ar.InferEncoding(wt)
	data, err := helper.HandleFileEncoding(enc, fileHandle)
	if err != nil {
		return err
	}
	breader, err := ar.HandleXMLfileHeader(enc, data)
	if err != nil {
		return err
	}
	af.Reader = breader

	baseNode, err := XMLgetOpenmediaBaseNode(af.Reader)
	if err != nil {
		return err
	}
	af.BaseNode = baseNode
	return nil
}

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
	err := extractor.ExtractTable(af.FilePath)
	if err != nil {
		return err
	}
	af.Extractor = extractor
	return nil
}

// ExtractByXMLquery
func (apf *ArchivePackageFile) ExtractByXMLquery(
	enc helper.FileEncodingCode, q *ArchiveFolderQuery) (*Extractor, error) {
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
	// extractor.TransformEmptyRowPart()
	// extractor.Transform(q.Transformer)
	// extractor.FiltersRun(q.FilterColumns)
	return &extractor, nil
}

// ExtractByParser
func (apf *ArchivePackageFile) ExtractByParser(
	enc helper.FileEncodingCode, q *ArchiveFolderQuery) error {
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
