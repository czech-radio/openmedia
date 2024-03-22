package internal

type PartPrefixCode int

const (
	FieldPrefix_HourlyHead PartPrefixCode = iota
	FieldPrefix_HourlyRec
	FieldPrefix_Subhead
	FieldPrefix_Subrec
	FieldPrefix_StoryHead
	FieldPrefix_StoryRec
	FieldPrefix_AudioClipHead
	FieldPrefix_AudioClipRec
	FieldPrefix_ContactItemHead
	FieldPrefix_ContactItemRec
)

type PartPrefix struct {
	Internal, External string
}

type PartsPrefixMap = map[PartPrefixCode]PartPrefix

var PartsPrefixMapProduction = PartsPrefixMap{
	FieldPrefix_HourlyHead:      {"Hourly-HED", "blok"},
	FieldPrefix_HourlyRec:       {"Hourly-REC", "blok"},
	FieldPrefix_Subhead:         {"Sub-HED", "SP"},
	FieldPrefix_Subrec:          {"Sub-REC", "SP"},
	FieldPrefix_StoryHead:       {"Story-HED", "P"},
	FieldPrefix_StoryRec:        {"Story-REC", "P"},
	FieldPrefix_AudioClipHead:   {"Audio-HED", "AUD"},
	FieldPrefix_AudioClipRec:    {"Audio-REC", "AUD"},
	FieldPrefix_ContactItemHead: {"Contact-HED", "KON"},
	FieldPrefix_ContactItemRec:  {"Contact-REC", "KON"},
}

type FieldsIDsNames map[string]string

var FieldsIDsNamesProduction = FieldsIDsNames{
	"8":    "Název",
	"1004": "Čas začátku",
	"1003": "Čas konce",
	"1005": "Stopáž",
	"321":  "Formát",
	"5081": "Stanice",
	"1036": "Audio stopáž",
	"1029": "Korekce",
	"1010": "Spočtená stopáž",
	"1002": "Plánovaná stopáž",
	"5079": "Cíl výroby",
	"16":   "Druh",
	"5082": "ItemCode",
	"5072": "IN_Code",
	"5016": "Téma",
	"5":    "Vytvořil",
	"6":    "Autor",
	"12":   "Redakce",
	"5071": "Schválil stanice",
	"5070": "Schválil redakce",
	"421":  "Jméno",
	"422":  "Příjmení",
	"423":  "Společnost",
	"424":  "Funkce",
	"5015": "Politická příslušnost",
	"5087": "CustomUniqueID2",
	"5088": "Gender",
}
