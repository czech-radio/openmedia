package internal

import (
	"fmt"
	"log/slog"

	"github.com/antchfx/xmlquery"
)

func ExpandTableRows(table CSVtable, extr OMextractor) (CSVtable, error) {
	objquery := XMLqueryFromPath(extr.ObjectPath)
	slog.Debug("object query", "query", objquery)

	var newTable CSVtable
	// for i, parentRow := range table {
	fmt.Println("")
	fmt.Println("NEW TABLE")
	slog.Debug("table length", "count", len(table.Rows))
	// for i := range table {
	for i, row := range table.Rows {
		subNodes := xmlquery.Find(row.Node, objquery)
		subNodesCount := len(subNodes)

		if subNodesCount == 0 {
			slog.Debug("no subnodes found", "row", i, "parentRow", row.CSVrow)
			newTable.Rows = append(newTable.Rows, row)
			continue
		}

		slog.Debug("subnodes found", "count", subNodesCount)
		// parentRowCopy := CopyRow(row.CSVrow)
		// Deep copy must be used here or at least in function which takes it as parameter and wants to modify it.
		// TODO: Try using CSVrowPart map[string]*CSVrowField insted of  map[string]CSVrowField. So the "copy of parent row" can be made parRow:=map[FieldID]&Field. Every row or its parts based on parent row will reference same value of common fields. So it can be changed/transformed globaly for whole table. Transforming operations must be done on whole column. If not the column will be contamineted and no furher global transform on column can be made easily without iterating over whole column. The pros of using field pointer is speed and less memory allocations.
		// subRows := ExtractNodesFields(parentRowCopy, subNodes, extr)
		subRows := ExtractNodesFields(row, subNodes, extr)
		if len(subRows.Rows) == 1 && extr.PreserveParentNode {
			subRows.Rows[0].Node = row.Node
			newTable.Rows = append(newTable.Rows, subRows.Rows[0])
			continue
		}

		if extr.KeepInputRows {
			newTable.Rows = append(newTable.Rows, subRows.Rows...)
			slog.Debug("appendig also input row")
			newTable.Rows = append(newTable.Rows, row)
			continue
		}

		if !extr.KeepInputRows {
			slog.Debug("replacing input row")
			newTable.Rows = append(newTable.Rows, subRows.Rows...)
			continue
		}
	}
	return newTable, nil
}

func CopyRow(inputRow CSVrow) CSVrow {
	newRow := make(CSVrow)
	for ai, a := range inputRow {
		newRow[ai] = make(CSVrowPart)
		for bi, b := range a {
			newRow[ai][bi] = b
		}
	}
	return newRow
}

func ExtractNodesFields(
	// parentRow CSVrow,
	parentRow *CSVrowNode,
	subNodes []*xmlquery.Node,
	extr OMextractor,
) CSVtable {
	var table CSVtable
	prefix := PartsPrefixMapProduction[extr.PartPrefixCode].Internal
	for _, subNode := range subNodes {
		parentRowCopy := CopyRow(parentRow.CSVrow)
		part := NodeToCSVrowPart(subNode, extr)
		rowNode := CSVrowNode{}
		rowNode.Node = subNode
		rowNode.Node = subNode
		rowNode.CSVrow = parentRowCopy
		rowNode.CSVrow[prefix] = part
		table.Rows = append(table.Rows, &rowNode)
	}
	return table
}

func NodeToCSVrowPart(node *xmlquery.Node, ext OMextractor) CSVrowPart {
	fieldCount := len(ext.ObjectAttrsNames) + len(ext.FieldIDs)
	part := make(CSVrowPart, fieldCount)
	// Object attibutes
	part = NodeGetAttributes(node, part, ext)
	// Object FieldIDs
	part = NodeGetFields(node, part, ext)
	return part
}

func NodeGetAttributes(
	node *xmlquery.Node, part CSVrowPart, ext OMextractor) CSVrowPart {
	for _, attrName := range ext.ObjectAttrsNames {
		attrVal, _ := GetFieldValueByName(node.Attr, attrName)
		field := CSVrowField{
			FieldID: attrName,
			// FieldName: fieldName,
			// or send it to map[Prefix]map[FieldID]FieldName so it does not take memory
			Value: attrVal,
		}
		part[attrName] = field
	}
	return part
}

func NodeGetFields(
	node *xmlquery.Node, part CSVrowPart, ext OMextractor) CSVrowPart {
	// Object FieldIDs
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
	for _, f := range fields {
		fieldID, _ := GetFieldValueByName(f.Attr, "FieldID")
		// fieldName, _ := GetFieldValueByName(f.Attr, "FieldID")
		field := CSVrowField{
			FieldID: fieldID,
			// FieldName: fieldName,
			// or send it to map[Prefix]map[FieldID]FieldName
			Value: f.InnerText(),
		}
		part[fieldID] = field
	}
	return part
}
