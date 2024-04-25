package extract

import ar "github/czech-radio/openmedia/internal/archive"

var EXTmock = OMextractors{
	{
		PartPrefixCode: RowPartCode_ComputedRID,
		FieldIDs:       []string{"RID"},
	},
	{
		ObjectPath:       "/Radio Rundown/<OM_RECORD>",
		ObjectAttrsNames: []string{"RecordID"},
		PartPrefixCode:   RowPartCode_RadioRec,
	},
	{
		ObjectPath:     "/*Hourly Rundown",
		FieldsPath:     ar.TemplateHeaderFieldPath,
		FieldIDs:       []string{"8"},
		PartPrefixCode: RowPartCode_HourlyHead,
	},
	{
		ObjectPath:       "/<OM_RECORD>",
		ObjectAttrsNames: []string{"RecordID"},
		PartPrefixCode:   RowPartCode_HourlyRec,
	},
	{
		ObjectPath:           "/Sub Rundown",
		ObjectAttrsNames:     []string{"TemplateName"},
		FieldsPath:           ar.TemplateHeaderFieldPath,
		FieldIDs:             ProductionFieldsSubRundown,
		PartPrefixCode:       RowPartCode_SubHead,
		KeepWhenZeroSubnodes: true,
	},
	{
		ObjectPath:           "/<OM_RECORD>",
		ObjectAttrsNames:     []string{"RecordID"},
		PartPrefixCode:       RowPartCode_SubRec,
		KeepWhenZeroSubnodes: true,
	},
	{
		ObjectPath:       "/Radio Story",
		FieldsPath:       ar.TemplateHeaderFieldPath,
		ObjectAttrsNames: []string{"TemplateName"},
		PartPrefixCode:   RowPartCode_StoryHead,
		FieldIDs:         []string{"8"},
		// FieldIDs:             ProductionFieldsRadioStory,
		KeepWhenZeroSubnodes: true,
	},
	{
		ObjectPath:           "/<OM_RECORD>",
		ObjectAttrsNames:     []string{"RecordID"},
		PartPrefixCode:       RowPartCode_StoryRec,
		KeepWhenZeroSubnodes: true,
	},
	{
		ObjectPath:           "Audioclip|Contact Item",
		ObjectAttrsNames:     []string{"TemplateName"},
		PartPrefixCode:       RowPartCode_StoryKategory,
		KeepWhenZeroSubnodes: true,
		ResultNodeGoUpLevels: 1,
	},
	{
		ObjectPath:           "Audioclip",
		FieldsPath:           ar.TemplateHeaderFieldPath,
		FieldIDs:             []string{"8"},
		PartPrefixCode:       RowPartCode_AudioClipHead,
		KeepWhenZeroSubnodes: true,
	},
	{
		ObjectPath:           "Contact Item",
		FieldsPath:           ar.TemplateHeaderFieldPath,
		FieldIDs:             []string{"1"},
		PartPrefixCode:       RowPartCode_ContactItemHead,
		KeepWhenZeroSubnodes: true,
	},
	{
		PartPrefixCode: RowPartCode_ComputedKON,
		FieldIDs:       []string{"ID"},
	},
}
