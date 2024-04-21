package internal

import (
	"fmt"
	"log/slog"
	"os"
	"testing"
)

func TestXMLqueryFromPath(t *testing.T) {
	path := "/Radio Rundown/<OM_RECORD>/*Hourly Rundown/<OM_RECORD>"
	res := XMLqueryFromPath(path)
	fmt.Println(res)
}

func TestTransformXML(t *testing.T) {
	testSubdir := "rundowns_valid"
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t, testSubdir)
	tp := testerConfig.TempSourcePathGeter(testSubdir)

	src_file := "RD_00-12_Pohoda_-_Fri_06_01_2023_orig.xml"
	fmt.Println("kek", tp(src_file))
	srcFile, err := os.Open(tp(src_file))
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
	validfiles := "rundowns_valid"
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t, validfiles)
	tp := testerConfig.TempSourcePathGeter(validfiles)
	srcPath := tp("")

	_, err := ValidateFilesInDirectory(srcPath, true)
	if err != nil {
		t.Error(err)
	}

	invalidfiles := "rundowns_invalid"
	tp = testerConfig.TempSourcePathGeter(invalidfiles)
	srcPath = tp("")
	_, err = ValidateFilesInDirectory(srcPath, true)
	if err == nil {
		t.Error("failed to catch error")
	}
}

func Test_XMLbuildAttrQuery(t *testing.T) {
	fmt.Println(XMLbuildAttrQuery("FieldID", []string{"1", "2", "3"}))
	fmt.Println(XMLbuildAttrQuery("TemplateName", []string{"Audioclip", "Contact Item"}))
}
