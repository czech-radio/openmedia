package internal

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
	Path         string
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
	for currentIndex, currentExt := range omoes {
	extractor:
		for followingExt := currentIndex + 1; currentIndex < lomoes-1; followingExt++ {
			if followingExt >= lomoes-1 {
				break extractor
			}
			if currentExt.Path == omoes[followingExt].Path {
				omoes[followingExt].DontReplaceParentObjectRow = true
			}
			continue
		}
	}
}

var EXTproduction = OMobjExtractors{
	{
		OmObject:   "Radio Rundown",
		Path:       "",
		FieldsPath: "/OM_HEADER/OM_FIELD",
		FieldIDs:   []string{"1", "8"},
	},
	{
		OmObject:   "Hourly Rundown",
		Path:       "/Radio Rundown",
		FieldsPath: "/OM_HEADER/OM_FIELD",
		FieldIDs:   []string{"1", "8", "9"},
		// FieldIDs:        []string{"*"}, // ALL fields
	},
	{
		OmObject:   "Sub Rundown",
		Path:       "/Radio Rundown/Hourly Rundown",
		FieldsPath: "/OM_HEADER/OM_FIELD",
		FieldIDs:   []string{"8"},
	},
	{
		OmObject:   "Radio Story",
		Path:       "/Radio Rundown/Hourly Rundown",
		FieldsPath: "/OM_HEADER/OM_FIELD",
		FieldIDs:   []string{"8"},
	},
	{
		OmObject:   "Makar",
		Path:       "/Radio Rundown/Hourly Rundown",
		FieldsPath: "/OM_HEADER/OM_FIELD",
		FieldIDs:   []string{"8"},
	},
}

// "//OM_OBJECT[@TemplateName='%s']/%s/*", ext.OM_type, ext.Path,
// FieldsPath: "/OM_RECORD/OM_FIELD",
var CSVproduction = []OMobjExtractor{
	{
		OmObject:                   "Radio Rundown",
		Path:                       "",
		FieldsPath:                 "/OM_HEADER/OM_FIELD",
		FieldIDs:                   []string{"1", "8"},
		DontReplaceParentObjectRow: true,
	},
	{
		OmObject:   "Hourly Rundown",
		Path:       "/Radio Rundown",
		FieldsPath: "/OM_HEADER/OM_FIELD",
		FieldIDs:   []string{"1", "8", "9"},
		// FieldIDs:        []string{"*"}, // ALL fields
		DontReplaceParentObjectRow: true,
	},
	{
		OmObject:                   "Sub Rundown",
		Path:                       "/Radio Rundown/Hourly Rundown",
		FieldsPath:                 "/OM_HEADER/OM_FIELD",
		FieldIDs:                   []string{"8"},
		DontReplaceParentObjectRow: false,
	},
	{
		OmObject:                   "Radio Story",
		Path:                       "/Radio Rundown/Hourly Rundown/Sub Rundown",
		FieldsPath:                 "/OM_HEADER/OM_FIELD",
		FieldIDs:                   []string{"8"},
		DontReplaceParentObjectRow: false,
	},
	// {
	// OmObject:        "Sub Rundown",
	// Path:            "/Radio Rundown/Hourly Rundown",
	// FieldsPath:      "/OM_HEADER/OM_FIELD",
	// FieldIDs:        []string{"8"},
	// ReplacePrevious: false,
	// },
	// {
	// OmObject:   "Sub Rundown",
	// Path:       "Radio Rundown/Hourly Rundown/Sub Rundown",
	// FieldsPath: "/OM_HEADER/OM_FIELD",
	// FieldIDs:   []string{"8"},
	// Level:      1,
	// },
}
