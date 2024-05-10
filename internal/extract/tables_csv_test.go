package extract

import (
	"path/filepath"
	"testing"
)

func TestCSVcompareRows(t *testing.T) {
	srcDir := "/tmp/test/"
	file1 := filepath.Join(srcDir, "all_filtr_day_2024-01-01_production1_eurovolby_woh.csv")
	file2 := filepath.Join(srcDir, "filtered.csv")
	type args struct {
		fileName1 string
		fileName2 string
	}
	tests := []struct {
		name string
		args args
	}{
		{"test1", args{file1, file2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CSVcompareRows(tt.args.fileName1, tt.args.fileName2)
		})
	}
}
