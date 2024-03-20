package internal

var EXTproduction = OMextractors{
	{
		ObjectPath:   "*Hourly Rundown",
		FieldsPath:   "/OM_HEADER/OM_FIELD",
		FieldIDs:     []string{"8"},
		FieldsPrefix: "HourlyR-HED",
	},
	{
		ObjectPath:   "*Sub Rundown",
		FieldsPath:   "/OM_HEADER/OM_FIELD",
		FieldIDs:     []string{"8", "1004", "1003", "1005", "321"},
		FieldsPrefix: "SubR-HED",
	},
	// {
	// ObjectPath:   "/*Radio Story",
	// FieldsPath:   "/OM_HEADER/OM_FIELD",
	// FieldIDs:     []string{"8", "1004", "1003", "1005", "321"},
	// FieldsPrefix: "Radio",
	// },

	// {
	// ObjectPath:   "/*Sub Rundown",
	// FieldsPath:   "/OM_HEADER/OM_FIELD",
	// FieldIDs:     []string{"8", "1004", "1003", "1005", "321"},
	// FieldsPrefix: "SubR-HED",
	// },
	// {
	// ObjectPath: "<OM_RECORD>/Radio Story",
	// FieldsPath: "OM_HEADER",
	// FieldIDs: []string{
	// "8", "5081", "1004", "1004", "1003", "1005", "1035", "1036",
	// "1029", "1010", "1002", "321", "5079", "16", "5082", "5072", "5016"
	// },
	// FieldsPrefix: "Story-HED",
	// },
	// {
	// ObjectPath:   "<OM_RECORD>/Contact Item",
	// FieldsPath:   "/OM_HEADER/OM_FIELD",
	// FieldIDs:     []string{"8"},
	// FieldsPrefix: "Contact-HED",
	// },
	// {
	// ObjectPath:   "*AudioClip",
	// FieldsPath:   "/OM_HEADER/OM_FIELD",
	// FieldIDs:     []string{"8"},
	// FieldsPrefix: "SubR-HED",
	// },
}

var EXTproductionRECandHED = OMextractors{
	{
		ObjectPath:   "/Radio Rundown",
		FieldsPath:   "/OM_HEADER/OM_FIELD",
		FieldIDs:     []string{"5081", "1", "8", "1003", "1004"},
		FieldsPrefix: "RadioR-HED",
	},
	{
		ObjectPath:   "/Radio Rundown/<OM_RECORD>",
		FieldsPath:   "/OM_FIELD",
		FieldIDs:     []string{"1", "8"},
		FieldsPrefix: "RadioR-REC",
	},
	{
		ObjectPath:   "/Radio Rundown/<OM_RECORD>/Hourly Rundown",
		FieldsPath:   "/OM_HEADER/OM_FIELD",
		FieldIDs:     []string{"1", "8"},
		FieldsPrefix: "HourlyR-HED",
	},
	{
		ObjectPath:   "/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>",
		FieldsPath:   "/OM_FIELD",
		FieldIDs:     []string{"8"},
		FieldsPrefix: "HourlyR-REC",
	},
	{
		ObjectPath: "/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>/Sub Rundown",
		FieldsPath: "/OM_HEADER/OM_FIELD",
		FieldIDs: []string{
			"5001", "8", "1004", "1003", "1005", "321", "5079"},
		FieldsPrefix: "SubR-HED",
	},
	{
		ObjectPath: "/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>/Sub Rundown/<OM_RECORD>",
		FieldsPath: "/OM_FIELD",
		FieldIDs: []string{
			"5001", "8", "1004", "1003", "1005", "321", "5079"},
		FieldsPrefix: "Sub-REC",
	},
	{
		ObjectPath: "/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>/Sub Rundown/<OM_RECORD>/Radio Story",
		FieldsPath: "/OM_HEADER/OM_FIELD",
		FieldIDs: []string{
			"5001", "8", "5081", "1004", "1003", "1005", "1036", "1029", "1010", "1002", "321", "5079", "16", "5082", "5072", "5016", "5", "6", "12", "5071", "5070"},
		FieldsPrefix: "RStory-HED",
	},
	{
		ObjectPath: "/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>/Sub Rundown/<OM_RECORD>/Radio Story/<OM_RECORD>",
		FieldsPath: "/OM_FIELD",
		FieldIDs: []string{
			"5001", "8", "5081", "1004", "1003", "1005", "1036", "1029", "1010", "1002", "321", "5079", "16", "5082", "5072", "5016", "5", "6", "12", "5071", "5070"},
		FieldsPrefix: "RStory-REC",
	},
	// {
	// ObjectPath:   "/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>/Sub Rundown/<OM_RECORD>/Radio Story",
	// FieldsPath:   "/OM_HEADER/OM_FIELD",
	// FieldIDs:     []string{"8"},
	// FieldsPrefix: "RStory-HED",
	// },
	// {
	// ObjectPath:   "/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>/Radio Story/Contact Item",
	// FieldsPath:   "/OM_HEADER/OM_FIELD",
	// FieldIDs:     []string{"8", "10"},
	// FieldsPrefix: "CI-HED",
	// },
}

