package extract

import (
	"fmt"
	"strings"
	"testing"
)

func TestPartPrintToCSV(t *testing.T) {
	var builder strings.Builder
	rowPart := CSVrowPart{
		"id_Kek": {"id_Kek", "n_Kek", "kardamon"},
		"id_Sek": {"id_Sek", "n_Sek", "cinnamon"},
		"id_Tak": {"id_Tak", "n_Tak", "vanilin"},
	}
	partsPos := CSVrowPartFieldsPositions{
		{"KAK", "id_Sek", "Nevim"},
		{"KAK", "id_Tak", "NoName"},
	}
	rowPart.CastToCSV(&builder, partsPos, CSVdelim)
	fmt.Println(builder.String())
}

func TestRowPrintToCSV(t *testing.T) {
	var builder strings.Builder
	rowPart1 := CSVrowPart{
		"id_Kek": {"id_Kek", "n_Kek", "kardamon"},
		"id_Sek": {"id_Sek", "n_Sek", "cinnamon"},
		"id_Tak": {"id_Tak", "n_Tak", "vanilin"},
	}
	rowPart2 := CSVrowPart{
		"id_Kek": {"id_Kek", "n_Kek", "cumin"},
		"id_Sek": {"id_Sek", "n_Sek", "peper"},
		"id_Tak": {"id_Tak", "n_Tak", "chilli"},
	}
	row := CSVrow{
		FieldPrefix_SubHead:   rowPart1,
		FieldPrefix_StoryHead: rowPart2,
	}

	partFieldPos1 := CSVrowPartFieldsPositions{
		{"A", "id_Sek", "Nevim"},
		{"A", "id_Tak", "NoName"},
	}
	partFieldPos2 := CSVrowPartFieldsPositions{
		{"B", "id_Sek", "Nevim"},
		{"B", "id_Tak", "NoName"},
	}

	partsFieldsPos := CSVrowPartsFieldsPositions{
		FieldPrefix_SubHead:   partFieldPos2,
		FieldPrefix_StoryHead: partFieldPos1,
	}
	partsPos := CSVrowPartsPositionsInternal{
		FieldPrefix_SubHead, FieldPrefix_StoryHead,
	}

	row.CastToCSV(&builder, partsPos, partsFieldsPos, CSVdelim)
	fmt.Println(builder.String())
}
