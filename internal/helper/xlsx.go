package helper

import "github.com/xuri/excelize/v2"

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
	RowHeader       []string
	RowHeaderMap    map[string]int // Name vs position
	ColumnHeader    []string
	ColumnHeaderMap map[string]int
	Rows            [][]string
}

func MapExcelTable(
	filePath, sheetName string, headerRow, headerColumn int,
) (*Table, error) {
	rows, err := ReadExcelFileSheetRows(filePath, sheetName)
	if err != nil {
		return nil, err
	}
	table := new(Table)
	table.RowHeader = rows[headerRow][headerColumn:]
	//

	return nil, nil
}
