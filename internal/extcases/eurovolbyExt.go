package extcases

import i "github/czech-radio/openmedia-archive/internal"

var EXTeuroVolby = i.OMextractors{
	// {
	// ObjectPath:       "/Radio Rundown/<OM_RECORD>",
	// ObjectAttrsNames: []string{"RecordID"},
	// PartPrefixCode:   FieldPrefix_RadioRec,
	// },
	{
		ObjectPath:     "/*Hourly Rundown",
		FieldsPath:     i.TemplateHeaderFieldPath,
		FieldIDs:       []string{"8"},
		PartPrefixCode: i.FieldPrefix_HourlyHead,
	},
	{
		ObjectPath: "/<OM_RECORD>",
		// ObjectAttrsNames: []string{"RecordID"},
		PartPrefixCode: i.FieldPrefix_HourlyRec,
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
		ObjectPath: "/<OM_RECORD>",
		// ObjectAttrsNames:     []string{"RecordID"},
		PartPrefixCode:       i.FieldPrefix_SubRec,
		KeepWhenZeroSubnodes: true,
	},
	{
		ObjectPath:       "/Radio Story",
		FieldsPath:       i.TemplateHeaderFieldPath,
		ObjectAttrsNames: []string{"TemplateName"},
		PartPrefixCode:   i.FieldPrefix_StoryHead,
		FieldIDs:         ProductionFieldsRadioStory,
		// KeepWhenZeroSubnodes: true,
	},
	// {
	// PartPrefixCode: FieldPrefix_ComputedID,
	// FieldIDs:       []string{"ID"},
	// },
}
