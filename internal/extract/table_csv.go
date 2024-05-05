package extract

import (
	"fmt"

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
	count := len(partsPos)
	for i, pos := range partsPos {
		fieldsPos := partsFieldsPos[pos]
		part := row[pos]
		part.CSVrowPartBuild(builder, fieldsPos, delim)
		if len(fieldsPos) == 0 {
			continue
		}
		if i < count-1 {
			fmt.Fprintf(builder, "%s", "\t")
		}
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
	// if e.TableXML.CSVheaderWriterLocal == nil {
	e.TableXML.CSVheaderWriterLocal = new(strings.Builder)
	// }
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
	switch len(rowsIndexes) {
	case 0:
		// Build all rows
		e.CSVheaderBuild(internalHeader, externalHeader)
		for i := 0; i < len(e.TableXML.Rows); i++ {
			e.TableXML.Rows[i].CSVrowBuild(
				sb, e.RowPartsPositions, e.RowPartsFieldsPositions, delim,
			)
			rowsCount++
		}
		// Build only specified rows in indexes slice
	case 1:
		e.CSVheaderBuild(internalHeader, externalHeader)
		for _, index := range rowsIndexes[0] {
			e.TableXML.Rows[index].CSVrowBuild(
				sb, e.RowPartsPositions, e.RowPartsFieldsPositions, delim,
			)
			rowsCount++
		}
	}
	return rowsCount
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
