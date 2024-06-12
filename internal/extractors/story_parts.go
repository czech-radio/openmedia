package extractors

import (
	"slices"

	a "github/czech-radio/openmedia/internal/archive"
	e "github/czech-radio/openmedia/internal/extract"
)

var EXTContacts = e.OMextractors{
	e.OMextractor{
		PartPrefixCode: e.RowPartCode_ComputedRID,
		FieldIDs:       e.ProductionFieldsComputedRID,
	},
	e.OMextractor{
		ObjectPath:           "/Radio Rundown",
		FieldsPath:           a.TemplateHeaderFieldPath,
		FieldIDs:             e.ProductionFieldsRadioRundown,
		PartPrefixCode:       e.RowPartCode_RadioHead,
		ResultNodeGoUpLevels: 1,
	},
	e.OMextractor{
		ObjectPath: "/Radio Rundown/<OM_RECORD>",
		// ObjectAttrsNames: []string{"RecordID"},
		PartPrefixCode: e.RowPartCode_RadioRec,
	},
	e.OMextractor{
		ObjectPath:       "/*Hourly Rundown",
		FieldsPath:       a.TemplateHeaderFieldPath,
		ObjectAttrsNames: []string{"ObjectID"},
		FieldIDs:         e.ProductionFieldsHourlyRundown,
		PartPrefixCode:   e.RowPartCode_HourlyHead,
	},
	e.OMextractor{
		ObjectPath: "/<OM_RECORD>",
		// ObjectAttrsNames: []string{"RecordID"},
		PartPrefixCode: e.RowPartCode_HourlyRec,
	},
	e.OMextractor{
		ObjectPath:           "/Sub Rundown",
		ObjectAttrsNames:     []string{"ObjectID"},
		FieldsPath:           a.TemplateHeaderFieldPath,
		FieldIDs:             e.ProductionFieldsSubRundown,
		PartPrefixCode:       e.RowPartCode_SubHead,
		KeepWhenZeroSubnodes: true,
	},
	e.OMextractor{
		ObjectPath: "/<OM_RECORD>",
		// ObjectAttrsNames:     []string{"RecordID"},
		PartPrefixCode:       e.RowPartCode_SubRec,
		KeepWhenZeroSubnodes: true,
	},
	e.OMextractor{
		ObjectPath:       "/*Radio Story",
		FieldsPath:       a.TemplateHeaderFieldPath,
		ObjectAttrsNames: []string{"ObjectID"},
		PartPrefixCode:   e.RowPartCode_StoryHead,
		// FieldIDs:         []string{"8"},
		FieldIDs: e.ProductionFieldsRadioStory,
	},

	// Unknow Record without OM_OBJECT insie
	e.OMextractor{
		ObjectPath:       "/<OM_RECORD>",
		ObjectAttrsNames: []string{"RecordID"},
		PartPrefixCode:   e.RowPartCode_StoryRec,
		FieldIDs: slices.Concat(
			[]string{"5001"},
			e.ProductionFieldsAudio, e.ProductionFieldsContactItems),
		KeepWhenZeroSubnodes: true,
	},
	// visidata: select rows with not null Story-REC_RecordID, select rows with null Story-Cat_ObjectID
	// Normal record
	e.OMextractor{
		ObjectPath:           "/Contact Item|Contact Bin",
		ObjectAttrsNames:     []string{"TemplateName", "ObjectID"},
		PartPrefixCode:       e.RowPartCode_StoryKategory,
		KeepWhenZeroSubnodes: true,
		ResultNodeGoUpLevels: 1,
	},
	e.OMextractor{
		ObjectPath: "/Contact Item|Contact Bin",
		FieldsPath: a.TemplateHeaderFieldPath,
		// FieldIDs:             []string{"1"},
		FieldIDs:             e.ProductionFieldsContactItems,
		PartPrefixCode:       e.RowPartCode_ContactItemHead,
		KeepWhenZeroSubnodes: true,
	},
	e.OMextractor{
		PartPrefixCode: e.RowPartCode_ComputedKON,
		// FieldIDs:       []string{"jmeno_spojene"},
	},
}

