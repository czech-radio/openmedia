package extract

import (
	ar "github/czech-radio/openmedia/internal/archive"
)

var EXTeuroVolby = OMextractors{
	// {
	// ObjectPath:       "/Radio Rundown/<OM_RECORD>",
	// ObjectAttrsNames: []string{"RecordID"},
	// PartPrefixCode:   FieldPrefix_RadioRec,
	// },
	{
		ObjectPath:     "/*Hourly Rundown",
		FieldsPath:     ar.TemplateHeaderFieldPath,
		FieldIDs:       []string{"8"},
		PartPrefixCode: FieldPrefix_HourlyHead,
	},
	{
		ObjectPath: "/<OM_RECORD>",
		// ObjectAttrsNames: []string{"RecordID"},
		PartPrefixCode: FieldPrefix_HourlyRec,
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
		ObjectPath: "/<OM_RECORD>",
		// ObjectAttrsNames:     []string{"RecordID"},
		PartPrefixCode:       FieldPrefix_SubRec,
		KeepWhenZeroSubnodes: true,
	},
	{
		ObjectPath:       "/Radio Story",
		FieldsPath:       ar.TemplateHeaderFieldPath,
		ObjectAttrsNames: []string{"TemplateName"},
		PartPrefixCode:   FieldPrefix_StoryHead,
		FieldIDs:         ProductionFieldsRadioStory,
	},
	// {
	// PartPrefixCode: FieldPrefix_ComputedID,
	// FieldIDs:       []string{"ID"},
	// },
}
