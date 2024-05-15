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
  # local odir="/mnt/remote/cro/R/GŘ/Strategický rozvoj/Analytická sekce/Analýzy/Produkce/Tests/2024_05_09_eurovolby/filtr"
  
  # local verbose="-4"
  local verbose="0"
  local fdf="2024-01-01"
  local fdt="2024-01-02"
  
  # local fdt="2024-01-02"
  # local fdf="2024-04-01"
  # local fdt="2024-05-01"
  # local fdt="2024-01-11"
  # local run_name="month_april"
  local run_name="day"
  local ofname="${frn:-all}_${run_name}_$fdf"
  local frdir="/home/jk/CRO/CRO_BASE/openmedia_backup/filters/"
  # local frn="Plus"
  local frn=""
  
  date > "${odir}/run_stat.txt"
  # old
  # local exsn="production1"
  # go run "$main_path" -v="$verbose" extractArchive \
  #   -fdf="$fdf" -fdt="$fdt" -frn="$frn" \
  #   -ofname="$ofname" -sdir="$sdir" -odir="$odir" \
  #   -frdir="$frdir" -exsn="$exsn"
  
  # new 
  # local exsn="production2"
  # go run "$main_path" -v="$verbose" extractArchive \
  #   -fdf="$fdf" -fdt="$fdt" -frn="$frn" \
  #   -ofname="$ofname" -sdir="$sdir" -odir="$odir" \
  #   -frdir="$frdir" -exsn="$exsn"

  # new filtered
  local exsn="production2"
  local ofname="${frn:-all}_${run_name}_filtered_$fdf"
  go run "$main_path" -v="$verbose" extractArchive \
    -fdf="$fdf" -fdt="$fdt" -frn="$frn" \
    -ofname="$ofname" -sdir="$sdir" -odir="$odir" \
    -frdir="$frdir" -exsn="$exsn" -frrec -frn="$frn"
  
  if [[ $? == 0 ]] ; then
    touch "${odir}/DONE"
  fi
  date >> "${odir}/run_stat.txt"
}

"$@"

