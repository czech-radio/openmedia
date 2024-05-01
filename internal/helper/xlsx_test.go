package helper

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	my := []string{"a", "b", "c", "d", "e"}
	mya := my[1:3]
	fmt.Println(mya[0:2])
	fmt.Println(len(mya))
	fmt.Println(len(my))
}

func TestTable_MapTableHeaders(t *testing.T) {
	filterFile := "/home/jk/CRO/CRO_BASE/openmedia_backup/filters/eurovolby - zadání.xlsx"
	rows, err := ReadExcelFileSheetRows(filterFile, "data")
	if err != nil {
		t.Error(err)
	}
	table := CreateTable(rows, 0, 0)
	fmt.Println(len(table.RowHeaderToColumnMap))
	fmt.Println(table.RowHeaderToColumnMap)
	fmt.Println(len(table.ColumnHeaderMap))
	// match := table.MatchRow("Pochman Stanislav", "navrhující strana", "KAN")
	match := table.MatchRow("Rohel Petr", "navrhující strana", "Levice")
	// match := table.MatchRow("Juřica Vojtěch", "navrhující strana", "Levice")
	// match := table.MatchRow("Široký Jan", "navrhující strana", "Levice")
	fmt.Println(match)
}
