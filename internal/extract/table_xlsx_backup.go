package extract

import (
	"fmt"
	"path/filepath"
	"strconv"
	"time"

	"github.com/triopium/go_utils/pkg/helper"
	"github.com/xuri/excelize/v2"
)

func XlsxFormatColumnsInDir(folder string) error {
	files, err := helper.ListDirFiles(folder)
	if err != nil {
		return err
	}
	for _, file := range files {
		ext := filepath.Ext(file)
		if ext != ".xlsx" {
			continue
		}
		err := XlsxFormatColumns(file)
		if err != nil {
			return err
		}
	}
	return nil
}

func XlsxFormatColumns(xlsxFile string) error {
	f, err := excelize.OpenFile(xlsxFile)
	if err != nil {
		return err
	}
	sheetName := "Sheet1"
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return err
	}
	if len(rows) == 0 {
		return nil
	}
	// Set column widths
	lastCol, err := excelize.ColumnNumberToName(len(rows[0]))
	if err != nil {
		return err
	}
	if err := f.SetColWidth(sheetName, "A", lastCol, 5); err != nil {
		return err
	}

	// Set column formats
	// table.RowHeaderToColumnMap
	if len(rows) < 2 {
		return fmt.Errorf("no rows to convert")
	}
	headerRow := 0
	for i := 0; i < len(rows[0]); i++ {
		fieldID, ok := FieldsIDsNames.GetByName(rows[0][i])
		if !ok {
			fieldID, ok = FieldsIDsNames.GetByName(rows[1][i])
			headerRow = 1
		}
		if !ok {
			continue
		}

		// column properties
		cname, err := excelize.ColumnNumberToName(i + 1)
		if err != nil {
			return err
		}
		firstCell := fmt.Sprintf("%s%d", cname, headerRow+1)
		lastCell := fmt.Sprintf("%s%d", cname, len(rows))

		// default column width
		if fieldID.Width > -1 {
			if err := f.SetColWidth(sheetName, cname, cname, fieldID.Width); err != nil {
				return err
			}
		}

		// Column format number
		styleDef := excelize.Style{}
		if fieldID.XLSXcolumnFormat != 0 {
			styleDef.NumFmt = fieldID.XLSXcolumnFormat
		}

		// Custom format
		if fieldID.XLSXcustomFormat != "" {
			styleDef.CustomNumFmt = &fieldID.XLSXcustomFormat
		}

		// exp := "[$-380A]dddd\\,\\ dd\" de \"mmmm\" de \"yyyy;@"
		// style, err := f.NewStyle(&excelize.Style{CustomNumFmt: &exp})

		// style, err := f.NewStyle(&excelize.Style{
		// 	Font: &excelize.Font{Bold: true},
		// },

		// Write style to column
		style, err := f.NewStyle(&styleDef)
		if err != nil {
			return err
		}
		err = f.SetCellStyle(sheetName, firstCell, lastCell, style)
		if err != nil {
			return err
		}
	}
	return f.Save()
}

func (e *Extractor) NewStyle() excelize.Style {
	st := excelize.Style{
		Border: []excelize.Border{},
		Fill: excelize.Fill{
			Type:    "",
			Pattern: 0,
			Color:   []string{},
			Shading: 0,
		},
		Font: &excelize.Font{
			Bold:         false,
			Italic:       false,
			Underline:    "",
			Family:       "",
			Size:         0.0,
			Strike:       false,
			Color:        "",
			ColorIndexed: 0,
			ColorTheme:   nil,
			ColorTint:    0.0,
			VertAlign:    "",
		},
		Alignment: &excelize.Alignment{
			Horizontal:      "",
			Indent:          0,
			JustifyLastLine: false,
			ReadingOrder:    0,
			RelativeIndent:  0,
			ShrinkToFit:     false,
			TextRotation:    0,
			Vertical:        "",
			WrapText:        false,
		},
		Protection: &excelize.Protection{
			Hidden: false,
			Locked: false,
		},
		NumFmt:        0,
		DecimalPlaces: nil,
		CustomNumFmt:  nil,
		NegRed:        false,
	}
	return st
}

