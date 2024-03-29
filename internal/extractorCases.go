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

var EXTproduction = OMextractors{
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
		ObjectPath: "/<OM_RECORD>",
		// ObjectAttrsNames:     []string{"RecordID"},
		PartPrefixCode:       FieldPrefix_StoryRec,
		KeepWhenZeroSubnodes: true,
	},
	{
		PartPrefixCode: FieldPrefix_ComputedKategory,
		FieldIDs:       []string{"kategory"},
	},
	{
		ObjectPath:       "/Audioclip",
		ObjectAttrsNames: []string{"TemplateName"},
		FieldsPath:       TemplateHeaderFieldPath,
		// FieldIDs:             ProductionSimpleTest,
		FieldIDs:             ProductionFieldsAudio,
		PartPrefixCode:       FieldPrefix_AudioClipHead,
		KeepWhenZeroSubnodes: true,
		PreserveParentNode:   true,
	},
	{
		ObjectPath:       "/Contact Item",
		ObjectAttrsNames: []string{"TemplateName"},
		FieldsPath:       TemplateHeaderFieldPath,
		// FieldIDs:             ProductionSimpleTest,
		FieldIDs:             ProductionFieldsContactItems,
		PartPrefixCode:       FieldPrefix_ContactItemHead,
		KeepWhenZeroSubnodes: true,
		PreserveParentNode:   true,
	},
	{
		PartPrefixCode: FieldPrefix_ComputedID,
		FieldIDs:       []string{"ID"},
	},
}

var EXTtest = OMextractors{
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
		ObjectPath: "/<OM_RECORD>",
		FieldIDs:   []string{"kategory"},
		OrExt: []OMextractor{
			{
				ObjectPath:     "/Audioclip",
				FieldsPath:     TemplateHeaderFieldPath,
				PartPrefixCode: FieldPrefix_AudioClipHead,
				FieldIDs:       ProductionFieldsAudio,
			},
			{
				ObjectPath:     "/Contact Item",
				FieldsPath:     TemplateHeaderFieldPath,
				PartPrefixCode: FieldPrefix_ContactItemHead,
			},
		}},
	{
		PartPrefixCode: FieldPrefix_ComputedKategory,
		FieldIDs:       []string{"kategory"},
	},
	{
		PartPrefixCode: FieldPrefix_ComputedID,
		FieldIDs:       []string{"ID"},
	},
}

// {
// ObjectPath:       "/Audioclip",
// ObjectAttrsNames: []string{"TemplateName"},
// FieldsPath:       TemplateHeaderFieldPath,
// FieldIDs:             ProductionFieldsAudio,
// PartPrefixCode:       FieldPrefix_AudioClipHead,
// KeepWhenZeroSubnodes: true,
// PreserveParentNode:   true,
// },
// {
// ObjectPath:       "/Contact Item",
// ObjectAttrsNames: []string{"TemplateName"},
// FieldsPath:       TemplateHeaderFieldPath,
// FieldIDs:             ProductionSimpleTest,
// FieldIDs:             ProductionFieldsContactItems,
// PartPrefixCode:       FieldPrefix_ContactItemHead,
// KeepWhenZeroSubnodes: true,
// PreserveParentNode:   true,
// },
