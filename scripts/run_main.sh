#!/bin/bash -u
SCRIPT_PATH="${BASH_SOURCE[0]:-$0}"
SCRIPT_DIR="${SCRIPT_PATH%/*}"

# COMMAND VARIABLES
MAIN_PATH="${SCRIPT_DIR}/../main.go"
SOURCE_DIR="/mnt/remote/cro/export-avo/Rundowns"

## FILTER FILES
FILTERS_DIR="${SCRIPT_DIR}/../../openmedia-filters"
FILE_VALIDATION="${FILTERS_DIR}/validace_new_ammended.xlsx"
FILE_OPOSITION="${FILTERS_DIR}/analýza opozice - zadání.xlsx"
FILE_EUROVOLBY="${FILTERS_DIR}/eurovolby - zadání.xlsx"

TouchDone(){
  local odir="$1"
  local result="$2"
  if [[ "$result" == 0 ]] ; then
    touch "${odir}/DONE_SUCCESS"
  else
    touch "${odir}/DONE_FAILED"
  fi
}

ArchiveExtractCommand(){
  local OUTPUT_DIR=${OUTPUT_DIR:-/tmp/test}
  local OUTPUT_FILENAME=${OUTPUT_FILENAME:-test}
  declare -a flags=(
    "-v=${VERBOSE:-0}"
    
    # IO specs   
    "-sdir=${SOURCE_DIR:-/mnt/remote/cro/export-avo/Rundowns}"
    "-sdirType=${RUNDOWN_TYPE:-MINIFIED.zip}"
    "-odir=${OUTPUT_DIR}"
    "-ofname=${OUTPUT_FILENAME}"
    
    # filter specs  
    "-frn=${RADIOS:-}"
    "-exsn=${EXTRACTOR:-production_all}"
    "-fdf=${FROM:-2024-01-01}"
    "-fdt=${TO:-2024-01-02}"
    "-valfn=${FILE_VALIDATION}"
    "-frfn=${FILE_FILTER:-}"
  ) 
  local logFile="${OUTPUT_DIR}/${OUTPUT_FILENAME}_run.log"
  go run "$MAIN_PATH" extractArchive "${flags[@]}" 2>&1 | tee "$logFile"
  TouchDone "${OUTPUT_DIR}" "$?"
}

ArchiveExtractKont(){
  local EXTRACTOR="production_contacts"
  local RADIOS="Plus"
  ArchiveExtractCommand
}


ArchiveExtractKontrolniHodinaProdukce(){
  local exsn="production_all"
  local sdir="/mnt/remote/cro/export-avo/Rundowns"
  local odir="/tmp/test/"
  local verbose="0"
  local sdirType="MINIFIED.zip"
  
  local frn="Plus"
  local run_name="kontrolni_hodina"
  local fdf="2024-05-02T13"
  local fdt="2024-05-02T14"
  local ofname="${run_name}-$fdf-$fdt"
  
  local valfn="${SCRIPT_DIR}/../../openmedia-filters/validace_new_ammended.xlsx"
  local frfn="${SCRIPT_DIR}/../../openmedia-filters/analýza opozice - zadání.xlsx"
  
  date > "${odir}/run_stat.txt"
  go run "$MAIN_PATH" -v="$verbose" extractArchive \
    -fdf="$fdf" -fdt="$fdt" -frn="$frn" \
    -ofname="$ofname" -sdir="$sdir" -odir="$odir" \
    -frfn="$frfn" -exsn="$exsn" \
    -sdirType="$sdirType" -valfn="$valfn"
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
  local fdt="2024-04-01"
  local ofname="${run_name}-$fdf-$fdt"
  
  local valfn="${SCRIPT_DIR}/../../openmedia-filters/validace_new_ammended.xlsx"
  local frfn="${SCRIPT_DIR}/../../openmedia-filters/analýza opozice - zadání.xlsx"
  
  date > "${odir}/run_stat.txt"
  go run "$MAIN_PATH" -v="$verbose" extractArchive \
    -fdf="$fdf" -fdt="$fdt" -frn="$frn" \
    -ofname="$ofname" -sdir="$sdir" -odir="$odir" \
    -frfn="$frfn" -exsn="$exsn" -valfn="$valfn"
  TouchDone "${odir}" "$?"
  date >> "${odir}/run_stat.txt"
}

ArchiveExtractKontrolni(){
  ArchiveExtractKontrolniHodinaProdukce
  ArchiveExtractKontrolniTydenProdukce
}

ArchiveExtractDateRange(){
  local exsn="production_all"
  local sdir="/mnt/remote/cro/export-avo/Rundowns"
  local odir="/tmp/test/"
  local verbose="0"
  local sdirType="MINIFIED.zip"
  local valfn="${SCRIPT_DIR}/../../openmedia-filters/validace_new_ammended.xlsx"
  local run_name="range"
  local fdf="2023-12-01"
  local fdt="2024-01-01"
  local ofname="${frn:-all}-${run_name}-$fdf-$fdt"
  
  date > "${odir}/run_stat.txt"
  go run "$MAIN_PATH" -v="$verbose" extractArchive \
    -fdf="$fdf" -fdt="$fdt" -frn="$frn" \
    -ofname="$ofname" -sdir="$sdir" -odir="$odir" \
    -frfn="$frfn" -exsn="$exsn" -valfn="$valfn"
  TouchDone "${odir}" "$?"
  date >> "${odir}/run_stat.txt"
}

ArchiveExtractEurovolby(){
  # local exsn="production_contacts"
  local exsn="production_all"
  local sdir="/mnt/remote/cro/export-avo/Rundowns"
  local odir="/tmp/test/"
  local verbose="0"
  local sdirType="MINIFIED.zip"
  
  # Pul roku
  # local fdf="2024-05-01"
  # local fdt="2024-06-01T01"
  local fdf="2024-05-01"
  local fdt="2024-05-02"
  # local run_name="opozice"
  local run_name="eurovolby"
  local ofname="${run_name}-$fdf-$fdt"
  
  local valfn="$FILE_VALIDATION"
  local frfn="$FILE_EUROVOLBY"
  
  date > "${odir}/run_stat.txt"
  go run "$MAIN_PATH" -v="$verbose" extractArchive \
    -fdf="$fdf" -fdt="$fdt" -frn="$frn" \
    -ofname="$ofname" -sdir="$sdir" -odir="$odir" \
    -frfn="$frfn" -exsn="$exsn"
  TouchDone "${odir}" "$?"
  date >> "${odir}/run_stat.txt"
}


ArchiveExtractValidate(){
  local exsn="production_all"
  local sdir="/mnt/remote/cro/export-avo/Rundowns"
  local odir="/tmp/test/"
  local verbose="0"
  local sdirType="MINIFIED.zip"
  
  # Prosinec
  local frn=""
  local fdf="2023-12-01"
  local fdt="2023-12-03"
  # local fdt="2024-01-01"

  local run_name="prosinec"
  local ofname="${run_name}-$fdf-$fdt"
  local valfn="$FILE_VALIDATION"
  
  date > "${odir}/run_stat.txt"
  go run "$MAIN_PATH" -v="$verbose" extractArchive \
    -fdf="$fdf" -fdt="$fdt" -frn="$frn" \
    -ofname="$ofname" -sdir="$sdir" -odir="$odir" \
    -frfn="$frfn" -valfn="$valfn" \
    -exsn="$exsn"
  TouchDone "${odir}" "$?"
  date >> "${odir}/run_stat.txt"
}

"$@"

