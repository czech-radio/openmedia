package cmd

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

var CommandArchivePresets = [][]string{
	{"help", "archive", "-h"},
	{"run err exit no", "-v=0", "archive"},
	// {"debug config", "-dc", "-v=0", "archive"},
	// {"dry run err exit", "-dr", "-v=0", "archive"},
	// {"dry run err exit no", "-dr", "-v=0", "archive", "-ifc"},
	// {"run err exit no", "-dc", "-v=-4", "archive", "-ifc"},
	// {"run err exit no", "-v=-4", "archive", "-ifc", "-R"},
	// {"run err exit no", "-v=0", "archive", "-ifc"},
	// {"run err exit no", "-v=-4", "archive", "-ifc"},
	// {"run err exit no", "-dc", "-v=-4", "archive", "-ifc", "-R"},
}

func TestRunCommandArchive(t *testing.T) {
	// testSubdir := "cmd"
	testSubdir := filepath.Join("cmd", "archive")
	// testSubdir := tpath
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t, testSubdir)
	for i, flags := range CommandArchivePresets {
		fn := func(t *testing.T) {
			// testerConfig.TempDestinationPathGeter(testSubdir)
			srcDir := filepath.Join(
				testerConfig.TempDataSource, testSubdir)
			dstDir := filepath.Join(
				testerConfig.TempDataOutput, testSubdir)
			flagss := append(flags, "-sdir="+srcDir)
			flagss = append(flagss, "-odir="+dstDir)
			command := append([]string{"run", "../main.go"}, flagss[1:]...)
			cmdexec := exec.Command("go", command...)
			resultLog, err := cmdexec.CombinedOutput()
			if err != nil {
				t.Error(err)
			}
			flagsJoined := strings.Join(flags[1:], " ")
			fmt.Printf("COMMAND_INPUT:\ngo run main.go %s\n", flagsJoined)
			fmt.Printf("openmedia %s\n", flagsJoined)
			fmt.Printf("COMMAND_OUTPUT_START:\n%s\n", string(resultLog))
		}
		testName := fmt.Sprintf("%d_%s", i, flags[0])
		t.Run(testName, fn)
	}
}
