package main

import (
	"path/filepath"
	"testing"
)

// TestProcessFolder should cover most of the code
func TestProcessFolder(t *testing.T) {

	input_path := filepath.Join("test", "input")
	output_path := filepath.Join("test", "output")

	err := ProcessFolder(input_path, output_path)
	if err != nil {
		t.Error(err.Error())
	}

}
