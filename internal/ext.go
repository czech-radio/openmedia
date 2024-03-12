package internal

import (
	"fmt"
	"log/slog"

	"github.com/antchfx/xmlquery"
)

type LinkedRow struct {
	FirstRow  *LinkedRow
	NextL     *LinkedRow
	PrevL     *LinkedRow
	Node      *xmlquery.Node
	CSVrow    CSVrow
	RowsCount *int
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
	firstRow := NewLinkedRow(node, nil, nil)
	firstRow.CSVrow = []CSVrowField{{1, "kek", "smek"}}
	lrow := NodeToCSVlinkedRow(node, CSVproduction[0], firstRow)
	fmt.Println("ka", lrow.CSVrow)
	// fmt.Println("ka", lrow.FirstRow.CSVrow)
	// fmt.Println(lrow.FirstRow.CSVrow)
	return nil
}

func NodeToCSVlinkedRow(
	objNode *xmlquery.Node, ext OMobjExtractor, lrow *LinkedRow,
) *LinkedRow {
	query := fmt.Sprintf("//OM_OBJECT[@TemplateName='%s']", ext.OmObject)
	curL := lrow
	for i := 0; i < *lrow.FirstRow.RowsCount; i++ {
		// Find objects
		nodes := xmlquery.Find(lrow.Node, query)
		slog.Debug("nodes found", "count", len(nodes))
		parentCsvRow := curL.CSVrow
		for _, node := range nodes {
			// Objet fields
			csvrow := NodeToCSVrow(node, ext)
			joined := append(parentCsvRow, csvrow...)
			fmt.Println(joined)
		}
	}
	return lrow.FirstRow
}

func NewLinkedRow(
	parentNode *xmlquery.Node, firstL *LinkedRow, csvRow CSVrow) *LinkedRow {
	newRow := new(LinkedRow)
	newRow.Node = parentNode
	if firstL == nil {
		newRow.FirstRow = newRow
		count := 1
		newRow.RowsCount = &count
	} else {
		newRow.FirstRow = firstL
	}
	if csvRow != nil {
		newRow.CSVrow = append(newRow.CSVrow, csvRow...)
	}
	// newRow.RowsCount = firstL.RowsCount
	return newRow
}

// curRow := lrow.FirstRow
// slog.Debug("rowsCount", "count", *curRow.RowsCount)
// for i := 0; i < *lrow.FirstRow.RowsCount; i++ {
// nodes := xmlquery.Find(curRow.Node, query)
// slog.Debug("nodes found", "count", len(nodes))
// for _, nd := range nodes {
// for _, node := range nodes {
// Object fields
// csvrow := NodeToCSVrow(node, ext)
// nl := NewLinkedRow(node, csvrow)
// fmt.Println(nl.CSVrow)

// NewLinkedRow(csvrow)
// newRow := new(LinkedRow)
// newRow.CSVrow = csvrow
// newRow.Node = node
// newRow.FirstRow = lrow.FirstRow
// newRow.RowsCount = lrow.RowsCount
// curRow.NextRow = newRow
// curRow = newRow
// }
// lrow.FirstRow.RowsCount=
// }
