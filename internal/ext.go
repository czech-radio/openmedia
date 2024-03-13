package internal

import (
	"fmt"
	"log/slog"

	"github.com/antchfx/xmlquery"
)

type LinkedRow struct {
	// Link internals
	// ** Double pointer address of address of object. Shared accross all links
	LinksCount **int
	Start      **LinkedRow
	End        **LinkedRow

	// Local variables to one link
	RowsCount   *int
	FirstL      *LinkedRow
	LastL       *LinkedRow
	NextL       *LinkedRow
	PrevL       *LinkedRow
	Initialized bool

	Payload LinkPayload
}

type LinkPayload struct {
	OmObject string
	NodePath string
	Node     *xmlquery.Node
	CSVrow   CSVrow
}

func (l *LinkedRow) NewNextLink(
	payload LinkPayload) *LinkedRow {
	newRow := new(LinkedRow)
	if l == nil || !l.Initialized {
		// Not initialized
		newRow.FirstL = newRow
		count := 1
		newRow.RowsCount = &count
		newRow.Start = &newRow
		// slog.Debug("new link sequence")
		newRow.End = &newRow
	} else {
		// Initialized
		// slog.Debug("new link add to sequence")
		newRow.Start = l.Start
		newRow.FirstL = l.FirstL
		l.NextL = newRow
		newRow.End = l.End
		*l.End = newRow
		newRow.PrevL = l
	}
	newRow.Payload = payload
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

func (l *LinkedRow) ExportCurrentLinkToCSV() {
}

func (l *LinkedRow) ExportAllLinksToCSV() {
}

func (l *LinkedRow) ReplaceLinkWithLinkSequence(newLink *LinkedRow) *LinkedRow {
	if l == nil {
		slog.Warn("replacing nil link")
		return newLink
	}
	var out *LinkedRow
	nlend := *newLink.End
	if l.PrevL == nil && l.NextL == nil {
		slog.Warn("replacing one link")
		return nlend
	}

	nlstart := *newLink.Start
	if l.PrevL == nil {
		slog.Debug("Replacing first link of links sequence")
		*l.Start = nlstart
		nlend.NextL = l.NextL
		*nlend.End = *l.End
		return nlend
	}

	if l.NextL == nil {
		slog.Debug("Replacing last link in sequence")
		lend := *l.End
		lend.NextL = nlstart
	}
	slog.Debug("Replacing link in the middle of links sequence")

	// Replace start and end in newLink sequence with from 'l'
	// lstart := *l.Start
	// lend := *l.End
	// nlstart := *newLink.Start
	// nend := newLink.End
	// *newLink.Start = lstart
	// *newLink.End = lend
	// newLink.Prev = l.PrevL
	// 1. replace prevl in newLink.Start
	// 2. replace nextL in newLink.End
	// 3. replace 'l' with new sequnce
	return out
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

	// First link construct
	firstRow := new(LinkedRow)
	payload := LinkPayload{
		NodePath: CSVproduction[1].Path,
		Node:     node,
		CSVrow:   CSVrow{{1, "testF1", "valueF1"}},
	}
	firstRow = firstRow.NewNextLink(payload)
	firstRow.Payload.CSVrow = []CSVrowField{{1, "testF1", "valueF1"}}

	// Expand rows
	// rows := firstRow.ExtractOMobjectsFields(CSVproduction[0])
	rows := firstRow.ExtractOMobjectsFields(CSVproduction[1])
	PrintLinks(rows)
	return nil
}

func (l *LinkedRow) ExtractOMobjectsFields(ext OMobjExtractor) *LinkedRow {
	query := fmt.Sprintf("//OM_OBJECT[@TemplateName='%s']", ext.OmObject)
	if l == nil {
		slog.Warn("querying nil node")
		return l
	}
	curL := l
	var index int

	var lrows []*LinkedRow
	var nodes []*xmlquery.Node
	var parentCsvRow CSVrow
	for {
		index++
		// Check the type of node
		if l.Payload.NodePath != ext.Path {
			slog.Warn("skipping node")
			goto checkNextLink
		}

		// Find objects
		nodes = xmlquery.Find(curL.Payload.Node, query)
		parentCsvRow = curL.Payload.CSVrow
		lrows = NodesExtractFieldsToRows(parentCsvRow, nodes, ext)
		slog.Debug("rows", "count", len(lrows))
		if len(lrows) > 0 {
			slog.Debug("replacing link", "count", len(lrows))
			curL = curL.ReplaceLinkWithLinkSequence(lrows[0])
		}
		// Checkout next link
	checkNextLink:
		slog.Debug("checking after index", "index", index)
		check := curL.NextL
		if check == nil {
			slog.Debug("breaking after index", "index", index)
			break
		}
		curL = curL.NextL
	}
	return curL
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
		payload := LinkPayload{
			// NodePath: ext.OmObject,
			OmObject: ext.OmObject,
			NodePath: ext.Path,
			Node:     node,
			CSVrow:   csvRowJoined,
		}
		nlr = nlr.NewNextLink(payload)
		nlr.Payload.CSVrow = csvRowJoined
		lrows[i] = nlr
	}
	return lrows
}

// func NodeExtractFields() *LinkedRow {
// }

func PrintLinks(link *LinkedRow) {
	// TODO: create test instead
	lnk := link
	fmt.Println()
	slog.Debug("printing input")
	fmt.Println(lnk)
	slog.Debug("printing start")
	fmt.Println(lnk.Start)
	start := *lnk.Start
	fmt.Println(start)
	slog.Debug("printing end")
	fmt.Println(*lnk.End)
	fmt.Println()
	slog.Debug("printing by previous")
	count := 10
	for i := 0; i < count; i++ {
		fmt.Println(i, "seq_prev", lnk)
		fmt.Println(i, "start_prev", *lnk.Start)
		fmt.Println(i, "end_prev", *lnk.End)
		lnktest := lnk.PrevL
		if lnktest == nil {
			slog.Debug("sequence end")
			break
		}
		lnk = lnk.PrevL
	}
	fmt.Println()
	slog.Debug("printing by next")
	for i := 0; i < count; i++ {
		fmt.Println(i, "seq_next", lnk)
		fmt.Println(i, "start_next", *lnk.Start)
		fmt.Println(i, "end_next", *lnk.End)
		lnk = lnk.NextL
		if lnk == nil {
			slog.Debug("sequence end")
			break
		}
	}
}

func PrintLinesSlice(lrows []*LinkedRow) {
	for _, r := range lrows {
		fmt.Println("current", r.Payload.CSVrow)
		if r.PrevL != nil {
			fmt.Println("previous", r.PrevL.Payload.CSVrow)
		}
	}
}
