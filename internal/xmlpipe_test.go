package internal

import (
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
	PipePrint(pr)
	if err != nil {
		t.Error(err)
	}
}

func Test_ValidateFilesInDirectory(t *testing.T) {
	srcDir := filepath.Join(TEMP_DIR_TEST_SRC)
	_, err := ValidateFilenamesInDirectory(srcDir)
	if err == nil {
		t.Error(err)
	}
	srcDir2 := filepath.Join(TEMP_DIR_TEST_SRC, "testdata", "rundowns_valid")
	_, err = ValidateFilenamesInDirectory(srcDir2)
	if err != nil {
		t.Error(err)
	}
}
