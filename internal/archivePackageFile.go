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

var CSVproduction = []OMobjExtractor{
	// {
	// OmObject:   "Radio Rundown",
	// FieldsPath: "/OM_HEADER/OM_FIELD",
	// FieldIDs:   []string{"1", "8"},
	// Level:      0,
	// },
	// {
	// OmObject:   "Radio Rundown",
	// FieldsPath: "/OM_RECORD/OM_FIELD",
	// FieldIDs:   []string{"8"},
	// Level:      0,
	// },
	{
		OmObject:   "Hourly Rundown",
		FieldsPath: "/OM_HEADER/OM_FIELD",
		FieldIDs:   []string{"1", "8"},
		Level:      0,
	},
	{
		OmObject:   "Sub Rundown",
		FieldsPath: "/OM_HEADER/OM_FIELD",
		FieldIDs:   []string{"8"},
		Level:      1,
	},
}

type CSVrowsIntMap map[int]CSVrow
type CSVrowsStringMap map[string]CSVrow

func (apf *ArchivePackageFile) ExtractByXMLqueryB(
	enc FileEncodingNumber, q *ArchiveFolderQuery) error {
	dataReader, err := ZipXmlFileDecodeData(apf.Reader, enc)
	if err != nil {
		return err
	}
	node, err := xmlquery.Parse(dataReader)
	if err != nil {
		return err
	}
	// var rows []CSVrow
	// var rowsm map[string]CSVrow
	var rowsm CSVrowsIntMap

	nodes := FindBaseNodes(node, CSVproduction[0])
	rowsm = NodesToCSVrows(nodes, CSVproduction[0], rowsm)
	PrintRows(rowsm)

	// for _, n := range nodes {
	// }
	return nil
}

// for i := range nodes {
// for i := range nodes {
// fmt.Println(nodes[i].Attr)
// nodules := FindSubNodes(nodes[i], CSVproduction[1])
// fmt.Println(len(nodules))
// rowsm := NodesToCSVrows(nodules, CSVproduction[1], rowsm)
// PrintRows(rowsm)
// }
// nodes := []xmlquery.Node{*node}
// CSVproduction[0].MapFields()
// rows = ExtractNodeToCSVrows(node, CSVproduction[0], rows)
// rows = ExtractNodeToCSVrows(node, CSVproduction[1], rows)
// fmt.Println("fuck", rows)
// fmt.Println(apf.Reader.Name)
// PrintRows(rowsm)

func PrintRows(rows map[int]CSVrow) {
	for i := 0; i < len(rows); i++ {
		fmt.Println(i, rows[i])
		fmt.Println()
	}
}

func FindBaseNodes(node *xmlquery.Node, ext OMobjExtractor) []*xmlquery.Node {
	query := fmt.Sprintf("//OM_OBJECT[@TemplateName='%s']", ext.OmObject)
	return xmlquery.Find(node, query)
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

func ExtractNodeToCSVrows(
	node *xmlquery.Node, ext OMobjExtractor, rows []CSVrow) []CSVrow {
	query := fmt.Sprintf("//OM_OBJECT[@TemplateName='%s']", ext.OmObject)

	// Objects query
	nodes := xmlquery.Find(node, query)
	// fmt.Println(len(nodes))
	// if len(ext.FieldIDs) == 0 {
	// return rows
	// }

	// // Fields query
	// // subquery := fmt.Sprintf("%s%s", query, ext.FieldsPath)
	if len(rows) == 0 {
		rowadd := make([]CSVrow, len(nodes))
		rows = append(rows, rowadd...)
	}
	// fmt.Println("LEN OBJECTS", len(nodes))
	for i, node := range nodes {
		row := NodeToCSVrow(node, ext)
		rows[i] = row
	}
	return rows
}

// func ExtractFields(node *xmlquery.Node, query string, fieldIDsMap map[string]bool) CSVrow {
// func ExtractFieldsB(node *xmlquery.Node, path string, fieldIDs []string) CSVrow {
// var row CSVrow
// query := path + BuildFieldsQuery(fieldIDs)
// fields := xmlquery.Find(node, query)
// for _, f := range fields {
// fieldID, _ := GetFieldValueByID(f.Attr, "FieldID")
// field := CSVrowField{
// FieldPosition: 0,
// FieldID:       fieldID,
// Value:         f.InnerText(),
// }
// row = append(row, field)
// }
// return row
// }

// query := fmt.Sprintf(
// "//OM_OBJECT[@TemplateName='Radio Rundown']",
// "//OM_OBJECT[@TemplateName='%s']/%s/*", ext.OM_type, ext.Path,
// "//OM_OBJECT[@TemplateName='%s']", ext.OM_type,
// )
