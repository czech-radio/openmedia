package internal

import (
	"log/slog"
	"path/filepath"

	"github.com/antchfx/xmlquery"
)

type ObjectAttributes = map[string]string
type UniqueValues = map[string]int // value vs count

type CSVrowFields = []CSVrowField

type OMextractor struct {
	ObjectPath   string
	FieldsPath   string
	FieldIDs     []string
	FieldsPrefix string

	// Internals
	// KeepInputRows bool
	KeepInputRows bool
	FieldIDsMap   map[string]bool
}

type OMextractors []OMextractor

type Extractor struct {
	OMextractors
	CSVrowPartsPositions
	CSVrowPartsFieldsPositions
	CSVrowHeader string
	CSVdelim     string
	BaseNode     *xmlquery.Node
	CSVtable
}

func (e *Extractor) Init(
	baseNode *xmlquery.Node,
	omextractors OMextractors,
	CSVdelim string) {
	e.OMextractors = omextractors
	e.CSVdelim = CSVdelim
	e.BaseNode = baseNode
	e.MapRowParts()
	e.MapRowPartsFieldsPositions()
	e.CSVheaderCreate(CSVdelim)
	e.OMextractors.KeepInputRowsChecker()
	e.CSVtable = []*CSVrowNode{{baseNode, CSVrow{}}}
}

func (e *Extractor) MapRowParts() {
	extCount := len(e.OMextractors)
	partsPos := make(CSVrowPartsPositions, extCount)
	for i, extr := range e.OMextractors {
		partsPos[i] = extr.FieldsPrefix
	}
	e.CSVrowPartsPositions = partsPos
}

func (e *Extractor) MapRowPartsFieldsPositions() {
	extCount := len(e.OMextractors)
	partsPos := make(CSVrowPartsFieldsPositions, extCount)
	for _, extr := range e.OMextractors {
		fp := GetPartFieldsPositions(extr)
		partsPos[extr.FieldsPrefix] = fp
	}
	e.CSVrowPartsFieldsPositions = partsPos
}

func (e *Extractor) ExtractTable() error {
	for i, extr := range e.OMextractors {
		rows, err := ExpandTableRows(e.CSVtable, extr) // : maybe wrong
		e.CSVtable = rows
		if err != nil {
			return err
		}
		slog.Debug("extractor", "position", i, "objectPath", extr)
		// e.PrintTableToCSV(true, "\t")
	}
	return nil
}

func GetPartFieldsPositions(extr OMextractor) CSVrowPartFieldsPositions {
	fieldsPositions := make(CSVrowPartFieldsPositions, 0, len(extr.FieldIDs))
	for _, fi := range extr.FieldIDs {
		fp := FieldPosition{
			FieldPrefix: extr.FieldsPrefix,
			FieldID:     fi,
			FieldName:   "",
		}
		fieldsPositions = append(fieldsPositions, fp)
	}
	return fieldsPositions
}

func (extr *OMextractor) MapFields() {
	extr.FieldIDsMap = make(map[string]bool, len(extr.FieldIDs))
	for _, id := range extr.FieldIDs {
		extr.FieldIDsMap[id] = true
	}
}

func (extrs OMextractors) KeepInputRowsChecker() {
	// Check if there is following extractor referencing same object as current extractor
	eCount := len(extrs)
	for eCurrent := 0; eCurrent < eCount; eCurrent++ {
		extr := extrs[eCurrent]
		if extr.KeepInputRows {
			continue
		}
		if eCurrent == eCount {
			// maybe not needed, also without allow the extr position to be independent insted to process sequentially
			extr.KeepInputRows = false
			break
		}
		eNext := eCurrent + 1
		for next := eNext; next < eCount; next++ {
			//TODO: Without it depends on manual input alone
			// fmt.Println("fek", eCurrent, next)
		}
	}
}

// for current, currentExtractor := range omoes {
// 	if current+1 > extractorsCount {
// 		continue
// 	}
// 	for next := current + 1; next < extractorsCount; next++ {
// 		currentParent := filepath.Dir(currentExtractor.ObjectPath)
// 		followingParent := filepath.Dir(omoes[next].ObjectPath)
// 		// fmt.Println("EF", current, currentExtractor.ObjectPath, currentParent, followingParent)
// 		if currentParent == followingParent {
// 			slog.Debug("wont be replaced", "extractor", currentExtractor.ObjectPath)
// 			omoes[current].KeepInputRows = true
// 		}
// 		continue
// 	}
// }

func GetLastPartOfObjectPath(path string) string {
	return filepath.Base(path)
}
