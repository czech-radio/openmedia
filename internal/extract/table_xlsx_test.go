package extract

import (
	"testing"

	"github.com/xuri/excelize/v2"
)

func TestXLSXsetStyle(t *testing.T) {
	filePath := "/tmp/test/test.xlsx"
	sheet := "Sheet1"
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	err = file.SetColWidth(sheet, "A", "B", 30)
	if err != nil {
		t.Error(err)
	}
	// err = file.Save()
	err = file.SaveAs(filePath)
	if err != nil {
		t.Error(err)
	}
}
