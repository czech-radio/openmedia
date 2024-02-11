package internal

import (
	"errors"
	"fmt"
	"log/slog"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/ncruces/go-strftime"
)

func Test_ErrorsMarshalLog(t *testing.T) {
	errs := []error{errors.New("hello"), errors.New("world")}
	err := ErrorsMarshalLog(errs)
	if err != nil {
		t.Error(err)
	}
}

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

func Test_ParseUplink(t *testing.T) {
	rgxPatt := `(\d*).xml$`
	regexpObject := regexp.MustCompile(rgxPatt)
	name := "ST_letohrad-zprovozneni-vodni-elektrarny-repo_2_18553718_20231212033507.xml"
	matches := regexpObject.FindStringSubmatch(name)
	date, err := strftime.Parse("%Y%m%d%H%M%S", matches[1])
	if err != nil {
		slog.Error(err.Error())
	}
	fmt.Printf("fuck %+v\n", date)
}

func Test_ProcessFolder(t *testing.T) {
	// srcDir := filepath.Join(TEMP_DIR_TEST_SRC, "contacts")
	// srcDir := filepath.Join(TEMP_DIR_TEST_SRC, "rundowns_mock")
	// srcDir := filepath.Join(TEMP_DIR_TEST_SRC, "invalid")
	// srcDir := filepath.Join(TEMP_DIR_TEST_SRC, "garbage")
	// srcDir := filepath.Join(TEMP_DIR_TEST_SRC)
	srcDir := filepath.Join(TEMP_DIR_TEST_SRC, "rundowns_mix")

	dstDir := filepath.Join(TEMP_DIR_TEST_DST)
	opts := ProcessOptions{
		SourceDirectory:      srcDir,
		DestinationDirectory: dstDir,
		InvalidFileRename:    false,
		// InvalidFileContinue:  false,
		InvalidFileContinue:    true,
		CompressionType:        "zip",
		RecurseSourceDirectory: true,
	}
	process := Process{Options: opts}
	err := process.Folder()
	fmt.Printf("%+v\n", process.Results)
	// Sleeper(1000, "s")
	if err != nil {
		t.Error(err)
	}
}

func Test_MapFilesInOldArchive(t *testing.T) {
	// No archive file present
	worker := new(ArchiveWorker)
	err := worker.MapFilesInOldArchive("some_path")
	if err != nil {
		t.Error(err)
	}
}

func Test_ProcessFolderInvalid(t *testing.T) {
	srcDir := filepath.Join(TEMP_DIR_TEST_SRC, "rundowns_invalid")
	dstDir := filepath.Join(TEMP_DIR_TEST_DST)
	opts := ProcessOptions{
		SourceDirectory:        srcDir,
		DestinationDirectory:   dstDir,
		InvalidFileRename:      false,
		InvalidFileContinue:    true,
		CompressionType:        "zip",
		RecurseSourceDirectory: true,
	}
	process := Process{Options: opts}
	err := process.Folder()
	fmt.Printf("%+v\n", process.Results)
	if err != nil {
		t.Error(err)
	}
}

func Test_ProcessFolderComplexNoDupes(t *testing.T) {
	srcDir := filepath.Join(TEMP_DIR_TEST_SRC, "rundowns_complex_nodupes")
	// srcDir := filepath.Join(TEMP_DIR_TEST_SRC, "rundowns_complex_nodupes")
	dstDir := filepath.Join(TEMP_DIR_TEST_DST)
	opts := ProcessOptions{
		SourceDirectory:          srcDir,
		DestinationDirectory:     dstDir,
		InvalidFileRename:        false,
		InvalidFileContinue:      true,
		CompressionType:          "zip",
		PreserveFoldersInArchive: false,
		RecurseSourceDirectory:   true,
		// PreserveFoldersInArchive: true,
	}
	process := Process{Options: opts}
	err := process.Folder()
	fmt.Printf("%+v\n", process.Results)

	// Sleeper(1000, "s")
	if err != nil {
		t.Error(err)
	}
}

func Test_ProcessFolderComplexDupes(t *testing.T) {
	srcDir := filepath.Join(TEMP_DIR_TEST_SRC, "rundowns_complex_dupes")
	dstDir := filepath.Join(TEMP_DIR_TEST_DST)
	// srcDir := filepath.Join(TEMP_DIR_TEST_SRC, "rundowns_complex_nodupes")
	opts := ProcessOptions{
		SourceDirectory:      srcDir,
		DestinationDirectory: dstDir,
		InvalidFileRename:    false,
		// InvalidFileContinue:  false,
		InvalidFileContinue:      true,
		CompressionType:          "zip",
		PreserveFoldersInArchive: false,
		// RecurseSourceDirectory:   false,
		RecurseSourceDirectory: true,
	}
	process := Process{Options: opts}
	err := process.Folder()
	fmt.Printf("%+v\n", process.Results)
	Sleeper(1000, "s")
	if err != nil {
		t.Error(err)
	}
}

func Test_ProcessFolderComplexDupesSame(t *testing.T) {
	srcDir := filepath.Join(TEMP_DIR_TEST_SRC, "rundowns_complex_dupes")
	// srcDir := filepath.Join(TEMP_DIR_TEST_SRC, "rundowns_complex_nodupes")
	// dstDir := filepath.Join(TEMP_DIR_TEST_DST)
	opts := ProcessOptions{
		SourceDirectory:          srcDir,
		DestinationDirectory:     srcDir,
		InvalidFileRename:        false,
		InvalidFileContinue:      true,
		CompressionType:          "zip",
		PreserveFoldersInArchive: false,
		RecurseSourceDirectory:   true,
	}
	process := Process{Options: opts}
	err := process.Folder()
	fmt.Printf("%+v\n", process.Results)
	// Sleeper(1000, "s")
	if err != nil {
		t.Error(err)
	}
}

func Test_ProcessFolderRundownsAppend(t *testing.T) {
	srcDir := filepath.Join(TEMP_DIR_TEST_SRC, "rundowns_append")
	subDirs := []string{
		"dir1",
		"dir1",
		// "dir2",
		// "dir3",
		// "dir4",
	}
	dstDir := filepath.Join(TEMP_DIR_TEST_DST, "rundowns_append")
	for i := range subDirs {
		srcSubDir := filepath.Join(srcDir, subDirs[i])
		fmt.Println("PROCESSING FOLDER", srcSubDir)
		opts := ProcessOptions{
			SourceDirectory:          srcSubDir,
			DestinationDirectory:     dstDir,
			InvalidFileRename:        false,
			InvalidFileContinue:      true,
			CompressionType:          "zip",
			PreserveFoldersInArchive: false,
			// PreserveFoldersInArchive: true,
			RecurseSourceDirectory: true,
		}
		process := Process{Options: opts}
		err := process.Folder()
		fmt.Printf("%+v\n", process.Results)
		if err != nil {
			t.Error(err)
		}
	}
	Sleeper(1000, "s")
}
