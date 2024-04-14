package extcases

import i "github/czech-radio/openmedia-archive/internal"

var EXTmock = i.OMextractors{
	{
		PartPrefixCode: i.FieldPrefix_ComputedRID,
		FieldIDs:       []string{"RID"},
	},
	{
		ObjectPath:       "/Radio Rundown/<OM_RECORD>",
		ObjectAttrsNames: []string{"RecordID"},
		PartPrefixCode:   i.FieldPrefix_RadioRec,
	},
	{
		ObjectPath:     "/*Hourly Rundown",
		FieldsPath:     i.TemplateHeaderFieldPath,
		FieldIDs:       []string{"8"},
		PartPrefixCode: i.FieldPrefix_HourlyHead,
	},
	{
		ObjectPath:       "/<OM_RECORD>",
		ObjectAttrsNames: []string{"RecordID"},
		PartPrefixCode:   i.FieldPrefix_HourlyRec,
	},
	{
		ObjectPath:           "/Sub Rundown",
		ObjectAttrsNames:     []string{"TemplateName"},
		FieldsPath:           i.TemplateHeaderFieldPath,
		FieldIDs:             ProductionFieldsSubRundown,
		PartPrefixCode:       i.FieldPrefix_SubHead,
		KeepWhenZeroSubnodes: true,
	},
	{
		ObjectPath:           "/<OM_RECORD>",
		ObjectAttrsNames:     []string{"RecordID"},
		PartPrefixCode:       i.FieldPrefix_SubRec,
		KeepWhenZeroSubnodes: true,
	},
	{
		ObjectPath:       "/Radio Story",
		FieldsPath:       i.TemplateHeaderFieldPath,
		ObjectAttrsNames: []string{"TemplateName"},
		PartPrefixCode:   i.FieldPrefix_StoryHead,
		FieldIDs:         []string{"8"},
		// FieldIDs:             ProductionFieldsRadioStory,
		KeepWhenZeroSubnodes: true,
	},
	{
		ObjectPath:           "/<OM_RECORD>",
		ObjectAttrsNames:     []string{"RecordID"},
		PartPrefixCode:       i.FieldPrefix_StoryRec,
		KeepWhenZeroSubnodes: true,
	},
	{
		ObjectPath:           "Audioclip|Contact Item",
		ObjectAttrsNames:     []string{"TemplateName"},
		PartPrefixCode:       i.FieldPrefix_StoryKategory,
		KeepWhenZeroSubnodes: true,
		PreserveParentNode:   true,
	},
	{
		ObjectPath:           "Audioclip",
		FieldsPath:           i.TemplateHeaderFieldPath,
		FieldIDs:             []string{"8"},
		PartPrefixCode:       i.FieldPrefix_AudioClipHead,
		KeepWhenZeroSubnodes: true,
	},
	{
		ObjectPath:           "Contact Item",
		FieldsPath:           i.TemplateHeaderFieldPath,
		FieldIDs:             []string{"1"},
		PartPrefixCode:       i.FieldPrefix_ContactItemHead,
		KeepWhenZeroSubnodes: true,
	},
	{
		PartPrefixCode: i.FieldPrefix_ComputedID,
		FieldIDs:       []string{"ID"},
	},
}
