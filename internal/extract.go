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
		OmObject: "",
		NodePath: "",
		Node:     baseNode,
		// CSVrowFields: CSVrowFields{{1, "tF1", "vF1"}},
	}
	var err error
	rows := []*ObjectRow{baseRow}
	extrs.ReplaceParentRowTrueChecker()
	for _, extr := range extrs {
		rows, err = ExpandObjectRows(rows, extr) // : maybe wrong
		if err != nil {
			return rows, err
		}
		// if i > 1 {
		// slog.Debug("break", "at", i)
		// break
		// }
	}
	return rows, nil
}

func QueryObject(objectName string) (string, error) {
	var XMLattrName string
	var XMLobjectName string
	var XMLattrValue string
	var attrquery string
	xmlTag, notObjTemplate := OmTagStructureMap[objectName]
	if notObjTemplate {
		XMLobjectName = xmlTag.XMLtagName
		XMLattrName = xmlTag.SelectorAttr
		if !notObjTemplate {
			return "", fmt.Errorf("unknown object: %s", objectName)
		}
		XMLattrValue = ""
	}
	if !notObjTemplate {
		XMLobjectName = "OM_OBJECT"
		XMLattrName = "TemplateName"
		XMLattrValue = objectName
		attrquery = XMLbuildAttrQuery(XMLattrName, []string{XMLattrValue})
	}
	fmt.Println("doprdlel", XMLobjectName, XMLattrName, XMLattrValue)
	objquery := "/" + XMLobjectName + attrquery
	return objquery, nil
}

func QueryFields(fieldsPath string, IDs []string) string {
	attrQuery := XMLbuildAttrQuery("FieldID", IDs)
	return fieldsPath + attrQuery
}

func ExpandObjectRows(rps []*ObjectRow, extr OMobjExtractor) ([]*ObjectRow, error) {
	objectType := GetLastPartOfObjectPath(extr.ObjectPath)
	objquery, err := QueryObject(objectType)
	if err != nil {
		return nil, err
	}
	slog.Debug("object query", "query", objquery)
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
		slog.Debug("sub rows", "count", len(subRows))
		subRowsCount += len(subRows)

		if !extr.DontReplaceParentObjectRow {
			slog.Debug("replacing previous row")
			result = append(result, subRows...)
		}
		if extr.DontReplaceParentObjectRow {
			slog.Debug("appending after previos row")
			result = append(result, subRows...)
			// also append the previous row
			result = append(result, rps[i])
		}
	}
	return result, nil
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
	slog.Debug("query fields", "query", query)
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
