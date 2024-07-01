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
	defer testerConfig.RecoverPanic(t)
	tests := []struct {
		name         string
		subDir       string
		wantErr      bool
		InvalidCount int
	}{
		{"valid files", "rundowns_valid", false, 0},
		{"invalid files", "rundowns_invalid", false, 1},
	}
	subDirs := []string{}
	for _, tt := range tests {
		subDirs = append(subDirs, tt.subDir)
	}
	testerConfig.InitTest(t, subDirs...)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tp := testerConfig.TempSourcePathGeter(tt.subDir)
			srcPath := tp(".")
			res, err := ValidateFilenamesInDirectory(srcPath, true)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateFilenamesInDirectory() error = %v, wantErr %v",
					err, tt.wantErr)
			}
			if res.FailureCount != tt.InvalidCount {
				t.Errorf("ValidateFilenamesInDirectory() res = %v, invalid files count found is not correct: %v vs %v",
					res, tt.InvalidCount, res.FailureCount)
			}
		})
	}
}

func Test_XMLbuildAttrQuery(t *testing.T) {
	fmt.Println(XMLbuildAttrQuery("FieldID", []string{"1", "2", "3"}))
	fmt.Println(XMLbuildAttrQuery("TemplateName", []string{"Audioclip", "Contact Item"}))
}
