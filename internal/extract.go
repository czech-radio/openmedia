package internal

import (
	"log/slog"
	"maps"

	"github.com/antchfx/xmlquery"
)

func ExpandTableRows(table CSVtable, extr OMextractor) (CSVtable, error) {
	objectType := GetLastPartOfObjectPath(extr.ObjectPath)
	objquery, err := QueryObject(objectType)
	if err != nil {
		return nil, err
	}
	slog.Debug("object query", "query", objquery)

	var result CSVtable
	// for i, parentRow := range table {
	for i := range table {
		slog.Debug("table length", "count", len(table))
		subNodes := xmlquery.Find(table[i].Node, objquery)
		subNodesCount := len(subNodes)

		if subNodesCount == 0 {
			slog.Debug("no subnodes found", "row", i, "parentRow", table[i].CSVrow)
			continue
		}

		slog.Debug("subnodes found", "count", subNodesCount)
		parentRowCopy := CSVrow{}
		maps.Copy(parentRowCopy, table[i].CSVrow)
		// Deep copy must be used here or at least in function which takes it as parameter and wants to modify it.
		subRows := ExtractNodesFields(parentRowCopy, subNodes, extr)

		if extr.KeepInputRows {
			result = append(result, subRows...)
			slog.Debug("appendig also input row")
			result = append(result, table[i])
		}

		if !extr.KeepInputRows {
			slog.Debug("replacing input row")
			result = append(result, subRows...)
		}
	}
	return result, nil
}

func ExtractNodesFields(
	parentRow CSVrow,
	subNodes []*xmlquery.Node,
	extr OMextractor,
) CSVtable {
	var table CSVtable
	for _, subNode := range subNodes {
		parentRowCopy := CSVrow{}
		maps.Copy(parentRowCopy, parentRow)
		part := NodeToCSVrowPart(subNode, extr)
		rowNode := CSVrowNode{}
		rowNode.Node = subNode
		rowNode.CSVrow = parentRowCopy
		rowNode.CSVrow[extr.FieldsPrefix] = part
		table = append(table, &rowNode)
	}
	return table
}

func NodeToCSVrowPart(node *xmlquery.Node, ext OMextractor) CSVrowPart {
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
