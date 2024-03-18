package internal

import (
	"log/slog"

	"github.com/antchfx/xmlquery"
)

func ExpandTableRows(table CSVtable, extr OMobjExtractor) (CSVtable, error) {
	objectType := GetLastPartOfObjectPath(extr.ObjectPath)
	objquery, err := QueryObject(objectType)
	if err != nil {
		return nil, err
	}
	slog.Debug("object query", "query", objquery)
	for i := range table {
		subNodes := xmlquery.Find(table[i].Node, objquery)
		subNodesCount := len(subNodes)
		if subNodesCount == 0 {
			slog.Debug("no subnodes found")
		}
		slog.Debug("subnodes found", "count", subNodesCount)
		// subRows := ExtractNodeFields()
	}
	return table, nil
}

// func ExtractRowNodesFields(nodes []*xmlquery.Node, extr OMobjExtractor) CSVrowFields {
func ExtractRowNodesFields(nodes []*xmlquery.Node, extr OMobjExtractor) CSVtable {
	var table CSVtable
	//"TODO
	// var nodeRows []*ObjectRow
	// for _, n := range nodes {
	// row := ExtractNodeFields(n, extr, parentRow)
	// nodeRows = append(nodeRows, row)
	// }
	// return nodeRows
	return table
}

func NodeToCSVrowPart(node *xmlquery.Node, ext OMobjExtractor) CSVrowPart {
	var part CSVrowPart
	attrQuery := XMLbuildAttrQuery("FieldID", ext.FieldIDs)
	if attrQuery == "" {
		return part //empty row
	}
	query := ext.FieldsPath + attrQuery
	slog.Debug("query fields", "query", query)
	fields := xmlquery.Find(node, query)
	if fields == nil {
		slog.Error("fields is nil")
		return part
	}
	if len(fields) == 0 {
		slog.Error("no fields found")
		return part
	}
	part = make(CSVrowPart, len(fields))
	for _, f := range fields {
		fieldID, _ := GetFieldValueByName(f.Attr, "FieldID")
		fieldName, _ := GetFieldValueByName(f.Attr, "FieldID")
		field := CSVrowField{
			FieldID:   fieldID,
			FieldName: fieldName, // or send it to map[Prefix]map[FieldID]FieldName
			Value:     f.InnerText(),
		}
		part[fieldID] = field
	}
	return part
}
