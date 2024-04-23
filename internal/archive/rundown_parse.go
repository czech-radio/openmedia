package internal

import (
	"regexp"
	"strings"
	"time"

	"github.com/ncruces/go-strftime"
)

const (
	DateLayout_RundownDate = "20060102T150405.000"
)

type OMmetaInfo struct {
	Date        time.Time
	Name        string
	HoursRange  string
	RundownType string
}

// OMmetaInfoParse
func (om OPENMEDIA) OMmetaInfoParse() (OMmetaInfo, error) {
	var err error
	var templateName string
	var omMetaInfo *OpenMediaFileType
	var omFileMetaInfo OMmetaInfo
	// Get values of main object attrs
	// TODO: optimize without loops
	// TODO: check empty parsed fields and parsing errors
	for _, attr := range om.OM_OBJECT.Attrs {
		switch attr.Name.Local {
		case "TemplateName":
			templateName = attr.Value
		}
	}
	omMetaInfo, err = GetOMtypeByTemplateName(templateName)
	if err != nil {
		return omFileMetaInfo, err
	}
	switch omMetaInfo.Code {
	case OmFileTypeRundown:
		omFileMetaInfo, err = om.ParseRadioRundnownFields()
	case OmFileTypeContact:
		omFileMetaInfo, err = om.ParseContactFields()
	}
	omFileMetaInfo.RundownType = templateName
	return omFileMetaInfo, err
}

func (om OPENMEDIA) ParseContactFields() (OMmetaInfo, error) {
	var err error
	var metaInfo OMmetaInfo
	var xmlFileNameValue string
	var dateString string
	fields := om.OM_OBJECT.OM_HEADER.Fields
	for _, field := range fields {
		for _, attr := range field.Attrs {
			switch attr.Value {
			case "Název":
				metaInfo.Name = strings.ReplaceAll(field.OM_STRING, " ", "_")
			}
		}
	}

	for _, attr := range om.OM_OBJECT.OM_UPLINK.Attrs {
		switch attr.Name.Local {
		case "FileName":
			// NOTE: Some contact files may have this field same even though they
			// filename in filesystem is different. It can result in dupes in zip
			// package. When there are dupes in package it does not mean it is
			// corupted or that the data cannot be extracted. Only fuse mount of
			// package cannot be used as it returns error there are duplicate files
			// in archive.
			xmlFileNameValue = attr.Value
		}
	}
	rgxPatt := `(\d*).xml$`
	regexpObject := regexp.MustCompile(rgxPatt)
	matches := regexpObject.FindStringSubmatch(xmlFileNameValue)
	if len(matches) == 2 {
		dateString = matches[1]
	}
	metaInfo.Date, err = strftime.Parse("%Y%m%d%H%M%S", dateString)
	return metaInfo, err
}

func (om OPENMEDIA) ParseRadioRundnownFields() (OMmetaInfo, error) {
	// Get values of main header fields
	var err error
	var metaInfo OMmetaInfo
	fields := om.OM_OBJECT.OM_HEADER.Fields
	for _, field := range fields {
		for _, attr := range field.Attrs {
			switch attr.Value {
			case "Čas začátku": // FieldID: 1004
				// case "Čas vytvoření": // FiledID: 1
				metaInfo.Date, err = strftime.Parse("%Y%m%dT%H%M%S", field.OM_DATETIME)
			case "Název":
				metaInfo.ParseRadioRundownName(field.OM_STRING)
			}
		}
	}
	return metaInfo, err
}

var RadioRundownNameRegex = regexp.MustCompile(`(\d\d-\d\d) ([\p{L}\s?]*)`)

func (r *OMmetaInfo) ParseRadioRundownName(rundownName string) {
	// e.g.: "05-09 ČRo Karlovy Vary - Wed, 04.03.2020"
	var radioName string
	var hoursRange string
	matches := RadioRundownNameRegex.FindStringSubmatch(rundownName)
	switch len(matches) {
	case 3:
		hoursRange = strings.TrimSpace(matches[1])
		radioName = strings.TrimSpace(matches[2])
		radioName = strings.ReplaceAll(radioName, " ", "_")
	}
	r.Name = radioName
	r.HoursRange = hoursRange
}
