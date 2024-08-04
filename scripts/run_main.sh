#!/bin/bash -u
SCRIPT_PATH="${BASH_SOURCE[0]:-$0}"
SCRIPT_DIR="${SCRIPT_PATH%/*}"

# COMMAND VARIABLES
MAIN_PATH="${SCRIPT_DIR}/../main.go"
SOURCE_DIR="/mnt/remote/cro/export-avo/Rundowns"

## FILTER FILES
FILTERS_DIR="${SCRIPT_DIR}/../../openmedia-filters"

### filter validation
FILE_VALIDATION="${FILTERS_DIR}/validace_new_ammended.xlsx"
### filter oposition
FILE_OPOSITION="${FILTERS_DIR}/filtr_opozice_2024-04-01_2024-05-31_v2.xlsx"
### filter eurovolby
FILE_EUROVOLBY="${FILTERS_DIR}/filtr_eurovolby_v1.xlsx"

TouchDone(){
  local odir="$1"
  local ofile="$2"
  local result="$3"
  local result_file="${odir}/${ofile}"
  if [[ "$result" == 0 ]] ; then
    touch "${result_file}_success.log"
  else
    touch "${result_file}_failed.log"
  fi
}
  
AddFlag(){
  flagName="$1"
  flagVar="${2:-}"
  if [[ -n "${flagVar}" ]]; then
    # echo "$flagName=$flagVar"
    flags+=("$flagName=$flagVar")
  fi
}

AddBoolFlag(){
  flagName="$1"
  flags+=("$flagName")
}

ArchiveExtractCommand(){
  local OUTPUT_DIR=${OUTPUT_DIR:-/tmp/test}
  local OUTPUT_FILENAME=${OUTPUT_FILENAME:-test}
  declare -a flags=(
    # IO specs   
    "-sdir=${SOURCE_DIR:-/mnt/remote/cro/export-avo/Rundowns}"
    "-sdirType=${RUNDOWN_TYPE:-MINIFIED.zip}"
    "-odir=${OUTPUT_DIR}"
    "-ofname=${OUTPUT_FILENAME}"
    # filter specs  
    "-frns=${RADIOS:-}"
    "-excode=${EXTRACTOR:-production_all}"
    "-valfn=${FILE_VALIDATION}"
    "-frfn=${FILE_FILTER:-${FILE_OPOSITION}}"
  )
  AddBoolFlag "-arn"
  AddFlag -fdf "${FROM:-}"
  AddFlag -fdt "${TO:-}"

  local logFile="${OUTPUT_DIR}/${OUTPUT_FILENAME}_run.log"
  go run "$MAIN_PATH" extractArchive "${flags[@]}" 2>&1 | tee "$logFile"
  # go run "$MAIN_PATH" -v="${VERBOSE:-0}" extractArchive "${flags[@]}" &> "$logFile"
  # tail -f "$logFile"
  TouchDone "${OUTPUT_DIR}" "${OUTPUT_FILENAME}" "$?"
}

ArchiveExtractTestTest(){
  local EXTRACTOR="production_contacts"
  local RADIOS=""
  local VERBOSE="0"
  ArchiveExtractCommand
}

ArchiveExtractConntrolProductionHour(){
  local EXTRACTOR="production_all"
  local FROM="2024-05-02T13"
  local TO="2024-05-02T14"
  local RADIOS="Plus"
  local OUTPUT_FILENAME="kontrolni_hodina"
  ArchiveExtractCommand
}


ArchiveExtractConntrolProductionWeek14(){
  local EXTRACTOR="production_all"
  local FROM="2024-04-01"
  local TO="2024-04-07"
  local OUTPUT_FILENAME="kontrolni_tyden_W14"
  ArchiveExtractCommand
}

ArchiveExtractConntrolProductionWeek(){
  local EXTRACTOR="production_all"
  local FROM="2024-03-25"
  local TO="2024-04-01"
  local OUTPUT_FILENAME="kontrolni_tyden_W13"
  ArchiveExtractCommand
}

ArchiveExtractControl(){
  ArchiveExtractConntrolProductionHour
  ArchiveExtractConntrolProductionWeek
}

ArchiveExtractConntrolProductionDecember(){
  local EXTRACTOR="production_all"
  local FROM="2023-12-01"
  local TO="2024-01-01"
  local OUTPUT_FILENAME="prosinec"
  ArchiveExtractCommand
}

ArchiveExtractControlValidation(){
  ArchiveExtractConntrolProductionHour
  ArchiveExtractConntrolProductionWeek
  ArchiveExtractConntrolProductionDecember
}

ArchiveExtractControlAddHocErr(){
  local EXTRACTOR="production_all"
  local FROM="2024-06-14"
  local TO="2024-06-20"
  local OUTPUT_FILENAME="baddhoc_err"
  ArchiveExtractCommand
}

ArchiveExtractRange(){
  local EXTRACTOR="production_all"
  local FROM="2023-12-01"
  local TO="2024-01-01"
  local OUTPUT_FILENAME="range"
  ArchiveExtractCommand
}

ArchiveExtractEurovolby(){
  local EXTRACTOR="production_contacts"
  local FROM="2024-06-01"
  local TO="2024-07-01"
  local OUTPUT_FILENAME="eurovolby"
  local FILE_FILTER="$FILE_EUROVOLBY"
  ArchiveExtractCommand
}

ArchiveExtractOpozice(){
  local EXTRACTOR="production_contacts"
  local FROM="2024-01-01"
  local TO="2024-02-01"
  # local RADIOS="Plus"
  local OUTPUT_FILENAME="opozice"
  local FILE_FILTER="$FILE_OPOSITION"
  ArchiveExtractCommand
}

ArchiveExtractLastWeek(){
  local EXTRACTOR="production_all"
  local OUTPUT_FILENAME="oneweek"
  ArchiveExtractCommand
}

"$@"
