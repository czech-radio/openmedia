#!/bin/bash
SCRIPT_PATH="${BASH_SOURCE[0]:-$0}"
SCRIPT_DIR="${SCRIPT_PATH%/*}"
main_path="${SCRIPT_DIR}/../main.go"

ArchiveExtractRadioDay(){
  local sdir="/mnt/remote/cro/export-avo/Rundowns"
  local odir="/tmp/test/"
  local fdf="2024-01-01"
  local fdt="2024-01-02"
  # local fdt="2024-02-02"
  local frn="Plus"
  local ofname="${frn}_day_$fdf"
  local frdir="/home/jk/CRO/CRO_BASE/openmedia_backup/filters/"
  # local sfn1="eurovolby - zadání.xlsx"
  go run "$main_path" -v=-4 extractArchive \
    -fdf="$fdf" -fdt="$fdt" -frn="$frn" \
    -ofname="$ofname" -sdir="$sdir" -odir="$odir" \
    -frdir="$frdir"
}

"$@"

