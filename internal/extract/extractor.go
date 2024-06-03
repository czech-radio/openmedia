package extract

import (
	"fmt"
	ar "github/czech-radio/openmedia/internal/archive"
	"log/slog"
	"path/filepath"

	"github.com/antchfx/xmlquery"
)

type ExtractorsMap map[ExtractorsPresetCode]OMextractors

type ExtractorsPresetCode string

const (
	ExtractorsProductionOmit     ExtractorsPresetCode = "production_omit"
	ExtractorsProductionAll      ExtractorsPresetCode = "production_all"
	ExtractorsProductionContacts ExtractorsPresetCode = "production_contacts"
)

var ExtractorsCodeMap = ExtractorsMap{
	ExtractorsProductionOmit:     EXTproductionOmit,
	ExtractorsProductionAll:      EXTproductionAll,
	ExtractorsProductionContacts: EXTproductionContacts,
}

// ObjectAttributes
type ObjectAttributes = map[string]string

// UniqueValues
type UniqueValues = map[string]int // value vs count
// CSVrowFields
type CSVrowFields = []RowField

// OMextractor
type OMextractor struct {
	OrExt []OMextractor

	ObjectPath       string
	ObjectAttrsNames []string
	FieldsPath       string
	FieldIDs         []string
	PartPrefixCode   RowPartCode
	FieldsPrefix     string

	// Internals
	KeepInputRow         bool
	ResultNodeGoUpLevels int
	KeepWhenZeroSubnodes bool
	FieldIDsMap          map[string]bool
}

// OMextractors
type OMextractors []OMextractor

// Extractor holds the data needed to extract specified XML tree to table/array and export the table as CSV, XLSX or print it to standard output
type Extractor struct {
	OMextractors
	BaseNode *xmlquery.Node

	RowPartsPositions
	RowPartsFieldsPositions

	TableXML
	HeaderInternal []string
	HeaderExternal []string

	CSVdelim string

	// Logging
	ValidationTableAccount
}

func (e *Extractor) Init(
	baseNode *xmlquery.Node,
	omextractors OMextractors,
	CSVdelim string) {
	e.OMextractors = omextractors
	e.CSVdelim = CSVdelim
	e.BaseNode = baseNode
	e.MapRowParts()
	e.MapRowPartsFieldsPositions()
	e.HeaderBuild()
	e.OMextractors.MapFieldsPath()
	if baseNode != nil {
		e.TableXML.Rows = []*RowNode{{baseNode, RowParts{}}}
	}
}

func (e *Extractor) MapRowParts() {
	var prefixesInternal []RowPartCode
	for _, extr := range e.OMextractors {
		prefixesInternal = append(
			prefixesInternal, extr.PartPrefixCode)
	}
	e.RowPartsPositions = prefixesInternal
}

func (e *Extractor) MapRowPartsFieldsPositions() {
	extCount := len(e.OMextractors)
	partsPos := make(RowPartsFieldsPositions, extCount)
	for _, extr := range e.OMextractors {
		fp := GetPartFieldsPositions(extr)
		partsPos[extr.PartPrefixCode] = append(partsPos[extr.PartPrefixCode], fp...)
	}
	e.RowPartsFieldsPositions = partsPos
}

func (e *Extractor) HeaderBuild() {
	var newInternal, newExternal []string
	for _, partPrefixCode := range e.RowPartsPositions {
		rowPart := e.RowPartsFieldsPositions[partPrefixCode]
		for i := range rowPart {
			internal := e.HeaderColumnInternalCreate(
				partPrefixCode, rowPart[i].FieldID, "")
			newInternal = append(newInternal, internal)

			external := e.HeaderColumnExternalCreate(
				partPrefixCode, rowPart[i].FieldID, "")
			newExternal = append(newExternal, external)
		}
	}
	e.HeaderInternal = newInternal
	e.HeaderExternal = newExternal
}

