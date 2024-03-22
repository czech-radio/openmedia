package internal

import (
	"log/slog"
	"path/filepath"

	"github.com/antchfx/xmlquery"
)

type ObjectAttributes = map[string]string
type UniqueValues = map[string]int // value vs count

type CSVrowFields = []CSVrowField

type OMextractor struct {
	ObjectPath       string
	ObjectAttrsNames []string
	FieldsPath       string
	FieldIDs         []string
	PartPrefixCode   PartPrefixCode
	FieldsPrefix     string

	// Internals
	KeepInputRows bool
	FieldIDsMap   map[string]bool
}

type OMextractors []OMextractor

type Extractor struct {
	OMextractors
	CSVrowPartsFieldsPositions
	CSVrowPartsPositionsInternal
	CSVrowPartsPositionsExternal
	CSVheaderInternal string
	CSVheaderExternal string
	CSVdelim          string
	BaseNode          *xmlquery.Node
	CSVtable
	CSVrowsFiltered []int
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
	e.CSVheaderCreate(CSVdelim)
	e.OMextractors.KeepInputRowsChecker()
	e.CSVtable = []*CSVrowNode{{baseNode, CSVrow{}}}
}

func (e *Extractor) MapRowParts() {
	var prefixesInternal, prefixesExternal []string
	for _, extr := range e.OMextractors {
		prefix := PartsPrefixMapProduction[extr.PartPrefixCode]
		prefixesInternal = append(prefixesInternal, prefix.Internal)
		prefixesExternal = append(prefixesExternal, prefix.External)
	}
	e.CSVrowPartsPositionsExternal = prefixesExternal
	e.CSVrowPartsPositionsInternal = prefixesInternal
}

func (e *Extractor) MapRowPartsFieldsPositions() {
	extCount := len(e.OMextractors)
	partsPos := make(CSVrowPartsFieldsPositions, extCount)
	for _, extr := range e.OMextractors {
		prefix := PartsPrefixMapProduction[extr.PartPrefixCode]
		fp := GetPartFieldsPositions(extr)
		partsPos[prefix.Internal] = fp
	}
	e.CSVrowPartsFieldsPositions = partsPos
}

func (e *Extractor) ExtractTable() error {
	for i, extr := range e.OMextractors {
		rows, err := ExpandTableRows(e.CSVtable, extr) // : maybe wrong
		e.CSVtable = rows
		if err != nil {
			return err
		}
		slog.Debug("extractor", "position", i, "objectPath", extr)
	}
	return nil
}

func GetPartFieldsPositions(extr OMextractor) CSVrowPartFieldsPositions {
	fieldsPositions := make(CSVrowPartFieldsPositions, 0, len(extr.FieldIDs))
	prefix := PartsPrefixMapProduction[extr.PartPrefixCode].Internal
	// Object Attributes
	for _, attr := range extr.ObjectAttrsNames {
		fp := FieldPosition{
			FieldPrefix: prefix,
			FieldID:     attr,
			FieldName:   "",
		}
		fieldsPositions = append(fieldsPositions, fp)
	}
	// Object FieldsID
	for _, fi := range extr.FieldIDs {
		fp := FieldPosition{
			FieldPrefix: prefix,
			FieldID:     fi,
			FieldName:   "",
		}
		fieldsPositions = append(fieldsPositions, fp)
	}
	return fieldsPositions
}

func (extr *OMextractor) MapFields() {
	extr.FieldIDsMap = make(map[string]bool, len(extr.FieldIDs))
	for _, id := range extr.FieldIDs {
		extr.FieldIDsMap[id] = true
	}
}

func (extrs OMextractors) KeepInputRowsChecker() {
	// Check if there is following extractor referencing same object as current extractor
	eCount := len(extrs)
	for eCurrent := 0; eCurrent < eCount; eCurrent++ {
		extr := extrs[eCurrent]
		if extr.KeepInputRows {
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

func GetLastPartOfObjectPath(path string) string {
	return filepath.Base(path)
}
