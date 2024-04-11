package internal

var EXTproduction = OMextractors{
	{
		PartPrefixCode: FieldPrefix_ComputedRID,
		FieldIDs:       []string{"FileName", "C-RID", "C-index"},
	},
	{
		ObjectPath:       "/Radio Rundown/<OM_RECORD>",
		ObjectAttrsNames: []string{"RecordID"},
		PartPrefixCode:   FieldPrefix_RadioRec,
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
		PartPrefixCode:   FieldPrefix_HourlyRec,
	},
	{
		ObjectPath:           "/Sub Rundown",
		ObjectAttrsNames:     []string{"TemplateName", "ObjectID"},
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
		ObjectPath:       "/*Radio Story",
		FieldsPath:       TemplateHeaderFieldPath,
		ObjectAttrsNames: []string{"TemplateName", "ObjectID"},
		PartPrefixCode:   FieldPrefix_StoryHead,
		// FieldIDs:         []string{"8"},
		FieldIDs: ProductionFieldsRadioStory,
	},
	{
		ObjectPath:           "/<OM_RECORD>/Audioclip|Contact Item|Contact Bin",
		ObjectAttrsNames:     []string{"TemplateName", "ObjectID"},
		PartPrefixCode:       FieldPrefix_StoryKategory,
		KeepWhenZeroSubnodes: true,
		PreserveParentNode:   true,
	},
	{
		ObjectPath: "/Audioclip",
		FieldsPath: TemplateHeaderFieldPath,
		// FieldIDs:   []string{"8"},
		FieldIDs:             ProductionFieldsAudio,
		PartPrefixCode:       FieldPrefix_AudioClipHead,
		KeepWhenZeroSubnodes: true,
	},
	{
		ObjectPath: "/Contact Item|Contact Bin",
		FieldsPath: TemplateHeaderFieldPath,
		// FieldIDs:             []string{"1"},
		FieldIDs:             ProductionFieldsContactItems,
		PartPrefixCode:       FieldPrefix_ContactItemHead,
		KeepWhenZeroSubnodes: true,
	},
	{
		PartPrefixCode: FieldPrefix_ComputedID,
		FieldIDs:       []string{"ID"},
	},
}

//"/Radio Rundown",
//"/Radio Rundown/<OM_RECORD>",
//"/Radio Rundown/<OM_RECORD>/Hourly Rundown",
//"/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>",
//"/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>/Sub Rundown",
//"/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>/Sub Rundown/<OM_RECORD>",
//"/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>/Sub Rundown/<OM_RECORD>/Radio Story",
//"/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>/Sub Rundown/<OM_RECORD>/Radio Story/<OM_RECORD>",
//"/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>/Sub Rundown/<OM_RECORD>/Radio Story",
//"/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>/Radio Story/Contact Item",
