package internal

var EXTeuroVolby = OMextractors{
	// {
	// ObjectPath:       "/Radio Rundown/<OM_RECORD>",
	// ObjectAttrsNames: []string{"RecordID"},
	// PartPrefixCode:   FieldPrefix_RadioRec,
	// },
	{
		ObjectPath:     "/*Hourly Rundown",
		FieldsPath:     TemplateHeaderFieldPath,
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
		FieldsPath:           TemplateHeaderFieldPath,
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
		FieldsPath:       TemplateHeaderFieldPath,
		ObjectAttrsNames: []string{"TemplateName"},
		PartPrefixCode:   FieldPrefix_StoryHead,
		FieldIDs:         ProductionFieldsRadioStory,
		// KeepWhenZeroSubnodes: true,
	},
	// {
	// PartPrefixCode: FieldPrefix_ComputedID,
	// FieldIDs:       []string{"ID"},
	// },
}
