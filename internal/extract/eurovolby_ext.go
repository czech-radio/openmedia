package extract

import (
	ar "github/czech-radio/openmedia/internal/archive"
)

var EXTeuroVolby = OMextractors{
	// {
	// ObjectPath:       "/Radio Rundown/<OM_RECORD>",
	// ObjectAttrsNames: []string{"RecordID"},
	// PartPrefixCode:   RowPartCode_RadioRec,
	// },
	{
		ObjectPath:     "/*Hourly Rundown",
		FieldsPath:     ar.TemplateHeaderFieldPath,
		FieldIDs:       []string{"8"},
		PartPrefixCode: RowPartCode_HourlyHead,
	},
	{
		ObjectPath: "/<OM_RECORD>",
		// ObjectAttrsNames: []string{"RecordID"},
		PartPrefixCode: RowPartCode_HourlyRec,
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
		ObjectPath: "/<OM_RECORD>",
		// ObjectAttrsNames:     []string{"RecordID"},
		PartPrefixCode:       RowPartCode_SubRec,
		KeepWhenZeroSubnodes: true,
	},
	{
		ObjectPath:       "/Radio Story",
		FieldsPath:       ar.TemplateHeaderFieldPath,
		ObjectAttrsNames: []string{"TemplateName"},
		PartPrefixCode:   RowPartCode_StoryHead,
		FieldIDs:         ProductionFieldsRadioStory,
	},
	// {
	// PartPrefixCode: RowPartCode_ComputedID,
	// FieldIDs:       []string{"ID"},
	// },
}
