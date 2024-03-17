package internal

var CSVdelim = "\t"

// func TestPartPrintToCSV(t *testing.T) {
// 	var builder strings.Builder
// 	rowPart := CSVrowPart{
// 		"10":   {1, "kek", "kkak"},
// 		"tek":  {1, "sek", "omoe"},
// 		"trak": {1, "seklo", "afda"},
// 	}
// 	partsPos := []string{"tek", "trak", "10"}
// 	rowPart.PrintToCSV(&builder, partsPos, CSVdelim)
// 	fmt.Println(builder.String())
// }

// func TestRowPrintToCSV(t *testing.T) {
// 	var builder strings.Builder
// 	rowPart1 := CSVrowPart{
// 		"10":   {1, "kek1", "kkak"},
// 		"tek":  {1, "sek1", "omoe"},
// 		"trak": {1, "set1", "afda"},
// 	}
// 	rowPart2 := CSVrowPart{
// 		"100":   {1, "kek2", "kkak2"},
// 		"tek2":  {1, "sek2", "omoe2"},
// 		"trak2": {1, "set2", "afda2"},
// 	}
// 	row := CSVrow{
// 		"RTS": rowPart1,
// 		"RTT": rowPart2,
// 	}

// 	partFieldPos1 := []string{"tek", "trak", "10"}
// 	partFieldPos2 := []string{"trak2", "tek2", "100"}

// 	partsFieldsPos := CSVrowPartsFieldsPositions{
// 		"RTT": partFieldPos2,
// 		"RTS": partFieldPos1,
// 	}
// 	partsPos := []string{"RTT", "RTS"}

// 	res := row.PrintToCSV(&builder, partsPos, partsFieldsPos, CSVdelim)
// 	fmt.Println(res)
// }

// func TestRowPrintTableToCSV(t *testing.T) {
// }
