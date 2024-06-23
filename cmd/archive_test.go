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

	{"dry run",
		"-dr", "-v=0", "archive"},

	{"exit on src file error or filename error",
		"-v=0", "archive", "-ifc", "-ifnc", "-R"},

	{"run exit on src file error",
		"-v=0", "archive", "-ifc", "-R"},

	{"do not exit on any file error",
		"-v=0", "archive", "-R"},

	{"do not recurse the source folder",
		"-v=0", "archive"},
}

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
		fn := ReturnTestFunc(
			len(CommandArchivePresets), i+1, commandName,
			testSubdir, flagss)
		t.Run(strconv.Itoa(i+1), fn)
	}
}
