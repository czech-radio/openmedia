package extract

import "strings"

// RowPartCode specifies part of table row (group of row fields) corresponing to xml OM_OBJECT attributes, OM_HEADER or OM_RECORD fields
type RowPartCode int

var CSVdelim = "\t"

// RowPartCodes
const (
	RowPartCode_RadioRec RowPartCode = iota
	RowPartCode_RadioHead
	RowPartCode_HourlyHead
	RowPartCode_HourlyRec
	RowPartCode_SubHead
	RowPartCode_SubRec
	RowPartCode_StoryHead
	RowPartCode_StoryRec
	RowPartCode_StoryRecContact
	RowPartCode_StoryRecAudio
	RowPartCode_StorySpec
	RowPartCode_AudioClipHead
	RowPartCode_AudioClipRec
	RowPartCode_ContactItemHead
	RowPartCode_ContactItemRec
	RowPartCode_StoryKategory
	RowPartCode_Record
	RowPartCode_ComputedKON
	RowPartCode_ComputedKategory
	RowPartCode_ComputedRID
)

// RowPartName
type RowPartName struct {
	Internal, External string
}

type PartsPrefixMap = map[RowPartCode]RowPartName

// RowPartsCodeMapProduction is a map which translates internal column header (prefix/sufix) name to name used in analytics. It represents the part of row which coresponds to one xml OM_OBJECT.
var RowPartsCodeMapProduction = PartsPrefixMap{
	RowPartCode_RadioRec:         {"Radio-REC", "RR"},
	RowPartCode_RadioHead:        {"Radio-HED", "RR"},
	RowPartCode_HourlyHead:       {"Hourly-HED", "HR"},
	RowPartCode_HourlyRec:        {"Hourly-REC", "HR"},
	RowPartCode_SubHead:          {"Sub-HED", "SR"},
	RowPartCode_SubRec:           {"Sub-REC", "SR"},
	RowPartCode_StoryHead:        {"Story-HED", ""},
	RowPartCode_StoryRec:         {"Story-REC", "sRec"},
	RowPartCode_StoryRecAudio:    {"Story-REC_AUD", "saRec"},
	RowPartCode_StoryRecContact:  {"Story-REC_CONT", "scRec"},
	RowPartCode_StorySpec:        {"Story-SPEC", ""},
	RowPartCode_StoryKategory:    {"Story-Cat", "CAST"},
	RowPartCode_AudioClipHead:    {"Audio-HED", "AUD"},
	RowPartCode_AudioClipRec:     {"Audio-REC", "AUDrec"},
	RowPartCode_ContactItemHead:  {"Contact-HED", "KON"},
	RowPartCode_ContactItemRec:   {"Contact-REC", "KONrec"},
	RowPartCode_ComputedKON:      {"Comp-KON", "KON"},
	RowPartCode_ComputedKategory: {"Comp-Cat", "kategory"},
	RowPartCode_ComputedRID:      {"Comp-RID", "HLP"},
}

type StoryPartCode string

const (
	StoryPartAudio       StoryPartCode = "Audio"
	StoryPartContactItem StoryPartCode = "Contact Item"
	StoryPartContactBin  StoryPartCode = "Contact Bin"
	StoryPartUnknown     StoryPartCode = "UNKNOWN"
)

type FieldID struct {
	NameShort        string
	NameLong         string
	XLSXcolumnFormat int
	XLSXcustomFormat string
	Width            float64
}

type FieldsIDsNamesMap map[string]FieldID

func (fi FieldsIDsNamesMap) GetByName(name string) (FieldID, bool) {
	for _, value := range fi {
		if value.NameShort == name {
			return value, true
		}
		out := strings.Split(name, "_")
		if len(out) == 1 {
			continue
		}
		cur := strings.Join(out[:len(out)-1], "_")
		if value.NameShort == cur {
			return value, true
		}
	}
	return FieldID{}, false
}

var FieldWidthDefault = float64(10)
var FieldWidthShort = float64(6)

