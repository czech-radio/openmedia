package internal

import (
	"encoding/json"
	"encoding/xml"
	"os"

	"fmt"
	"path/filepath"
	"testing"
)

func Test_XmlValidateWithXSD(t *testing.T) {
	valid_files := []string{
		"RD_00-12_Pohoda_-_Fri_06_01_2023_utf8_formated.xml",
	}
	for _, vf := range valid_files {
		file := filepath.Join(TEMP_DIR_TEST_SRC, "rundowns_valid", vf)
		_, err := XmlValidateWithXSD(file)
		if err != nil {
			t.Error(err)
		}
	}
}

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

func TestXmlMarshalOM_FIELDomitEmpty(t *testing.T) {
	test_fields := `
<OM_HEADER>
<OM_FIELD FieldID="1" FieldType="3" FieldName="Čas vytvoření" IsEmpty="no">
	<OM_DATETIME>20221223T010414,000</OM_DATETIME>
</OM_FIELD>
<OM_FIELD FieldID="2" FieldType="3" FieldName="Aktualizováno kdy" IsEmpty="no">
	<OM_DATETIME>20221223T010414,000</OM_DATETIME>
</OM_FIELD>
<OM_FIELD FieldID="3" FieldType="1" FieldName="Owner Name" IsEmpty="no">
	<OM_STRING>admin</OM_STRING>
</OM_FIELD>
<OM_FIELD FieldID="62" FieldType="2" FieldName="OnAirStatus" IsEmpty="yes">
	<OM_INT32/>
</OM_FIELD>
<OM_FIELD FieldID="128" FieldType="3" FieldName="AiredTime" IsEmpty="yes">
	<OM_DATETIME/>
</OM_FIELD>
</OM_HEADER>`
	header := OM_HEADER{}
	err := xml.Unmarshal([]byte(test_fields), &header)
	if err != nil {
		t.Error(err)
	}
	res, err := xml.MarshalIndent(header, "", "\t")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(res))
}

func Test_XmlMarshalSaveToFile(t *testing.T) {
	valid_files := []string{
		"RD_00-12_Pohoda_-_Fri_06_01_2023_utf8_formated.xml",
	}
	for _, vf := range valid_files {
		file := filepath.Join(TEMP_DIR_TEST_SRC, "rundowns_valid", vf)
		om, err := RundownUnmarshal(file)
		if err != nil {
			t.Error(err)
		}
		res, err := xml.MarshalIndent(om, "", "  ")
		if err != nil {
			t.Error(err)
		}
		outputFile := filepath.Join(TEMP_DIR_TEST_DST, vf)
		err = os.WriteFile(outputFile, res, 0700)
		if err != nil {
			t.Error(err)
		}
	}
	Sleeper(1000, "s")
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
