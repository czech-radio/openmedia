package internal

import (
	"log/slog"
	"os"
	"path/filepath"
	"testing"
)

func TestTransformXML(t *testing.T) {
	// src_file := filepath.Join(TEMP_DIR_TEST_SRC, "rundowns_valid", "RD_00-12_Pohoda_-_Fri_06_01_2023_orig.xml")
	src_file := filepath.Join(TEMP_DIR_TEST_SRC, "rundowns_valid", "RD_00-12_Pohoda_-_Fri_06_01_2023_orig_wo_header.xml")
	srcFile, err := os.Open(src_file)
	if err != nil {
		t.Error(err)
	}
	defer srcFile.Close()
	pr := PipeUTF16leToUTF8(srcFile)
	pr = PipeRundownHeaderAmmend(pr)
	// PrintReadrPipe(pr)
	om, err := PipeRundownMinfiy(pr) // treat EOF error
	if err != nil {
		slog.Error(err.Error())
		t.Error(err)
	}
	err = RundownMarshal(om, "kek")
	// fmt.Printf("%+v\n", om)
	// res, err := om.FileDate()
	if err != nil {
		t.Error(err)
	}
	Sleeper(100, "s")
	// fmt.Println(res.ISOWeek())
}
