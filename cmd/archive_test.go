package cmd

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

var CommandArchivePresets = [][]string{
	{"help",
		"archive", "-h"},

	{"debug config",
		"-v=0", "-dc", "archive", "-ifc", "-R"},

	{"run exit on src file error",
		"-v=0", "archive", "-ifc"},

	{"run do not exit on src file error",
		"-v=0", "archive"},

	{"dry run, no exit on src file error",
		"-dr", "-v=0", "archive"},

	// {"do not recurse source directory",
	// "-v=0", "archive"},

	// {"recurse source directory",
	// "-v=0", "archive", "-R"},
}

// TODO: report number of archived files

func TestRunCommand_Archive(t *testing.T) {
	commandName := "archive"
	testSubdir := filepath.Join("cmd", commandName)
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t, testSubdir)
	for i, flags := range CommandArchivePresets {
		srcDir := filepath.Join(
			testerConfig.TempDataSource, testSubdir)
		dstDir := filepath.Join(
			testerConfig.TempDataOutput, testSubdir, strconv.Itoa(i+1))
		err := os.Mkdir(dstDir, 0700)
		if err != nil {
			panic(err)
		}
		flagss := append(flags, "-sdir="+srcDir)
		flagss = append(flagss, "-odir="+dstDir)
		fn := ReturnTestFunc(i+1, commandName, testSubdir, flagss)
		t.Run(strconv.Itoa(i+1), fn)
	}
}
