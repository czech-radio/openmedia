package internal

type FieldsPrefixesCodes int

const (
	FieldPrefix_Hourly FieldsPrefixesCodes = iota
	// FieldPrefix_HourlyRundownRec
	FieldPrefix_SubRundownHed
	FieldPrefix_SubRundownRec
	FieldPrefix_RadioStoryHed
	FieldPrefix_RadioStoryRec
	FieldPrefix_AudioClipHed
	FieldPrefix_AudioClipRec
	FieldPrefix_ContactItemHed
	FieldPrefix_ContactItemRec
)

var InternalFieldPrefixes = map[FieldsPrefixesCodes]string{
	// FieldPrefix_HourlyRundown: "HourlyR-HED",
	// FieldPrefix_SubRundown:    "SubR-HED",
	// FieldPrefix_RadioStory:    "Rstory-HED",
	// FieldPrefix_AudioClip
	// FieldPrefix_ContactItem
}
