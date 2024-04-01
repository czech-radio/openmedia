package internal

var EXTtest = OMextractors{
	{
		PartPrefixCode: FieldPrefix_ComputedRID,
		FieldIDs:       []string{"RID"},
	},
	{
		ObjectPath:     "/*Hourly Rundown",
		FieldsPath:     TemplateHeaderFieldPath,
		FieldIDs:       []string{"8"},
		PartPrefixCode: FieldPrefix_HourlyHead,
	},
	{
		ObjectPath:       "/<OM_RECORD>",
		ObjectAttrsNames: []string{"RecordID"},
		// FieldIDs:         []string{"8"},
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
		ObjectPath:           "/<OM_RECORD>",
		ObjectAttrsNames:     []string{"RecordID"},
		PartPrefixCode:       FieldPrefix_SubRec,
		KeepWhenZeroSubnodes: true,
	},
	{
		ObjectPath:     "/Radio Story",
		FieldsPath:     TemplateHeaderFieldPath,
		PartPrefixCode: FieldPrefix_StoryHead,
		FieldIDs:       []string{"8"},
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
		PartPrefixCode:       FieldPrefix_AudioClipHead,
		KeepWhenZeroSubnodes: true,
	},
	// {
	// 	PartPrefixCode: FieldPrefix_ComputedKategory,
	// 	FieldIDs:       []string{"kategory"},
	// },
	// {
	// 	PartPrefixCode: FieldPrefix_ComputedID,
	// 	FieldIDs:       []string{"ID"},
	// },
}
