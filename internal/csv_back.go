package internal

import (
	"fmt"
	"log/slog"
	"strings"
)

func (e *Extractor) CSVheaderPrint() {
	fmt.Println(e.CSVheaderInternal)
	fmt.Println(e.CSVheaderExternal)
}

func (e *Extractor) PrintTableRowsToCSV(
	header bool, delim string, rowsIndexes ...[]int) {
	var sb strings.Builder
	// Print header
	if header {
		e.CSVheaderPrint()
	}

	if len(rowsIndexes) > 1 {
		slog.Error("not implemented multiple indexes' slices")
	}
	var count int
	// Print specified rows
	if len(rowsIndexes) == 1 {
		for _, index := range rowsIndexes[0] {
			// e.CSVtable[index].PrintRowToCSV(
			e.CSVtable.Rows[index].CastToCSV(
				&sb, e.CSVrowPartsPositionsInternal,
				e.CSVrowPartsFieldsPositions,
				delim,
			)
			count++
		}
		fmt.Print(sb.String())
		slog.Debug("lines printed", "count", count)
		return
	}

	// Print whole table
	for _, row := range e.CSVtable.Rows {
		row.CastToCSV(
			&sb, e.CSVrowPartsPositionsInternal,
			e.CSVrowPartsFieldsPositions,
			delim,
		)
		count++
	}
	fmt.Print(sb.String())
	slog.Debug("lines printed", "count", count)
}
