package internal

import (
	"fmt"
	"log/slog"

	"github.com/antchfx/xmlquery"
)

type LinkedRow struct {
	FirstL    *LinkedRow
	NextL     *LinkedRow
	PrevL     *LinkedRow
	Node      *xmlquery.Node
	CSVrow    CSVrow
	RowsCount *int
	// LinkNumber int ??? without function?
}

func NewLinkedRow(
	parentNode *xmlquery.Node, firstL, prevL *LinkedRow, csvRow CSVrow) *LinkedRow {
	newRow := new(LinkedRow)
	newRow.Node = parentNode
	if firstL == nil {
		newRow.FirstL = newRow
		count := 1
		newRow.RowsCount = &count
	} else {
		newRow.FirstL = firstL
	}
	if csvRow != nil {
		newRow.CSVrow = append(newRow.CSVrow, csvRow...)
	}
	if prevL != nil {
		newRow.PrevL = prevL
	}
	// newRow.RowsCount = firstL.RowsCount
	return newRow
}

func (apf *ArchivePackageFile) ExtractByXMLquery(
	enc FileEncodingNumber, q *ArchiveFolderQuery) error {
	dataReader, err := ZipXmlFileDecodeData(apf.Reader, enc)
	if err != nil {
		return err
	}
	node, err := xmlquery.Parse(dataReader)
	if err != nil {
		return err
	}
	firstRow := NewLinkedRow(node, nil, nil, nil)
	firstRow.CSVrow = []CSVrowField{{1, "kek", "smek"}}
	lrow := NodeToCSVlinkedRow(node, CSVproduction[0], firstRow)
	fmt.Println("ka", lrow.CSVrow)
	return nil
}

func NodeToCSVlinkedRow(
	objNode *xmlquery.Node, ext OMobjExtractor, lrow *LinkedRow,
) *LinkedRow {
	query := fmt.Sprintf("//OM_OBJECT[@TemplateName='%s']", ext.OmObject)
	if lrow == nil {
		return lrow
	}
	curL := lrow
	for {
		// Find objects
		nodes := xmlquery.Find(lrow.Node, query)
		slog.Debug("nodes found", "count", len(nodes))
		parentCsvRow := curL.CSVrow
		lrows := NodesExtractFieldsToRows(parentCsvRow, nodes, ext)
		slog.Debug("rows", "count", len(lrows))
		// Checkout next link
		curL = curL.NextL
		if curL == nil {
			break
		}
	}
	return lrow.FirstL
}

func NodesExtractFieldsToRows(
	parentCsvRow CSVrow, nodes []*xmlquery.Node, ext OMobjExtractor,
) []*LinkedRow {
	lrows := make([]*LinkedRow, len(nodes))
	var first, cur *LinkedRow
	for i, node := range nodes {
		// Object fields
		csvrow := NodeToCSVrow(node, ext)
		cvsRowJoined := append(parentCsvRow, csvrow...)
		nlr := NewLinkedRow(node, first, cur, cvsRowJoined)
		if first == nil {
			first = nlr
		}
		cur = nlr
		lrows[i] = nlr
	}
	PrintLinesSlice(lrows)
	return lrows
}

func PrintLinesSlice(lrows []*LinkedRow) {
	for _, r := range lrows {
		fmt.Println("current", r.CSVrow)
		if r.PrevL != nil {
			fmt.Println("previous", r.PrevL.CSVrow)
		}
	}
}
