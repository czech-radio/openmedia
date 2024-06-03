package extract

import (
	"encoding/csv"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/triopium/go_utils/pkg/helper"
	"github.com/xuri/excelize/v2"
)

// TODO: treat maximum rows when exporting to xlsx

// f.NewStyle()
// f.NewConditionalStyle()
// f.SetColStyle(sheetName, column, styleId)
// f.SetColWidth()
// f.SetCellStyle()
// sw.SetRow()
// func (f *File) SetDefaultFont(fontName string)
// func (f *File) GetStyle(idx int) (*Style, error)
// err = f.SetColStyle("Sheet1", "H", style)
// // err = f.SetRowStyle
// err = f.SetColWidth(sheetName, "A", "B", 90)

func (e *Extractor) XLSXstreamTableSave(
	filePath, sheetName string, overWrite bool,
	internalHeader, externalHeader, columnFormat bool) (lastRow int, result error) {
	el := helper.ErrList{}
	defer el.Handle(&result)
	var err error

	// Add data
	f, fileClose, err := XLSXopenFile(filePath, overWrite)
	el.ErrorRaise(err)
	defer fileClose(f)
	sw, err := f.NewStreamWriter(sheetName) // will overwrite sheet
	el.ErrorRaise(err)
	lastRow, err = e.XLSXstreamAddHeader(sw, internalHeader, externalHeader)
	el.ErrorRaise(err)
	lastRow, err = e.XLSXstreamAddRows(sw, lastRow)
	el.ErrorRaise(err)
	err = sw.Flush()
	el.ErrorRaise(err)
	err = f.SaveAs(filePath)
	el.ErrorRaise(err)

	// Format table
	f, fileClose, err = XLSXopenFile(filePath, false)
	el.ErrorRaise(err)
	defer fileClose(f)
	err = e.XLSXstreamTableFormat(f, sheetName, lastRow)
	el.ErrorRaise(err)
	err = e.XLSXstreamTableSetColumnsStyle(f, sheetName, columnFormat)
	el.ErrorRaise(err)
	err = e.XLSXstreamTableHeaderStyle(f, sheetName)
	el.ErrorRaise(err)

	return lastRow, f.SaveAs(filePath)
}

func XLSXopenFile(
	filePath string, overWrite bool,
) (*excelize.File, func(*excelize.File), error) {
	var file *excelize.File
	deferFunc := func(*excelize.File) {
		err := file.Close()
		if err != nil {
			slog.Error("error closing file", "err", err.Error())
		}
	}
	ok, err := helper.FileExists(filePath)
	if err != nil {
		return file, deferFunc, err
	}

	if !ok || overWrite {
		file = excelize.NewFile()
	} else {
		file, err = excelize.OpenFile(filePath)
		if err != nil {
			return file, deferFunc, err
		}
	}
	return file, deferFunc, err
}

func ConvertStringSliceToInterface(slice []string) []interface{} {
	newslice := make([]interface{}, len(slice))
	for i, e := range slice {
		newslice[i] = e
	}
	return newslice
}

func (e *Extractor) XLSXstreamAddHeader(
	sw *excelize.StreamWriter, internalHeader, externalHeader bool) (int, error) {
	lastRow := 0
	if internalHeader {
		lastRow++
		rowStart, _ := excelize.CoordinatesToCellName(1, lastRow)
		data := ConvertStringSliceToInterface(e.HeaderInternal)
		err := sw.SetRow(rowStart, data)
		if err != nil {
			return lastRow, err
		}
	}
	if externalHeader {
		lastRow++
		rowStart, _ := excelize.CoordinatesToCellName(1, lastRow)
		data := ConvertStringSliceToInterface(e.HeaderExternal)
		err := sw.SetRow(rowStart, data)
		if err != nil {
			return lastRow, err
		}
	}
	return lastRow, nil
}

func (e *Extractor) XLSXstreamAddRows(
	sw *excelize.StreamWriter, lastRow int) (int, error) {
	// Write rows
	currentRow := lastRow
	for _, row := range e.Rows {
		currentRow++
		rowStart, err := excelize.CoordinatesToCellName(1, currentRow)
		if err != nil {
			return currentRow, err
		}
		xlsxrow := row.RowParts.PartsToXLSXrow(
			e.RowPartsPositions, e.RowPartsFieldsPositions)
		err = sw.SetRow(rowStart, xlsxrow)
		if err != nil {
			return currentRow, err
		}
	}
	return currentRow, nil
}

