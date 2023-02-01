package main


import (
 // "fmt"
  "testing"
  "path/filepath"
)


func TestProcessFolder(t *testing.T) {

  input_path := filepath.Join("data","input")
  output_path := filepath.Join("data","output")

  err := ProcessFolder(input_path, output_path)
  if err != nil {
    t.Error("run ProcessFloder FAILED!")
  }

}
