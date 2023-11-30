#!/bin/bash
SCRIPT_PATH="${BASH_SOURCE[0]:-$0}"
SCRIPT_DIR="${SCRIPT_PATH%/*}"
# TEST_CMD="go test -v "
# TEST_CMD="${HOME}/go/bin/gotestsum --format testname "

Go_test_debug(){
  local path
  local path="$1"
  local path="${SCRIPT_DIR}/../${path}/..."
  local test_pattern
  test_pattern="${2:-''}"
  GOLOGLEVEL=-4 go test -v "$path" -run "$test_pattern"
}

Go_bench(){
  local path
  local path="$1"
  local path="${SCRIPT_DIR}/../${path}/..."
  local test_pattern
  test_pattern="${2:-''}"

	local test_opts=(
		"$path"
		-bench .
		-run=^$
		-benchmem
    -cpu 1,2,4
		# -benchtime 20s
    "$test_pattern"
	)
	# go test "$path" -bench . -run=^$ -benchmem -benchtime 5s
	echo go test "${test_opts[@]}"
	go test "${test_opts[@]}"
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
