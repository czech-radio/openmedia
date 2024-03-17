package internal

import (
	"fmt"
	"log/slog"

	"github.com/antchfx/xmlquery"
)

type ObjectRow struct {
	OmObject     string
	FieldsPrefix string
	NodePath     string
	Node         *xmlquery.Node
	CSVrowFields
}

func ExtractBaseObjectRows(baseNode *xmlquery.Node, extrs OMobjExtractors) ([]*ObjectRow, error) {
	// Main extract
	baseRow := &ObjectRow{
		OmObject:     "",
		NodePath:     "",
		Node:         baseNode,
		CSVrowFields: CSVrowFields{{1, "tF1", "vF1"}},
	}
	rows := []*ObjectRow{baseRow}
	extrs.ReplaceParentRowTrueChecker()
	for _, e := range extrs {
		rows = ExpandObjectRows(rows, e)
	}
	return rows, nil
}

func ExpandObjectRows(rps []*ObjectRow, extr OMobjExtractor) []*ObjectRow {
	objquery := fmt.Sprintf("//OM_OBJECT[@TemplateName='%s']", extr.OmObject)
	subRowsCount := len(rps)
	var result []*ObjectRow
	for i := range rps {
		subNodes := xmlquery.Find(rps[i].Node, objquery)
		subNodesCount := len(subNodes)
		if subNodesCount == 0 {
			slog.Debug("no subnodes found")
		}
		slog.Debug("subnodes found", "count", subNodesCount)
		subRows := ExtractNodesFields(subNodes, extr, rps[i].CSVrowFields)
		subRowsCount += len(subRows)

		if extr.DontReplaceParentObjectRow {
			slog.Debug("replacing previous row")
			result = append(result, subRows...)
		}
		if !extr.DontReplaceParentObjectRow {
			slog.Debug("appending after previos row")
			result = append(result, subRows...)
			// also append the previous row
			result = append(result, rps[i])
		}
	}
	return result
}

func ExtractNodesFields(
	nodes []*xmlquery.Node, extr OMobjExtractor, parentRow CSVrowFields,
) []*ObjectRow {
	var nodeRows []*ObjectRow
	for _, n := range nodes {
		row := ExtractNodeFields(n, extr, parentRow)
		nodeRows = append(nodeRows, row)
	}
	return nodeRows
}

func ExtractNodeFields(
	node *xmlquery.Node, extr OMobjExtractor, parentRow CSVrowFields,
) *ObjectRow {
	csvrow := NodeToCSVrow(node, extr)
	return &ObjectRow{
		OmObject:     extr.OmObject,
		FieldsPrefix: extr.FieldsPrefix,
		NodePath:     "",
		Node:         node,
		CSVrowFields: append(parentRow, csvrow...),
	}
}

func NodeToCSVrow(node *xmlquery.Node, ext OMobjExtractor) CSVrowFields {
	var csvrow CSVrowFields
	attrQuery := XMLbuildAttrQuery("FieldID", ext.FieldIDs)
	if attrQuery == "" {
		return csvrow //empty row
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