func (e *Extractor) XLSXstreamTableFormat(
	file *excelize.File, sheetName string, lastRow int) (result error) {
	el := helper.ErrList{}
	defer el.Handle(&result)
	// Set pane split (split first row)
	err := file.SetPanes(sheetName, &excelize.Panes{
		Freeze: true,
		YSplit: 1,
		XSplit: 1,
		// TopLeftCell: "A10",
		// ActivePane:  "bottomLeft",
	})
	el.ErrorRaise(err)

	// Set table dimensions and format
	endCell, err := excelize.CoordinatesToCellName(
		len(e.HeaderInternal), lastRow)
	el.ErrorRaise(err)
	cellRange := strings.Join([]string{"A1", endCell}, ":")
	rowStripes := true
	table := excelize.Table{
		Range:             cellRange,
		Name:              "table",
		StyleName:         "TableStyleMedium2",
		ShowColumnStripes: false,
		ShowFirstColumn:   false,
		ShowLastColumn:    false,
		ShowRowStripes:    &rowStripes,
		// ShowRowStripes:    true,
	}
	err = file.AddTable(sheetName, &table)
	el.ErrorRaise(err)
	return nil
}

func (e *Extractor) XLSXstreamTableSetColumnStyle(
	file *excelize.File, sheetName string,
	columName string, fieldID FieldID) error {

	style := excelize.Style{
		// NumFmt: 1, DecimalPlaces: nil,
		// NegRed:       false,
	}
	if fieldID.XLSXcolumnFormat != 0 {
		style.NumFmt = fieldID.XLSXcolumnFormat
	}
	if fieldID.XLSXcustomFormat != "" {
		// customNumFmt := "[$$-409]#,##0.00"
		style.CustomNumFmt = &fieldID.XLSXcustomFormat
	}
	styleIndex, err := file.NewStyle(&style)
	if err != nil {
		return err
	}
	slog.Info("column style applied", "column", columName, "style", styleIndex)
	return file.SetColStyle(sheetName, columName, styleIndex)
}

func (e *Extractor) XLSXstreamTableSetColumnsStyle(
	file *excelize.File, sheetName string, format bool) (output error) {
	el := helper.ErrList{}
	defer el.Handle(&output)
	if !format {
		return nil
	}
	var columnIndex int
	// var err error
	for _, partPos := range e.RowPartsPositions {
		fieldsPos := e.RowPartsFieldsPositions[partPos]
		for _, fp := range fieldsPos {
			columnIndex++
			col, err := excelize.ColumnNumberToName(columnIndex)
			el.ErrorRaise(err)
			pos, ok := FieldsIDsNamesProduction2[fp.FieldID]
			if !ok {
				continue
			}
			err = file.SetColWidth(sheetName, col, col, pos.Width)
			el.ErrorRaise(err)
			err = e.XLSXstreamTableSetColumnStyle(
				file, sheetName, col, pos)
			el.ErrorRaise(err)
		}
	}

	// err = file.SetColWidth(sheetName, "A", "B", 90)
	// if err != nil {
	// return err
	// }
	// err = file.Save()
	return nil
}

func (e *Extractor) XLSXstreamTableHeaderStyle(
	file *excelize.File, sheetName string) error {
	style := excelize.Style{
		Border: []excelize.Border{},
		Fill: excelize.Fill{
			Type: "", Pattern: 1, Color: []string{}, Shading: 0,
		},
		Font: &excelize.Font{
			Bold: true, Italic: false, Underline: "",
			// Family: "Times New Roman", Size: 10, Strike: false,
			Family: "", Size: 10, Strike: false,
			// Color: "777777", ColorIndexed: 0, ColorTheme: nil,
			Color: "", ColorIndexed: 0, ColorTheme: nil,
			ColorTint: 0.0, VertAlign: "",
		},
		Alignment: &excelize.Alignment{
			Horizontal: "", Indent: 0, JustifyLastLine: false, ReadingOrder: 0,
			RelativeIndent: 0, ShrinkToFit: false, TextRotation: 0, Vertical: "",
			WrapText: false,
		},
		Protection: &excelize.Protection{
			Hidden: false, Locked: false,
		},
		NumFmt: 0, DecimalPlaces: nil,
		// customNumFmt := "[$$-409]#,##0.00"
		CustomNumFmt: nil,
		NegRed:       false,
	}
	styleIndex, err := file.NewStyle(&style)
	if err != nil {
		return err
	}
	colName, err := excelize.ColumnNumberToName(len(e.HeaderInternal))
	if err != nil {
		return err
	}
	err = file.SetColWidth(sheetName, "A", colName, 15)
	if err != nil {
		return err
	}
	slog.Info("header style applied", "sheet", sheetName, "cols", colName)
	return file.SetRowStyle(sheetName, 1, 1, styleIndex)
}

func (r RowParts) PartsToXLSXrow(
	partsPos RowPartsPositions,
	partsFieldsPos RowPartsFieldsPositions) []interface{} {
	var out []interface{}
	for _, partPos := range partsPos {
		fieldsPos := partsFieldsPos[partPos]
		if len(fieldsPos) == 0 {
			continue
		}
		part := r[partPos]
		outPart := part.PartToXSLXrow(fieldsPos)
		out = append(out, outPart...)
	}
	return out
}

