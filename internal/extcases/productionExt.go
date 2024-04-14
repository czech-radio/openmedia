package extcases

import (
	i "github/czech-radio/openmedia-archive/internal"
)

var EXTproduction = i.OMextractors{
	i.OMextractor{
		PartPrefixCode: i.FieldPrefix_ComputedRID,
		FieldIDs:       []string{"FileName", "C-RID", "C-index"},
	},
	i.OMextractor{
		ObjectPath:       "/Radio Rundown/<OM_RECORD>",
		ObjectAttrsNames: []string{"RecordID"},
		PartPrefixCode:   i.FieldPrefix_RadioRec,
	},
	i.OMextractor{
		ObjectPath:     "/*Hourly Rundown",
		FieldsPath:     i.TemplateHeaderFieldPath,
		FieldIDs:       []string{"8"},
		PartPrefixCode: i.FieldPrefix_HourlyHead,
	},
	i.OMextractor{
		ObjectPath:       "/<OM_RECORD>",
		ObjectAttrsNames: []string{"RecordID"},
		PartPrefixCode:   i.FieldPrefix_HourlyRec,
	},
	i.OMextractor{
		ObjectPath:           "/Sub Rundown",
		ObjectAttrsNames:     []string{"TemplateName", "ObjectID"},
		FieldsPath:           i.TemplateHeaderFieldPath,
		FieldIDs:             ProductionFieldsSubRundown,
		PartPrefixCode:       i.FieldPrefix_SubHead,
		KeepWhenZeroSubnodes: true,
	},
	i.OMextractor{
		ObjectPath:           "/<OM_RECORD>",
		ObjectAttrsNames:     []string{"RecordID"},
		PartPrefixCode:       i.FieldPrefix_SubRec,
		KeepWhenZeroSubnodes: true,
	},
	i.OMextractor{
		ObjectPath:       "/*Radio Story",
		FieldsPath:       i.TemplateHeaderFieldPath,
		ObjectAttrsNames: []string{"TemplateName", "ObjectID"},
		PartPrefixCode:   i.FieldPrefix_StoryHead,
		// FieldIDs:         []string{"8"},
		FieldIDs: ProductionFieldsRadioStory,
	},
	i.OMextractor{
		ObjectPath:           "/<OM_RECORD>/Audioclip|Contact Item|Contact Bin",
		ObjectAttrsNames:     []string{"TemplateName", "ObjectID"},
		PartPrefixCode:       i.FieldPrefix_StoryKategory,
		KeepWhenZeroSubnodes: true,
		PreserveParentNode:   true,
	},
	i.OMextractor{
		ObjectPath: "/Audioclip",
		FieldsPath: i.TemplateHeaderFieldPath,
		// FieldIDs:   []string{"8"},
		FieldIDs:             ProductionFieldsAudio,
		PartPrefixCode:       i.FieldPrefix_AudioClipHead,
		KeepWhenZeroSubnodes: true,
	},
	i.OMextractor{
		ObjectPath: "/Contact Item|Contact Bin",
		FieldsPath: i.TemplateHeaderFieldPath,
		// FieldIDs:             []string{"1"},
		FieldIDs:             ProductionFieldsContactItems,
		PartPrefixCode:       i.FieldPrefix_ContactItemHead,
		KeepWhenZeroSubnodes: true,
	},
	i.OMextractor{
		PartPrefixCode: i.FieldPrefix_ComputedID,
		FieldIDs:       []string{"ID"},
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
