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
	srcFile, err := os.Open(src_file)
	if err != nil {
		t.Error(err)
	}
	defer srcFile.Close()
	pr := PipeUTF16leToUTF8(srcFile)
	pr = PipeRundownheaderAmmend(pr)
	// PrintReadrPipe(pr)
	om, err := PipeRundownMinfiy(pr)
	if err != nil {
		slog.Error(err.Error())
		t.Error(err)
	}
	fmt.Printf("%+v\n", om)
	// res, err := om.FileDate()
	// if err != nil {
	// t.Error(err)
	// }
	// fmt.Println(res.ISOWeek())
}
