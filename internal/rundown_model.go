package internal

import (
	"encoding/json"
	"encoding/xml"
)

var lineEnd = []byte("\n")
var openMediaXmlHeader []byte = append([]byte(
	`<?xml version="1.0" encoding="UTF-8" standalone="no" ?>
<!DOCTYPE OPENMEDIA SYSTEM "ann_objects.dtd">`,
), lineEnd...)

type OPENMEDIA struct {
	XMLName   xml.Name  `xml:"OPENMEDIA"`
	OM_SERVER OM_SERVER `xml:"OM_SERVER,omitempty"`
	OM_OBJECT OM_OBJECT `xml:"OM_OBJECT,omitempty"`
}

type OM_OBJECT struct {
	Attrs      []xml.Attr  `xml:",any,attr"`
	OM_HEADER  OM_HEADER   `xml:"OM_HEADER,omitempty"`
	OM_UPLINK  OM_UPLINK   `xml:"OM_UPLINK,omitempty"`
	OM_RECORDS []OM_RECORD `xml:"OM_RECORD,omitempty"`
}

type OM_SERVER struct {
	Attrs []xml.Attr `xml:",any,attr"`
}

type OM_RECORD struct {
	Attrs      []xml.Attr  `xml:",any,attr"`
	Fields     []OM_FIELD  `xml:"OM_FIELD,omitempty"`
	OM_RECORDS []OM_RECORD `xml:"OM_RECORD,omitempty"`
	OM_OBJECTS OM_OBJECT   `xml:"OM_OBJECT"`
}

type OM_HEADER struct {
	Attrs  []xml.Attr `xml:",any,attr"`
	Fields []OM_FIELD `xml:"OM_FIELD,omitempty"`
}

type OM_UPLINK struct {
	Attrs []xml.Attr `xml:",any,attr"`
}

// OM_FIELD contais various nested tag names.
// Custom unmarshalXML method must be used.
type OM_FIELD struct {
	Attrs       []xml.Attr `xml:",any,attr"`
	OM_STRING   string     `xml:"OM_STRING,omitempty"`
	OM_DATETIME string     `xml:"OM_DATETIME,omitempty"`
	OM_TIMESPAN string     `xml:"OM_TIMESPAN,omitempty"`
	OM_INT32    string     `xml:"OM_INT32,omitempty"`
}

// OM_FIELD.MarshalXml filter empty, when it contains empty tags
func (omf *OM_FIELD) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(omf)
	err := json.Unmarshal(inrec, &inInterface)
	if err != nil {
		return err
	}
	maxEmptyCount := len(inInterface) - 1 // Filed attrs should never be empty
	var emptyCount int
	for _, val := range inInterface {
		if val == "" {
			emptyCount++
		}
	}
	if emptyCount == maxEmptyCount {
		return e.EncodeElement(nil, start)
	}
	return e.EncodeElement(*omf, start)
}

func (omh *OM_HEADER) ExtractFieldsByFieldID(ids []int) CSVrowFields {
	var row CSVrowFields
	// for _, field := range omh.Fields {
	// field.OM
	// row[field.]
	// }
	return row
}
