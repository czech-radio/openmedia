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

func (e *Extractor) XLSXheaderStreamSave(
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

func (e *Extractor) XLSXrowsStreamSave(
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

func (e *Extractor) XLSXtableStreamSave(
	filePath, sheetName string,
	internalHeader, externalHeader, overWrite bool) (int, error) {
	f, fileClose, err := XLSXopenFile(filePath, overWrite)
	if err != nil {
		return 0, err
	}
	defer fileClose(f)
	sw, err := f.NewStreamWriter(sheetName) // will overwrite sheet
	if err != nil {
		return 0, err
	}
	lastRow, err := e.XLSXheaderStreamSave(sw, externalHeader, internalHeader)
	if err != nil {
		return lastRow, err
	}
	lastRow, err = e.XLSXrowsStreamSave(sw, lastRow)
	if err != nil {
		return lastRow, err
	}
	err = sw.Flush()
	if err != nil {
		return lastRow, err
	}
	return lastRow, f.SaveAs(filePath)
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
