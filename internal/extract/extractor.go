package extract

import (
	ar "github/czech-radio/openmedia/internal/archive"
	"log/slog"
	"path/filepath"

	"github.com/antchfx/xmlquery"
)

type ObjectAttributes = map[string]string
type UniqueValues = map[string]int // value vs count

type CSVrowFields = []CSVrowField

type OMextractor struct {
	OrExt []OMextractor

	ObjectPath       string
	ObjectAttrsNames []string
	FieldsPath       string
	FieldIDs         []string
	PartPrefixCode   PartPrefixCode
	FieldsPrefix     string

	// Internals
	KeepInputRow            bool
	PreserveParentNode      bool // Add fields to input row
	PreserveParentNodeLevel int
	KeepWhenZeroSubnodes    bool
	FieldIDsMap             map[string]bool
}

type OMextractors []OMextractor

type Extractor struct {
	OMextractors
	BaseNode *xmlquery.Node
	CSVdelim string

	CSVrowPartsFieldsPositions
	CSVrowPartsPositionsInternal
	CSVrowPartsPositionsExternal

	CSVheaders          map[string]CSVrow
	CSVheadersPositions []CSVheaderCodeName
	CSVheaderInternal   string
	CSVheaderExternal   string

	CSVtable
	CSVtables
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
	e.CreateTablesHeader(CSVdelim)
	e.OMextractors.KeepInputRowsChecker()
	e.OMextractors.MapFieldsPath()
	e.CSVtable.Rows = []*CSVrowNode{{baseNode, CSVrow{}}}
}

func (e *Extractor) MapRowParts() {
	var prefixesInternal, prefixesExternal []PartPrefixCode
	for _, extr := range e.OMextractors {
		// prefix := PartsPrefixMapProduction[extr.PartPrefixCode]
		prefixesInternal = append(
			prefixesInternal, extr.PartPrefixCode)
		prefixesExternal = append(
			prefixesExternal, extr.PartPrefixCode)
		if len(extr.OrExt) > 0 {
			for _, subExt := range extr.OrExt {
				prefixesInternal = append(
					prefixesInternal, subExt.PartPrefixCode)
				prefixesExternal = append(
					prefixesExternal, subExt.PartPrefixCode)
			}
		}
	}
	e.CSVrowPartsPositionsExternal = prefixesExternal
	e.CSVrowPartsPositionsInternal = prefixesInternal
}

func (e *Extractor) MapRowPartsFieldsPositions() {
	extCount := len(e.OMextractors)
	partsPos := make(CSVrowPartsFieldsPositions, extCount)
	for _, extr := range e.OMextractors {
		fp := GetPartFieldsPositions(extr)
		partsPos[extr.PartPrefixCode] = append(partsPos[extr.PartPrefixCode], fp...)
		if len(extr.OrExt) > 0 {
			for _, subExt := range extr.OrExt {
				fp := GetPartFieldsPositions(subExt)
				partsPos[subExt.PartPrefixCode] = append(
					partsPos[subExt.PartPrefixCode], fp...)
			}
		}
	}
	e.CSVrowPartsFieldsPositions = partsPos
}

func (e *Extractor) ExtractTable(fileName string) error {
	for i, extr := range e.OMextractors {
		if extr.ObjectPath == "" {
			slog.Debug("extractor not extracted", "cause", "empty object")
			continue
		}
		rows, err := ExpandTableRows(e.CSVtable, extr)
		e.CSVtable = rows
		e.CSVtable.SrcFilePath = fileName
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

func GetLastPartOfObjectPath(path string) string {
	return filepath.Base(path)
}
