package cmd

import (
	c "github.com/triopium/go_utils/pkg/configure"
	"github.com/triopium/go_utils/pkg/helper"
)

func OptionsCommonExtractFilter() {
	add := SubcommandConfig.AddOption
	add("ExtractorsCode", "excode", "production_all", "string", c.NotNil,
		"Name of extractor which specifies the parts of xml to be extracted", nil, nil)
	add("FilterRadioNames", "frns", "", "[]string", "",
		"Filter data corresponding to radio names", nil, nil)

	add("FilterDateFrom", "fdf",
		helper.ISOweekStartLocal(-1).String(), "date", c.NotNil,
		"Filter rundowns from date. Format of the date is given in form 'YYYY-mm-ddTHH:mm:ss' e.g. 2024, 2024-02-01 or 2024-02-01T10. The precission of date given is arbitrary.", nil, nil)
	add("FilterDateTo", "fdt",
		helper.ISOweekStartLocal(0).String(), "date", c.NotNil,
		"Filter rundowns to date", nil, nil)

	add("FilterIsoWeeks", "fisow", "", "[]int", "",
		"Filter data corresponding to specified ISO weeks", nil, nil)
	add("FilterMonths", "fmonths", "", "[]int", "",
		"Filter data corresponding to specified months", nil, nil)
	add("FilterWeekDays", "fwdays", "", "[]int", "",
		"Filter data corresponding to specified weekdays", nil, nil)

	// Special columns
	add("AddRecordsNumbers", "arn", "false", "bool", "",
		"Add record numbers columns and dependent columns", "", nil)

	// Validation
	add("ValidatorFileName", "valfn", "", "string", "",
		"xlsx file containing validation receipe", nil, CheckFileExistsIfNotNull)

	// Special filters
	add("FilterFileName", "frfn", "", "string", "",
		"Special filters filename. The filter filename specifies how the file is parsed and how it is used", nil, CheckFileExistsIfNotNull)
	add("FilterSheetName", "frsn", "data", "string", "",
		"Special filters sheetname", nil, nil)
}

func OptionsCommonExtractPath() {
	add := SubcommandConfig.AddOption
	add("SourceFilePath", "sfp",
		"", "string", c.NotNil,
		"Source rundown file.", nil, helper.FileExists)
	add("OutputDirectory", "odir",
		"", "string", c.NotNil,
		"Output file path for extracted data.", nil,
		helper.DirectoryExists)
	add("OutputFileName", "ofn",
		"", "string", c.NotNil,
		"Output file path for extracted data.", nil,
		nil)
	add("CSVdelim", "csvD", "\t", "string", "",
		"csv column field delimiter", []string{"\t", ";"}, nil)
}