// // problem azerty
// var EXTproductionContacts2 = OMextractors{
// 	OMextractor{
// 		PartPrefixCode: RowPartCode_ComputedRID,
// 		FieldIDs:       ProductionFieldsComputedRID,
// 	},
// 	OMextractor{
// 		ObjectPath:           "/Radio Rundown",
// 		FieldsPath:           ar.TemplateHeaderFieldPath,
// 		FieldIDs:             ProductionFieldsRadioRundown,
// 		PartPrefixCode:       RowPartCode_RadioHead,
// 		ResultNodeGoUpLevels: 1,
// 	},
// 	OMextractor{
// 		ObjectPath: "/Radio Rundown/<OM_RECORD>",
// 		// ObjectAttrsNames: []string{"RecordID"},
// 		PartPrefixCode: RowPartCode_RadioRec,
// 	},
// 	OMextractor{
// 		ObjectPath:       "/*Hourly Rundown",
// 		FieldsPath:       ar.TemplateHeaderFieldPath,
// 		ObjectAttrsNames: []string{"ObjectID"},
// 		FieldIDs:         ProductionFieldsHourlyRundown,
// 		PartPrefixCode:   RowPartCode_HourlyHead,
// 	},
// 	OMextractor{
// 		ObjectPath: "/<OM_RECORD>",
// 		// ObjectAttrsNames: []string{"RecordID"},
// 		PartPrefixCode: RowPartCode_HourlyRec,
// 	},
// 	OMextractor{
// 		ObjectPath:           "/Sub Rundown",
// 		ObjectAttrsNames:     []string{"ObjectID"},
// 		FieldsPath:           ar.TemplateHeaderFieldPath,
// 		FieldIDs:             ProductionFieldsSubRundown,
// 		PartPrefixCode:       RowPartCode_SubHead,
// 		KeepWhenZeroSubnodes: true,
// 	},
// 	OMextractor{
// 		ObjectPath: "/<OM_RECORD>",
// 		// ObjectAttrsNames:     []string{"RecordID"},
// 		PartPrefixCode:       RowPartCode_SubRec,
// 		KeepWhenZeroSubnodes: true,
// 	},
// 	OMextractor{
// 		ObjectPath:       "/*Radio Story",
// 		FieldsPath:       ar.TemplateHeaderFieldPath,
// 		ObjectAttrsNames: []string{"ObjectID"},
// 		PartPrefixCode:   RowPartCode_StoryHead,
// 		// FieldIDs:         []string{"8"},
// 		FieldIDs: ProductionFieldsRadioStory,
// 	},

// 	// Unknow Record without OM_OBJECT insie
// 	OMextractor{
// 		ObjectPath:       "/<OM_RECORD>",
// 		ObjectAttrsNames: []string{"RecordID"},
// 		PartPrefixCode:   RowPartCode_StoryRec,
// 		FieldIDs: slices.Concat(
// 			[]string{"5001"}, ProductionFieldsAudio, ProductionFieldsContactItems),
// 		KeepWhenZeroSubnodes: true,
// 	},
// 	// visidata: select rows with not null Story-REC_RecordID, select rows with null Story-Cat_ObjectID
// 	// Normal record
// 	OMextractor{
// 		ObjectPath:           "/Contact Item|Contact Bin",
// 		ObjectAttrsNames:     []string{"TemplateName", "ObjectID"},
// 		PartPrefixCode:       RowPartCode_StoryKategory,
// 		KeepWhenZeroSubnodes: true,
// 		ResultNodeGoUpLevels: 1,
// 	},
// 	OMextractor{
// 		ObjectPath: "/Contact Item|Contact Bin",
// 		FieldsPath: ar.TemplateHeaderFieldPath,
// 		// FieldIDs:             []string{"1"},
// 		FieldIDs:             ProductionFieldsContactItems,
// 		PartPrefixCode:       RowPartCode_ContactItemHead,
// 		KeepWhenZeroSubnodes: true,
// 	},
// 	OMextractor{
// 		PartPrefixCode: RowPartCode_ComputedKON,
// 		// FieldIDs:       []string{"jmeno_spojene"},
// 	},
// }
