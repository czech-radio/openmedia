package internal

import (
	"fmt"
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

func (omo *OMobjExtractor) MapFields() {
	omo.FieldIDsMap = make(map[string]bool, len(omo.FieldIDs))
	for _, id := range omo.FieldIDs {
		omo.FieldIDsMap[id] = true
	}
}

type OMobjExtractor struct {
	OmObject     string
	ObjectPath   string
	FieldsPath   string
	FieldsPrefix string
	// Internals
	FieldIDs                   []string
	FieldIDsMap                map[string]bool
	DontReplaceParentObjectRow bool
}

type OMobjExtractors []OMobjExtractor

func (omoes OMobjExtractors) ReplaceParentRowTrueChecker() {
	// Check if there is following extractor referencing same object as current extractor
	lomoes := len(omoes)
	fmt.Println("EFcount", lomoes)
	for currentIndex, currentExt := range omoes {
	extractor:
		for followingIndex := currentIndex + 1; followingIndex < lomoes; followingIndex++ {
			if followingIndex > lomoes {
				fmt.Println(currentIndex, "ich break")
				break extractor
			}
			currentParent := filepath.Dir(currentExt.ObjectPath)
			followingParent := filepath.Dir(omoes[followingIndex].ObjectPath)
			fmt.Println("EF", currentIndex, currentExt.ObjectPath, currentParent, followingParent)
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

var EXTproduction = OMobjExtractors{
	{
		OmObject:   "Radio Rundown",
		ObjectPath: "/Radio Rundown",
		FieldsPath: "/OM_HEADER/OM_FIELD",
		FieldIDs:   []string{"1", "8"},
	},
	{
		OmObject:   "Hourly Rundown",
		ObjectPath: "/Radio Rundown/Hourly Rundown",
		FieldsPath: "/OM_HEADER/OM_FIELD",
		FieldIDs:   []string{"1", "8", "9"},
		// FieldIDs:        []string{"*"}, // ALL fields
	},
	{
		OmObject:   "Radio Story",
		ObjectPath: "/Radio Rundown/Hourly Rundown/Sub Rundown",
		FieldsPath: "/OM_HEADER/OM_FIELD",
		FieldIDs:   []string{"8"},
	},
	{
		OmObject:   "Radio Story",
		ObjectPath: "/Radio Rundown/Hourly Rundown/Radio Story",
		FieldsPath: "/OM_HEADER/OM_FIELD",
		FieldIDs:   []string{"8"},
	},
	{
		OmObject:   "Radio Story",
		ObjectPath: "//AudioClip",
		FieldsPath: "/OM_HEADER/OM_FIELD",
		FieldIDs:   []string{"8"},
	},
	{
		OmObject:   "Radio Story",
		ObjectPath: "//Contact Item",
		FieldsPath: "/OM_HEADER/OM_FIELD",
		FieldIDs:   []string{"8"},
	},
}

// "//OM_OBJECT[@TemplateName='%s']/%s/*", ext.OM_type, ext.Path,
