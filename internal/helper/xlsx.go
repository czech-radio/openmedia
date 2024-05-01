package helper

import (
	"github.com/xuri/excelize/v2"
)

// ReadExcelFileSheetRows
func ReadExcelFileSheetRows(filePath, sheetName string) (
	rows [][]string, err error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			rows = nil
		}
	}()
	// cell, err := f.GetCellValue("Sheet1", "B2")
	// Get all the rows in the Sheet1.
	return f.GetRows(sheetName)
}

// MapExcelSheetColumn reads specified specified excel file sheet and creates map from the specified column. Useful to check whether column contains specific value(s).
func MapExcelSheetColumn(
	filePath, sheetName string, columnNumber int,
) (map[string]bool, error) {
	res := make(map[string]bool)
	rows, err := ReadExcelFileSheetRows(filePath, sheetName)
	if err != nil {
		return nil, err
	}
	for i, row := range rows {
		if i < 1 {
			// omit header
			continue
		}
		res[row[columnNumber]] = true
	}
	return res, nil
}

type Table struct {
	Rows                 [][]string
	RowHeaderToColumnMap map[string][]string
	ColumnHeaderMap      map[string]int
	ColumnHeader         []string
}

func CreateTable(rows [][]string,
	columnHeaderRow, rowHeaderColumn int) *Table {
	table := new(Table)
	table.MapTableHeaders(rows, columnHeaderRow, rowHeaderColumn)
	table.Rows = rows[columnHeaderRow+1:][rowHeaderColumn+1:]
	return table
}

func (t *Table) MapTableHeaders(
	rows [][]string, columnsHeaderRow, rowsHeaderColumn int) {
	t.RowHeaderToColumnMap = make(map[string][]string)
	// Map rows
	r := rows
	i := columnsHeaderRow + 1
	for k := i; k < len(r); k++ {
		t.RowHeaderToColumnMap[r[k][rowsHeaderColumn]] = r[k][rowsHeaderColumn+1:]
	}

	t.ColumnHeaderMap = make(map[string]int)
	// Map columns header columnName vs position
	for j, val := range r[rowsHeaderColumn] {
		t.ColumnHeaderMap[val] = j
	}
}

func (t *Table) MatchRow(
	rowHeaderValue, columnName, columnValue string) bool {
	row, ok := t.RowHeaderToColumnMap[rowHeaderValue]
	if !ok {
		return false
	}
	_ = row
	columnIndex := t.ColumnHeaderMap[columnName] - 1
	// slog.Warn("fuck", "name", rowHeaderValue, "colname", columnName, "colindex", columnIndex, "row", row)
	colVal := row[columnIndex]
	// slog.Warn("fuck", "rowval", rowHeaderValue, "val", value, "row", row)
	return colVal == columnValue
}
