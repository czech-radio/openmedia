package extract

import (
	"log/slog"
	"strconv"
	"strings"

	"github.com/triopium/go_utils/pkg/files"
)

type Validator struct {
	ValidatorFileName string
}

// ValidateStopaz
func ValidateStopaz(stopaz string) (string, error) {
	milliSeconds, err := strconv.ParseInt(stopaz, 10, 64)
	specVal := RowFieldSpecialValueCodeMap[RowFieldValueNotValid]
	if err != nil {
		return stopaz, err
	}
	if milliSeconds < 0 {
		// negative stopaz is invalid
		return specVal, err
	}
	return stopaz, nil
}

func (e *Extractor) ValidateColumns(xlsxValidationReceipeFile string) {
	// err := e.ValidateColumnValues(
	// xlsxValidationReceipeFile, "cil_vyroby", RowPartCode_StoryHead, "5079")
	// if err != nil {
	// slog.Error(err.Error())
	// }
}

type ValidateColumnAddressParams struct {
	SheetName string
	RowPartCode
	FieldID string
}

func (e *Extractor) ValidateColumnsValues(
	xlsxValidationReceipeFile string) {
	columnsParams := []ValidateColumnAddressParams{
		{"stanice_RR", RowPartCode_RadioHead, "5081"},
		{"format", RowPartCode_SubHead, "321"},
		{"redakce", RowPartCode_StoryHead, "12"},
		{"cil_vyroby", RowPartCode_StoryHead, "5079"},
		{"pohlavi_KON", RowPartCode_ContactItemHead, "5088"},
	}

	for _, c := range columnsParams {
		err := e.ValidateColumnValues(xlsxValidationReceipeFile, c)
		if err != nil {
			slog.Error("fuck validation error")
		}
	}
}

func (e *Extractor) ValidateColumnValues(
	xlsxValidationReceipeFile string,
	vp ValidateColumnAddressParams,
	// sheetName string, rowPartCode RowPartCode, fieldID string,
) error {
	sheetRows, err := files.ReadExcelFileSheetRows(
		xlsxValidationReceipeFile, vp.SheetName)
	if err != nil {
		return err
	}
	// sheetTableMapped := files.CreateTable(
	// sheetRows, 0, 0)
	sheetTableMapped := files.CreateTableTransformColumn(
		sheetRows, 0, 0, FormatFieldBeforeValidation)
	valueNC := RowFieldSpecialValueCodeMap[RowFieldValueNotValid]

	rs := e.TableXML.Rows
	for _, r := range rs {
		part, field, ok := GetRowPartAndField(
			r.RowParts, vp.RowPartCode, vp.FieldID)
		if !ok {
			slog.Error("field not present in row", "filedID", field.FieldID)
			continue
		}
		if CheckIfFieldValueIsSpecialValue(field.Value) {
			continue
		}
		valueToValidate := FormatFieldBeforeValidation(field.Value)
		_, ok = sheetTableMapped.RowHeaderToColumnMap[valueToValidate]

		// Value is not valid
		if !ok {
			part[field.FieldID] = RowField{
				FieldID:   vp.FieldID,
				FieldName: "",
				// Value:     valueNC + " " + valueToValidate + " " + field.Value,
				Value: valueNC,
			}
		}

		// Value is valid
		if ok {
			continue
			// part[field.FieldID] = RowField{
			// FieldID:   vp.FieldID,
			// FieldName: "",
			// Value:     "valid+" + valueToValidate + " " + field.Value,
			// }
		}

		// TODO: account the rownumber, rowPartCode, fieldID, origvalue
		// e.MarkField(r.RowParts, RowPartCode_ContactItemHead, newColumnName, mark)
		// _, ok = sheetTableMapped.RowHeaderToColumnMap[valueTransformed]
		// mark := MarkValue(ok, field.Value, valueNP)
		// mark := MarkValue(ok, valueTransformed, valueNP)
		// e.MarkField(r.RowParts, RowPartCode_ContactItemHead, newColumnName, mark)
	}
	return nil
}

func FormatFieldBeforeValidation(input string) string {
	out := strings.ToLower(input)                // all lower letters
	out = strings.ReplaceAll(out, "\u00a0", " ") // non-breaking space (NBSP) not needed if strings.Join(strings.Fields(out)," ") is used as it removes NBSP also
	// out = strings.Replace(out, " ", "", -1)      // replace space
	out = strings.Replace(out, "\t", "", -1)
	out = strings.Replace(out, "\n", "", -1)
	out = strings.Join(strings.Fields(out), " ")
	return out
}
