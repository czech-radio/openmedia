package internal

import (
	"encoding/json"
	"encoding/xml"
	"time"

	"github.com/ncruces/go-strftime"
)

var openMediaXmlHeader []byte = []byte(
	`<?xml version="1.0" encoding="UTF-8" standalone="no" ?>
<!DOCTYPE OPENMEDIA SYSTEM "ann_objects.dtd">`,
)

type OPENMEDIA struct {
	XMLName   xml.Name  `xml:"OPENMEDIA"`
	OM_SERVER OM_SERVER `xml:"OM_SERVER"`
	OM_OBJECT OM_OBJECT `xml:"OM_OBJECT"`
}

// RundownDate: get date from xml rundown
// TODO: optimize without loops
func (om OPENMEDIA) RundownDate() (time.Time, error) {
	fields := om.OM_OBJECT.OM_HEADER.Fields
	var fieldValue string
loopfields:
	for _, field := range fields {
		for _, attr := range field.Attrs {
			// field.FieldID == 1004
			if attr.Value == "Čas začátku" {
				fieldValue = field.OM_DATETIME
				break loopfields
			}
		}
	}
	return strftime.Parse("%Y%m%dT%H%M%S", fieldValue)
}

type OM_OBJECT struct {
	Attrs      []xml.Attr  `xml:",any,attr"`
	OM_HEADER  OM_HEADER   `xml:"OM_HEADER"`
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
// Custom unmarshalXML method must be used. It is faster to use map for attributes then usign struct fields then. (Reflect must be used, when iterating over struct fields)
type OM_FIELD struct {
	Attrs       []xml.Attr `xml:",any,attr"`
	OM_STRING   string     `xml:"OM_STRING,omitempty"`
	OM_DATETIME string     `xml:"OM_DATETIME,omitempty"`
	OM_TIMESPAN string     `xml:"OM_TIMESPAN,omitempty"`
	OM_INT32    string     `xml:"OM_INT32,omitempty"`
	// Attrs map[string]string `xml:"-"`
	// Value string `xml:",any"` // This will parse all remaining tags (it can parse all tags)
	// Value map[string]string
	// `xml:",any"`
}

// V1: OM_FIELD filter empty, when it contains empty tags
func (omf *OM_FIELD) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(omf)
	json.Unmarshal(inrec, &inInterface)
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

// V2: OM_FIELD alternative filter out empty. When parsing all nested tags as OM_FIELD.Value. 2 times faster than V1
// func (omf *OM_FIELD) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
// res := OM_FIELD{}
// if omf.Value != "" {
// res.Attrs = omf.Attrs
// res.Value = omf.Value
// return e.EncodeElement(*omf, start)
// }
// return e.EncodeElement(nil, start)
// }
