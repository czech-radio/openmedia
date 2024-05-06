#!/bin/bash
SCRIPT_PATH="${BASH_SOURCE[0]:-$0}"
SCRIPT_DIR="${SCRIPT_PATH%/*}"
main_path="${SCRIPT_DIR}/../main.go"
LOG_DIR="/tmp/test/"
OUTPUT_DIR="${1:-/tmp/test}"

ArchiveExtractRadioDay(){
  local sdir="/mnt/remote/cro/export-avo/Rundowns"
  local odir="/tmp/test/"
  local fdf="2024-01-01"
  local fdt="2024-01-02"
  local exsn="production1"
  # local fdt="2024-02-02"
  local frn="Plus"
  local ofname="${frn}_day_$fdf"
  local frdir="/home/jk/CRO/CRO_BASE/openmedia_backup/filters/"
  
  go run "$main_path" -v=-4 extractArchive \
    -fdf="$fdf" -fdt="$fdt" -frn="$frn" \
    -ofname="$ofname" -sdir="$sdir" -odir="$odir" \
    -frdir="$frdir" -exsn="$exsn"
}


ArchiveExtractWeek(){
  local sdir="/mnt/remote/cro/export-avo/Rundowns"
  local odir="/tmp/test/"
  
  # local verbose="-4"
  local verbose="0"
  local fdf="2024-01-01"
  # local fdt="2024-01-02"
  local fdt="2024-04-01"
  # local fdt="2024-01-11"
  local ofname="${frn:-all}_3months_$fdf"
  local frdir="/home/jk/CRO/CRO_BASE/openmedia_backup/filters/"
  
  date > "${odir}/run_stat.txt"
  
  local exsn="production1"
  go run "$main_path" -v="$verbose" extractArchive \
    -fdf="$fdf" -fdt="$fdt" -frn="$frn" \
    -ofname="$ofname" -sdir="$sdir" -odir="$odir" \
    -frdir="$frdir" -exsn="$exsn"
  
  local exsn="production2"
  go run "$main_path" -v="$verbose" extractArchive \
    -fdf="$fdf" -fdt="$fdt" -frn="$frn" \
    -ofname="$ofname" -sdir="$sdir" -odir="$odir" \
    -frdir="$frdir" -exsn="$exsn"
  
  if [[ $? == 0 ]] ; then
    touch "${odir}/DONE"
  fi
  date >> "${odir}/run_stat.txt"
}

"$@"

