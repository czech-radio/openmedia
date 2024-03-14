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

func JoinLinksSequences(A, B *LinkedRow) *LinkedRow {
	A.NextL = *B.End
	B.PrevL = *A.End
	*A.End = *B.End
	*B.Start = *A.Start
	return nil
}

func CrossLinkSequences(A, B *LinkedRow) *LinkedRow {
	// Must define intersection of links
	return nil
}

func (l *LinkedRow) ReplaceLinkWithLinkSequence(
	newLink *LinkedRow) *LinkedRow {
	if l == nil {
		slog.Warn("replacing nil link")
		return newLink
	}
	nlend := *newLink.End

	if l.PrevL == nil && l.NextL == nil {
		slog.Warn("replacing one link")
		return nlend
	}

	if l.PrevL == nil {
		slog.Warn("replacing the first link of the links sequence")
		*l.Start = *newLink.Start    // Replace start link in all original links
		l.NextL.PrevL = *newLink.End // Go to next current l-link and replace its previous link with last of newLink. (Effectively omiting current l-link)
		*newLink.End = *l.End        // Rplace end link in all new links
		nlend.NextL = l.NextL        // Join current l.link to the end of newLink
		return nlend
	}

	if l.NextL == nil {
		slog.Debug("replacing the last link in sequence")
		// PrintLinks("KUKD", newLink)
		return newLink
	}

	slog.Debug("replacing link in the middle of links sequence")
	nlstart := *newLink.Start
	nlstart.PrevL = l.PrevL
	nlend.NextL = l.NextL
	// PrintLinks("KUKCn", newLink)
	return newLink
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
		Node: node,
		// CSVrow: CSVrow{{1, "testF1", "valueF1"}},
	}
	startLV := 0
	levels := 3
	firstRow = firstRow.NextLinkAdd(payload)
	for i := startLV; i < levels; i++ {
		slog.Debug("extracting object", "number", i, "object", CSVproduction[i].OmObject)
		firstRow = firstRow.ExtractOMobjectsFields(CSVproduction[i])
		firstRow = *firstRow.Start
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
	var index int
	var lrows *LinkedRow
	var nodes []*xmlquery.Node
	var parentCsvRow CSVrow
	loopSequenceLink := l
	query := fmt.Sprintf("//OM_OBJECT[@TemplateName='%s']", ext.OmObject)
	// Maybe rewrite with appending resulting sequences to new sequence.
	for {
		index++
		if loopSequenceLink.Payload.NodePath != ext.Path {
			// Check the type of node
			slog.Warn("skipping extraction", "ext", ext)
			slog.Warn("skipping based on", "nodePath", l.Payload.NodePath, "extPath", ext.Path)
			goto checkNextLink
		}
		// Find objects
		slog.Warn("extracting", "ext", ext)
		nodes = xmlquery.Find(loopSequenceLink.Payload.Node, query)
		parentCsvRow = loopSequenceLink.Payload.CSVrow

		if len(nodes) == 0 {
			slog.Warn("subnodes not found")
			goto checkNextLink
		}
		slog.Debug("subnodes found", "count", len(nodes))
		lrows = NodesExtractFieldsToRows(parentCsvRow, nodes, ext)
		loopSequenceLink = loopSequenceLink.ReplaceLinkWithLinkSequence(*lrows.End)

	checkNextLink:
		// Checkout next link
		slog.Debug("checking after index", "index", index)
		check := loopSequenceLink.NextL
		if check == nil {
			slog.Debug("breaking after index", "index", index)
			break
		}
		loopSequenceLink = loopSequenceLink.NextL
	}
	return loopSequenceLink
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
	return nlr // Last link of sequence
}
