package cmd

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"

	"github.com/triopium/go_utils/pkg/helper"
)

var testerConfig = helper.TesterConfig{
	TempDirName:    "openmedia",
	TestDataSource: "../test/testdata",
}

func TestMain(m *testing.M) {
	testerConfig.TesterMain(m)
}

var CommandRootPresets = [][]string{
	{"help", "-h"},
	{"print version", "-V"},
	{"print config", "-dc"},
	{"print version and config",
		"-V", "-dc"},
	{"test log [err]",
		"-logt=plain", "-logts", "-v=6"},
	{"test log [err,warn]",
		"-logt=plain", "-logts", "-v=4"},
	{"test log [err,warn,info]",
		"-logt=plain", "-logts", "-v=0"},
	{"terst log [err,warn,info,debug]",
		"-logt=plain", "-logts", "-v=-4"},
	{"test log json", "-logt=json", "-logts", "-v=-4"},
}

func TestRunCommandRoot(t *testing.T) {
	testSubdir := "cmd"
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t, testSubdir)
	for i, flags := range CommandRootPresets {
		fn := func(t *testing.T) {
			command := append([]string{"run", "../main.go"}, flags[1:]...)
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
