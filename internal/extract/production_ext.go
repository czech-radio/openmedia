package extract

import (
	ar "github/czech-radio/openmedia/internal/archive"
)

var EXTproduction = OMextractors{
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
	// OMextractor{
	// ObjectPath:           "/<OM_RECORD>",
	// ObjectAttrsNames:     []string{"RecordID"},
	// PartPrefixCode:       RowPartCode_StoryRec,
	// KeepWhenZeroSubnodes: true,
	// ResultNodeGoUpLevels: 1,
	// },
	OMextractor{
		ObjectPath: "/<OM_RECORD>/Audioclip|Contact Item|Contact Bin",
		// ObjectPath: "/*Audioclip|Contact Item|Contact Bin",
		// ObjectPath:           "/Audioclip|Contact Item|Contact Bin",
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

//"/Radio Rundown",
//"/Radio Rundown/<OM_RECORD>",
//"/Radio Rundown/<OM_RECORD>/Hourly Rundown",
//"/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>",
//"/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>/Sub Rundown",
//"/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>/Sub Rundown/<OM_RECORD>",
//"/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>/Sub Rundown/<OM_RECORD>/Radio Story",
//"/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>/Sub Rundown/<OM_RECORD>/Radio Story/<OM_RECORD>",
//"/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>/Sub Rundown/<OM_RECORD>/Radio Story",
//"/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>/Radio Story/Contact Item",
