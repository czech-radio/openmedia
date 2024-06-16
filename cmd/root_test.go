package cmd

import (
	"fmt"
	"os/exec"
	"strconv"
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

var CommandRootCommandPresets = [][]string{
	{"-V"},
	{"-dc"},
	{"-V", "-dc"},
	{"-logt=plain", "-v=1"},
	{"-logt=plain", "-v=0"},
	{"-logt=plain", "-v=-2"},
	{"-logt=plain", "-v=-4"},
	{"-logt=json", "-v=-4"},
}

func TestRunCommandRoot(t *testing.T) {
	testSubdir := "cmd"
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t, testSubdir)
	for i, pars := range CommandRootCommandPresets {
		fn := func(t *testing.T) {
			pars = append([]string{"run", "../main.go"}, pars...)
			cmd := exec.Command("go", pars...)
			res, err := cmd.CombinedOutput()
			if err != nil {
				t.Error(err)
			}
			fmt.Printf("COMMAND_INPUT:\n%v\n", pars)
			fmt.Printf("COMMAND_OUTPUT_START:\n%s\n", string(res))
			fmt.Printf("COMMAND_OUTPUT_END\n\n")
		}
		t.Run(strconv.Itoa(i), fn)
	}
}
