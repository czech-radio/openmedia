package internal

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"strings"

	"github.com/antchfx/xmlquery"
)

type ArchiveFile struct {
	Reader      *bytes.Reader
	Tables      map[WorkerTypeCode]CSVtable
	Encoding    string
	RundownType string
	FilePath    string
	BaseNode    *xmlquery.Node
	Extractor
}

// switch enc {
// case UTF8:
// data, err = io.ReadAll(fileHandle)
// case UTF16le:
// utf8reader := enc_unicode.UTF16(enc_unicode.LittleEndian, enc_unicode.IgnoreBOM).NewDecoder().Reader(fileHandle)
// data, err = io.ReadAll(utf8reader)
// default:
// err = fmt.Errorf("unknown encoding")
// }

func (af *ArchiveFile) Init(wt WorkerTypeCode, filePath string) error {
	wr := WorkerTypeMap[wt]
	instructions := strings.Split(wr, "_")
	af.RundownType = instructions[0]
	af.Encoding = instructions[2]
	af.FilePath = filePath
	fileHandle, err := os.Open(af.FilePath)
	if err != nil {
		return err
	}

	// switch file encoding type
	enc := InferEncoding(wt)
	data, err := HandleFileEncoding(enc, fileHandle)
	if err != nil {
		return err
	}
	breader, err := HandleXMLfileHeader(enc, data)
	if err != nil {
		return err
	}
	// utf8reader := enc_unicode.UTF16(enc_unicode.LittleEndian, enc_unicode.IgnoreBOM).NewDecoder().Reader(fileHandle)
	// data, err := io.ReadAll(utf8reader)
	// if err != nil {
	// return err
	// }
	// bytesReader := bytes.NewReader(data)
	// bytesReader, err = XmlAmendUTF16header(bytesReader)
	// if err != nil {
	// return err
	// }
	af.Reader = breader

	openMedia, err := XmlFindBaseOpenMediaNode(af.Reader)
	if err != nil {
		return err
	}
	af.BaseNode = openMedia
	return nil
}

type ArchivePackageFile struct {
	Reader *zip.File
	Tables map[WorkerTypeCode]CSVtable
}

func (af *ArchiveFile) ExtractByXMLquery(extrs OMextractors) error {
	// Extract specfied object fields
	var extractor Extractor
	extractor.Init(af.BaseNode, extrs, CSVdelim)
	err := extractor.ExtractTable()
	if err != nil {
		return err
	}
	af.Extractor = extractor
	return nil
}

func (apf *ArchivePackageFile) ExtractByXMLquery(
	enc FileEncodingNumber, q *ArchiveFolderQuery) error {
	var err error
	// Extract file from zip
	dataReader, err := ZipXmlFileDecodeData(apf.Reader, enc)
	if err != nil {
		return err
	}
	// Parse base xml node
	openMedia, err := XmlFindBaseOpenMediaNode(dataReader)
	if err != nil {
		return err
	}
	// Extract specfied object fields
	var extractor Extractor
	extractor.Init(openMedia, q.Extractors, CSVdelim)
	err = extractor.ExtractTable()
	if err != nil {
		return err
	}
	if q.ComputeUniqueRows {
		extractor.UniqueRows()
	}
	extractor.TransformProduction()
	extractor.PrintTableRowsToCSV(q.PrintHeader, q.CSVdelim)
	return nil
}

func (apf *ArchivePackageFile) ExtractByParser(
	enc FileEncodingNumber, q *ArchiveFolderQuery) error {
	dr, err := ZipXmlFileDecodeData(apf.Reader, enc)
	if err != nil {
		return err
	}
	var OM OPENMEDIA
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
