package main

import (
	// "fmt"
	"path/filepath"
	"testing"
)

// TestProcessFolder should cover most of the code
func TestProcessFolder(t *testing.T) {

	input_path := filepath.Join("data", "input")
	output_path := filepath.Join("data", "output")

	err := ProcessFolder(input_path, output_path)
	if err != nil {
		t.Error("run ProcessFloder FAILED!")
	}

}
