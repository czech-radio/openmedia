package internal

import (
	"encoding/json"
	"encoding/xml"
	"time"

	"github.com/ncruces/go-strftime"
)

var lineEnd = []byte("\n")
var openMediaXmlHeader []byte = append([]byte(
	`<?xml version="1.0" encoding="UTF-8" standalone="no" ?>
<!DOCTYPE OPENMEDIA SYSTEM "ann_objects.dtd">`,
), lineEnd...)

type OPENMEDIA struct {
	XMLName   xml.Name  `xml:"OPENMEDIA"`
	OM_SERVER OM_SERVER `xml:"OM_SERVER"`
	OM_OBJECT OM_OBJECT `xml:"OM_OBJECT"`
}

type RundownMetaInfo struct {
	Date      time.Time
	RadioName string
}

// RundownDate: get date from xml rundown
// TODO: optimize without loops
func (om OPENMEDIA) RundownMetaInfoParse() (RundownMetaInfo, error) {
	fields := om.OM_OBJECT.OM_HEADER.Fields
	var metaInfo RundownMetaInfo
	var err error
	for _, field := range fields {
		for _, attr := range field.Attrs {
			switch attr.Value {
			case "Čas vytvoření":
				// field.FieldID == 1004
				metaInfo.Date, err = strftime.Parse("%Y%m%dT%H%M%S", field.OM_DATETIME)
			case "Název":
				metaInfo.RadioName = field.OM_STRING
			}
			if err != nil {
				return metaInfo, err
			}
		}
	}
	return metaInfo, err
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
