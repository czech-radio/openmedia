package internal

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"log/slog"

	"github.com/antchfx/xmlquery"
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
	OmObject    string
	Path        string
	FieldsPath  string
	FieldIDs    []string
	Level       int
	FieldIDsMap map[string]bool
}

func (omo *OMobjExtractor) MapFields() {
	omo.FieldIDsMap = make(map[string]bool, len(omo.FieldIDs))
	for _, id := range omo.FieldIDs {
		omo.FieldIDsMap[id] = true
	}
}

var FieldName = map[int]string{
	8: "Název",
}

// "//OM_OBJECT[@TemplateName='%s']/%s/*", ext.OM_type, ext.Path,
// FieldsPath: "/OM_HEADER/OM_FIELD",
// FieldsPath: "/OM_RECORD/OM_FIELD",
var CSVproduction = []OMobjExtractor{
	{
		OmObject:   "Radio Rundown",
		Path:       "Radio Rundown",
		FieldsPath: "/OM_HEADER/OM_FIELD",
		FieldIDs:   []string{"1", "8"},
		Level:      0,
	},
	{
		OmObject:   "Hourly Rundown",
		Path:       "Radio Rundown/Hourly Rundown",
		FieldsPath: "/OM_HEADER/OM_FIELD",
		FieldIDs:   []string{"1", "8"},
		Level:      0,
	},
	{
		OmObject:   "Sub Rundown",
		Path:       "Radio Rundown/Hourly Rundown/Sub Rundown",
		FieldsPath: "/OM_HEADER/OM_FIELD",
		FieldIDs:   []string{"8"},
		Level:      1,
	},
	{
		OmObject:   "Sub Rundown",
		Path:       "Radio Rundown/Hourly Rundown/Sub Rundown",
		FieldsPath: "/OM_HEADER/OM_FIELD",
		FieldIDs:   []string{"8"},
		Level:      1,
	},
}

type CSVrowsIntMap map[int]CSVrow
type CSVrowsStringMap map[string]CSVrow

func PrintRows(rows map[int]CSVrow) {
	for i := 0; i < len(rows); i++ {
		fmt.Println(i, rows[i])
		fmt.Println()
	}
}

func FindSubNodes(node *xmlquery.Node, ext OMobjExtractor) []*xmlquery.Node {
	query := fmt.Sprintf("//OM_OBJECT[@TemplateName='%s']", ext.OmObject)
	return xmlquery.Find(node, query)
}

// func NodesToCSVrows(nodes []*xmlquery.Node, ext OMobjExtractor, rows []CSVrow) []CSVrow {
func NodesToCSVrows(nodes []*xmlquery.Node, ext OMobjExtractor, rows CSVrowsIntMap) CSVrowsIntMap {
	if len(rows) == 0 {
		// rows = make(map[int]CSVrow, len(nodes))
		rows = make(CSVrowsIntMap, len(nodes))
	}
	for i, node := range nodes {
		row := NodeToCSVrow(node, ext)
		rows[i] = append(rows[i], row...)
	}
	return rows
}

func NodeToCSVrow(node *xmlquery.Node, ext OMobjExtractor) CSVrow {
	var csvrow CSVrow
	// query := ext.FieldsPath + BuildFieldsQuery(ext.FieldIDs)
	query := ext.FieldsPath + XMLbuildAttrQuery("FieldID", ext.FieldIDs)
	fields := xmlquery.Find(node, query)
	if len(fields) == 0 {
		slog.Error("nothing found")
		return csvrow
	}
	for _, f := range fields {
		fieldID, _ := GetFieldValueByID(f.Attr, "FieldID")
		field := CSVrowField{
			FieldPosition: 0,
			FieldID:       fieldID,
			Value:         f.InnerText(),
		}
		csvrow = append(csvrow, field)
	}
	return csvrow
}
