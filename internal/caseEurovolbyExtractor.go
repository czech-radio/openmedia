package internal

var EXTeuroVolbyRID = OMextractors{
	{
		ObjectPath:     "/*Hourly Rundown",
		FieldsPath:     TemplateHeaderFieldPath,
		FieldIDs:       []string{"8"},
		PartPrefixCode: FieldPrefix_HourlyHead,
		KeepInputRow:   false,
	},
	{
		ObjectPath:       "/<OM_RECORD>",
		ObjectAttrsNames: []string{"RecordID"},
		// FieldIDs:         []string{"8"},
		PartPrefixCode: FieldPrefix_HourlyRec,
		KeepInputRow:   false,
	},
	{
		ObjectPath:           "/Sub Rundown",
		ObjectAttrsNames:     []string{"TemplateName"},
		FieldsPath:           TemplateHeaderFieldPath,
		FieldIDs:             ProductionFieldsSubRundown,
		PartPrefixCode:       FieldPrefix_SubHead,
		PreserveParentNode:   true,
		KeepWhenZeroSubnodes: true,
	},
	{
		ObjectPath: "/Sub Rundown/<OM_RECORD>",
		// ObjectAttrsNames:     []string{"RecordID"},
		PartPrefixCode:       FieldPrefix_SubRec,
		KeepWhenZeroSubnodes: true,
	},
	{
		ObjectPath:     "/Radio Story",
		FieldsPath:     TemplateHeaderFieldPath,
		PartPrefixCode: FieldPrefix_StoryHead,
		// FieldIDs:       []string{"8"},
		FieldIDs: ProductionFieldsRadioStory,
		// KeepWhenZeroSubnodes: true,
	},
	{
		ObjectPath:           "/<OM_RECORD>",
		ObjectAttrsNames:     []string{"RecordID"},
		PartPrefixCode:       FieldPrefix_StoryRec,
		KeepWhenZeroSubnodes: true,
	},
	{
		PartPrefixCode: FieldPrefix_ComputedID,
		FieldIDs:       []string{"ID"},
	},
}

var EXTeuroVolby = OMextractors{
	{
		ObjectPath:     "/*Hourly Rundown",
		FieldsPath:     TemplateHeaderFieldPath,
		FieldIDs:       []string{"8"},
		PartPrefixCode: FieldPrefix_HourlyHead,
		KeepInputRow:   false,
	},
	{
		ObjectPath: "/<OM_RECORD>",
		// ObjectAttrsNames: []string{"RecordID"},
		// FieldIDs:         []string{"8"},
		PartPrefixCode: FieldPrefix_HourlyRec,
		KeepInputRow:   false,
	},
	{
		ObjectPath:           "/Sub Rundown",
		ObjectAttrsNames:     []string{"TemplateName"},
		FieldsPath:           TemplateHeaderFieldPath,
		FieldIDs:             ProductionFieldsSubRundown,
		PartPrefixCode:       FieldPrefix_SubHead,
		PreserveParentNode:   true,
		KeepWhenZeroSubnodes: true,
	},
	{
		ObjectPath: "/Sub Rundown/<OM_RECORD>",
		// ObjectAttrsNames:     []string{"RecordID"},
		PartPrefixCode:       FieldPrefix_SubRec,
		KeepWhenZeroSubnodes: true,
	},
	{
		ObjectPath:     "/Radio Story",
		FieldsPath:     TemplateHeaderFieldPath,
		PartPrefixCode: FieldPrefix_StoryHead,
		// FieldIDs:       []string{"8"},
		FieldIDs: ProductionFieldsRadioStory,
		// KeepWhenZeroSubnodes: true,
	},
	{
		PartPrefixCode: FieldPrefix_ComputedID,
		FieldIDs:       []string{"ID"},
	},
	// {
	// ObjectPath:           "/<OM_RECORD>",
	// ObjectAttrsNames:     []string{"RecordID"},
	// PartPrefixCode:       FieldPrefix_StoryRec,
	// KeepWhenZeroSubnodes: true,
	// },
}