var FieldsIDsNames = FieldsIDsNamesMap{
	"1":                {"cas_vytvoreni", "Čas vytvoření", 0, "", 20},
	"1000":             {"datum", "", 14, "DD. MM. R", FieldWidthDefault},
	"1002":             {"planovana_stopaz", "Plánovaná stopáž", 0, "@", FieldWidthDefault},
	"1003":             {"cas_konce", "Čas konce", 0, "@", FieldWidthDefault},
	"1004":             {"cas_zacatku", "Čas začátku", 0, "@", FieldWidthDefault},
	"1005":             {"stopaz", "Stopáž", 0, "@", FieldWidthDefault},
	"1010":             {"spoctena_stopaz", "Spočtená stopáž", 0, "@", FieldWidthDefault},
	"1029":             {"korekce", "Korekce", 0, "@", FieldWidthDefault},
	"1035":             {"cas_textu", "Čas textu", 0, "@", FieldWidthDefault},
	"1036":             {"audio_stopaz", "Audio Stopáž", 0, "@", FieldWidthDefault},
	"12":               {"redakce", "Redakce", 0, "@", FieldWidthDefault},
	"16":               {"druh", "Druh", 0, "@", FieldWidthDefault},
	"321":              {"format", "Formát", 0, "@", FieldWidthDefault},
	"38":               {"stopaz", "Stopáž", 0, "@", FieldWidthDefault},
	"406":              {"region", "Region", 0, "@", FieldWidthDefault},
	"421":              {"jmeno", "Jméno", 0, "@", FieldWidthDefault},
	"422":              {"prijmeni", "Příjmení", 0, "@", FieldWidthDefault},
	"423":              {"spolecnost", "Společnost", 0, "@", FieldWidthDefault},
	"424":              {"funkce", "Funkce", 0, "@", FieldWidthDefault},
	"5":                {"vytvoril", "Vytovřil", 0, "@", FieldWidthDefault},
	"5015":             {"strana", "Politická příslušnost", 0, "@", FieldWidthDefault},
	"5016":             {"tema", "Téma", 0, "@", FieldWidthDefault},
	"5068":             {"ID_5068", "ContactContainerFieldID", 0, "@", FieldWidthDefault},
	"5070":             {"schvalil_redakce", "Schválil redakce", 0, "@", FieldWidthDefault},
	"5071":             {"schvalil_stanice", "Schválil stanice", 0, "@", FieldWidthDefault},
	"5072":             {"incode", "IN_code", 0, "@", FieldWidthDefault},
	"5079":             {"cil_vyroby", "Cíl výroby", 0, "@", FieldWidthDefault},
	"5081":             {"stanice", "Stanice", 1, "@", FieldWidthDefault},
	"5082":             {"itemcode", "ItemCode", 0, "@", FieldWidthDefault},
	"5087":             {"ID_5087", "CustomUniqueID2", 0, "@", FieldWidthDefault},
	"5088":             {"pohlavi", "Gender", 0, "@", FieldWidthDefault},
	"6":                {"autor", "Autor", 0, "@", FieldWidthDefault},
	"8":                {"nazev", "Název", 0, "@", 22},
	"CRID":             {"CRID", "Computed RecordID", 0, "@", FieldWidthDefault},
	"Cindex":           {"index", "", 0, "@", 24},
	"FileName":         {"FileName", "FileName", 0, "@", FieldWidthDefault},
	"Filename_CRID":    {"Filename_CRID", "Flename w Computed RecordIDs", 0, "@", FieldWidthDefault},
	"ID":               {"compID", "Computed ID", 0, "@", FieldWidthDefault},
	"ObjectID":         {"ObjectID", "ObjectID", 0, "@", FieldWidthDefault},
	"RecordID":         {"RID", "RecordID", 0, "@", FieldWidthDefault},
	"TemplateName":     {"kategorie", "Kategorie", 0, "@", FieldWidthDefault},
	"datum":            {"datum", "Datum", 0, "@", FieldWidthDefault},
	"jmeno_spojene":    {"jmeno_spojene", "Jméno Spojené", 0, "@", FieldWidthDefault},
	"kategory":         {"kategory", "Kategory", 0, "@", FieldWidthDefault},
	"name&party_match": {"kontrola_strany", "Kontrola strany", 0, "@", FieldWidthShort},
	"name_match":       {"name_match", "Name match", 0, "@", FieldWidthShort},
	"referencni_jmeno": {"referencni_jmeno", "Referencni jmeno", 0, "@", FieldWidthShort},
	"vysoka_politika":  {"vysoka_politika", "Vysoká politika", 0, "@", FieldWidthShort},
}

type RowFieldSpecialValueCode int

const (
	RowFieldValueEmptyString = iota
	RowFieldValueNotPossible
	RowFieldValueNotContain
	RowFieldValueNotValid
	RowFieldValueChildNotFound
	RowFieldValueParentNotFound
)

var RowFieldSpecialValueCodeMap = map[RowFieldSpecialValueCode]string{
	RowFieldValueEmptyString:    "(NS)", // (NOT SPECIFIED), (NEUVEDENO)
	RowFieldValueNotPossible:    "(NP)", // (NOT POSSIBLE), (NELZE)
	RowFieldValueNotContain:     "(NC)", // (NOT CONTAIN), (NEOBSAHUJE)
	RowFieldValueNotValid:       "(NV)", // (NOT VALID), (INVALID)
	RowFieldValueChildNotFound:  "(NP)", // (NOT POSIBLE), (NELZE)
	RowFieldValueParentNotFound: "(NC)", // (NOT CONTAIN), (NEOBSAHUJE)
}

