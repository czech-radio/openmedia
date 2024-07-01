package extract

import (
	"fmt"
	"strings"
	"testing"
)

func TestPartPrintToCSV(t *testing.T) {
	var builder strings.Builder
	rowPart := RowPart{
		"id_Kek": {"id_Kek", "n_Kek", "kardamon"},
		"id_Sek": {"id_Sek", "n_Sek", "cinnamon"},
		"id_Tak": {"id_Tak", "n_Tak", "vanilin"},
	}
	partsPos := RowPartFieldsPositions{
		{"KAK", "id_Sek", "Nevim"},
		{"KAK", "id_Tak", "NoName"},
	}
	rowPart.CSVrowPartBuild(&builder, partsPos, CSVdelim)
	fmt.Println(builder.String())
}

func TestRowPrintToCSV(t *testing.T) {
	var builder strings.Builder
	rowPart1 := RowPart{
		"id_Kek": {"id_Kek", "n_Kek", "kardamon"},
		"id_Sek": {"id_Sek", "n_Sek", "cinnamon"},
		"id_Tak": {"id_Tak", "n_Tak", "vanilin"},
	}
	rowPart2 := RowPart{
		"id_Kek": {"id_Kek", "n_Kek", "cumin"},
		"id_Sek": {"id_Sek", "n_Sek", "peper"},
		"id_Tak": {"id_Tak", "n_Tak", "chilli"},
	}
	row := RowParts{
		RowPartCode_SubHead:   rowPart1,
		RowPartCode_StoryHead: rowPart2,
	}

	partFieldPos1 := RowPartFieldsPositions{
		{"A", "id_Sek", "Nevim"},
		{"A", "id_Tak", "NoName"},
	}
	partFieldPos2 := RowPartFieldsPositions{
		{"B", "id_Sek", "Nevim"},
		{"B", "id_Tak", "NoName"},
	}

	partsFieldsPos := RowPartsFieldsPositions{
		RowPartCode_SubHead:   partFieldPos2,
		RowPartCode_StoryHead: partFieldPos1,
	}
	partsPos := RowPartsPositions{
		RowPartCode_SubHead, RowPartCode_StoryHead,
	}

	row.CSVrowBuild(&builder, partsPos, partsFieldsPos, CSVdelim)
	fmt.Println(builder.String())
}

func TestRowPartsPositionsFromCSVfile(t *testing.T) {
	testSubdir := "exports"
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t, testSubdir)
	tp := testerConfig.TempSourcePathGeter(testSubdir)
	filePath := tp(
		"production_production_contacts_base_wh.csv")
	type args struct {
		fileName string
		delim    rune
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"one", args{filePath, '\t'}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0, got1, got2, err := RowPartsPositionsFromCSVfile(tt.args.fileName, tt.args.delim)
			if (err != nil) != tt.wantErr {
				t.Errorf("RowPartsPositionsFromCSVfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got0[0])
			fmt.Println(got1)
			fmt.Println(got2)
		})
	}
}
