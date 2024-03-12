package internal

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"testing"
)

func TestTransformXML(t *testing.T) {
	src_file := filepath.Join(TEMP_DIR_TEST_SRC, "rundowns_valid", "RD_00-12_Pohoda_-_Fri_06_01_2023_orig.xml")
	// src_file := filepath.Join(TEMP_DIR_TEST_SRC, "rundowns_valid", "RD_00-12_Pohoda_-_Fri_06_01_2023_orig_wo_header.xml")
	srcFile, err := os.Open(src_file)
	if err != nil {
		t.Error(err)
	}
	defer srcFile.Close()
	pr := PipeUTF16leToUTF8(srcFile)
	pr = PipeRundownHeaderAmmend(pr)
	om, err := PipeRundownUnmarshal(pr) // treat EOF error
	if err != nil {
		slog.Error(err.Error())
		t.Error(err)
	}
	pr = PipeRundownMarshal(om)
	pr = PipeRundownHeaderAdd(pr)
	_ = pr
	// PipePrint(pr)
	if err != nil {
		t.Error(err)
	}
}

func Test_ValidateFilesInDirectory(t *testing.T) {
	valid := "rundowns_valid"
	srcPath := filepath.Join(TEMP_DIR_TEST_SRC, valid)
	_, err := ValidateFilesInDirectory(srcPath, true)
	if err != nil {
		t.Error(err)
	}

	invalid := "rundowns_invalid"
	srcPath = filepath.Join(TEMP_DIR_TEST_SRC, invalid)
	_, err = ValidateFilesInDirectory(srcPath, true)
	if err == nil {
		t.Error("failed to catch error")
	}
}

func Test_XMLbuildAttrQuery(t *testing.T) {
	fmt.Println(XMLbuildAttrQuery("FieldID", []string{"1", "2", "3"}))
}
