package extract

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"fmt"
	ar "github/czech-radio/openmedia/internal/archive"
	"time"

	"github.com/triopium/go_utils/pkg/helper"

	"github.com/antchfx/xmlquery"
)

type ArchiveQueryCommon struct {
	ExtractorsCode    ExtractorsPresetCode
	FilterRadioNames  map[string]bool
	FilterDateFrom    time.Time
	FilterDateTo      time.Time
	DateRange         [2]time.Time
	FilterIsoWeeks    map[int]bool
	FilterMonths      map[int]bool
	FilterWeekDays    map[int]bool
	ValidatorFileName string

	AddRecordNumbers  bool
	ComputeUniqueRows bool
}

type ArchiveIO struct {
	SourceDirectory     string
	SourceDirectoryType ar.WorkerType
	SourceFilePath      string
	SourceCharEncoding  helper.CharEncoding
	OutputDirectory     string
	OutputFileName      string
	CSVdelim            string
}

type ArchiveFile struct {
	ArchiveQueryCommon
	ArchiveIO
	FilterFile

	Tables   map[ar.WorkerType]TableXML
	Reader   *bytes.Reader
	BaseNode *xmlquery.Node
	Extractor
}

// Init
func (af *ArchiveFile) Init() error {
	data, enc, err := helper.FileReadAllHandleEncoding(af.SourceFilePath)
	if err != nil {
		return fmt.Errorf("%w %q: ", err, af.SourceFilePath)
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
	Tables map[ar.WorkerType]TableXML
}

func (af *ArchiveFile) ExtractByXMLquery(extrs OMextractors) error {
	// Extract specfied object fields
	var extractor Extractor
	extractor.Init(af.BaseNode, extrs, CSVdelim)
	if af.AddRecordNumbers {
		extractor.AddRecordsColumns()
	}
	err := extractor.ExtractTable(af.SourceDirectory)
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
	extractor.Init(baseNode, q.Extractors, q.CSVdelim)

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
