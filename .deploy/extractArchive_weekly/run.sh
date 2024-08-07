#!/bin/bash
SCRIPT_PATH="${BASH_SOURCE[0]:-$0}"
SCRIPT_DIR="${SCRIPT_PATH%/*}"
FILTERS_DIR="${SCRIPT_DIR}/../../../openmedia-filters"

OUTPUT_DIR='/mnt/remote/cro/R/GŘ/Strategický rozvoj/Analytická sekce/_Exports/test/'
OUTPUT_DIR='/tmp/test/'
SOURCE_DIR='/mnt/remote/cro/export-avo/Rundowns/'
FILE_VALIDATION="${FILTERS_DIR}/validace_new_ammended.xlsx"

declare -a cmd=(
  "-sdir=$SOURCE_DIR"
  "-odir=$OUTPUT_DIR"
  "-valfn=${FILE_VALIDATION}"
  "-arn"
)

go run ../../main.go extractArchive ${cmd[*]}
# ./openmedia ${cmd[*]}

