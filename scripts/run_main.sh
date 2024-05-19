#!/bin/bash
SCRIPT_PATH="${BASH_SOURCE[0]:-$0}"
SCRIPT_DIR="${SCRIPT_PATH%/*}"
main_path="${SCRIPT_DIR}/../main.go"
LOG_DIR="/tmp/test/"
OUTPUT_DIR="${1:-/tmp/test}"


TouchDone(){
  local odir="$1"
  local result="$2"
  if [[ "$result" == 0 ]] ; then
    touch "${odir}/DONE"
  else
    touch "${odir}/FAIL"
  fi
}

ArchiveExtractKontrolniHodinaProdukce(){
  local exsn="production_all"
  local sdir="/mnt/remote/cro/export-avo/Rundowns"
  local odir="/tmp/test/"
  local verbose="0"
  
  # 2024-01-02T13
  # kontrolni den/hodina posunute o 2 hodiny
  local frn="Plus"
  local fdf="2024-01-02T15"
  local fdt="2024-01-02T17"
  local run_name="kontrolni_hodina"
  local ofname="${frn:-all}-${run_name}-$fdf-$fdt"
  local sdirType="ORIGINAL.zip"
  
  date > "${odir}/run_stat.txt"
  go run "$main_path" -v="$verbose" extractArchive \
    -fdf="$fdf" -fdt="$fdt" -frn="$frn" \
    -ofname="$ofname" -sdir="$sdir" -odir="$odir" \
    -frdir="$frdir" -frfn="$frfn" -exsn="$exsn" -frrec \
    -sdirType="$sdirType"
  date >> "${odir}/run_stat.txt"
}

ArchiveExtractKontrolniWeekProdukce(){
  local exsn="production_all"
  local sdir="/mnt/remote/cro/export-avo/Rundowns"
  local odir="/tmp/test/"
  local verbose="0"
  
  local run_name="kontrolni_hodina"
  local fdf="2024-03-25"
  local fdt="2024-03-31"
  local run_name="kontrolni_W13"
  local ofname="${frn:-all}-${run_name}-$fdf-$fdt"
  date > "${odir}/run_stat.txt"
  go run "$main_path" -v="$verbose" extractArchive \
    -fdf="$fdf" -fdt="$fdt" -frn="$frn" \
    -ofname="$ofname" -sdir="$sdir" -odir="$odir" \
    -frdir="$frdir" -frfn="$frfn" -exsn="$exsn" -frrec
  date >> "${odir}/run_stat.txt"
}



ArchiveExtractContacts(){
  local sdir="/mnt/remote/cro/export-avo/Rundowns"
  local odir="/tmp/test/"
  local verbose="0"
  local exsn="contacts"
  # local exsn="production2"
  # local fdf="2024-01-01"
  # local fdt="2024-01-02"
  
  # Pul roku
  local fdf="2023-10-01"
  local fdt="2024-04-01"
  local run_name="kontrolni_hodina"
  local ofname="${frn:-all}-${run_name}-$fdf-$fdt"
  local frdir="/home/jk/CRO/CRO_BASE/openmedia-filters/"
  local frfn="analýza opozice - zadání.xlsx"
  
  date > "${odir}/run_stat.txt"
  go run "$main_path" -v="$verbose" extractArchive \
    -fdf="$fdf" -fdt="$fdt" -frn="$frn" \
    -ofname="$ofname" -sdir="$sdir" -odir="$odir" \
    -frdir="$frdir" -frfn="$frfn" -exsn="$exsn" -frrec
  TouchDone "${odir}" "$?"
  date >> "${odir}/run_stat.txt"
}

"$@"

