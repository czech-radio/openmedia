package internal

import (
	"fmt"
	"path/filepath"
	"testing"
)

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

func Test_ZipArchive(t *testing.T) {
	// srcDir := filepath.Join(TEMP_DIR_TEST_SRC, "rundowns_additional")
	srcDir := filepath.Join(TEMP_DIR_TEST_SRC)
	dstFile := filepath.Join(TEMP_DIR_TEST_DST, "kek.tar.gz")
	err, results := ZipArchive(srcDir, dstFile)
	if err != nil {
		fmt.Printf("%+v\n", results)
	} else {
		t.Error(err.Error())
	}
	Sleeper(100, "s")
}

func Test_TarGzArchive(t *testing.T) {
	// srcDir := filepath.Join(TEMP_DIR_TEST_SRC, "rundowns_additional")
	srcDir := filepath.Join(TEMP_DIR_TEST_SRC)
	dstFile := filepath.Join(TEMP_DIR_TEST_DST, "kek.tar.gz")
	err, results := TarGzArchive(srcDir, dstFile)
	if err != nil {
		fmt.Printf("%+v\n", results)
	} else {
		t.Error(err.Error())
	}
}
