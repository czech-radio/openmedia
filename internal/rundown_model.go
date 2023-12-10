package internal

import (
	"encoding/json"
	"encoding/xml"
	"regexp"
	"strings"
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

type RundownMetaInfo struct {
	Date        time.Time
	RadioName   string
	HoursRange  string
	RundownType string
}

var RadioNameRegex = regexp.MustCompile(`([\d[:ascii:]]*)([\p{L}\ ]*)`)

func (r *RundownMetaInfo) ParseRundownName(rundownName string) error {
	// e.g.: "05-09 ČRo Karlovy Vary - Wed, 04.03.2020"
	//TODO: match result against map of code vs radio name
	var unparsedFields []string
	var radioName string
	var hoursRange string
	matches := RadioNameRegex.FindStringSubmatch(rundownName)
	switch len(matches) {
	case 3:
		hoursRange = strings.TrimSpace(matches[1])
		radioName = strings.TrimSpace(matches[2])
	}
	if hoursRange == "" {
		// errs = append(errs, fmt.Errorf("cannot parse hours range from: %s", rundownName))
	}
	r.RadioName = radioName
	r.HoursRange = hoursRange
	// if radioName == "" {
	// errs = append(errs, fmt.Errorf("cannot parse radio name from: %s", rundownName))
	// }
	return nil
}

// RundownMetaInfoParse
// TODO: optimize without loops
func (om OPENMEDIA) RundownMetaInfoParse() (RundownMetaInfo, error) {
	fields := om.OM_OBJECT.OM_HEADER.Fields
	// var noName = fmt.Errorf("cannot parse radio name")
	var metaInfo RundownMetaInfo
	var err error
	var errs []error
	// var rundownName string

	// Get values of main object attrs
	for _, attr := range om.OM_OBJECT.Attrs {
		switch attr.Name.Local {
		case "TemplateName":
			metaInfo.RundownType = attr.Value
		}
	}

	// Get values of main header fields
	for _, field := range fields {
		for _, attr := range field.Attrs {
			switch attr.Value {
			// case "Čas vytvoření": // FiledID: 1
			case "Čas začátku": // FieldID: 1004
				metaInfo.Date, err = strftime.Parse("%Y%m%dT%H%M%S", field.OM_DATETIME)
				// errs = ErrorAppend(errs, err)
				// errs = append(errs, fmt.Errorf("cannot parse date from: %s", field.OM_DATETIME))
			case "Název":
				rundownName = field.OM_STRING
				// errs = ErrorAppend()
				// errs = append(errs, metaInfo.ParseRundownName(rundownName)...)
			}
		}
	}
	// if len(errs)
	return metaInfo, err
}
