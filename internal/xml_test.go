package internal

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"path/filepath"
	"testing"
)

func Test_FileUTF16leToUTF8(t *testing.T) {
	src_file := filepath.Join(TEMP_DIR_TEST_SRC, "rundowns_valid", "RD_00-12_Pohoda_-_Fri_06_01_2023_orig.xml")
	dst_file := filepath.Join(TEMP_DIR_TEST_DST, "convert.xml")
	err := FileUTF16leToUTF8(src_file, dst_file)
	if err != nil {
		t.Error(err)
	}
	Sleeper(1000, "s")
}

func Test_XmlUnmarshal(t *testing.T) {
	valid_files := []string{
		"RD_00-12_Pohoda_-_Fri_06_01_2023_utf8_formated.xml",
	}
	for _, vf := range valid_files {
		file := filepath.Join(TEMP_DIR_TEST_SRC, "rundowns_valid", vf)
		om, err := RundownUnmarshal(file)
		if err != nil {
			t.Error(err)
		}
		js, err := json.MarshalIndent(om, "", "  ")
		if err != nil {
			t.Error(err.Error())
		}
		fmt.Println(string(js))
	}
}

func BenchmarkXmlUnmarshal(b *testing.B) {
	valid_files := []string{
		"RD_00-12_Pohoda_-_Fri_06_01_2023_utf8_formated.xml",
		// "RD_00-12_Pohoda_-_Fri_06_01_2023_utf8.xml",
	}
	file := filepath.Join(TEMP_DIR_TEST_SRC, "rundowns_valid", valid_files[0])
	// var res *OPENMEDIA
	for i := 0; i < b.N; i++ {
		omu, err := RundownUnmarshal(file)
		// _, err := RundownUnmarshal(file)
		if err != nil {
			b.Error(err)
		}
		_, err = xml.MarshalIndent(omu, "", "  ")
		if err != nil {
			b.Error(err)
		}
	}
}
