package internal

var EXTproduction = OMobjExtractors{
	{
		ObjectPath:   "/Radio Rundown",
		FieldsPath:   "/OM_HEADER/OM_FIELD",
		FieldIDs:     []string{"1", "8"},
		FieldsPrefix: "R-He",
	},
	{
		ObjectPath:   "/Radio Rundown/<OM_RECORD>",
		FieldsPath:   "/OM_FIELD",
		FieldIDs:     []string{"1", "8"},
		FieldsPrefix: "R-Re",
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
	{
		ObjectPath:   "/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>/Sub Rundown",
		FieldsPath:   "OM_HEADER/OM_FIELD",
		FieldIDs:     []string{"1", "8"},
		FieldsPrefix: "S-He",
	},
}

// OmObject:   "Radio Story",
// ObjectPath: "//AudioClip",
// OmObject:   "Radio Story",
// ObjectPath: "//Contact Item",

// "//OM_OBJECT[@TemplateName='%s']/%s/*", ext.OM_type, ext.Path,
