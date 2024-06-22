package cmd

import (
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"testing"
)

var CommandArchivePresets = [][]string{
	{"help", "archive", "-h"},
	{"debug config", "-dc", "-v=0", "archive", "-ifc", "-R"},
	{"run exit on src file error", "-v=0", "archive", "-ifc"},
	{"run do not exit on src file error", "-v=0", "archive"},
	{"dry run, no exit on src file error", "-dr", "-v=0", "archive"},
	{"recurse source directory", "-v=0", "archive", "-ifc", "-R"},
}

// TODO: find out why not recurse directory
// TODO: report number of archived files

func TestRunCommand_Archive(t *testing.T) {
	commandName := "archive"
	testSubdir := filepath.Join("cmd", commandName)
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t, testSubdir)
	for i, flags := range CommandArchivePresets {
		fn := func(t *testing.T) {
			// testerConfig.TempDestinationPathGeter(testSubdir)
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
			command := append([]string{
				"run", "../main.go"}, flagss[1:]...)
			cmdexec := exec.Command("go", command...)
			resultLog, err := cmdexec.CombinedOutput()
			PrintOutput(len(CommandArchivePresets), i, commandName,
				flagss, resultLog)
		}
		t.Run(strconv.Itoa(i), fn)
	}
}
