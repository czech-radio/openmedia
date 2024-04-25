package extract

import (
	"fmt"
	ar "github/czech-radio/openmedia/internal/archive"
	"log/slog"
	"path/filepath"

	// github
	"github.com/antchfx/xmlquery"
)

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

	CSVdelim          string
	CSVheaderInternal string
	CSVheaderExternal string
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
	e.OMextractors.KeepInputRowsChecker()
	e.OMextractors.MapFieldsPath()
	e.TableXML.Rows = []*RowNode{{baseNode, RowParts{}}}
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
	for _, partPrefixCode := range e.RowPartsPositions {
		rowPart := e.RowPartsFieldsPositions[partPrefixCode]
		for _, field := range rowPart {
			internal := HeaderColumnInternalCreate(
				partPrefixCode, field.FieldID, "")
			e.HeaderInternal = append(e.HeaderInternal, internal)
			external := HeaderColumnExternalCreate(
				partPrefixCode, field.FieldID, "",
			)
			e.HeaderExternal = append(e.HeaderExternal, external)
		}
	}
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

// KeepInputRowsChecker
func (extrs OMextractors) KeepInputRowsChecker() {
	// Check if there is following extractor referencing same object as current extractor
	eCount := len(extrs)
	for eCurrent := 0; eCurrent < eCount; eCurrent++ {
		extr := extrs[eCurrent]
		if extr.KeepInputRow {
			continue
		}
		if eCurrent == eCount {
			// NOTE
			// maybe not needed, also without it allow the extr position to be independent insted to process sequentially
			// extr.KeepInputRows = false
			break
		}
		eNext := eCurrent + 1
		for next := eNext; next < eCount; next++ {
			//TODO: Without it depends on manual input alone
			// fmt.Println("fek", eCurrent, next)
		}
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
			}
			if !ok && len(extr.FieldIDs) > 0 {
				panic("fields path not given from which to extract")
			}
		}
	}
}

// GetLastPartOfObjectPath
func GetLastPartOfObjectPath(path string) string {
	return filepath.Base(path)
}

// HeaderColumnInternalCreate
func HeaderColumnInternalCreate(
	rowPartCode RowPartCode, fieldID string, delim string) string {
	prefix := RowPartsCodeMapProduction[rowPartCode]
	return fmt.Sprintf("%s_%s%s", prefix.Internal, fieldID, delim)
}

// HeaderColumnExternalCreate
func HeaderColumnExternalCreate(
	rowPartCode RowPartCode, fieldID, delim string) string {
	prefix := RowPartsCodeMapProduction[rowPartCode]
	if prefix.External == "" {
		return fieldID
	}
	fieldName := FieldsIDsNamesProduction[fieldID]
	return fmt.Sprintf("%s_%s%s", fieldName, prefix.External, delim)
}
