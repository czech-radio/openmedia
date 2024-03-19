package internal

import (
	"log/slog"
	"path/filepath"

	"github.com/antchfx/xmlquery"
)

type ObjectAttributes = map[string]string
type Fields = map[int]string       // FieldID/FieldName vs value
type UniqueValues = map[string]int // value vs count

type CSVrowsIntMap map[int]CSVrowFields
type CSVrowFields = []CSVrowField

type OMobjExtractor struct {
	ObjectPath   string
	FieldsPath   string
	FieldIDs     []string
	FieldsPrefix string

	// Internals
	DontReplaceParentObjectRow bool
	FieldIDsMap                map[string]bool
}

type OMobjExtractors []OMobjExtractor

type Extractor struct {
	OMobjExtractors
	CSVrowPartsPositions
	CSVrowPartsFieldsPositions
	CSVrowHeader string
	CSVdelim     string
	BaseNode     *xmlquery.Node
	CSVtable
}

func (e *Extractor) Init(
	baseNode *xmlquery.Node,
	omextractors OMobjExtractors,
	CSVdelim string) {
	e.OMobjExtractors = omextractors
	e.CSVdelim = CSVdelim
	e.BaseNode = baseNode
	e.MapRowParts()
	e.MapRowPartsFieldsPositions()
	e.CSVheaderCreate(CSVdelim)
	e.OMobjExtractors.ReplaceParentRowTrueChecker()
	e.CSVtable = []*CSVrowNode{{baseNode, CSVrow{}}}
}

func (e *Extractor) MapRowParts() {
	extCount := len(e.OMobjExtractors)
	partsPos := make(CSVrowPartsPositions, extCount)
	for i, extr := range e.OMobjExtractors {
		partsPos[i] = extr.FieldsPrefix
	}
	e.CSVrowPartsPositions = partsPos
}

func (e *Extractor) MapRowPartsFieldsPositions() {
	extCount := len(e.OMobjExtractors)
	partsPos := make(CSVrowPartsFieldsPositions, extCount)
	for _, extr := range e.OMobjExtractors {
		fp := GetPartFieldsPositions(extr)
		partsPos[extr.FieldsPrefix] = fp
	}
	e.CSVrowPartsFieldsPositions = partsPos
}

func (e *Extractor) ExtractTable() error {
	for _, extr := range e.OMobjExtractors {
		rows, err := ExpandTableRows(e.CSVtable, extr) // : maybe wrong
		if err != nil {
			return err
		}
		e.CSVtable = rows
	}
	return nil
}

func GetPartFieldsPositions(extr OMobjExtractor) CSVrowPartFieldsPositions {
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

func (omo *OMobjExtractor) MapFields() {
	omo.FieldIDsMap = make(map[string]bool, len(omo.FieldIDs))
	for _, id := range omo.FieldIDs {
		omo.FieldIDsMap[id] = true
	}
}

func (omoes OMobjExtractors) ReplaceParentRowTrueChecker() {
	// Check if there is following extractor referencing same object as current extractor
	lomoes := len(omoes)
	if lomoes == 1 {
		return
	}
	for currentIndex, currentExt := range omoes {
	extractor:
		for followingIndex := currentIndex + 1; followingIndex < lomoes; followingIndex++ {
			if followingIndex > lomoes {
				break extractor
			}
			currentParent := filepath.Dir(currentExt.ObjectPath)
			followingParent := filepath.Dir(omoes[followingIndex].ObjectPath)
			// fmt.Println("EF", currentIndex, currentExt.ObjectPath, currentParent, followingParent)
			if currentParent == followingParent {
				slog.Debug("wont be replaced", "extractor", currentExt.ObjectPath)
				omoes[currentIndex].DontReplaceParentObjectRow = true
			}
			continue
		}
	}
}

func GetLastPartOfObjectPath(path string) string {
	return filepath.Base(path)
}
