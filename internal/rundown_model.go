package internal

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type OPENMEDIA struct {
	XMLName    xml.Name    `xml:"OPENMEDIA"`
	OM_OBJECTS []OM_OBJECT `xml:"OM_OBJECT"`
}

type OM_OBJECT struct {
	OM_HEADER OM_HEADER `xml:"OM_HEADER"`
}

type OM_HEADER struct {
	// Fields []OM_FIELD `xml:"OM_FIELD"`
	Fields []Om_field_value `xml:"OM_FIELD"`
}

type OM_RECORD struct {
	// Fields []OM_FIELD `xml:"OM_FIELD"`
	Fields []Om_field_value `xml:"OM_FIELD"`
}

// type OM_FIELD struct {
// Type      string         `xml:"FieldID,attr"`
// FieldType string         `xml:"FieldType,attr"`
// FieldName string         `xml:"FieldName,attr"`
// IsEmpty   string         `xml:"IsEmpty,attr"`
// Value     Om_field_value `xml:""`
// Value     string `xml:"OM_DATETIME,omitempty"`
// Value string `xml:"OM_DATETIME,OM_STRING,omitempty"`
// Date      OM_DATETIME `xml:"OM_DATETIME"`
// }

// OM_FIELD contais various nested tag name as field value. Must be parsed differently
type Om_field_value struct {
	Value string
}

func (omf *Om_field_value) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var tagValue strings.Builder
	var start_count int = 0
	var end_count int = 0
	var errUnexpectedTagStructure = fmt.Errorf("unexpected xml tag structure")
	for {
		token, err := d.Token()
		if err != nil {
			break
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

// type OM_UPLINK struct {
// }
