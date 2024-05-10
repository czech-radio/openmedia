package extract

import (
	"fmt"
	"path/filepath"

	"log/slog"
	"os"
	"strings"

	"github.com/triopium/go_utils/pkg/helper"
)

// csv.NewWriter() would be alternative

// CSVrowBuild
func (row RowParts) CSVrowBuild(
	builder *strings.Builder,
	partsPos RowPartsPositions,
	partsFieldsPos RowPartsFieldsPositions,
	delim string,
) {
	for i, partPos := range partsPos {
		fieldsPos := partsFieldsPos[partPos]
		if len(fieldsPos) == 0 {
			continue
		}
		if i != 0 {
			fmt.Fprintf(builder, "%s", "\t")
		}
		part := row[partPos]
		part.CSVrowPartBuild(builder, fieldsPos, delim)
	}
	// end the line
	fmt.Fprintf(builder, "%s", "\n")
}

// CSVrowPartBuild
func (part RowPart) CSVrowPartBuild(
	builder *strings.Builder,
	fieldsPosition RowPartFieldsPositions,
	delim string,
) {
	specValEmpty := RowFieldSpecialValueCodeMap[RowFieldValueEmptyString]
	count := len(fieldsPosition)
	formatUse := "%s" + delim
	for i, pos := range fieldsPosition {
		if i == count-1 || count == 1 {
			formatUse = "%s"
		}
		field, ok := part[pos.FieldID]
		if !ok {
			value := specValEmpty
			fmt.Fprintf(builder, formatUse, value)
			continue
		}
		value := strings.TrimSpace(field.Value)
		value = TransformEmptyString(value)
		value = helper.EscapeCSVdelim(value)
		fmt.Fprintf(builder, formatUse, value)
	}
}

func (e *Extractor) CSVheaderBuild(internal, external bool) {
	e.TableXML.CSVheaderWriterLocal = new(strings.Builder)
	sb := e.TableXML.CSVheaderWriterLocal
	if internal {
		iheader := strings.Join(e.HeaderInternal, e.CSVdelim)
		fmt.Fprintln(sb, iheader)
	}
	if external {
		eheader := strings.Join(e.HeaderExternal, e.CSVdelim)
		fmt.Fprintln(sb, eheader)
	}
}

func (e *Extractor) CSVtableBuild(
	internalHeader, externalHeader bool,
	delim string, clearBuilder bool, rowsIndexes ...[]int) int {
	//TODO: remove header booleans and make wrapper function for building table and header
	var rowsCount int
	if e.TableXML.CSVtableWriterLocal == nil {
		e.TableXML.CSVtableWriterLocal = new(strings.Builder)
	}
	if clearBuilder {
		e.TableXML.CSVtableWriterLocal = new(strings.Builder)
	}
	sb := e.TableXML.CSVtableWriterLocal
	if len(rowsIndexes) > 1 {
		panic("not implemented for multiple indexes' slices")
	}
	if len(rowsIndexes) == 0 || rowsIndexes[0] == nil {
		e.CSVheaderBuild(internalHeader, externalHeader)
		for i := 0; i < len(e.TableXML.Rows); i++ {
			e.TableXML.Rows[i].CSVrowBuild(
				sb, e.RowPartsPositions, e.RowPartsFieldsPositions, delim,
			)
			rowsCount++
		}
		return rowsCount
	}
	// Build only specified rows in indexes slice
	e.CSVheaderBuild(internalHeader, externalHeader)
	for _, index := range rowsIndexes[0] {
		e.TableXML.Rows[index].CSVrowBuild(
			sb, e.RowPartsPositions, e.RowPartsFieldsPositions, delim,
		)
		rowsCount++
	}
	return rowsCount
}

func (e *Extractor) CSVtableOutputs(dstDir, fileName, extractorsName, preset string, internal bool) {
	if internal {
		// csv with internal header
		name := strings.Join(
			[]string{fileName, extractorsName, preset, "wh.csv"}, "_")
		dstFile1 := filepath.Join(
			dstDir, name)
		e.CSVheaderBuild(true, true)
		e.CSVheaderWrite(dstFile1, true)
		e.CSVtableWrite(dstFile1, false)
	}

	// csv without internal header
	name := strings.Join(
		[]string{fileName, extractorsName, preset, "woh.csv"}, "_")
	dstFile2 := filepath.Join(
		dstDir, name)
	e.CSVheaderBuild(false, true)
	e.CSVheaderWrite(dstFile2, true)
	e.CSVtableWrite(dstFile2, false)

	// xlsx with internal header
	name = strings.Join(
		[]string{fileName, extractorsName, preset, "woh.xlsx"}, "_")
	dstFile3 := filepath.Join(
		dstDir, name)
	lastRow, err := e.XLSXstreamTableSave(
		dstFile3, "Sheet1", false, true, true)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Info("last row written", "number", lastRow)
}

func (e *Extractor) CSVtableOutput(dstFile string) {
	if dstFile == "" {
		fmt.Print(e.TableXML.CSVtableWriterLocal)
		return
	}
}

func FileOverwritePermissions(overWrite bool) int {
	perms := os.O_RDWR | os.O_CREATE | os.O_APPEND
	if overWrite {
		perms = os.O_RDWR | os.O_CREATE | os.O_TRUNC
	}
	return perms
}

// CSVtableWrite
func (e *Extractor) CSVtableWrite(dstFilePath string, overWrite bool) {
	if dstFilePath == "" {
		fmt.Println(e.TableXML.CSVtableWriterLocal)
		return
	}
	perms := FileOverwritePermissions(overWrite)
	outputFile, err := os.OpenFile(dstFilePath, perms, 0600)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	n, err := outputFile.WriteString(e.TableXML.CSVtableWriterLocal.String())
	if err != nil {
		panic(err)
	}
	slog.Warn("written bytes to file", "fileName", dstFilePath, "bytesCount", n)
}

func (e *Extractor) CSVheaderWrite(dstFilePath string, overWrite bool) {
	if dstFilePath == "" {
		fmt.Println(e.TableXML.CSVheaderWriterLocal)
		return
	}
	perms := FileOverwritePermissions(overWrite)
	outputFile, err := os.OpenFile(dstFilePath, perms, 0600)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	n, err := outputFile.WriteString(e.TableXML.CSVheaderWriterLocal.String())
	if err != nil {
		panic(err)
	}
	slog.Warn("written bytes to file", "fileName", dstFilePath, "bytesCount", n)
}
