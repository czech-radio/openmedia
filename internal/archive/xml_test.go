package internal

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"testing"
)

func TestXMLqueryFromPath(t *testing.T) {
	path := "/Radio Rundown/<OM_RECORD>/*Hourly Rundown/<OM_RECORD>"
	res := XMLqueryFromPath(path)
	fmt.Println(res)
}

func TestTransformXML(t *testing.T) {
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t, true)

	src_file := filepath.Join(
		testerConfig.TestDataSource, "rundowns_valid",
		"RD_00-12_Pohoda_-_Fri_06_01_2023_orig.xml")
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
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t, true)

	valid := "rundowns_valid"
	srcPath := filepath.Join(testerConfig.TempDataSource, valid)

	_, err := ValidateFilesInDirectory(srcPath, true)
	if err != nil {
		t.Error(err)
	}

	invalid := "rundowns_invalid"
	srcPath = filepath.Join(testerConfig.TempDataSource, invalid)
	_, err = ValidateFilesInDirectory(srcPath, true)
	if err == nil {
		t.Error("failed to catch error")
	}
}

func Test_XMLbuildAttrQuery(t *testing.T) {
	fmt.Println(XMLbuildAttrQuery("FieldID", []string{"1", "2", "3"}))
	fmt.Println(XMLbuildAttrQuery("TemplateName", []string{"Audioclip", "Contact Item"}))
}
