package internal

import (
	"fmt"
	"log/slog"

	"github.com/antchfx/xmlquery"
)

func (apf *ArchivePackageFile) ExtractByXMLquery(
	enc FileEncodingNumber, q *ArchiveFolderQuery) error {
	dataReader, err := ZipXmlFileDecodeData(apf.Reader, enc)
	if err != nil {
		return err
	}
	baseNode, err := xmlquery.Parse(dataReader)
	if err != nil {
		return err
	}
	pay, err := ExtractFromBaseObject(baseNode, CSVproduction)
	PrintRowPayloads("RESULT", pay)
	return err
}

type RowPayload struct {
	OmObject     string
	FieldsPrefix string
	NodePath     string
	Node         *xmlquery.Node
	CSVrow
}

func ExtractFromBaseObject(baseNode *xmlquery.Node, extrs []OMobjExtractor) ([]*RowPayload, error) {
	baseRow := &RowPayload{
		OmObject: "",
		NodePath: "",
		Node:     baseNode,
		CSVrow:   CSVrow{{1, "tF1", "vF1"}},
	}
	rows := []*RowPayload{baseRow}
	for _, e := range extrs {
		rows = ExpandRows(rows, e)
	}
	return rows, nil
}

func ExpandRows(rps []*RowPayload, extr OMobjExtractor) []*RowPayload {
	objquery := fmt.Sprintf("//OM_OBJECT[@TemplateName='%s']", extr.OmObject)
	subRowsCount := len(rps)
	var result []*RowPayload
	for i := range rps {
		subNodes := xmlquery.Find(rps[i].Node, objquery)
		subNodesCount := len(subNodes)
		if subNodesCount == 0 {
			slog.Debug("no subnodes found")
		}
		slog.Debug("subnodes found", "count", subNodesCount)
		subRows := ExtractNodesFields(subNodes, extr, rps[i].CSVrow)
		subRowsCount += len(subRows)

		if extr.ReplacePrevious {
			slog.Debug("replacing previous row")
			result = append(result, subRows...)
		}
		if !extr.ReplacePrevious {
			slog.Debug("appending to previos row")
			result = append(result, subRows...)
			// also append the previous row
			result = append(result, rps[i])
		}
	}
	return result
}

func ExtractNodesFields(
	nodes []*xmlquery.Node, extr OMobjExtractor, parentRow CSVrow,
) []*RowPayload {
	var nodeRows []*RowPayload
	for _, n := range nodes {
		row := ExtractNodeFields(n, extr, parentRow)
		nodeRows = append(nodeRows, row)
	}
	return nodeRows
}

func ExtractNodeFields(
	node *xmlquery.Node, extr OMobjExtractor, parentRow CSVrow,
) *RowPayload {
	csvrow := NodeToCSVrow(node, extr)
	return &RowPayload{
		OmObject:     extr.OmObject,
		FieldsPrefix: extr.FieldsPrefix,
		NodePath:     "",
		Node:         node,
		CSVrow:       append(parentRow, csvrow...),
	}
}

func NodeToCSVrow(node *xmlquery.Node, ext OMobjExtractor) CSVrow {
	var csvrow CSVrow
	attrQuery := XMLbuildAttrQuery("FieldID", ext.FieldIDs)
	if attrQuery == "" {
		return csvrow
	}
	query := ext.FieldsPath + attrQuery
	fields := xmlquery.Find(node, query)
	if fields == nil {
		return csvrow
	}
	if len(fields) == 0 {
		slog.Error("nothing found")
		return csvrow
	}
	for pos, f := range fields {
		fieldID, _ := GetFieldValueByName(f.Attr, "FieldID")
		field := CSVrowField{
			FieldPosition: pos,
			FieldID:       fieldID,
			Value:         f.InnerText(),
		}
		csvrow = append(csvrow, field)
	}
	return csvrow
}

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

func FindSubNodes(node *xmlquery.Node, ext OMobjExtractor) []*xmlquery.Node {
	query := fmt.Sprintf("//OM_OBJECT[@TemplateName='%s']", ext.OmObject)
	return xmlquery.Find(node, query)
}
