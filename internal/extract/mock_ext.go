package extract

import ar "github/czech-radio/openmedia/internal/archive"

var EXTmock = OMextractors{
	{
		PartPrefixCode: FieldPrefix_ComputedRID,
		FieldIDs:       []string{"RID"},
	},
	{
		ObjectPath:       "/Radio Rundown/<OM_RECORD>",
		ObjectAttrsNames: []string{"RecordID"},
		PartPrefixCode:   FieldPrefix_RadioRec,
	},
	{
		ObjectPath:     "/*Hourly Rundown",
		FieldsPath:     ar.TemplateHeaderFieldPath,
		FieldIDs:       []string{"8"},
		PartPrefixCode: FieldPrefix_HourlyHead,
	},
	{
		ObjectPath:       "/<OM_RECORD>",
		ObjectAttrsNames: []string{"RecordID"},
		PartPrefixCode:   FieldPrefix_HourlyRec,
	},
	{
		ObjectPath:           "/Sub Rundown",
		ObjectAttrsNames:     []string{"TemplateName"},
		FieldsPath:           ar.TemplateHeaderFieldPath,
		FieldIDs:             ProductionFieldsSubRundown,
		PartPrefixCode:       FieldPrefix_SubHead,
		KeepWhenZeroSubnodes: true,
	},
	{
		ObjectPath:           "/<OM_RECORD>",
		ObjectAttrsNames:     []string{"RecordID"},
		PartPrefixCode:       FieldPrefix_SubRec,
		KeepWhenZeroSubnodes: true,
	},
	{
		ObjectPath:       "/Radio Story",
		FieldsPath:       ar.TemplateHeaderFieldPath,
		ObjectAttrsNames: []string{"TemplateName"},
		PartPrefixCode:   FieldPrefix_StoryHead,
		FieldIDs:         []string{"8"},
		// FieldIDs:             ProductionFieldsRadioStory,
		KeepWhenZeroSubnodes: true,
	},
	{
		ObjectPath:           "/<OM_RECORD>",
		ObjectAttrsNames:     []string{"RecordID"},
		PartPrefixCode:       FieldPrefix_StoryRec,
		KeepWhenZeroSubnodes: true,
	},
	{
		ObjectPath:           "Audioclip|Contact Item",
		ObjectAttrsNames:     []string{"TemplateName"},
		PartPrefixCode:       FieldPrefix_StoryKategory,
		KeepWhenZeroSubnodes: true,
		PreserveParentNode:   true,
	},
	{
		ObjectPath:           "Audioclip",
		FieldsPath:           ar.TemplateHeaderFieldPath,
		FieldIDs:             []string{"8"},
		PartPrefixCode:       FieldPrefix_AudioClipHead,
		KeepWhenZeroSubnodes: true,
	},
	{
		ObjectPath:           "Contact Item",
		FieldsPath:           ar.TemplateHeaderFieldPath,
		FieldIDs:             []string{"1"},
		PartPrefixCode:       FieldPrefix_ContactItemHead,
		KeepWhenZeroSubnodes: true,
	},
	{
		PartPrefixCode: FieldPrefix_ComputedID,
		FieldIDs:       []string{"ID"},
	},
}
