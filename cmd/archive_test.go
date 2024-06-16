package cmd

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestRunCommandArchive(t *testing.T) {
	testSubdir := "cmd"
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t, testSubdir)
	mainOpts := []string{
		"run", "../main.go", "-logt=plain", "-v=-4"}
	subOpts := []string{"archive"}
	command := append(mainOpts, subOpts...)
	cmd := exec.Command("go", command...)
	res, err := cmd.CombinedOutput()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("\nCOMMAND_OUTPUT_START:\n%s\nCOMMAND_OUTPUT_END\n\n", string(res))
}
