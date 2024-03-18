package internal

var EXTproduction = OMobjExtractors{
	{
		ObjectPath:   "/Radio Rundown",
		FieldsPath:   "/OM_HEADER/OM_FIELD",
		FieldIDs:     []string{"1", "8"},
		FieldsPrefix: "R-He",
	},
	// {
	// ObjectPath:   "/Radio Rundown/<OM_RECORD>",
	// FieldsPath:   "/OM_FIELD",
	// FieldIDs:     []string{"1", "8"},
	// FieldsPrefix: "R-Re",
	// },
	// {
	// ObjectPath: "/Radio Rundown/<OM_RECORD>/Hourly Rundown",
	// FieldsPath: "/OM_HEADER/OM_FIELD",
	// FieldIDs:     []string{"1", "8"},
	// FieldIDs:     []string{"1", "2", "3", "4", "5", "8", "5088"},
	// FieldsPrefix: "H-He",
	// },
	// {
	// ObjectPath:   "/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>",
	// FieldsPath:   "/OM_FIELD",
	// FieldIDs:     []string{"1", "8"},
	// FieldsPrefix: "H-Re",
	// },
	// {
	// ObjectPath:   "/Radio Rundown/<OM_RECORD>/Hourly Rundown/<OM_RECORD>/Sub Rundown",
	// FieldsPath:   "OM_HEADER/OM_FIELD",
	// FieldIDs:     []string{"1", "2", "3", "4"},
	// FieldsPrefix: "S-He",
	// FieldIDs:     []string{"*"}, -> No ID, NO VALUE
	// },
}

var EXTall = OMobjExtractors{
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
