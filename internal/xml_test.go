package internal

import (
	"fmt"
	"path/filepath"
	"testing"
)

func Test_FileIsValidXmlToMinify(t *testing.T) {
	filename := "RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431.xml"
	valid, err := FileIsValidXmlToMinify(filepath.Join(TEMP_DIR_TEST_SRC, filename))
	if err != nil {
		t.Error(err)
	}
	if !valid {
		t.Error("file not valid: ", filename)
	}
}

func Test_XmlFileLinesValidate(t *testing.T) {
	// Sleeper(10, "s")
	filename_valid := "RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431.xml"
	filename_invalid := "RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml"

	type testTuple struct {
		input  string
		result bool
	}

	tests := []testTuple{
		{filename_valid, true},
		{filename_invalid, false},
	}

	for tcase := range tests {
		valid, err := XmlFileLinesValidate(filepath.Join(TEMP_DIR_TEST_SRC, tests[tcase].input))
		if valid != tests[tcase].result {
			if err != nil {
				fmt.Println("error: ", err.Error())
			}
			t.Fatalf("expectd: %v, got %v", tests[tcase], valid)
		}
	}
}
