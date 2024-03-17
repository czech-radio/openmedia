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

func (omo *OMobjExtractor) MapFields() {
	omo.FieldIDsMap = make(map[string]bool, len(omo.FieldIDs))
	for _, id := range omo.FieldIDs {
		omo.FieldIDsMap[id] = true
	}
}

type OMobjExtractor2 struct {
	ObjectPath     string
	ObjectSelector string // OM_OBJEC,OM_RECORD...
	SelectorName   string
	FieldsPath     string
	FieldsPrefix   string
	// Internals
	FieldIDs                   []string
	FieldIDsMap                map[string]bool
	DontReplaceParentObjectRow bool
}

type OMobjExtractors []OMobjExtractor

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

// XMLtag/path <OM_OBJECT> or Radio Rundown, name without <> will be OM_OBJECT and AttrName "TemplateName"
// AttrName TemplateName
// AttrValue Radio Rundown

var ObjectXMLnameMap = map[string]string{
	"OM_OBJECT": "TemplateName",
	"OM_RECORD": "RecordID",
	"OM_FIELD":  "FieldID",
}

type OMobjExtractor struct {
	ObjectPath     string
	ObjectSelector string // OM_OBJEC,OM_RECORD...
	SelectorName   string
	FieldsPath     string
	FieldsPrefix   string
	// Internals
	FieldIDs                   []string
	FieldIDsMap                map[string]bool
	DontReplaceParentObjectRow bool
}

var EXTproduction = OMobjExtractors{
	{
		ObjectPath: "/Radio Rundown",
		FieldsPath: "/OM_HEADER/OM_FIELD",
		FieldIDs:   []string{"1", "8"},
	},
	// {
	// ObjectPath: "/Radio Rundown/<OM_RECORD>",
	// AttrName:   "RecordID",
	// },
}

// attr name, values
var EXTproduction2 = OMobjExtractors{
	{
		ObjectPath: "/Radio Rundown",
		// ObjectSelector: "//OM_OBJECT[@TemplateName='%s']",
		ObjectSelector: "//OM_OBJECT[@TemplateName='%s']",
		SelectorName:   "Radio Rundown",
		FieldsPath:     "/OM_HEADER/OM_FIELD",
		FieldIDs:       []string{"8"},
	},
	{
		ObjectPath:     "/Radio Rundown/<OM_RECORD>",
		ObjectSelector: "/OM_RECORD[@RecordID='%s']",
	},
	// ObjectSelector: "/OM_RECORD/[@RecordID='%s']",
	// SelectorName: "RecordID"
	// FieldsPath:     "/OM_RECORD/OM_FIELD",
	// FieldIDs:       []string{"1", "8"},
	// },
	// FieldIDs: []string{"8"},
	// },
	// {
	// ObjectPath: "/Radio Rundown/Hourly Rundown",
	// FieldsPath: "/OM_HEADER/OM_FIELD",
	// FieldIDs:   []string{"1", "8", "9"},
	// FieldIDs:        []string{"*"}, // ALL fields
	// },
	// {
	// ObjectPath: "/Radio Rundown/Hourly Rundown/Sub Rundown",
	// FieldsPath: "/OM_HEADER/OM_FIELD",
	// FieldIDs:   []string{"8"},
	// },
	// {
	// ObjectPath: "/Radio Rundown/Hourly Rundown/Radio Story",
	// ObjectPath: "/Radio Rundown/Hourly Rundown/RECORD",
	// FieldsPath: "/OM_HEADER/OM_FIELD",
	// FieldIDs:   []string{"8"},
	// },
	// {
	// OmObject:   "Radio Story",
	// ObjectPath: "//AudioClip",
	// FieldsPath: "/OM_FIELD",
	// FieldIDs:   []string{"8"},
	// },
	// {
	// OmObject:   "Radio Story",
	// ObjectPath: "//Contact Item",
	// FieldsPath: "/OM_HEADER/OM_FIELD",
	// FieldIDs:   []string{"8"},
	// },
}

// "//OM_OBJECT[@TemplateName='%s']/%s/*", ext.OM_type, ext.Path,
