package internal

import (
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func Test_RundownNameParse(t *testing.T) {
	type Case struct {
		input  string
		result string
		error  error
	}
	regexPattern := `([\d[:ascii:]]*)([\p{L}\ ]*)`
	// \p{L} unicode letter
	regexpObject := regexp.MustCompile(regexPattern)
	Cases := []Case{
		{input: "05-09 ČRo Region SC - Středa 04.03.2020",
			result: "ČRo Region SC",
			error:  nil},
		{input: "05-09 ČRo Sever - Wed, 04.03.2020",
			result: "ČRo Sever",
			error:  nil},
	}
	for _, c := range Cases {
		matches := regexpObject.FindStringSubmatch(c.input)
		var name string = ""
		if len(matches) == 3 {
			name = strings.TrimSpace(matches[2])
		}
		if name != c.result {
			t.Error("does not match", c.input)
		}
	}
}

func Test_ProcessFolder(t *testing.T) {
	srcDir := filepath.Join(TEMP_DIR_TEST_SRC, "rundowns_additional")
	// srcDir := filepath.Join(TEMP_DIR_TEST_SRC, "rundowns_mock")
	// srcDir := filepath.Join(TEMP_DIR_TEST_SRC, "rundowns_valid")
	// srcDir := filepath.Join(TEMP_DIR_TEST_SRC)
	dstDir := filepath.Join(TEMP_DIR_TEST_DST)
	opts := ProcessOptions{
		SourceDirectory:        srcDir,
		DestinationDirectory:   dstDir,
		InputEncoding:          "",
		OutputEncoding:         "",
		ValidateWithDefaultXSD: false,
		ValidateWithXSD:        "",
		ValidatePre:            false,
		ValidatePost:           false,
		ArchiveType:            "zip",
		InvalidFileRename:      false,
		// InvalidFileContinue:    false,
		InvalidFileContinue: true,
	}
	process := Process{Options: opts}
	err := process.Folder()
	// fmt.Printf("%+v\n", process.Results)
	Sleeper(1000, "s")
	if err != nil {
		t.Error(err)
	}
}

// func Test_RundownFileNameNormalize(t *testing.T) {
// srcDir := filepath.Join(TEMP_DIR_TEST_SRC, "rundowns_additional")
// ValidateFilenamesInDirectory(srcDir)
// }
