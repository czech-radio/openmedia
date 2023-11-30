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
	// fmt.Printf("%+v\n", om)
	// res, err := om.FileDate()
	if err != nil {
		t.Error(err)
	}
	// Sleeper(100, "s")
	// fmt.Println(res.ISOWeek())
}

func Test_ProcessFolder(t *testing.T) {
	// srcDir := filepath.Join(TEMP_DIR_TEST_SRC, "rundowns_additional")
	srcDir := filepath.Join(TEMP_DIR_TEST_SRC)
	dstDir := filepath.Join(TEMP_DIR_TEST_DST)
	opts := ProcessFolderOptions{
		SourceDirectory:        srcDir,
		DestinationDirectory:   dstDir,
		InputEncoding:          "",
		OutputEncoding:         "",
		ValidateWithDefaultXSD: false,
		ValidateWithXSD:        "",
		ValidatePre:            false,
		ValidatePost:           false,
		ArchiveType:            "",
		InvalidFileRename:      false,
		// InvalidFileContinue:    false,
		InvalidFileContinue: true,
	}
	_, err := ProcessFolder(opts)
	// Sleeper(100, "s")
	if err != nil {
		t.Error(err)
	}
}
