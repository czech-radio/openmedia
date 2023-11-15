#!/bin/bash
SCRIPT_PATH="${BASH_SOURCE[0]:-$0}"
SCRIPT_DIR="${SCRIPT_PATH%/*}"
# TEST_CMD="go test -v "
# TEST_CMD="${HOME}/go/bin/gotestsum --format testname "

Go_test_with_debug(){
  local path
  path="$1"
  path="${SCRIPT_DIR}/../${path}/..."
  local test_pattern
  test_pattern="${2:-''}"
  # GOLOGLEVEL=-4 go test -v "$path" -run "$test_pattern"
  GOLOGLEVEL=-4 go test -v "$path" -run "$test_pattern"
}

Go_test_normal(){
  local path
  path="$1"
  local test_pattern
  test_pattern="${2:-''}"
  path="${SCRIPT_DIR}/../${path}/..."
  go test "$path" -run "$test_pattern"
}

"$@"
