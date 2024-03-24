package internal

var ProductionFieldsSubRundown = []string{
	"8",
	"1004",
	"1003",
	"1005",
	"321",
}

var ProductionFieldsRadioStory = []string{
	"8",
	// "5081",
	// "1004",
	// "1003",
	// "1005",
	// "1035",
	// "1036",
	// "1029",
	// "1010",
	// "1002",
	// "321",
	// "5079",
	// "16",
	// "5082",
	// "5072",
	// "5016",
	// "5",
	// "6",
	// "12",
	// "5071",
	// "5070",
}

var ProductionFieldsAudio = []string{
	"8",
	// "1005",
	// "5082",
}

var ProductionFieldsContactItems = []string{
	"421",
	// "422",
	// "423",
	// "424",
	// "5015",
	// "5016",
	// "5087",
	// "5088",
}

var EXTproduction = OMextractors{
	{
		ObjectPath:     "/*Hourly Rundown",
		FieldsPath:     TemplateHeaderFieldPath,
		FieldIDs:       []string{"8"},
		PartPrefixCode: FieldPrefix_HourlyHead,
		KeepInputRows:  false,
	},
	// {
	// ObjectPath: "/<OM_RECORD>",
	// ObjectPath: "/<OM_READER>/<OM_HEADER>",
	// ObjectAttrsNames: []string{"RecordID"},
	// FieldIDs:         []string{"8"},
	// PartPrefixCode:   FieldPrefix_HourlyRec,
	// KeepInputRows:    false,
	// },
	{
		ObjectPath:         "/*Sub Rundown",
		ObjectAttrsNames:   []string{"TemplateName"},
		FieldsPath:         TemplateHeaderFieldPath,
		FieldIDs:           ProductionFieldsSubRundown,
		PartPrefixCode:     FieldPrefix_SubHead,
		PreserveParentNode: true,
		// KeepInputRows:    true,
	},
	{
		ObjectPath:       "/*Sub Rundown",
		ObjectAttrsNames: []string{"TemplateName"},
		FieldsPath:       TemplateHeaderFieldPath,
		FieldIDs:         []string{"321"},
		// PartPrefixCode:   FieldPrefix_StoryKategory,
		PartPrefixCode: FieldPrefix_SubHead,
		// PreserveParentNode: true,
		// KeepInputRows:    true,
	},
	// {
	// ObjectPath:       "/<OM_RECORD>/Radio Story",
	// ObjectAttrsNames: []string{"TemplateName"},
	// FieldsPath:       HeaderFieldPath,
	// FieldIDs:         ProductionFieldsRadioStory,
	// PartPrefixCode:   FieldPrefix_StoryHead,
	// KeepInputRows:    false,
	// },
	// {
	// ObjectPath:       "/<OM_RECORD>",
	// ObjectAttrsNames: []string{"RecordID"},
	// FieldsPath:       RecordFieldPath,
	// FieldIDs:         ProductionFieldsRadioStory,
	// PartPrefixCode:   FieldPrefix_StoryHead,
	// KeepInputRows:    false,
	// },
	// {
	// ObjectPath:     "<OM_RECORD>/Audioclip",
	// FieldsPath:     "/OM_HEADER/OM_FIELD",
	// FieldIDs:       ProductionFieldsAudio,
	// PartPrefixCode: FieldPrefix_AudioClipHead,
	// KeepInputRows:  true,
	// },
	// {
	// ObjectPath:     "<OM_RECORD>/Contact Item",
	// FieldsPath:     "/OM_HEADER/OM_FIELD",
	// FieldIDs:       ProductionFieldsContactItems,
	// PartPrefixCode: FieldPrefix_ContactItemHead,
	// KeepInputRows:  false,
	// },
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
