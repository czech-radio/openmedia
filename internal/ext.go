package internal

import (
	"fmt"
	"log/slog"

	"github.com/antchfx/xmlquery"
)

type LinkedRow struct {
	Start     **LinkedRow
	FirstL    *LinkedRow
	NextL     *LinkedRow
	PrevL     *LinkedRow
	Node      *xmlquery.Node
	RowsCount *int
	CSVrow    CSVrow
	// LinkNumber int ??? without function?
}

func (l *LinkedRow) NewNextLink(
	parentNode *xmlquery.Node, csvRow CSVrow,
) *LinkedRow {
	newRow := new(LinkedRow)
	newRow.Node = parentNode
	if l == nil {
		newRow.FirstL = newRow
		count := 1
		newRow.RowsCount = &count
		newRow.Start = &newRow
		slog.Debug("added")
	} else {
		newRow.Start = l.Start
		newRow.FirstL = l.FirstL
	}
	if csvRow != nil {
		newRow.CSVrow = append(newRow.CSVrow, csvRow...)
	}
	newRow.PrevL = l
	return newRow
}

func (l *LinkedRow) ReplaceLink(starting, end *LinkedRow) {
	// firstl := l.FirstL
	// prevl := l.PrevL
	// nextl := l.NextL
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
	// firstRow := NewLinkedRow(node, nil, nil, nil)
	firstRow := new(LinkedRow)
	firstRow = firstRow.NewNextLink(node, nil)
	firstRow.CSVrow = []CSVrowField{{1, "testF1", "valueF1"}}
	// lrow := NodeToCSVlinkedRow(node, CSVproduction[0], firstRow)
	firstRow = NodeToCSVlinkedRow(node, CSVproduction[0], firstRow)
	fmt.Println("ka", firstRow)
	return nil
}

func NodeToCSVlinkedRow(
	objNode *xmlquery.Node, ext OMobjExtractor, lrow *LinkedRow,
) *LinkedRow {
	query := fmt.Sprintf("//OM_OBJECT[@TemplateName='%s']", ext.OmObject)
	if lrow == nil {
		slog.Debug("querying nil node")
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
		start := *lrows[1].Start
		fmt.Println("starter", start.CSVrow)
		// Checkout next link
		curL = curL.NextL
		if curL == nil {
			break
		}
	}
	// fmt.Println("six", lrow.Start)
	return lrow.FirstL
	// return curL.FirstL
}
func NodesExtractFieldsToRows(
	parentCsvRow CSVrow, nodes []*xmlquery.Node, ext OMobjExtractor,
) []*LinkedRow {
	lrows := make([]*LinkedRow, len(nodes))
	var nlr *LinkedRow
	for i, node := range nodes {
		// Object fields
		csvrow := NodeToCSVrow(node, ext)
		csvRowJoined := append(parentCsvRow, csvrow...)
		nlr = nlr.NewNextLink(node, csvRowJoined)
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
