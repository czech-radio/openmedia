package internal

import (
	"os"
	"testing"
)

// TestMain setup, run tests, and teadrdown (cleanup after tests)
func TestMain(m *testing.M) {
	// Tests setup
	DetectLinuxSytemOrPanic()
	level := os.Getenv("GOLOGLEVEL")
	SetLogLevel(level)
	temp_directory := DirectoryCreateInRam()
	// Run tests
	code := m.Run()
	// Testst teardown
	DirectoryDelete(temp_directory)
	defer os.Exit(code)
}

func Test_DetectLinuxSystemPanic(t *testing.T) {
	DetectLinuxSytemOrPanic()
}

func Test_DirectoryCreateInRam(t *testing.T) {
	directory := DirectoryCreateInRam()
	defer os.RemoveAll(directory)
}
