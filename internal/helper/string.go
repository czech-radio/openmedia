package helper

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"
)

// func PrintRow(input CSVrow) {
// for ai, a := range input {
// fmt.Println(ai, a)
// }
// }

// func PrintRows(rows map[int]CSVrowFields) {
// for i := 0; i < len(rows); i++ {
// fmt.Println(i, rows[i])
// fmt.Println()
// }
// }

func PrintMap(input map[string]map[string]string) {
	for ai, a := range input {
		fmt.Println(ai, a)
	}
}

func PrintObjJson(mark string, input any) {
	res, err := json.MarshalIndent(input, "", "\t")
	if err != nil {
		slog.Error("cannot marshal structure", "mark", mark, "input", input, "err", err.Error())
		return
	}
	fmt.Println(mark, string(res))
}

func JoinObjectPath(oldpath, newpath string) string {
	return oldpath + "/" + newpath
}

func EscapeCSVdelim(value string) string {
	// out := strings.TrimSpace(value)
	out := strings.ReplaceAll(value, "\t", "\\t")
	out = strings.ReplaceAll(out, "\n", "\\n")
	return out
}
