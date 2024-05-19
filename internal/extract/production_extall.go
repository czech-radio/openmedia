package extract

import (
	ar "github/czech-radio/openmedia/internal/archive"
	"slices"
)

// EXTproductionAll extracts all story parts and also parts which does not contain OM_OBJECT which holds attribute TemlateName
var EXTproductionAll = OMextractors{
	OMextractor{
		PartPrefixCode: RowPartCode_ComputedRID,
		FieldIDs:       ProductionFieldsComputedRID,
	},
	OMextractor{
		ObjectPath:           "/Radio Rundown",
		FieldsPath:           ar.TemplateHeaderFieldPath,
		FieldIDs:             ProductionFieldsRadioRundown,
		PartPrefixCode:       RowPartCode_RadioHead,
		ResultNodeGoUpLevels: 1,
	},
	OMextractor{
		ObjectPath: "/Radio Rundown/<OM_RECORD>",
		// ObjectAttrsNames: []string{"RecordID"},
		PartPrefixCode: RowPartCode_RadioRec,
	},
	OMextractor{
		ObjectPath:       "/*Hourly Rundown",
		FieldsPath:       ar.TemplateHeaderFieldPath,
		ObjectAttrsNames: []string{"ObjectID"},
		FieldIDs:         ProductionFieldsHourlyRundown,
		PartPrefixCode:   RowPartCode_HourlyHead,
	},
	OMextractor{
		ObjectPath: "/<OM_RECORD>",
		// ObjectAttrsNames: []string{"RecordID"},
		PartPrefixCode: RowPartCode_HourlyRec,
	},
	OMextractor{
		ObjectPath:           "/Sub Rundown",
		ObjectAttrsNames:     []string{"ObjectID"},
		FieldsPath:           ar.TemplateHeaderFieldPath,
		FieldIDs:             ProductionFieldsSubRundown,
		PartPrefixCode:       RowPartCode_SubHead,
		KeepWhenZeroSubnodes: true,
	},
	OMextractor{
		ObjectPath: "/<OM_RECORD>",
		// ObjectAttrsNames:     []string{"RecordID"},
		PartPrefixCode:       RowPartCode_SubRec,
		KeepWhenZeroSubnodes: true,
	},
	OMextractor{
		ObjectPath:       "/*Radio Story",
		FieldsPath:       ar.TemplateHeaderFieldPath,
		ObjectAttrsNames: []string{"ObjectID"},
		PartPrefixCode:   RowPartCode_StoryHead,
		// FieldIDs:         []string{"8"},
		FieldIDs: ProductionFieldsRadioStory,
	},

	// Unknow Record without OM_OBJECT insie
	OMextractor{
		ObjectPath:       "/<OM_RECORD>",
		ObjectAttrsNames: []string{"RecordID"},
		PartPrefixCode:   RowPartCode_StoryRec,
		FieldIDs: slices.Concat(
			[]string{"5001"}, ProductionFieldsAudio, ProductionFieldsContactItems),
		KeepWhenZeroSubnodes: true,
	},
	// visidata: select rows with not null Story-REC_RecordID, select rows with null Story-Cat_ObjectID
	// Normal record
	OMextractor{
		ObjectPath:           "/Audioclip|Contact Item|Contact Bin",
		ObjectAttrsNames:     []string{"TemplateName", "ObjectID"},
		PartPrefixCode:       RowPartCode_StoryKategory,
		KeepWhenZeroSubnodes: true,
		ResultNodeGoUpLevels: 1,
	},
	OMextractor{
		ObjectPath: "/Audioclip",
		FieldsPath: ar.TemplateHeaderFieldPath,
		// FieldIDs:   []string{"8"},
		FieldIDs:             ProductionFieldsAudio,
		PartPrefixCode:       RowPartCode_AudioClipHead,
		KeepWhenZeroSubnodes: true,
	},
	OMextractor{
		ObjectPath: "/Contact Item|Contact Bin",
		FieldsPath: ar.TemplateHeaderFieldPath,
		// FieldIDs:             []string{"1"},
		FieldIDs:             ProductionFieldsContactItems,
		PartPrefixCode:       RowPartCode_ContactItemHead,
		KeepWhenZeroSubnodes: true,
	},
	OMextractor{
		PartPrefixCode: RowPartCode_ComputedKON,
		// FieldIDs:       []string{"jmeno_spojene"},
	},
}
