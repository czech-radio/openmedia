package helper

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {
	my := []string{"a", "b", "c", "d", "e"}
	mya := my[1:3]
	fmt.Println(mya[0:2])
}

func TestMapExcelTable(t *testing.T) {
	type args struct {
		filePath     string
		sheetName    string
		headerRow    int
		headerColumn int
	}
	tests := []struct {
		name    string
		args    args
		want    *Table
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MapExcelTable(tt.args.filePath, tt.args.sheetName, tt.args.headerRow, tt.args.headerColumn)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapExcelTable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapExcelTable() = %v, want %v", got, tt.want)
			}
		})
	}
}