// HeaderColumnInternalCreate
func (e *Extractor) HeaderColumnInternalCreate(
	rowPartCode RowPartCode, fieldID string, delim string) string {
	prefix := RowPartsCodeMapProduction[rowPartCode]
	return fmt.Sprintf("%s_%s%s", prefix.Internal, fieldID, delim)
}

// HeaderColumnExternalCreate
func (e *Extractor) HeaderColumnExternalCreate(
	rowPartCode RowPartCode, fieldID, delim string) string {
	prefix := RowPartsCodeMapProduction[rowPartCode]
	part := e.RowPartsFieldsPositions[rowPartCode]
	var resName string
	for _, field := range part {
		if field.FieldID != fieldID {
			continue
		}
		if field.FieldName != "" {
			resName = field.FieldName
			break
		}
	}
	if resName == "" {
		resName = FieldsIDsNamesProduction[fieldID]
	}
	if prefix.External == "" {
		// NOTE!!!: maybe add delim too
		return resName
	}
	return fmt.Sprintf("%s_%s%s", resName, prefix.External, delim)
}

func GetColumnHeaderExternal(rowPartCode RowPartCode, fieldID string) string {
	resName := FieldsIDsNamesProduction[fieldID]
	prefix := RowPartsCodeMapProduction[rowPartCode]
	if prefix.External == "" {
		return resName
	}
	return fmt.Sprintf("%s_%s", resName, prefix.External)
}

func (e *Extractor) ExtractTable(fileName string) error {
	for i, extr := range e.OMextractors {
		if extr.ObjectPath == "" {
			slog.Debug("extractor not extracted", "cause", "empty object")
			continue
		}
		rows, err := ExpandTableRows(e.TableXML, extr)
		e.TableXML = rows
		e.TableXML.SrcFilePath = fileName
		if err != nil {
			return err
		}
		slog.Debug("extractor", "position", i, "objectPath", extr)
	}
	return nil
}

// GetPartFieldsPositions
func GetPartFieldsPositions(extr OMextractor) RowPartFieldsPositions {
	fieldsPositions := make(RowPartFieldsPositions, 0, len(extr.FieldIDs))
	prefix := RowPartsCodeMapProduction[extr.PartPrefixCode].Internal
	// Object Attributes
	for _, attr := range extr.ObjectAttrsNames {
		fp := RowPartFieldPosition{
			RowPartName: prefix,
			FieldID:     attr,
			FieldName:   "",
		}
		fieldsPositions = append(fieldsPositions, fp)
	}
	// Object FieldsID
	for _, fi := range extr.FieldIDs {
		fp := RowPartFieldPosition{
			RowPartName: prefix,
			FieldID:     fi,
			FieldName:   "",
		}
		fieldsPositions = append(fieldsPositions, fp)
	}
	return fieldsPositions
}

// MapFields
func (extr *OMextractor) MapFields() {
	extr.FieldIDsMap = make(map[string]bool, len(extr.FieldIDs))
	for _, id := range extr.FieldIDs {
		extr.FieldIDsMap[id] = true
	}
}

// MapFieldsPath
func (extrs OMextractors) MapFieldsPath() {
	for i, extr := range extrs {
		if extr.ObjectPath == "" {
			continue
		}
		objectName := ar.GetObjectNameFromPath(extr.ObjectPath)
		if extr.FieldsPath == "" {
			tag, ok := ar.OmTagStructureMap[objectName]
			if ok {
				extrs[i].FieldsPath = tag.FieldsPath
				continue
			}
			if objectName == "." {
				extrs[i].FieldsPath = tag.FieldsPath
				continue
			}
			if !ok && len(extr.FieldIDs) > 0 {
				panic(fmt.Errorf("fields path not given from which to extract, path: %q", objectName))
			}
		}
	}
}

// GetLastPartOfObjectPath
func GetLastPartOfObjectPath(path string) string {
	return filepath.Base(path)
}
