package internal

import (
	"fmt"
	"path/filepath"
	"strings"
)

type ObjectAttributes = map[string]string
type Fields = map[int]string       // FieldID/FieldName vs value
type UniqueValues = map[string]int // value vs count

type CSVrowField struct {
	FieldPosition int
	FieldID       string
	Value         string
}

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
}

func (e *Extractor) Init(omextractors OMobjExtractors) {
	e.OMobjExtractors = omextractors
	e.MapRowParts()
	e.MapRowPartsFieldsPositions()
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

func (e *Extractor) CreateHeader(delim string) {
	var builder strings.Builder
	for _, i := range e.CSVrowPartsPositions {
		pfp := e.CSVrowPartsFieldsPositions[i]
		fmt.Println(pfp)
		for _, j := range pfp {
			fmt.Fprintf(&builder, "%s_%s%s", j.FieldPrefix, j.FieldID, delim)
		}
	}
	e.CSVrowHeader = builder.String()
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

type XMLomTagStructure struct {
	XMLtagName   string
	SelectorAttr string
}

var OmTagStructureMap = map[string]XMLomTagStructure{
	"<OM_OBJCET>": {"OM_OBJECT", "TemplateName"},
	"<OM_RECORD>": {"OM_RECORD", "RecorddID"},
}

var ObjectXMLnameMap = map[string]string{
	"OM_OBJECT": "TemplateName",
	"OM_RECORD": "RecordID",
	"OM_FIELD":  "FieldID",
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
				omoes[currentIndex].DontReplaceParentObjectRow = true
			}
			continue
		}
	}
}

func (omoes OMobjExtractors) GetRowParts() {
	// for currentIndex, currentExt := range omoes {
}

func GetLastPartOfObjectPath(path string) string {
	return filepath.Base(path)
}
