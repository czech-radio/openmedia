#!/bin/bash
SCRIPT_PATH="${BASH_SOURCE[0]:-$0}"
SCRIPT_DIR="${SCRIPT_PATH%/*}"
main_path="${SCRIPT_DIR}/../main.go"

ArchiveExtractRadioDay(){
  sdir="/mnt/remote/cro/export-avo/Rundowns"
  odir="/tmp/test/"
  fdf="2024-01-01"
  fdt="2024-01-02"
  frn="Plus"
  ofname="${frn}_day_$fdf"
  go run "$main_path" -v=-4 extractArchive \
    -fdf="$fdf" -fdt="$fdt" -frn="$frn" \
    -ofname="$ofname" -sdir="$sdir" -odir="$odir"
}

"$@"

