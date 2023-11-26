package internal

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/fs"
	"log/slog"
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

func Test_FileIsValidXmlToMinify(t *testing.T) {
	var mustBeValid bool = true // wheter the results of subsequent tests should result true or false. I.e. if mustBeValid == true all files should result as valid.

	//Test valid files
	validFilesDir := filepath.Join(TEMP_DIR_TEST_SRC, "rundowns_valid")
	walkFunc := func(directory string, de fs.DirEntry) error {
		filePath := filepath.Join(directory, de.Name())
		slog.Debug("validating", "file", filePath)
		validity, err := FileIsValidXmlToMinify(filePath)
		if mustBeValid != validity {
			t.Errorf("file: %s, err: %s", filePath, err)
		}
		return nil
	}
	err := DirectoryTraverse(validFilesDir, walkFunc, false)
	if err != nil {
		t.Error(err)
	}

	//Test invalid files
	invalidFile := filepath.Join(TEMP_DIR_TEST_SRC, "rundowns_invalid", "bad_tags.xml")
	validity, err := FileIsValidXmlToMinify(invalidFile)
	if validity == true {
		t.Errorf("file should be invalid: %s, invalidity cause: %s", invalidFile, err)
	}
}

func Test_XmlFileLinesValidate(t *testing.T) {
	var mustBeValid bool = true // wheter the results of subsequent tests should result true or false. I.e. if mustBeValid == true all files should result as valid.

	//Test valid files
	validFilesDir := filepath.Join(TEMP_DIR_TEST_SRC, "rundowns_valid")
	walkFunc := func(directory string, de fs.DirEntry) error {
		filePath := filepath.Join(directory, de.Name())
		slog.Debug("validating", "file", filePath)
		validity, err := XmlFileLinesValidate(filePath)
		if mustBeValid != validity {
			t.Errorf("file: %s, err: %s", filePath, err)
		}
		return nil
	}
	err := DirectoryTraverse(validFilesDir, walkFunc, false)
	if err != nil {
		t.Error(err)
	}

	//Test invalid files
	invalidFile := filepath.Join(TEMP_DIR_TEST_SRC, "rundowns_invalid", "bad_tags.xml")
	validity, err := XmlFileLinesValidate(invalidFile)
	if validity == true {
		t.Errorf("file should be invalid: %s, invalidity cause: %s", invalidFile, err)
	}
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
		// fmt.Printf("%+v", om)
		// fmt.Printf("%+v\n", om.OM_OBJECTS[0].OM_HEADER)
		// for _, i := range om.OM_OBJECTS[0].OM_HEADER.Fields {
		// fmt.Printf("%+v\n", i)
		// }
		// for _, i := range om.OM_OBJECTS[0].OM_RECORD.Fields {
		// for n, i := range om.OM_OBJECTS[0].OM_RECORDS {
		// fmt.Printf("RECORD NUMBER: %d\n", n)
		// fmt.Printf("%+v\n", i)
		// }
		// for _, i := range om.OM_OBJECTS {
		// fmt.Printf("%v\n", (i))
		// }
		// json.MarshallIndent(om)
		js, err := json.MarshalIndent(om, "", "  ")
		fmt.Println(string(js))
	}
}

// var result *OPENMEDIA

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
	// result = res

}
