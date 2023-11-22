package internal

import (
	"fmt"
	"io/fs"
	"log/slog"
	"path/filepath"
	"testing"
)

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
		om, err := RundownUnmarshall(file)
		if err != nil {
			t.Error(err)
		}
		// fmt.Printf("%+v\n", om.OM_OBJECTS[0].OM_HEADER)
		for _, i := range om.OM_OBJECTS[0].OM_HEADER.Fields {
			fmt.Printf("%+v\n", i)
		}
		// for _, i := range om.OM_OBJECTS[0].OM_RECORD.Fields {
		// fmt.Printf("%+v\n", i)
		// }
	}
}
