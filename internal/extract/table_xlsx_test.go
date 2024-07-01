package extract

import (
	"testing"

	"github.com/xuri/excelize/v2"
)

func TestXLSXsetStyle(t *testing.T) {
	testSubdir := "exports"
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t, testSubdir)
	tp := testerConfig.TempSourcePathGeter(testSubdir)
	fileName := "production_production_contacts_base_wh.xlsx"
	fileSrcPath := tp(fileName)

	tp = testerConfig.TempDestinationPathGeter(testSubdir)
	fileDstPath := tp(fileName)
	sheet := "Sheet1"
	file, err := excelize.OpenFile(fileSrcPath)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	err = file.SetColWidth(sheet, "A", "B", 30)
	if err != nil {
		t.Error(err)
	}
	// err = file.Save()
	err = file.SaveAs(fileDstPath)
	if err != nil {
		t.Error(err)
	}
}