func XlsxStyle() {
	_ = excelize.Style{
		Border: []excelize.Border{},
		Fill: excelize.Fill{
			Type:    "",
			Pattern: 0,
			Color:   []string{},
			Shading: 0,
		},
		Font: &excelize.Font{
			Bold:         false,
			Italic:       false,
			Underline:    "",
			Family:       "",
			Size:         0.0,
			Strike:       false,
			Color:        "",
			ColorIndexed: 0,
			ColorTheme:   nil,
			ColorTint:    0.0,
			VertAlign:    "",
		},
		Alignment: &excelize.Alignment{
			Horizontal:      "",
			Indent:          0,
			JustifyLastLine: false,
			ReadingOrder:    0,
			RelativeIndent:  0,
			ShrinkToFit:     false,
			TextRotation:    0,
			Vertical:        "",
			WrapText:        false,
		},
		Protection: &excelize.Protection{
			Hidden: false,
			Locked: false,
		},
		NumFmt:        0,
		DecimalPlaces: nil,
		CustomNumFmt:  nil,
		NegRed:        false,
	}
}

func CSVtoXLSXB(csvFile string, csvDelim rune) error {
	// read csv
	records, err := CSVreadRows(csvFile, csvDelim)
	if err != nil {
		return err
	}
	if len(records) < 2 {
		return nil
	}

	// Create a new Excel file
	xlsxFile := excelize.NewFile()
	sheetName := "Sheet1"
	// Create a new sheet
	index, err := xlsxFile.NewSheet(sheetName)
	if err != nil {
		return err
	}
	xlsxFile.SetActiveSheet(index)

	// Set style
	lastCol, err := excelize.ColumnNumberToName(len(records[0]))
	if err != nil {
		return err
	}
	if err := xlsxFile.SetColWidth(sheetName, "A", lastCol, 5); err != nil {
		return err
	}
	err = XLSXsetColumnStyle(xlsxFile, sheetName, records)
	if err != nil {
		return err
	}

	for i, row := range records {
		for j, cell := range row {
			cellRef, _ := excelize.CoordinatesToCellName(j+1, i+1)
			err := xlsxFile.SetCellValue(sheetName, cellRef, cell)
			if err != nil {
				return err
			}

			// styleDef := excelize.Style{}
			// // st := "DD MM R"
			// st := "@"
			// styleDef.CustomNumFmt = &st
			// nst, err := xlsxFile.NewStyle(&styleDef)
			// if err != nil {
			// 	return err
			// }
			// err = xlsxFile.SetCellStyle(sheetName, cellRef, cellRef, nst)
			// if err != nil {
			// 	return err
			// }
		}
	}
	// err = XLSXsetColumnStyleText(xlsxFile, sheetName, records)
	// if err != nil {
	// return err
	// }
	name := helper.FilenameWithoutExtension(csvFile)
	dir := filepath.Dir(csvFile)
	xlsxFilePath := filepath.Join(dir, name+".xlsx")
	return xlsxFile.SaveAs(xlsxFilePath)
}

func XLSXsetColumnStyleText(
	f *excelize.File, sheetName string, rows [][]string) error {
	headerRow := 0
	for i := 0; i < len(rows[0]); i++ {
		_, ok := FieldsIDsNames.GetByName(rows[0][i])
		if !ok {
			_, ok = FieldsIDsNames.GetByName(rows[1][i])
			headerRow = 1
		}
		if !ok {
			continue
		}
		// column properties
		cname, err := excelize.ColumnNumberToName(i + 1)
		if err != nil {
			return err
		}
		firstCell := fmt.Sprintf("%s%d", cname, headerRow+1)
		lastCell := fmt.Sprintf("%s%d", cname, len(rows))

		// Column format number
		styleDef := excelize.Style{}
		styleDef.NumFmt = 0

		// Custom format
		stext := "@"
		styleDef.CustomNumFmt = &stext

		// Write style to column
		style, err := f.NewStyle(&styleDef)
		if err != nil {
			return err
		}

		err = f.SetCellStyle(sheetName, firstCell, lastCell, style)
		if err != nil {
			return err
		}
	}
	return nil
}

func XLSXsetCellValue(
	file *excelize.File, sheetName string,
	fd FieldID, cellRef, cell string) error {
	switch fd.XLSXcolumnFormat {
	case 14:
		date, err := ParseXMLdate(cell)
		if err == nil {
			return file.SetCellValue(sheetName, cellRef, date)
		}
		date, err = time.Parse("2006-01-02 15:04:05", cell)
		if err != nil {
			goto text
		}
		return file.SetCellValue(sheetName, cellRef, date)
	case 1:
		res, err := strconv.Atoi(cell)
		if err != nil {
			goto text
		}
		return file.SetCellValue(sheetName, cellRef, res)
	}
text:
	return file.SetCellValue(sheetName, cellRef, cell)
}
