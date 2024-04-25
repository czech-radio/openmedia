// Package extract manages extraction of data from Openmedia rundowns archive (zip packaged rundown xml files) or standalone xml rundown files.
package extract

import (
	ar "github/czech-radio/openmedia/internal/archive"
	"github/czech-radio/openmedia/internal/helper"
	"io"
	"log/slog"
	"strings"

	"github.com/antchfx/xmlquery"
)

// ExpandTableRows parse additional data from xml for each row. Multiple rows none,one or more rows may be created from one original row.
// Deep copy must be used here or at least in function which takes it as parameter and wants to modify it.
// TODO: Try using CSVrowPart map[string]*CSVrowField insted of  map[string]CSVrowField. So the "copy of parent row" can be made parRow:=map[FieldID]&Field. Every row or its parts based on parent row will reference same value of common fields. So it can be changed/transformed globaly for whole table. Transforming operations must be done on whole column. If not the column will be contamineted and no furher global transform on column can be made easily without iterating over whole column. The pros of using field pointer is speed and less memory allocations.
func ExpandTableRows(table TableXML, extr OMextractor) (TableXML, error) {
	var newTable TableXML
	objquery := ar.XMLqueryFromPath(extr.ObjectPath)
	slog.Debug("object query", "query", objquery)
	slog.Debug("table length", "count", len(table.Rows))
	for _, row := range table.Rows {
		// Find main object: exmp OM_RECORD
		subNodes := xmlquery.Find(row.Node, objquery)
		subNodesCount := len(subNodes)

		prow := CopyRow(row.RowParts)
		newRow := &RowNode{row.Node, prow}
		if len(subNodes) == 0 && extr.KeepWhenZeroSubnodes {
			slog.Debug("subnodes not_found",
				"objquery", objquery, "parent", row.Node.Data)
			newTable.Rows = append(newTable.Rows, newRow)
			continue
		}
		// Extract objects fields
		slog.Debug("subnodes found", "count", subNodesCount)
		subRows := ExtractNodesFields(row, subNodes, extr)
		newTable.Rows = append(newTable.Rows, subRows.Rows...)
	}
	return newTable, nil
}

// CopyRow
func CopyRow(inputRow RowParts) RowParts {
	newRow := make(RowParts)
	for ai, a := range inputRow {
		newRow[ai] = make(RowPart)
		for bi, b := range a {
			newRow[ai][bi] = b
		}
	}
	return newRow
}

// ExtractNodesFields
func ExtractNodesFields(
	parentRow *RowNode,
	subNodes []*xmlquery.Node,
	extr OMextractor,
) TableXML {
	var newTable TableXML
	for _, subNode := range subNodes {
		parentRowCopy := CopyRow(parentRow.RowParts)
		part := NodeToCSVrowPart(subNode, extr)
		newRowNode := RowNode{}
		parentNode, _ := helper.XMLnodeLevelUp(
			subNode, extr.ResultNodeGoUpLevels,
		)
		newRowNode.Node = parentNode
		newRowNode.RowParts = parentRowCopy
		newRowNode.RowParts[extr.PartPrefixCode] = part
		newTable.Rows = append(newTable.Rows, &newRowNode)
	}
	return newTable
}

// NodeToCSVrowPart
func NodeToCSVrowPart(node *xmlquery.Node, ext OMextractor) RowPart {
	fieldCount := len(ext.ObjectAttrsNames) + len(ext.FieldIDs)
	part := make(RowPart, fieldCount)
	// Object attibutes
	part = NodeGetAttributes(node, part, ext)
	// Object FieldIDs
	part = NodeGetFields(node, part, ext)
	return part
}

// NodeGetAttributes
func NodeGetAttributes(
	node *xmlquery.Node, part RowPart, ext OMextractor) RowPart {
	for _, attrName := range ext.ObjectAttrsNames {
		attrVal, _ := ar.GetFieldValueByName(node.Attr, attrName)
		field := RowField{
			FieldID: attrName,
			// FieldName: fieldName,
			// or send it to map[Prefix]map[FieldID]FieldName so it does not take memory
			Value: strings.TrimSpace(attrVal),
		}
		part[attrName] = field
	}
	return part
}

func NodeGetFields(
	node *xmlquery.Node, part RowPart, ext OMextractor) RowPart {
	// Query object FieldIDs
	attrQuery := ar.XMLbuildAttrQuery("FieldID", ext.FieldIDs)
	if attrQuery == "" { // No fields found
		return part
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
		fieldID, _ := ar.GetFieldValueByName(f.Attr, "FieldID")
		field := RowField{
			FieldID: fieldID,
			// FieldName: fieldName,
			// or send it to map[Prefix]map[FieldID]FieldName
			Value: strings.TrimSpace(f.InnerText()),
		}
		part[fieldID] = field
	}
	return part
}

// RowPartMarkNotPossible
// func RowPartMarkNotPossible(part CSVrowPart, ext OMextractor) {
// }

func XMLparalelQuery(extractors []OMextractor) string {
	var query string
	objects := []string{}
	for _, e := range extractors {
		objects = append(objects, e.ObjectPath)
	}
	query = ar.XMLbuildAttrQuery("TemplateName", objects)
	return query
}

func XMLgetOpenmediaBaseNode(reader io.Reader) (*xmlquery.Node, error) {
	return helper.XMLgetBaseNode(reader, "/OPENMEDIA")
}
