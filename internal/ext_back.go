package internal

import (
	"fmt"
	"log/slog"

	"github.com/antchfx/xmlquery"
)

// NOTE: Deprecated

func PrintLinks(name string, link *LinkedRow) {
	// TODO: create test instead
	lnk := link
	fmt.Println()
	slog.Debug(name, "output", "printing input")
	fmt.Println(name, lnk)
	slog.Debug(name, "output", "printing start")
	fmt.Println(name, lnk.Start)
	start := *lnk.Start
	end := *lnk.End
	fmt.Println(name, "start", start)
	slog.Debug(name, "output", "printing end")
	fmt.Println(name, "end", *lnk.End)
	fmt.Println()
	count := 10
	lnk = start
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
	fmt.Println()
	slog.Debug(name, "output", "printing by previous")
	lnk = end
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
}

func NodeToCSVrow(node *xmlquery.Node, ext OMobjExtractor) CSVrow {
	var csvrow CSVrow
	// query := ext.FieldsPath + BuildFieldsQuery(ext.FieldIDs)
	query := ext.FieldsPath + XMLbuildAttrQuery("FieldID", ext.FieldIDs)
	fields := xmlquery.Find(node, query)
	if len(fields) == 0 {
		slog.Error("nothing found")
		return csvrow
	}
	for _, f := range fields {
		fieldID, _ := GetFieldValueByID(f.Attr, "FieldID")
		field := CSVrowField{
			FieldPosition: 0,
			FieldID:       fieldID,
			Value:         f.InnerText(),
		}
		csvrow = append(csvrow, field)
	}
	return csvrow
}

func PrintLinesSlice(lrows []*LinkedRow) {
	for _, r := range lrows {
		fmt.Println("current", r.Payload.CSVrow)
		if r.PrevL != nil {
			fmt.Println("previous", r.PrevL.Payload.CSVrow)
		}
	}
}

func PrintRows(rows map[int]CSVrow) {
	for i := 0; i < len(rows); i++ {
		fmt.Println(i, rows[i])
		fmt.Println()
	}
}

func FindSubNodes(node *xmlquery.Node, ext OMobjExtractor) []*xmlquery.Node {
	query := fmt.Sprintf("//OM_OBJECT[@TemplateName='%s']", ext.OmObject)
	return xmlquery.Find(node, query)
}

func NodesToCSVrows(nodes []*xmlquery.Node, ext OMobjExtractor, rows CSVrowsIntMap) CSVrowsIntMap {
	if len(rows) == 0 {
		// rows = make(map[int]CSVrow, len(nodes))
		rows = make(CSVrowsIntMap, len(nodes))
	}
	for i, node := range nodes {
		row := NodeToCSVrow(node, ext)
		rows[i] = append(rows[i], row...)
	}
	return rows
}
