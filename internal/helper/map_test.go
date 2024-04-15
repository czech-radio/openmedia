package helper

import (
	"maps"
	"testing"
)

func TestMapCopy(t *testing.T) {
	originalMap := map[string]map[string]string{
		"key1": {"innerKey1": "value1", "innerKey2": "value2"},
		"key2": {"innerKey3": "value3", "innerKey4": "value4"},
	}
	copiedMap := make(map[string]map[string]string)
	maps.Copy(copiedMap, originalMap)
	originalMap["key1"]["innerKey1"] = "ekk"
	PrintMap(originalMap)
	PrintMap(copiedMap)
}

// func TestMapCopy2(t *testing.T) {
// originalRow := CSVrow{
// "PartA": {"field1": {"Ahoj", "Hello", "Hi"}},
// "PartB": {"field1": {"Ahoj", "Hello", "Hi"}},
// }
// newRow := CopyRow(originalRow)
// originalRow["PartA"]["field1"] = CSVrowField{
// "Kek", "Mek", "sek"}
// PrintRow(originalRow)
// PrintRow(newRow)
// }
