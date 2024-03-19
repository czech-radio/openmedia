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
	DontReplaceParentObjectRow bool
	FieldIDsMap                map[string]bool
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
	e.OMextractors.ReplaceParentRowTrueChecker()
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
	for _, extr := range e.OMextractors {
		rows, err := ExpandTableRows(e.CSVtable, extr) // : maybe wrong
		if err != nil {
			return err
		}
		e.CSVtable = rows
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

func (omo *OMextractor) MapFields() {
	omo.FieldIDsMap = make(map[string]bool, len(omo.FieldIDs))
	for _, id := range omo.FieldIDs {
		omo.FieldIDsMap[id] = true
	}
}

func (omoes OMextractors) ReplaceParentRowTrueChecker() {
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
