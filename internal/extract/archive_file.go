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

type ArchiveFile struct {
	Reader      *bytes.Reader
	Tables      map[ar.WorkerTypeCode]CSVtable
	Encoding    string
	RundownType string
	FilePath    string
	BaseNode    *xmlquery.Node
	Extractor
}

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

	openMedia, err := ar.XmlFindBaseOpenMediaNode(af.Reader)
	if err != nil {
		return err
	}
	af.BaseNode = openMedia
	return nil
}

type ArchivePackageFile struct {
	Reader *zip.File
	Tables map[ar.WorkerTypeCode]CSVtable
}

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

func (apf *ArchivePackageFile) ExtractByXMLquery(
	enc helper.FileEncodingNumber, q *ArchiveFolderQuery) error {
	var err error
	// Extract file from zip
	dataReader, err := ar.ZipXmlFileDecodeData(apf.Reader, enc)
	if err != nil {
		return err
	}
	// Parse base xml node
	openMedia, err := ar.XmlFindBaseOpenMediaNode(dataReader)
	if err != nil {
		return err
	}
	// Extract specfied object fields
	var extractor Extractor
	extractor.Init(openMedia, q.Extractors, CSVdelim)

	// Add default row
	// extractor.CSVtable.Rows[0].CSVrow[FieldPrefix_ComputedRID] = make(CSVrowPart)
	// extractor.CSVtable.Rows[0].CSVrow[FieldPrefix_ComputedRID]["FileName"] = field

	err = extractor.ExtractTable(apf.Reader.Name)
	if err != nil {
		return err
	}
	if q.ComputeUniqueRows {
		extractor.UniqueRows()
	}
	extractor.Transform(q.Transformer)
	extractor.FiltersRun(q.FilterColumns)
	extractor.PrintTableRowsToCSV(q.PrintHeader, q.CSVdelim)
	return nil
}

func (apf *ArchivePackageFile) ExtractByParser(
	enc helper.FileEncodingNumber, q *ArchiveFolderQuery) error {
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
