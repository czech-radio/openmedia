package extract

import (
	ar "github/czech-radio/openmedia/internal/archive"
)

var EXTproduction = OMextractors{
	OMextractor{
		PartPrefixCode: FieldPrefix_ComputedRID,
		FieldIDs:       []string{"FileName", "C-RID", "C-index"},
	},
	OMextractor{
		ObjectPath:           "/Radio Rundown",
		FieldsPath:           ar.TemplateHeaderFieldPath,
		FieldIDs:             []string{"5081"},
		PartPrefixCode:       FieldPrefix_RadioHead,
		ResultNodeGoUpLevels: 1,
	},
	OMextractor{
		ObjectPath:       "/Radio Rundown/<OM_RECORD>",
		ObjectAttrsNames: []string{"RecordID"},
		PartPrefixCode:   FieldPrefix_RadioRec,
	},
	OMextractor{
		ObjectPath:       "/*Hourly Rundown",
		FieldsPath:       ar.TemplateHeaderFieldPath,
		ObjectAttrsNames: []string{"ObjectID"},
		FieldIDs:         []string{"1000", "8"},
		PartPrefixCode:   FieldPrefix_HourlyHead,
	},
	OMextractor{
		ObjectPath:       "/<OM_RECORD>",
		ObjectAttrsNames: []string{"RecordID"},
		PartPrefixCode:   FieldPrefix_HourlyRec,
	},
	OMextractor{
		ObjectPath:           "/Sub Rundown",
		ObjectAttrsNames:     []string{"TemplateName", "ObjectID"},
		FieldsPath:           ar.TemplateHeaderFieldPath,
		FieldIDs:             ProductionFieldsSubRundown,
		PartPrefixCode:       FieldPrefix_SubHead,
		KeepWhenZeroSubnodes: true,
	},
	OMextractor{
		ObjectPath:           "/<OM_RECORD>",
		ObjectAttrsNames:     []string{"RecordID"},
		PartPrefixCode:       FieldPrefix_SubRec,
		KeepWhenZeroSubnodes: true,
	},
	OMextractor{
		ObjectPath:       "/*Radio Story",
		FieldsPath:       ar.TemplateHeaderFieldPath,
		ObjectAttrsNames: []string{"TemplateName", "ObjectID"},
		PartPrefixCode:   FieldPrefix_StoryHead,
		// FieldIDs:         []string{"8"},
		FieldIDs: ProductionFieldsRadioStory,
	},
	OMextractor{
		ObjectPath:           "/<OM_RECORD>/Audioclip|Contact Item|Contact Bin",
		ObjectAttrsNames:     []string{"TemplateName", "ObjectID"},
		PartPrefixCode:       FieldPrefix_StoryKategory,
		KeepWhenZeroSubnodes: true,
		ResultNodeGoUpLevels: 1,
	},
	OMextractor{
		ObjectPath: "/Audioclip",
		FieldsPath: ar.TemplateHeaderFieldPath,
		// FieldIDs:   []string{"8"},
		FieldIDs:             ProductionFieldsAudio,
		PartPrefixCode:       FieldPrefix_AudioClipHead,
		KeepWhenZeroSubnodes: true,
	},
	OMextractor{
		ObjectPath: "/Contact Item|Contact Bin",
		FieldsPath: ar.TemplateHeaderFieldPath,
		// FieldIDs:             []string{"1"},
		FieldIDs:             ProductionFieldsContactItems,
		PartPrefixCode:       FieldPrefix_ContactItemHead,
		KeepWhenZeroSubnodes: true,
	},
	OMextractor{
		PartPrefixCode: FieldPrefix_ComputedID,
		FieldIDs:       []string{"ID", "jmeno_spojene"},
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
