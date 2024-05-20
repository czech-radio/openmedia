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
  # local exsn="production_contacts"
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
  local sdirType="MINIFIED.zip"
  local frfn="${SCRIPT_DIR}/../../openmedia-filters/analýza opozice - zadání.xlsx"
  
  date > "${odir}/run_stat.txt"
  go run "$main_path" -v="$verbose" extractArchive \
    -fdf="$fdf" -fdt="$fdt" -frn="$frn" \
    -ofname="$ofname" -sdir="$sdir" -odir="$odir" \
    -frfn="$frfn" -exsn="$exsn" \
    -sdirType="$sdirType"
  date >> "${odir}/run_stat.txt"
}

ArchiveExtractKontrolniTydenProdukce(){
  local exsn="production_all"
  local sdir="/mnt/remote/cro/export-avo/Rundowns"
  local odir="/tmp/test/"
  local verbose="0"
  local sdirType="MINIFIED.zip"
  
  local run_name="kontrolni_tyden_W13"
  local fdf="2024-03-25"
  local fdt="2024-03-31"
  local ofname="${frn:-all}-${run_name}-$fdf-$fdt"
  date > "${odir}/run_stat.txt"
  go run "$main_path" -v="$verbose" extractArchive \
    -fdf="$fdf" -fdt="$fdt" -frn="$frn" \
    -ofname="$ofname" -sdir="$sdir" -odir="$odir" \
    -frfn="$frfn" -exsn="$exsn"
  TouchDone "${odir}" "$?"
  date >> "${odir}/run_stat.txt"
}



ArchiveExtractContacts(){
  local sdir="/mnt/remote/cro/export-avo/Rundowns"
  local odir="/tmp/test/"
  local verbose="0"
  local sdirType="MINIFIED.zip"
  local exsn="production_contacts"
  # local fdf="2024-01-01"
  # local fdt="2024-01-02"
  
  # Pul roku
  local fdf="2023-10-01"
  local fdt="2024-04-01"
  local run_name="opozice"
  local ofname="${frn:-all}-${run_name}-$fdf-$fdt"
  local frdir="/home/jk/CRO/CRO_BASE/openmedia-filters/"
  local frfn="${SCRIPT_DIR}/../../openmedia-filters/analýza opozice - zadání.xlsx"
  
  date > "${odir}/run_stat.txt"
  go run "$main_path" -v="$verbose" extractArchive \
    -fdf="$fdf" -fdt="$fdt" -frn="$frn" \
    -ofname="$ofname" -sdir="$sdir" -odir="$odir" \
    -frfn="$frfn" -exsn="$exsn"
  TouchDone "${odir}" "$?"
  date >> "${odir}/run_stat.txt"
}

"$@"

