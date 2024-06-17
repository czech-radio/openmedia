package internal

import (
	"errors"
	"fmt"

	"github.com/triopium/go_utils/pkg/helper"

	"log/slog"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/ncruces/go-strftime"
)

var testerConfig = helper.TesterConfig{
	TestDataSource: "../../test/testdata",
}

func TestMain(m *testing.M) {
	testerConfig.TesterMain(m)
}

func TestAABnofail(t *testing.T) {
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t)
	helper.Sleeper(2, "s")
}

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
	_ = date
}

func Test_ProcessFolder(t *testing.T) {
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t, "rundowns_mix")

	subDir := "rundowns_mix"
	srcDir := filepath.Join(
		testerConfig.TempDataSource, subDir)
	dstDir := filepath.Join(
		testerConfig.TempDataOutput, subDir)
	opts := ArchiveOptions{
		SourceDirectory:   srcDir,
		OutputDirectory:   dstDir,
		InvalidFileRename: false,
		// InvalidFileContinue:  false,
		InvalidFileContinue:    true,
		CompressionType:        "zip",
		RecurseSourceDirectory: true,
	}
	process := Archive{Options: opts}
	err := process.Folder()
	if err != nil {
		t.Error(err)
	}
}

func Test_ProcessFolderInvalid(t *testing.T) {
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t)

	subDir := "rundowns_invalid"
	srcDir := filepath.Join(
		testerConfig.TempDataSource, subDir)
	dstDir := filepath.Join(
		testerConfig.TempDataOutput, subDir)

	opts := ArchiveOptions{
		SourceDirectory:        srcDir,
		OutputDirectory:        dstDir,
		InvalidFileRename:      false,
		InvalidFileContinue:    true,
		CompressionType:        "zip",
		RecurseSourceDirectory: true,
	}
	process := Archive{Options: opts}
	err := process.Folder()
	if err != nil {
		t.Error(err)
	}
}

func Test_ProcessFolderComplexNoDupes(t *testing.T) {
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t)

	subDir := "rundowns_complex_nodupes"
	srcDir := filepath.Join(
		testerConfig.TempDataSource, subDir)
	dstDir := filepath.Join(
		testerConfig.TempDataOutput, subDir)

	opts := ArchiveOptions{
		SourceDirectory:          srcDir,
		OutputDirectory:          dstDir,
		InvalidFileRename:        false,
		InvalidFileContinue:      true,
		CompressionType:          "zip",
		PreserveFoldersInArchive: false,
		RecurseSourceDirectory:   true,
		// PreserveFoldersInArchive: true,
	}
	process := Archive{Options: opts}
	err := process.Folder()
	if err != nil {
		t.Error(err)
	}
}

func Test_ProcessFolderComplexDupes(t *testing.T) {
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t)

	subDir := "rundowns_complex_dupes"
	srcDir := filepath.Join(
		testerConfig.TempDataSource, subDir)
	dstDir := filepath.Join(
		testerConfig.TempDataOutput, subDir)

	opts := ArchiveOptions{
		SourceDirectory:   srcDir,
		OutputDirectory:   dstDir,
		InvalidFileRename: false,
		// InvalidFileContinue:  false,
		InvalidFileContinue:      true,
		CompressionType:          "zip",
		PreserveFoldersInArchive: false,
		// RecurseSourceDirectory:   false,
		RecurseSourceDirectory: true,
	}
	process := Archive{Options: opts}
	err := process.Folder()
	if err != nil {
		t.Error(err)
	}
}

func Test_ProcessFolderComplexDupesSame(t *testing.T) {
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t)

	subDir := "rundowns_complex_dupes"
	srcDir := filepath.Join(
		testerConfig.TempDataSource, subDir)

	opts := ArchiveOptions{
		SourceDirectory:          srcDir,
		OutputDirectory:          srcDir,
		InvalidFileRename:        false,
		InvalidFileContinue:      true,
		CompressionType:          "zip",
		PreserveFoldersInArchive: false,
		RecurseSourceDirectory:   true,
	}
	process := Archive{Options: opts}
	err := process.Folder()
	fmt.Printf("%+v\n", process.Results)
	if err != nil {
		t.Error(err)
	}
}

func Test_ProcessFolderRundownsAppend(t *testing.T) {
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t)

	subDir := "rundowns_append"
	srcDir := filepath.Join(
		testerConfig.TempDataSource, subDir)
	subDirs := []string{
		"dir1",
		// "dir2",
		// "dir3",
		// "dir4",
	}
	dstDir := filepath.Join(
		testerConfig.TempDataOutput, subDir)
	for i := range subDirs {
		srcSubDir := filepath.Join(srcDir, subDirs[i])
		fmt.Println("PROCESSING FOLDER", srcSubDir)
		opts := ArchiveOptions{
			SourceDirectory:          srcSubDir,
			OutputDirectory:          dstDir,
			InvalidFileRename:        false,
			InvalidFileContinue:      true,
			CompressionType:          "zip",
			PreserveFoldersInArchive: false,
			// PreserveFoldersInArchive: true,
			ProcessedFileRename:    true,
			RecurseSourceDirectory: true,
		}
		process := Archive{Options: opts}
		err := process.Folder()
		fmt.Printf("%+v\n", process.Results)
		if err != nil {
			t.Error(err)
		}
	}
}

func Test_ProcessFolderDate(t *testing.T) {
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t)

	subDir := "rundowns_date"
	srcDir := filepath.Join(
		testerConfig.TempDataSource, subDir)
	dstDir := filepath.Join(
		testerConfig.TempDataOutput, subDir)

	opts := ArchiveOptions{
		SourceDirectory:          srcDir,
		OutputDirectory:          dstDir,
		InvalidFileRename:        false,
		InvalidFileContinue:      true,
		CompressionType:          "zip",
		PreserveFoldersInArchive: false,
		RecurseSourceDirectory:   true,
	}
	process := Archive{Options: opts}
	err := process.Folder()
	if err != nil {
		t.Error(err)
	}
}
