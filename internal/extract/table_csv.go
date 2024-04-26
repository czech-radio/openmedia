package extract

import (
	"fmt"
	"github/czech-radio/openmedia/internal/helper"
	"log/slog"
	"os"
	"strings"
)

// csv.NewWriter() would be alternative

// CSVrowBuild
func (row RowParts) CSVrowBuild(
	builder *strings.Builder,
	partsPos RowPartsPositions,
	partsFieldsPos RowPartsFieldsPositions,
	delim string,
) {
	for _, pos := range partsPos {
		fieldsPos := partsFieldsPos[pos]
		part := row[pos]
		part.CSVrowPartBuild(builder, fieldsPos, delim)
	}
	fmt.Fprintf(builder, "%s", "\n")
}

// CSVrowPartBuild
func (part RowPart) CSVrowPartBuild(
	builder *strings.Builder,
	fieldsPosition RowPartFieldsPositions,
	delim string,
) {
	specValEmpty := RowFieldValueCodeMap[RowFieldValueEmptyString]
	formatUse := "%s" + delim
	for _, pos := range fieldsPosition {
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
	if e.TableXML.CSVwriterLocal == nil {
		e.TableXML.CSVwriterLocal = new(strings.Builder)
	}
	sb := e.TableXML.CSVwriterLocal
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
	if e.TableXML.CSVwriterLocal == nil {
		e.TableXML.CSVwriterLocal = new(strings.Builder)
	}
	if clearBuilder {
		e.TableXML.CSVwriterLocal = new(strings.Builder)
	}
	sb := e.TableXML.CSVwriterLocal
	if len(rowsIndexes) > 1 {
		panic("not implemented for multiple indexes' slices")
	}
	switch len(rowsIndexes) {
	case 0:
		e.CSVheaderBuild(internalHeader, externalHeader)
		for i := 0; i < len(e.TableXML.Rows); i++ {
			e.TableXML.Rows[i].CSVrowBuild(
				sb, e.RowPartsPositions, e.RowPartsFieldsPositions, delim,
			)
			rowsCount++
		}
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
		fmt.Print(e.TableXML.CSVwriterLocal)
		return
	}
}

// CSVtableWrite
func (e *Extractor) CSVtableWrite(dstFilePath string) {
	if dstFilePath == "" {
		fmt.Println(e.TableXML.CSVwriterLocal)
		return
	}
	outputFile, err := os.OpenFile(dstFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	n, err := outputFile.WriteString(e.TableXML.CSVwriterLocal.String())
	if err != nil {
		panic(err)
	}
	slog.Warn("written bytes to file", "fileName", dstFilePath, "bytesCount", n)
}
