package internal

import (
	"encoding/xml"
	"fmt"
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
	string
}

func (c *Om_field_value) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// var token_value string
	// var parts []string

	for {
		token, err := d.Token()
		if err == nil {
			switch t := token.(type) {
			case xml.CharData:
				// fmt.Println(string(t))
				c.string = string(t)
				fmt.Println(c)
			}
		}
		if err != nil {
			return nil
		}
	}
	// switch t := token.(type) {
	// case xml.CharData:
	// parts = append(parts, string(t))
	// case xml.EndElement:
	// if t == start.End() {
	// End of the custom element, break the loop
	// break
	// }
	// default:
	// break
	// }
	// }

	// Combine the parts into a single value
	// c.Value = strings.Join(parts, "")
	// return nil
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
}
