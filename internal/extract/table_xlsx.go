package extract

import (
	"log/slog"
	"strings"

	"github.com/triopium/go_utils/pkg/helper"
	"github.com/xuri/excelize/v2"
)

// f.NewStyle()
// f.NewConditionalStyle()
// f.SetColStyle(sheetName, column, styleId)
// f.SetColWidth()
// set
// f.SetCellStyle()
// sw.SetRow()

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

// func (f *File) SetDefaultFont(fontName string)
// func (f *File) GetStyle(idx int) (*Style, error)
// err = f.SetColStyle("Sheet1", "H", style)
// // err = f.SetRowStyle
// err = f.SetColWidth(sheetName, "A", "B", 90)

func (e *Extractor) XLSXstreamTableSave(
	filePath, sheetName string,
	internalHeader, externalHeader, overWrite bool) (lastRow int, result error) {
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
	err = e.XLSXstreamTableSetColumnsStyle(f, sheetName)
	el.ErrorRaise(err)
	err = e.XLSXstreamTableHeaderStyle(f, sheetName)
	el.ErrorRaise(err)

	return lastRow, f.SaveAs(filePath)
}

func (e *Extractor) XLSXstreamTableFormat(
	file *excelize.File, sheetName string, lastRow int) error {
	// Set pane split (split first row)
	err1 := file.SetPanes(sheetName, &excelize.Panes{
		Freeze: true,
		YSplit: 1,
		XSplit: 1,
		// TopLeftCell: "A10",
		// ActivePane:  "bottomLeft",
	})
	if err1 != nil {
		return err1
	}

	// Set table
	endCell, err2 := excelize.CoordinatesToCellName(
		len(e.HeaderInternal), lastRow)
	cellRange := strings.Join([]string{"A1", endCell}, ":")
	if err2 != nil {
		return err2
	}
	rowStripes := true
	table := excelize.Table{
		Range:             cellRange,
		Name:              "table",
		StyleName:         "TableStyleMedium2",
		ShowColumnStripes: true,
		ShowFirstColumn:   true,
		ShowLastColumn:    true,
		ShowRowStripes:    &rowStripes,
		// ShowRowStripes: true,
	}
	err3 := file.AddTable(sheetName, &table)
	if err3 != nil {
		return err3
	}

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
	file *excelize.File, sheetName string) error {
	var columnIndex int
	// var err error
	for _, partPos := range e.RowPartsPositions {
		fieldsPos := e.RowPartsFieldsPositions[partPos]
		for _, fp := range fieldsPos {
			columnIndex++
			col, err := excelize.ColumnNumberToName(columnIndex)
			if err != nil {
				return err
			}
			pos, ok := FieldsIDsNamesProduction2[fp.FieldID]
			// _, ok := FieldsIDsNamesProduction2[fp.FieldID]
			if ok {
				continue
			}
			err = file.SetColWidth(sheetName, col, col, 30)
			if err != nil {
				return err
			}
			err = e.XLSXstreamTableSetColumnStyle(
				file, sheetName, col, pos)
			if err != nil {
				return err
			}
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
	err = file.SetColWidth(sheetName, "A", colName, 30)
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
