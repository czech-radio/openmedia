package internal

import (
	"fmt"
	"log/slog"

	"github.com/antchfx/xmlquery"
)

type LinkedRow struct {
	Initialized bool
	Start       **LinkedRow
	End         **LinkedRow
	FirstL      *LinkedRow
	LastL       *LinkedRow
	NextL       *LinkedRow
	PrevL       *LinkedRow
	Node        *xmlquery.Node
	RowsCount   *int
	CSVrow      CSVrow
	// LinkNumber int ??? without function?
}

// func NewLink(xmlquery.Node,csvRow CSVrow) *LinkedRow {
// }

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
		slog.Debug("new sequence")
		newRow.End = &newRow
	} else {
		slog.Debug("new link")
		newRow.Start = l.Start
		newRow.FirstL = l.FirstL
		l.NextL = newRow
	}
	if csvRow != nil {
		newRow.CSVrow = append(newRow.CSVrow, csvRow...)
	}
	// *newRow.End = &newRow
	newRow.PrevL = l
	newRow.Initialized = true
	slog.Debug("initialized")
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
	// newLink.Start
	// newLink.
	// newLink.
	// firstl := l.FirstL
	// prevl := l.PrevL
	// nextl := l.NextL
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
		// slog.Debug("nodes found", "count", len(nodes))
		parentCsvRow := curL.CSVrow
		lrows := NodesExtractFieldsToRows(parentCsvRow, nodes, ext)
		slog.Debug("rows", "count", len(lrows))
		// start := *lrows[1].Start
		// fmt.Println("starter", start.CSVrow)
		// fmt.Println("six: end", lrows[0].GoToEndLink().CSVrow)
		// fmt.Println("six: end", lrows[0].GoToEndLink())
		// Checkout next link
		curL = curL.NextL
		if curL == nil {
			break
		}
	}
	// fmt.Println("six", lrow.Start.End)
	// start = lrow.GoToStartLink()
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
	// fmt.Println(lrows[3])
	// fmt.Println(lrows[3].PrevL)
	// fmt.Println(lrows[3].PrevL.PrevL)
	PrintLinks(lrows[len(lrows)-1])
	// PrintLinesSlice(lrows)
	return lrows
}

func PrintLinks(link *LinkedRow) {
	lnk := link
	slog.Debug("printing input")
	fmt.Println(lnk)
	slog.Debug("printing start")
	fmt.Println(lnk.Start)
	slog.Debug("printing end")
	fmt.Println(lnk.End)
	slog.Debug("printing by previous")
	count := 10
	for i := 0; i < count; i++ {
		fmt.Println(lnk)
		lnktest := lnk.PrevL
		if lnktest == nil {
			slog.Debug("sequence end")
			break
		}
		lnk = lnk.PrevL
	}
	// lnk = link
	slog.Debug("printing by next")
	for i := 0; i < count; i++ {
		fmt.Println(lnk)
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
