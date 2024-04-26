package helper

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"

	"github.com/antchfx/xmlquery"
	"github.com/go-xmlfmt/xmlfmt"
)

// PrintMap
func PrintMap(input map[string]map[string]string) {
	for ai, a := range input {
		fmt.Println(ai, a)
	}
}

// PrintObjectJson marshals provided input object. prefix is string which is added as prefix to resulting json.
func PrintObjectJson(prefix string, input any) {
	res, err := json.MarshalIndent(input, "", "\t")
	if err != nil {
		slog.Error("cannot marshal structure", "mark", prefix, "input", input, "err", err.Error())
		return
	}
	fmt.Println(prefix, string(res))
}

// JoinObjectPath joins object path in defined way
func JoinObjectPath(oldpath, newpath string) string {
	return oldpath + "/" + newpath
}

// EscapeCSVdelim escapes value such that the value does not colide with csv tab delimiter.
func EscapeCSVdelim(value string) string {
	out := strings.ReplaceAll(value, "\t", "\\t")
	out = strings.ReplaceAll(out, "\n", "\\n")
	return out
}

// XMLprint prints selected node as xml
func XMLprint(node *xmlquery.Node) {
	ex := xmlfmt.FormatXML(node.OutputXML(true), "", "\t")
	fmt.Println(ex)
}

func FirstLetterToLowercase(input string) string {
	return strings.ToLower(input[0:1]) + input[1:]
}
