package internal

import (
	"fmt"
	"log/slog"

	"github.com/antchfx/xmlquery"
)

type LinkedRow struct {
	// ** Double pointer address of address of object. Global variable for list of objects
	// Link internals
	Initialized bool
	RowsCount   *int
	LinkCount   **int
	Start       **LinkedRow
	End         **LinkedRow
	FirstL      *LinkedRow
	LastL       *LinkedRow
	NextL       *LinkedRow
	PrevL       *LinkedRow
	Node        *xmlquery.Node

	// Payload
	CSVrow CSVrow
}

func (l *LinkedRow) NewNextLink(
	parentNode *xmlquery.Node, csvRow CSVrow,
) *LinkedRow {
	newRow := new(LinkedRow)
	newRow.Node = parentNode
	if l == nil || !l.Initialized {
		newRow.FirstL = newRow
		count := 1
		newRow.RowsCount = &count
		newRow.Start = &newRow
		slog.Debug("new link sequence")
		newRow.End = &newRow
	} else {
		slog.Debug("new link add to sequence")
		newRow.Start = l.Start
		newRow.FirstL = l.FirstL
		l.NextL = newRow
		newRow.End = l.End
		*l.End = newRow
	}
	if csvRow != nil {
		newRow.CSVrow = append(newRow.CSVrow, csvRow...)
	}
	newRow.PrevL = l
	newRow.Initialized = true
	slog.Debug("initialized new link")
	return newRow
}

func (l *LinkedRow) GoToNextLink() (*LinkedRow, bool) {
	if l.NextL == nil {
		return l, false
	}
	l = l.NextL
	return l.NextL, true
}

func (l *LinkedRow) GoToStartLink() *LinkedRow {
	l = *l.Start
	return *l.Start
}

func (l *LinkedRow) GoToEndLink() *LinkedRow {
	l = *l.End
	return *l.End
}

func (l *LinkedRow) ReplaceLinkWithLinkSequence(newLink *LinkedRow) *LinkedRow {
	if l == nil {
		slog.Warn("replacing nil link")
		return newLink
	}
	return l
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
	firstRow := new(LinkedRow)
	firstRow = firstRow.NewNextLink(node, nil)
	firstRow.CSVrow = []CSVrowField{{1, "testF1", "valueF1"}}
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
	var nlr *LinkedRow
	for i, node := range nodes {
		// Object fields
		csvrow := NodeToCSVrow(node, ext)
		csvRowJoined := append(parentCsvRow, csvrow...)
		nlr = nlr.NewNextLink(node, csvRowJoined)
		lrows[i] = nlr
	}
	PrintLinks(lrows[len(lrows)-1])
	return lrows
}

func PrintLinks(link *LinkedRow) {
	// TODO: create test instead
	lnk := link
	slog.Debug("printing input")
	fmt.Println(lnk)
	slog.Debug("printing start")
	fmt.Println(lnk.Start)
	start := *lnk.Start
	fmt.Println(start)
	slog.Debug("printing end")
	fmt.Println(*lnk.End)
	slog.Debug("printing by previous")
	count := 10
	for i := 0; i < count; i++ {
		fmt.Println(lnk)
		fmt.Println("start_prev", *lnk.Start)
		fmt.Println("end_prev", *lnk.End)
		lnktest := lnk.PrevL
		if lnktest == nil {
			slog.Debug("sequence end")
			break
		}
		lnk = lnk.PrevL
	}
	slog.Debug("printing by next")
	for i := 0; i < count; i++ {
		fmt.Println(lnk)
		fmt.Println("start_next", *lnk.Start)
		fmt.Println("end_next", *lnk.End)
		lnk = lnk.NextL
		if lnk == nil {
			slog.Debug("sequence end")
			break
		}
	}
}

func PrintLinesSlice(lrows []*LinkedRow) {
	for _, r := range lrows {
		fmt.Println("current", r.CSVrow)
		if r.PrevL != nil {
			fmt.Println("previous", r.PrevL.CSVrow)
		}
	}
}
