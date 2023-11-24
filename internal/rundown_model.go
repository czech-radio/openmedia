package internal

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type OPENMEDIA struct {
	XMLName    xml.Name    `xml:"OPENMEDIA"`
	OM_SERVER  OM_SERVER   `xml:"OM_SERVER"`
	OM_OBJECTS []OM_OBJECT `xml:"OM_OBJECT"`
}

type OM_OBJECT struct {
	Attrs      []xml.Attr  `xml:",any,attr"`
	OM_HEADER  OM_HEADER   `xml:"OM_HEADER"`
	OM_UPLINK  OM_UPLINK   `xml:"OM_UPLINK"`
	OM_RECORDS []OM_RECORD `xml:"OM_RECORD"`
}

type OM_SERVER struct {
	Attrs []xml.Attr `xml:",any,attr"`
}

type OM_RECORD struct {
	Attrs      []xml.Attr  `xml:",any,attr"`
	Fields     []OM_FIELD  `xml:"OM_FIELD"`
	OM_RECORDS []OM_RECORD `xml:"OM_RECORD"`
	OM_OBJECTS OM_OBJECT   `xml:"OM_OBJECT"`
}

type OM_HEADER struct {
	Attrs []xml.Attr `xml:",any,attr"`
	// Type      string     `xml:"FieldID,attr"`
	// FieldType string     `xml:"FieldType,attr"`
	// FieldName string     `xml:"FieldName,attr"`
	// IsEmpty   string     `xml:"IsEmpty,attr"`
	Fields []OM_FIELD `xml:"OM_FIELD"`
}

// type OM_FIELD struct {
// Value     Om_field_value `xml:"any"`
// FieldID   string         `xml:"FieldID,attr"`
// FieldType string         `xml:"FieldType,attr"`
// FieldName string         `xml:"FieldName,attr"`
// IsEmpty   string         `xml:"IsEmpty,attr"`
// }

// OM_FIELD contais various nested tag names.
// Custom unmarshalXML method must be used. It is faster to use map for attributes then usign struct fields then. (Reflect must be used, when iterating over struct fields)
type OM_FIELD struct {
	Attrs []xml.Attr `xml:",any,attr"`
	Value string     `xml:",omitempty"`
	// Attrs map[string]string `xml:"-"`
}

// Much faster to use map then iterate over struct fields.
var OM_FIELD_ATTRS_NAMES = map[string]string{
	"FieldID":   "",
	"FieldType": "",
	"FieldName": "",
	"IsEmpty":   "",
}

func (omf *OM_FIELD) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var tagValue strings.Builder
	var start_count int = 0
	var end_count int = 0
	var errUnexpectedTagStructure = fmt.Errorf("unexpected xml tag structure")
	// omf.Attrs = XmlTagAttributesMap(start, OM_FIELD_ATTRS_NAMES)
	// start.
	omf.Attrs = start.Attr
loop1:
	for {
		token, err := d.Token()
		if err != nil {
			break loop1
		}
		if err == nil {
			switch t := token.(type) {
			case xml.CharData:
				content := strings.TrimSpace(string(t))
				if content != "" {
					if tagValue.Len() == 0 {
						tagValue.WriteString(content)
					} else {
						tagValue.WriteString("\n" + content)
					}
				}
				// Following lines validates xml so that it does not contain unexpected strucute of OM_FIELD, xsd will be used for that
			case xml.StartElement:
				if start_count > 1 {
					return errUnexpectedTagStructure
				}
				start_count++
				continue
			case xml.EndElement:
				if end_count > 1 {
					return errUnexpectedTagStructure
				}
				end_count++
				continue
			default:
				return fmt.Errorf("unknown token type: %T", t)
			}
		}
	}
	omf.Value = tagValue.String()
	return nil
}

// type OM_DATETIME struct {
// time.Time
// Date time.Time `xml:"OM_DATETIME"`
// }

// func (omd *OM_DATETIME) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
// 	const layout = "20060102T150405,"
// 	var timeString string
// 	err := d.DecodeElement(&timeString, &start)
// 	if err != nil {
// 		return err
// 	}
// 	parsed_date, err := time.Parse(layout, timeString)
// 	if err != nil {
// 		return err
// 	}
// 	*omd = OM_DATETIME{parsed_date}
// 	return nil
// }

type OM_UPLINK struct {
	Attrs []xml.Attr `xml:",any,attr"`
}
