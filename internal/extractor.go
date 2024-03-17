package internal

import (
	"path/filepath"
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

// NEW STRUCTURE
type CSVrowPart map[string]CSVrowField // ObjectPrefix:CSVrowField
type CSVrowParts map[string]CSVrowPart // Whole CSV line
type CSVrow []*CSVrowParts
type CSVtable []*CSVrow
type CSVtables map[string]*CSVtable

type OMobjExtractor struct {
	ObjectPath string
	FieldsPath string
	FieldIDs   []string

	// Internals
	FieldsPrefix               string
	DontReplaceParentObjectRow bool
	FieldIDsMap                map[string]bool
}

type OMobjExtractors []OMobjExtractor
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

func GetLastPartOfObjectPath(path string) string {
	return filepath.Base(path)
}
