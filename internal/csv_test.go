package internal

import (
	"fmt"
	"strings"
	"testing"
)

func TestPartPrintToCSV(t *testing.T) {
	var builder strings.Builder
	rowPart := CSVrowPart{
		"10":   {1, "kek", "kkak"},
		"tek":  {1, "sek", "omoe"},
		"trak": {1, "seklo", "afda"},
	}
	partsPos := CSVrowPartFieldsPositions{
		{"KAK", "10", "Nevim"},
		{"KAK", "tek", "NoName"},
	}
	rowPart.PrintToCSV(&builder, partsPos, CSVdelim)
	fmt.Println(builder.String())
}

func TestRowPrintToCSV(t *testing.T) {
	var builder strings.Builder
	rowPart1 := CSVrowPart{
		"10":   {1, "kek1", "kkak"},
		"tek":  {1, "sek1", "omoe"},
		"trak": {1, "set1", "afda"},
	}
	rowPart2 := CSVrowPart{
		"100":   {1, "kek2", "kkak2"},
		"tek2":  {1, "sek2", "omoe2"},
		"trak2": {1, "set2", "afda2"},
	}
	row := CSVrow{
		"RTS": rowPart1,
		"RTT": rowPart2,
	}

	partFieldPos1 := CSVrowPartFieldsPositions{
		{"RTT", "10", "Nevim"},
		{"RTT", "trak", "Nevim"},
		{"RTT", "tek", "Nevim"},
	}
	partFieldPos2 := CSVrowPartFieldsPositions{
		{"RTS", "100", "Nevim"},
		{"RTS", "trak2", "Nevim"},
		{"RTS", "tek2", "Nevim"},
	}

	partsFieldsPos := CSVrowPartsFieldsPositions{
		"RTT": partFieldPos2,
		"RTS": partFieldPos1,
	}
	partsPos := CSVrowPartsPositions{
		"RTS", "RTT",
	}

	row.PrintToCSV(&builder, partsPos, partsFieldsPos, CSVdelim)
	fmt.Println(builder.String())
}