var EXTproductionTest = OMextractors{
	{
		ObjectPath:   "/Radio Rundown",
		FieldsPath:   "/OM_HEADER/OM_FIELD",
		FieldIDs:     []string{"5081", "1", "8", "1003", "1004"},
		FieldsPrefix: "RadioR-HED",
	},
	{
		ObjectPath:   "/Radio Rundown/<OM_RECORD>",
		FieldsPath:   "/OM_FIELD",
		FieldIDs:     []string{"1", "8"},
		FieldsPrefix: "RadioR-REC",
	},
	{
		ObjectPath:   "/Radio Rundown/<OM_RECORD>/Hourly Rundown",
		FieldsPath:   "/OM_HEADER/OM_FIELD",
		FieldIDs:     []string{"1", "8"},
		FieldsPrefix: "HourlyR-HED",
	},
	{
		ObjectPath:   "/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>",
		FieldsPath:   "/OM_FIELD",
		FieldIDs:     []string{"8"},
		FieldsPrefix: "HourlyR-REC",
	},
	{
		ObjectPath:   "/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>/Sub Rundown",
		FieldsPath:   "/OM_HEADER/OM_FIELD",
		FieldIDs:     []string{"5001", "8", "1004", "1003", "1005", "321", "5079"},
		FieldsPrefix: "SubR-HED",
	},
	{
		ObjectPath:   "/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>/Sub Rundown/<OM_RECORD>",
		FieldsPath:   "/OM_FIELD",
		FieldIDs:     []string{"5001", "8", "1004", "1003", "1005", "321", "5079"},
		FieldsPrefix: "Sub-REC",
	},
	{
		ObjectPath:   "/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>/Sub Rundown/<OM_RECORD>/Radio Story",
		FieldsPath:   "/OM_HEADER/OM_FIELD",
		FieldIDs:     []string{"5001", "8", "5081", "1004", "1003", "1005", "1036", "1029", "1010", "1002", "321", "5079", "16", "5082", "5072", "5016", "5", "6", "12", "5071", "5070"},
		FieldsPrefix: "RStory-HED",
	},
	{
		ObjectPath:   "/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>/Sub Rundown/<OM_RECORD>/Radio Story/<OM_RECORD>",
		FieldsPath:   "/OM_FIELD",
		FieldIDs:     []string{"5001", "8", "5081", "1004", "1003", "1005", "1036", "1029", "1010", "1002", "321", "5079", "16", "5082", "5072", "5016", "5", "6", "12", "5071", "5070"},
		FieldsPrefix: "RStory-REC",
	},
	// {
	// ObjectPath:   "/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>/Sub Rundown/<OM_RECORD>/Radio Story",
	// FieldsPath:   "/OM_HEADER/OM_FIELD",
	// FieldIDs:     []string{"8"},
	// FieldsPrefix: "RStory-HED",
	// },
	// {
	// ObjectPath:   "/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>/Radio Story/Contact Item",
	// FieldsPath:   "/OM_HEADER/OM_FIELD",
	// FieldIDs:     []string{"8", "10"},
	// FieldsPrefix: "CI-HED",
	// },
}

