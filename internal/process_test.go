package internal

import (
	"fmt"
	"path/filepath"
	"regexp"
	"testing"
)

func Test_RundownNameParse(t *testing.T) {
	type Case struct {
		input  string
		result string
		error  error
	}
	// regexPattern := `(\p{L}+\s?\p{L}*\s?\p{L}*)`
	// regexPattern := `([[^alpha]-\d])*`
	regexPattern := `([\d[:ascii:]]*)`
	regexpObject := regexp.MustCompile(regexPattern)
	Cases := []Case{
		Case{
			input:  "05-09 ČRo Region SC - Středa 04.03.2020",
			result: "ČRo Region SC",
			error:  nil,
		},
		Case{
			input:  "05-09 ČRo Sever - Wed, 04.03.2020",
			result: "ČRo Sever",
			error:  nil,
		},
	}
	for _, c := range Cases {
		matches := regexpObject.FindStringSubmatch(c.input)
		if len(matches) > 0 {
			fmt.Printf("%q\n", c.input)
			fmt.Printf("%q\n", matches[1])
			fmt.Printf("%q\n", matches)
			if matches[1] != c.result {
				t.Error("does not match")
			}

		} else {
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
