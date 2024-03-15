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
	subl := len(rps)
	i := 0
	for i < subl {
		slog.Debug("slice range", "range", len(rps))
		slog.Debug("index", "index", i)
		subNodes := xmlquery.Find(rps[i].Node, objquery)
		subNodesCount := len(subNodes)
		if subNodesCount == 0 {
			slog.Debug("no subnodes found")
			i++
			continue
		}
		slog.Debug("subnodes found", "count", subNodesCount)
		subRows := ExtractNodesFields(subNodes, extr, rps[i].CSVrow)
		if extr.ReplacePrevious {
			slog.Debug("replacing previous row")
			subl += len(subRows)
			PrintRowPayloads(fmt.Sprintf("ORIG: %0d", i), rps)
			PrintRowPayloads(fmt.Sprintf("SUB: %0d", i), subRows)
			rps = subRows
			PrintRowPayloads(fmt.Sprintf("JOIN: %0d", i), rps)
			curIndex := 1
			i = curIndex + subl - 1
			slog.Debug("index", "current", curIndex, "new", i, "newrows", subl)
			// fmt.Println("new index", i)
			continue
		}
		if !extr.ReplacePrevious {
			slog.Debug("appending to previos row")
			subl += len(subRows)
			i = len(subRows) + 1
			rps = append(rps, subRows...)
		}
	}
	return rps
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
		OmObject: "",
		NodePath: "",
		Node:     node,
		CSVrow:   append(parentRow, csvrow...),
	}
}