var EXTproductionWorks = OMextractors{
	{
		ObjectPath:   "/Radio Rundown",
		FieldsPath:   "/OM_HEADER/OM_FIELD",
		FieldIDs:     []string{"1", "8"},
		FieldsPrefix: "RR-He",
	},
	{
		ObjectPath:   "/Radio Rundown/<OM_RECORD>",
		FieldsPath:   "/OM_FIELD",
		FieldIDs:     []string{"8"},
		FieldsPrefix: "HR-Re",
	},
	{
		ObjectPath:   "/Radio Rundown/<OM_RECORD>/Hourly Rundown",
		FieldsPath:   "/OM_HEADER/OM_FIELD",
		FieldIDs:     []string{"1", "8"},
		FieldsPrefix: "H-He",
	},
	{
		ObjectPath:   "/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>",
		FieldsPath:   "/OM_FIELD",
		FieldIDs:     []string{"1", "8"},
		FieldsPrefix: "H-Re",
	},
	// {
	// ObjectPath:   "/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>/Sub Rundown",
	// FieldsPath:   "OM_HEADER/OM_FIELD",
	// FieldIDs:     []string{"1", "2", "3", "4"},
	// FieldsPrefix: "S-He",
	// FieldIDs:     []string{"*"}, -> No ID, NO VALUE
	// },
}

var EXTall = OMextractors{
	{
		ObjectPath:   "/Radio Rundown",
		FieldsPath:   "/OM_HEADER/OM_FIELD",
		FieldIDs:     []string{"*"},
		FieldsPrefix: "R-He",
	},
	{
		ObjectPath:   "/Radio Rundown/<OM_RECORD>",
		FieldsPath:   "/OM_FIELD",
		FieldsPrefix: "R-Re",
	},
	{
		ObjectPath:   "/Radio Rundown/<OM_RECORD>/Hourly Rundown",
		FieldsPath:   "/OM_HEADER/OM_FIELD",
		FieldsPrefix: "H-He",
	},
	{
		ObjectPath:   "/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>",
		FieldsPath:   "/OM_FIELD",
		FieldsPrefix: "H-Re",
	},
	{
		ObjectPath:   "/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>/Sub Rundown",
		FieldsPath:   "OM_HEADER/OM_FIELD",
		FieldsPrefix: "S-He",
	},
	// {
	// ObjectPath:   "/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>/Sub Rundown/<OM_RECORD>",
	// FieldsPath:   "OM_FIELD",
	// FieldIDs:     []string{"1", "8"},
	// FieldsPrefix: "S-He",
	// },

	// VAR1
	// {
	// ObjectPath:   "/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>/Sub Rundown/<OM_RECORD>/Radio Story",
	// FieldsPath:   "OM_HEADER/OM_FIELD",
	// FieldIDs:     []string{"1", "8"},
	// FieldsPrefix: "S-He",
	// },

	// VAR2
	// {
	// ObjectPath:   "/Radio Story",
	// FieldsPath:   "OM_HEADER/OM_FIELD",
	// FieldIDs:     []string{"1", "8"},
	// FieldsPrefix: "S-He",
	// },
	// {
	// ObjectPath:   "/Radio Story/<OM_RECORD>",
	// FieldsPath:   "OM_HEADER/OM_FIELD",
	// FieldIDs:     []string{"1", "8"},
	// FieldsPrefix: "S-He",
	// },
	// {
	// ObjectPath:   "/Radio Story/<OM_RECORD>/Contact Item",
	// FieldsPath:   "OM_HEADER/OM_FIELD",
	// FieldIDs:     []string{"1", "8"},
	// FieldsPrefix: "S-He",
	// },
	// ObjectPath:   "/Radio Story/<OM_RECORD>/AudioClip",
	// FieldsPath:   "OM_HEADER/OM_FIELD",
	// FieldIDs:     []string{"1", "8"},
	// FieldsPrefix: "S-He",
	// },
}

// OmObject:   "Radio Story",
// ObjectPath: "//AudioClip",
// OmObject:   "Radio Story",
// ObjectPath: "//Contact Item",

// "//OM_OBJECT[@TemplateName='%s']/%s/*", ext.OM_type, ext.Path,