func CheckIfFieldValueIsSpecialValue(fieldValue string) bool {
	for _, sval := range RowFieldSpecialValueCodeMap {
		if sval == fieldValue {
			return true
		}
	}
	return false
}

func CheckIfMapContainsKeyValue(inMap map[RowFieldSpecialValueCode]string, value string) bool {
	for _, spec := range inMap {
		if spec == value {
			return true
		}
	}
	return false
}

type RadioSationIDs struct {
	Openmedia_stanice string
	Openmedia_ID      string
	Croapp_code       string
	Croapp_shortTitle string
	Croapp_ID         string
}

type RadioCodesMap map[string]RadioSationIDs

var RadioCodes = RadioCodesMap{
	"5": {"CRo-Český rozhlas", "5", "", "", ""},
	"11": {"RZ-Radiožurnál", "11", "radiozurnal", "Radiožurnál",
		"4082f63f-30e8-375d-a326-b32cf7d86e02"},
	"13": {"PS-Plus", "13", "plus", "Plus",
		"c639d927-f37b-3db8-a64f-1d64ee1469b2"},
	"15": {"DV-Dvojka", "15", "dvojka", "Dvojka",
		"17821883-be2d-3880-b971-eceb794388fa"},
	"17": {"VL-Vltava", "17", "vltava", "Vltava",
		"0134ce01-8684-3556-b568-f208392ac0bd"},
	"19": {"WA-Wave", "19", "radiowave", "Radio Wave",
		"6ab28be7-cdc8-3222-bd6e-c229553125fb"},
	"21": {"RJ-Rádio Junior", "21", "radiojunior", "Rádio Junior",
		"598a62af-86b5-3485-b89d-65379562694a"},
	"23": {"ZV-Radio Prague International", "23", "cro7", "Radio Prague Int.",
		"6731e7ee-98e7-36c3-993b-3122ad1317d1"},
	"31": {"RD-Radio DAB Praha", "31", "", "", ""},
	"33": {"SC-Region Střední Čechy", "33", "strednicechy", "Region",
		"f8b5ee83-c786-3c67-8e37-4fdee1007147"},
	"35": {"PN-Plzeň", "35", "plzen", "Plzeň",
		"bbe67c0f-4848-355e-a0ab-28f40ade3d38"},
	"37": {"KV-Karlovy Vary", "37", "kv", "Karlovy Vary",
		"663e7133-bb52-3477-8d7f-635e43371962"},
	"39": {"SE-Sever", "39", "sever", "Sever",
		"5e94e06f-5435-3291-87c5-aeca8bc9c884"},
	"41": {"LB-Liberec", "41", "liberec", "Liberec",
		"ad541211-198e-30b7-995c-10bf08c3aea0"},
	"43": {"HK-Hradec Králové", "43", "hradec", "Hradec Králové",
		"a831a457-9b80-3271-b153-ddf0ee63a18c"},
	"45": {"PC-Pardubice", "45", "pardubice", "Pardubice",
		"67e26b4a-af56-3153-bb09-eaaaa61a7fc7"},
	"47": {"CB-České Budějovice", "47", "cb", "České Budějovice",
		"1e6c0c0b-aba4-3357-95fc-c64edc9e75e6"},
	"49": {"VY-Vysočina", "49", "vysocina", "Vysočina",
		"36dc9675-1fe8-3c69-b362-ee6784b04ef0"},
	"51": {"BO-Brno", "51", "brno", "Brno",
		"184a888f-3d06-3a88-a07f-901bd9b09396"},
	"53": {"OL-Olomouc", "53", "olomouc", "Olomouc",
		"c49f6468-acd3-3f98-9db9-92d5e5c0f038"},
	"55": {"OV-Ostrava", "55", "ostrava", "Ostrava",
		"318fc506-4dc9-33c6-b0f4-88df17cafa20"},
	"57": {"ZL-Zlín", "57", "zlín", "Zlín",
		"b0f03203-0809-3363-bb3d-ccda436d6760"},
	"73": {"REG-všem regionům", "73", "", "", ""},
	"75": {"REGIONY-NOC", "75", "", "", ""},
	"25": {"SP-RŽ Sport", "25", "", "", ""},
	"27": {"PO-Pohoda", "27", "", "", ""},
}

var GenderCodes = map[string]string{
	"0": "(NEUTR-0)",
	"1": "Male",
	"2": "Female",
}
