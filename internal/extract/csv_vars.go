package extract

type PartPrefixCode int

var CSVdelim = "\t"

type CSVheaderCodeName int

const (
	CSVheaderInternal CSVheaderCodeName = iota
	CSVheaderExternal
)

const (
	FieldPrefix_RadioRec PartPrefixCode = iota
	FieldPrefix_HourlyHead
	FieldPrefix_HourlyRec
	FieldPrefix_SubHead
	FieldPrefix_SubRec
	FieldPrefix_StoryHead
	FieldPrefix_StoryRec
	FieldPrefix_AudioClipHead
	FieldPrefix_AudioClipRec
	FieldPrefix_ContactItemHead
	FieldPrefix_ContactItemRec
	FieldPrefix_StoryKategory
	FieldPrefix_Record
	FieldPrefix_ComputedID
	FieldPrefix_ComputedKategory
	FieldPrefix_ComputedRID
)

type PartPrefix struct {
	Internal, External string
}

type PartsPrefixMap = map[PartPrefixCode]PartPrefix

var PartsPrefixMapProduction = PartsPrefixMap{
	FieldPrefix_RadioRec:         {"Radio-REC", "radio"},
	FieldPrefix_HourlyHead:       {"Hourly-HED", "blok"},
	FieldPrefix_HourlyRec:        {"Hourly-REC", "blok"},
	FieldPrefix_SubHead:          {"Sub-HED", "SP"},
	FieldPrefix_SubRec:           {"Sub-REC", "SP-REC"},
	FieldPrefix_StoryHead:        {"Story-HED", ""},
	FieldPrefix_StoryRec:         {"Story-REC", ""},
	FieldPrefix_StoryKategory:    {"Story-Cat", "CAST"},
	FieldPrefix_AudioClipHead:    {"Audio-HED", "AUD"},
	FieldPrefix_AudioClipRec:     {"Audio-REC", "AUD"},
	FieldPrefix_ContactItemHead:  {"Contact-HED", "KON"},
	FieldPrefix_ContactItemRec:   {"Contact-REC", "KON"},
	FieldPrefix_ComputedID:       {"Comp-ID", "ID"},
	FieldPrefix_ComputedKategory: {"Comp-Cat", "kategory"},
	FieldPrefix_ComputedRID:      {"Comp-RID", ""},
}

type FieldsIDsNames map[string]string

var FieldsIDsNamesProduction = FieldsIDsNames{
	"1":            "cas_vytvoreni",
	"1002":         "planovana_stopaz",
	"1003":         "cas_konce",
	"1004":         "cas_zacatku",
	"1005":         "stopaz",
	"1010":         "spoctena_stopaz",
	"1029":         "korekce",
	"1035":         "cas_textu",
	"1036":         "audio_stopaz",
	"12":           "redakce",
	"16":           "druh",
	"321":          "format",
	"38":           "stopaz",
	"421":          "jmeno",
	"422":          "prijmeni",
	"423":          "spolecnost",
	"424":          "funkce",
	"5":            "vytvoril",
	"5015":         "strana",
	"5016":         "tema",
	"5070":         "schvalil_redakce",
	"5071":         "schvalil_stanice",
	"5072":         "incode",
	"5079":         "cil_vyroby",
	"5081":         "stanice",
	"5082":         "itemcode",
	"5087":         "ID",
	"5088":         "pohlavi",
	"6":            "autor",
	"8":            "nazev",
	"ID":           "compID",
	"RecordID":     "RID",
	"TemplateName": "kategorie",
	"datum":        "datum",
	"kategory":     "kategory",
	"C-RID":        "RID",
	"C-index":      "index",
	"ObjectID":     "ObjectID",
	"FileName":     "FileName",
	"filtered":     "filtered",
}

var FieldsIDsNamesProductionLong = FieldsIDsNames{
	"1":            "Čas vytvoření",
	"TemplateName": "kategorie",
	"RecordID":     "RID",
	"8":            "Název",
	"1004":         "Čas začátku",
	"1003":         "Čas konce",
	"1005":         "Stopáž",
	"321":          "Formát",
	"5081":         "Stanice",
	"1036":         "Audio stopáž",
	"1029":         "Korekce",
	"1010":         "Spočtená stopáž",
	"1002":         "Plánovaná stopáž",
	"5079":         "Cíl výroby",
	"16":           "Druh",
	"5082":         "ItemCode",
	"5072":         "IN_Code",
	"5016":         "Téma",
	"5":            "Vytvořil",
	"6":            "Autor",
	"12":           "Redakce",
	"5071":         "Schválil stanice",
	"5070":         "Schválil redakce",
	"421":          "Jméno",
	"422":          "Příjmení",
	"423":          "Společnost",
	"424":          "Funkce",
	"5015":         "Politická příslušnost",
	"5087":         "CustomUniqueID2",
	"5088":         "Gender",
	"1035":         "Čas textu",
}

type CSVspecialValueCode int

// (NS)	Všechny buňky, které je možné vyplnit skrze formuláře (příspěvku, audia, kontaktu, subrundownu) a které ovšem vyplněny nejsou, přepíšeme hodnotou "(NS)", tj. "not specified".
// (NC)	Jestliže příspěvek neobsahuje žádnou část, sloupec [kategorie_CAST]bude obsahovat hodnotu "(NC)".
// Jestliže příspěvek není součástí žádného subrundownu, sloupec [kategorie_SP]bude obsahovat hodnotu "(NC)".
// (NP)	Jestliže příspěvek neobsahuje žádnou část, pak všechny další sloupce, které se týkají částí příspěvku (audia i kontaktu) budou obsahovat hodnotu "(NP)". (S výjimkou sloupce [kategorie_CAST], který bude obsahovat hodnotu "(NC)" - viz výše).
// Jestliže příspěvek obsahuje audio, sloupce týkající se kontaktu budou vyplněny hodnotou "(NP)". Jestliže příspěvek obsahuje kontakt, sloupce týkající se audia budou vyplněny hodnotou "(NP)".
// Jestliže příspěvek není součástí žádného subrundownu, všechny sloupce týkající se subrundownu budou vyplněny hodnotou "(NP)". (S výjimkou sloupce [kategorie_SP], který bude obsahovat hodnotu "(NC)" - viz výše).

// Později
// (NV)	Pro potřeby validace dat pak v další fázi práce na analýze produkce přibyde hodnota (NV), tj. "not valid" či "invalid". Ta bude označovat buňky vyplněné neplatnou hodntodou, takže místo [redakce] = "banán" tam pak bude [redakce] = "(NV)". Vedle toho bude seznam těch přepsaných hodnot s jejich četnostmi, takže tam pak bude k proměnné [stanice] údaj "banán" = 1krát.

const (
	CSVspecialValueEmptyString = iota
	CSVspecialValueChildNotFound
	CSVspecialValueParentNotFound
)

var CSVspecialValues = map[CSVspecialValueCode]string{
	CSVspecialValueEmptyString:    "(NS)", // (NEUVEDENO)
	CSVspecialValueChildNotFound:  "(NP)", // (NELZE)
	CSVspecialValueParentNotFound: "(NC)", // (NEOBSAHUJE)
}

type Radio struct {
	Openmedia_stanice string
	Openmedia_ID      string
	Croapp_code       string
	Croapp_shortTitle string
	Croapp_ID         string
}

type RadioCodesMap map[string]Radio

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