func (p RowPart) PartToXSLXrow(
	fieldsPosition RowPartFieldsPositions) []interface{} {
	specValEmpty := RowFieldSpecialValueCodeMap[RowFieldValueEmptyString]
	var out []interface{}
	for _, pos := range fieldsPosition {
		var value string
		field, ok := p[pos.FieldID]
		if !ok {
			value = specValEmpty
			out = append(out, value)
			continue
		}
		value = strings.TrimSpace(field.Value)
		value = TransformEmptyString(value)
		value = helper.EscapeCSVdelim(value)
		out = append(out, value)
	}
	return out
}

func CSVreadRows(csvFileName string, csvDelim rune) ([][]string, error) {
	file, err := os.Open(csvFileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 0 // must be same as first row
	// reader.FieldsPerRecord = -1 // do not check record length
	reader.Comma = csvDelim
	// reader.LazyQuotes = true
	return reader.ReadAll()
}

func GetHeaderRowIndex(rows [][]string) int {
	headerRow := 0
	for i := 0; i < len(rows[0]); i++ {
		_, ok := FieldsIDsNamesProduction2.GetByName(rows[0][i])
		if ok {
			break
		}
		_, ok = FieldsIDsNamesProduction2.GetByName(rows[1][i])
		if ok {
			headerRow = 1
			break
		}
	}
	return headerRow
}

func XLSXsetColumnStyle(
	f *excelize.File, sheetName string, rows [][]string) error {
	headerRow := GetHeaderRowIndex(rows)
	for i := 0; i < len(rows[0]); i++ {
		fieldID, ok := FieldsIDsNamesProduction2.GetByName(rows[headerRow][i])
		// column properties
		cname, err := excelize.ColumnNumberToName(i + 1)
		if err != nil {
			return err
		}
		firstCell := fmt.Sprintf("%s%d", cname, headerRow+1)
		lastCell := fmt.Sprintf("%s%d", cname, len(rows))

		// default column width
		if ok && fieldID.Width > -1 {
			if err := f.SetColWidth(
				sheetName, cname, cname, fieldID.Width); err != nil {
				return err
			}
		}

		// Column format number
		styleDef := excelize.Style{}
		styleDef.NumFmt = fieldID.XLSXcolumnFormat

		// Custom format
		if ok && fieldID.XLSXcustomFormat != "" {
			styleDef.CustomNumFmt = &fieldID.XLSXcustomFormat
		}
		if !ok {
			stext := "@"
			styleDef.CustomNumFmt = &stext
		}

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

func CSVtoXLSX(csvFile string, csvDelim rune) error {
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
	if err := xlsxFile.SetColWidth(sheetName, "A", lastCol, 12); err != nil {
		return err
	}
	err = XLSXsetColumnStyle(xlsxFile, sheetName, records)
	if err != nil {
		return err
	}
	ncol := len(records[0])
	nrow := len(records)
	headerRow := GetHeaderRowIndex(records)
	for coli := 0; coli < ncol; coli++ {
		fieldID, _ := FieldsIDsNamesProduction2.GetByName(records[headerRow][coli])
		for rowi := 0; rowi < nrow; rowi++ {
			cellRef, _ := excelize.CoordinatesToCellName(coli+1, rowi+1)
			// fmt.Println("INDEX", coli, rowi)
			// fmt.Println("ROW_LEN", len(records[rowi]))
			// if len(records[rowi]) < coli {
			// continue
			// }
			// fmt.Println("FUCITL", nrow, ncol)
			cell := records[rowi][coli]
			// err := XLSXsetCellValueTry(xlsxFile, sheetName, cellRef, cell)
			err := XLSXsetCellValue(xlsxFile, sheetName, fieldID, cellRef, cell)
			if err != nil {
				return err
			}
		}
	}

	name := helper.FilenameWithoutExtension(csvFile)
	dir := filepath.Dir(csvFile)
	xlsxFilePath := filepath.Join(dir, name+".xlsx")
	// fmt.Println("FUCIT")
	return xlsxFile.SaveAs(xlsxFilePath)
}

func XLSXsetCellValueTry(
	file *excelize.File, sheetName string, cellRef, cell string,
) error {
	// DATE
	date, err := ParseXMLdate(cell)
	if err == nil {
		return file.SetCellValue(sheetName, cellRef, date)
	}
	date2, err := time.Parse("2006-01-02 15:04:05", cell)
	if err == nil {
		return file.SetCellValue(sheetName, cellRef, date2)
	}
	// NUMBER
	res, err := strconv.Atoi(cell)
	if err == nil {
		return file.SetCellValue(sheetName, cellRef, res)
	}
	return file.SetCellValue(sheetName, cellRef, cell)
}

func CSVdirToXLSX(csvFolder string, csvDelim rune) error {
	files, err := helper.ListDirFiles(csvFolder)
	if err != nil {
		return err
	}
	for _, f := range files {
		ext := filepath.Ext(f)
		if ext == ".csv" {
			err = CSVtoXLSX(f, csvDelim)
			if err != nil {
				return fmt.Errorf("%w file: %s", err, f)
			}
		}
	}
	return nil
}
