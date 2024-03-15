package internal

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
)

type ObjectAttributes = map[string]string
type Fields = map[int]string       // FieldID/FieldName vs value
type UniqueValues = map[string]int // value vs count

type CSVrowField struct {
	FieldPosition int
	FieldID       string
	Value         string
}

type CSVrow = []CSVrowField

type CSVheaderField = map[string]string
type CSVheader []CSVheaderField
type CSVtable struct {
	Header []CSVheaderField
	Rows   []CSVrow
}

type ArchivePackageFile struct {
	Reader *zip.File
	Tables map[WorkerTypeCode]CSVtable
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
		fmt.Println(i.OM_OBJECTS.OM_HEADER)
	}
	return nil
}

type OMobjExtractor struct {
	OmObject        string
	Path            string
	FieldsPath      string
	FieldIDs        []string
	FieldIDsMap     map[string]bool
	ReplacePrevious bool
}

type OMobjExtractors struct {
	Extractors []OMobjExtractor
}

func JoinObjectPath(oldpath, newpath string) string {
	return oldpath + "/" + newpath
}

func (omo *OMobjExtractor) MapFields() {
	omo.FieldIDsMap = make(map[string]bool, len(omo.FieldIDs))
	for _, id := range omo.FieldIDs {
		omo.FieldIDsMap[id] = true
	}
}

// "//OM_OBJECT[@TemplateName='%s']/%s/*", ext.OM_type, ext.Path,
// FieldsPath: "/OM_HEADER/OM_FIELD",
// FieldsPath: "/OM_RECORD/OM_FIELD",
var CSVproduction = []OMobjExtractor{
	{
		OmObject:        "Radio Rundown",
		Path:            "",
		FieldsPath:      "/OM_HEADER/OM_FIELD",
		FieldIDs:        []string{"1", "8"},
		ReplacePrevious: true,
	},
	{
		OmObject:        "Hourly Rundown",
		Path:            "/Radio Rundown",
		FieldsPath:      "/OM_HEADER/OM_FIELD",
		FieldIDs:        []string{"1", "8", "9"},
		ReplacePrevious: true,
	},
	// {
	// OmObject:   "Sub Rundown",
	// Path:       "/Radio Rundown/Hourly Rundown",
	// FieldsPath: "/OM_HEADER/OM_FIELD",
	// FieldIDs:   []string{"8"},
	// },
	// {
	// OmObject:   "Sub Rundown",
	// Path:       "Radio Rundown/Hourly Rundown/Sub Rundown",
	// FieldsPath: "/OM_HEADER/OM_FIELD",
	// FieldIDs:   []string{"8"},
	// Level:      1,
	// },
}

type CSVrowsIntMap map[int]CSVrow
type CSVrowsStringMap map[string]CSVrow
