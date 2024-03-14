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
	Index    int
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

func NewLinkSequence(payload LinkPayload) *LinkedRow {
	newRow := new(LinkedRow)
	newRowE := new(LinkedRow) // Just for another random addr
	newRow.Start = &newRow
	newRow.End = &newRowE
	newRow.Payload = payload
	return newRow
}

func (l *LinkedRow) NextLinkAdd(payload LinkPayload) *LinkedRow {
	newRow := new(LinkedRow)
	newRow.Start = l.Start
	newRow.FirstL = l.FirstL
	l.NextL = newRow
	newRow.End = l.End
	*l.End = newRow
	newRow.PrevL = l
	newRow.Payload = payload
	return newRow
}

func (l *LinkedRow) NewNextLink(
	payload LinkPayload) *LinkedRow {
	newRow := new(LinkedRow)
	if l == nil || !l.Initialized {
		// if l == nil {
		// Not initialized
		slog.Debug("initializing first link in sequence")
		newRow.FirstL = newRow
		count := 1
		newRow.RowsCount = &count
		// newRow.Start = nil
		// fmt.Println("krax", *newRow.Start)
		// *newRow.Start = newRow
		newRow.Start = &newRow
		newRow.End = &newRow
		newRow.Payload = payload
	} else {
		// Initialized
		slog.Debug("initializing new link in sequence")
		newRow.Start = l.Start
		// *newRow.Start = *l.PrevL.Start
		newRow.FirstL = l.FirstL
		l.NextL = newRow
		newRow.End = l.End
		*l.End = newRow
		newRow.PrevL = l
	}
	newRow.Payload = payload
	newRow.Initialized = true
	return newRow
}

// func KerujDebug() {
// fmt.Println("keruj", newRow.Start, l.Start)
// fmt.Println("keruj", *newRow.Start)
// chkn := *newRow.Start
// chkl := *l.Start
// fmt.Println("keruj payload", chkn.Payload, chkl.Payload)
// }

func (l *LinkedRow) ReplaceLinkWithLinkSequence(
	newLink *LinkedRow) *LinkedRow {
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
		// NodePath: "", // [0]
		NodePath: "/Radio Rundown", // [1]
		Node:     node,
		CSVrow:   CSVrow{{1, "testF1", "valueF1"}},
	}
	firstRow = firstRow.NewNextLink(payload)
	chk := *firstRow.Start
	fmt.Println("keruje", chk)
	firstRow.Payload.CSVrow = []CSVrowField{{1, "testF1", "valueF1"}}

	// Expand rows
	// slog.Debug("fist object")
	// firstRow = firstRow.ExtractOMobjectsFields(CSVproduction[0])
	slog.Debug("second object")
	firstRow = firstRow.ExtractOMobjectsFields(CSVproduction[1])
	// slog.Debug("third object")
	// firstRow = firstRow.ExtractOMobjectsFields(CSVproduction[2])
	PrintLinks("LAST RESULT", firstRow)
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
			slog.Warn("skipping extraction", "ext", ext)
			slog.Warn("skipping based on", "nodePath", l.Payload.NodePath, "extPath", ext.Path)
			goto checkNextLink
		}
		slog.Warn("extracting", "ext", ext)

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
	return *curL.Start
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
			OmObject: ext.OmObject,
			NodePath: JoinObjectPath(ext.Path, ext.OmObject),
			Node:     node,
			CSVrow:   csvRowJoined,
		}
		nlr = nlr.NewNextLink(payload)
		nlr.Payload.CSVrow = csvRowJoined
		nlr.Payload.Index = i
		lrows[i] = nlr
	}

	PrintLinks("FUCK", lrows[0])
	PrintLinks("FUCK1", lrows[1])
	return lrows
}

func PrintLinks(name string, link *LinkedRow) {
	// TODO: create test instead
	lnk := link
	fmt.Println()
	slog.Debug(name, "output", "printing input")
	fmt.Println(name, lnk)
	slog.Debug(name, "output", "printing start")
	fmt.Println(name, lnk.Start)
	start := *lnk.Start
	fmt.Println(name, "start", start)
	slog.Debug(name, "output", "printing end")
	fmt.Println(name, "end", *lnk.End)
	fmt.Println()
	slog.Debug(name, "output", "printing by previous")
	count := 10
	for i := 0; i < count; i++ {
		fmt.Println(name, i, "seq_prev", lnk)
		fmt.Println(name, i, "start_prev", *lnk.Start)
		fmt.Println(name, i, "end_prev", *lnk.End)
		lnktest := lnk.PrevL
		if lnktest == nil {
			slog.Debug(name, "output", "sequence end")
			break
		}
		lnk = lnk.PrevL
	}
	fmt.Println()
	slog.Debug(name, "output", "printing by next")
	for i := 0; i < count; i++ {
		fmt.Println(name, i, "seq_next", lnk)
		fmt.Println(name, i, "start_next", *lnk.Start)
		fmt.Println(name, i, "end_next", *lnk.End)
		lnk = lnk.NextL
		if lnk == nil {
			slog.Debug(name, "output", "sequence end")
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
