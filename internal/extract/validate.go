package extract

import (
	"log/slog"
	"regexp"
	"strconv"
	"strings"

	"github.com/triopium/go_utils/pkg/files"
)

type ValidateColumnAddressParams struct {
	SheetName string
	RowPartCode
	FieldID     string
	FormatRegex string
}

// VALIDATE UNIT FUNCTION
func FormatFieldBeforeValidation(input string) string {
	out := strings.ToLower(input)                // all lower letters
	out = strings.ReplaceAll(out, "\u00a0", " ") // non-breaking space (NBSP) not needed if strings.Join(strings.Fields(out)," ") is used as it removes NBSP also
	// out = strings.Replace(out, " ", "", -1)      // replace space
	out = strings.Replace(out, "\t", "", -1)
	out = strings.Replace(out, "\n", "", -1)
	out = strings.Join(strings.Fields(out), " ")
	return out
}

func FormatFieldValuesListBeforeValidation(input string) []string {
	// split list by various delimiters to slice
	out := strings.ToLower(input)                // all lower letters
	out = strings.ReplaceAll(out, "\u00a0", " ") // non-breaking space (NBSP) not needed if strings.Join(strings.Fields(out)," ") is used as it removes NBSP also
	out = strings.Replace(out, "\t", " ", -1)
	out = strings.Replace(out, "\n", " ", -1)
	out = strings.Replace(out, ";", " ", -1)
	out = strings.Replace(out, ",", " ", -1)
	return strings.Fields(out)
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

func (e *Extractor) ValidateAllColumns(xlsxValidationReceipeFile string) {
	e.ValidateColumnsValues(xlsxValidationReceipeFile)
	e.ValidateColumnsValuesList(xlsxValidationReceipeFile)
	e.ValidateColumnsFormat(xlsxValidationReceipeFile)
}

// VALIDATE COLUMNS VALUES LIST
func ValidateValuesList(
	allovedVals map[string][]string, valuesList string) bool {
	// valueList is string containing values delimited by delimiter
	values := FormatFieldValuesListBeforeValidation(valuesList)
	for _, val := range values {
		_, ok := allovedVals[val+";"]
		// NOTE: in allovedVals the values are "01;" from source xlsx file
		if !ok {
			return false
		}
	}
	return true
}

func (e *Extractor) ValidateColumnsValuesList(xlsxValidationReceipeFile string) {
	columnsParams := []ValidateColumnAddressParams{
		{"druh", RowPartCode_StoryHead, "16", ""},
		{"tema", RowPartCode_StoryHead, "5016", ""},
		{"tema", RowPartCode_ContactItemHead, "5016", ""},
	}
	for _, c := range columnsParams {
		err := e.ValidateColumnValuesList(xlsxValidationReceipeFile, c)
		if err != nil {
			slog.Error("fuck validation error")
		}
	}
}

func (e *Extractor) ValidateColumnValuesList(
	xlsxValidationReceipeFile string,
	vp ValidateColumnAddressParams,
) error {
	sheetRows, err := files.ReadExcelFileSheetRows(
		xlsxValidationReceipeFile, vp.SheetName)
	if err != nil {
		return err
	}
	sheetTableMapped := files.CreateTableTransformRowHeader(
		sheetRows, 0, 0, FormatFieldBeforeValidation)
	valueNC := RowFieldSpecialValueCodeMap[RowFieldValueNotValid]
	allovedVals := sheetTableMapped.RowHeaderToColumnMap

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
		ok = ValidateValuesList(allovedVals, field.Value)
		// Value is not valid
		if !ok {
			part[field.FieldID] = RowField{
				FieldID:   vp.FieldID,
				FieldName: "",
				Value:     valueNC,
			}
		}

		// Value is valid
		if ok {
			continue
		}
	}
	return nil
}

// VALIDATE COLUMNS VALUES
func (e *Extractor) ValidateColumnsValues(
	xlsxValidationReceipeFile string) {
	columnsParams := []ValidateColumnAddressParams{
		{"stanice_RR", RowPartCode_RadioHead, "5081", ""},
		{"format_SR", RowPartCode_SubHead, "321", ""},
		{"format", RowPartCode_StoryHead, "321", ""},
		{"redakce", RowPartCode_StoryHead, "12", ""},
		{"cil_vyroby", RowPartCode_StoryHead, "5079", ""},
		{"pohlavi_KON", RowPartCode_ContactItemHead, "5088", ""},
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
	sheetTableMapped := files.CreateTableTransformRowHeader(
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

// VALIDATE COLUMNS FORMAT
func (e *Extractor) ValidateColumnsFormat(xlsxValidationReceipeFile string) {
	rgx_ID := `^[a-fA-F\d]{8}(-[a-fA-F\d]{4}){3}-[a-fA-F\d]{8}$`
	// UUID variand eg.: F763E219-4759-4698-B2D6-5E6DFDBC
	rgx_Stopaz := `^\d{1,16}$`         // only numbers
	rgx_Datum := `^\d{8}T\d{6},\d{3}$` // 20240102T130010,333
	rgx_Incode := `^OM\d{3,}$`         // OM + 3 or more digits
	rgx_Itemcode := `^\d{4,}`          // 4 or more digits

	columnsParams := []ValidateColumnAddressParams{
		// ID
		{"ID_KON", RowPartCode_ContactItemHead, "5087", rgx_ID},
		{"ID_KON", RowPartCode_ContactItemHead, "5068", rgx_ID},

		// STOPAZ
		{"stopaz", RowPartCode_SubHead, "1005", rgx_Stopaz},
		{"stopaz", RowPartCode_SubHead, "1002", rgx_Stopaz},
		{"stopaz", RowPartCode_StoryHead, "1005", rgx_Stopaz},
		{"stopaz", RowPartCode_StoryHead, "1002", rgx_Stopaz},
		{"stopaz", RowPartCode_AudioClipHead, "38", rgx_Stopaz},

		// DATUM 4x
		{"datum", RowPartCode_RadioHead, "1000", rgx_Datum},
		{"datum", RowPartCode_HourlyHead, "1000", rgx_Datum},
		{"datum", RowPartCode_SubHead, "1004", rgx_Datum},
		{"datum", RowPartCode_StoryHead, "1004", rgx_Datum},

		{"incode", RowPartCode_StoryHead, "5072", rgx_Incode},
		{"itemcode", RowPartCode_AudioClipHead, "5082", rgx_Itemcode},
	}
	for _, c := range columnsParams {
		err := e.ValidateColumnFormat(xlsxValidationReceipeFile, c)
		if err != nil {
			slog.Error("fuck validation error")
		}
	}
}

func (e *Extractor) ValidateColumnFormat(
	xlsxValidationReceipeFile string,
	vp ValidateColumnAddressParams,
) error {
	valueNC := RowFieldSpecialValueCodeMap[RowFieldValueNotValid]
	formatRegex := regexp.MustCompile(vp.FormatRegex)
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
		ok = formatRegex.MatchString(field.Value)
		// Value is not valid
		if !ok {
			part[field.FieldID] = RowField{
				FieldID:   vp.FieldID,
				FieldName: "",
				Value:     valueNC,
			}
		}

		// Value is valid
		if ok {
			continue
		}
	}
	return nil
}
