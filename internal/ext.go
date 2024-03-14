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

	// Payload
	Payload LinkPayload
}

type LinkPayload struct {
	OmObject string
	NodePath string
	Node     *xmlquery.Node
	CSVrow   CSVrow
	Index    int
	IndexStr string
	ExtCount int
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
	slog.Debug("new link sequence created")
	return newRow
}

func (l *LinkedRow) NextLinkAdd(payload LinkPayload) *LinkedRow {
	if l == nil {
		res := NewLinkSequence(payload)
		*res.End = res
		return res
	}
	newRow := new(LinkedRow)
	newRow.Start = l.Start
	newRow.FirstL = l.FirstL
	l.NextL = newRow
	newRow.End = l.End
	*l.End = newRow
	newRow.PrevL = l
	newRow.Payload = payload
	slog.Debug("next link created in sequence")
	return newRow
}

func (l *LinkedRow) ReplaceLinkWithLinkSequence(
	newLink *LinkedRow) *LinkedRow {
	// PrintLinks("KEK", newLink)
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
		slog.Warn("Replacing first link of links sequence")
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
	var firstRow *LinkedRow
	payload := LinkPayload{
		NodePath: "", // [0]
		// NodePath: "/Radio Rundown", // [1]
		Node:   node,
		CSVrow: CSVrow{{1, "testF1", "valueF1"}},
	}
	startLV := 0
	levels := 1
	firstRow = firstRow.NextLinkAdd(payload)
	for i := startLV; i <= levels; i++ {
		slog.Debug("extracting object", "number", i)
		// firstRow = firstRow.NextLinkAdd(payload)
		firstRow = firstRow.ExtractOMobjectsFields(CSVproduction[i])
		// firstRow = firstRow.ExtractOMobjectsFields(CSVproduction[1])
		// firstRow = firstRow.ExtractOMobjectsFields(CSVproduction[2])
	}
	PrintLinks("LAST RESULT", firstRow)
	return nil
}

func (l *LinkedRow) ExtractOMobjectsFields(ext OMobjExtractor) *LinkedRow {
	// Chekcs
	if l == nil {
		slog.Warn("querying nil node")
		return l
	}
	if l.Payload.NodePath != ext.Path {
		slog.Warn("skipping extraction", "ext", ext)
		slog.Warn("skipping based on", "nodePath", l.Payload.NodePath, "extPath", ext.Path)
	}
	var index int
	var lrows *LinkedRow
	var nodes []*xmlquery.Node
	var parentCsvRow CSVrow
	curL := l
	query := fmt.Sprintf("//OM_OBJECT[@TemplateName='%s']", ext.OmObject)
	// var tak *LinkedRow
	for {
		index++
		// Check the type of node
		// if l.Payload.NodePath != ext.Path {
		// slog.Warn("skipping extraction", "ext", ext)
		// slog.Warn("skipping based on", "nodePath", l.Payload.NodePath, "extPath", ext.Path)
		// goto checkNextLink
		// }
		// Find objects
		slog.Warn("extracting", "ext", ext)
		nodes = xmlquery.Find(curL.Payload.Node, query)
		parentCsvRow = curL.Payload.CSVrow
		if len(nodes) == 0 {
			slog.Warn("subnodes not found")
			goto checkNextLink
		}
		slog.Debug("subnodes found", "count", len(nodes))
		lrows = NodesExtractFieldsToRows(parentCsvRow, nodes, ext)
		curL = curL.ReplaceLinkWithLinkSequence(*lrows.End)

	checkNextLink:
		// Checkout next link
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
) *LinkedRow {
	var nlr *LinkedRow
	slog.Debug("finding fields for obj", "count", len(nodes))
	for i, node := range nodes {
		// Object fields
		csvrow := NodeToCSVrow(node, ext)

		csvRowJoined := append(parentCsvRow, csvrow...)
		payload := LinkPayload{
			OmObject: ext.OmObject,
			NodePath: JoinObjectPath(ext.Path, ext.OmObject),
			Node:     node,
			CSVrow:   csvRowJoined,
			Index:    i,
			ExtCount: len(nodes),
		}
		nlr = nlr.NextLinkAdd(payload)
		nlr.Payload.Index = i
	}
	PrintLinks("SEX", nlr)
	return nlr
}
