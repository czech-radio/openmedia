package internal

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
	rowPart.PrintToCSV(&builder, partsPos, CSVdelim)
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
		"A": rowPart1,
		"B": rowPart2,
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
		"A": partFieldPos2,
		"B": partFieldPos1,
	}
	partsPos := CSVrowPartsPositions{
		"A", "B",
	}

	row.PrintToCSV(&builder, partsPos, partsFieldsPos, CSVdelim)
	fmt.Println(builder.String())
}
