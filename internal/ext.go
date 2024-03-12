package internal

import (
	"fmt"

	"github.com/antchfx/xmlquery"
)

type LinkedRow struct {
	FirstRow  *LinkedRow
	NextRow   *LinkedRow
	Node      *xmlquery.Node
	CSVrow    CSVrow
	RowsCount int
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
	var lrow LinkedRow
	lrow = NodeToCSVlinkedRow(node, CSVproduction[0], lrow)
	// fmt.Println(lrow.CSVrow)
	// fmt.Println(lrow.FirstRow.CSVrow)
	return nil
}

func NodeToCSVlinkedRow(node *xmlquery.Node, ext OMobjExtractor, lrow LinkedRow) LinkedRow {
	query := fmt.Sprintf("//OM_OBJECT[@TemplateName='%s']", ext.OmObject)
	if lrow.FirstRow == nil {
		nrow := new(LinkedRow)
		nrow.Node = node
		nrow.RowsCount = 1
		lrow.FirstRow = nrow
	}
	crow := lrow.FirstRow
	for i := 0; i < lrow.FirstRow.RowsCount; i++ {
		nodes := xmlquery.Find(crow.Node, query)
		for _, nd := range nodes {
			fmt.Println(nd.Attr)
		}
		// fmt.Println(i)
		// fmt.Println(nodes[0])
		// crow = &LinkedRow{}
	}

	// for _, n := range nodes {
	// csvrow := NodeToCSVrow(n, ext)
	// lrow.CSVrow = csvrow
	// if lrow.FirstRow == nil {
	// nrow := new(LinkedRow)
	// nrow.CSVrow = csvrow
	// nrow.Node = n
	// }
	// }
	return lrow
}
