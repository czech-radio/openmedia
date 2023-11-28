#!/bin/bash -u
SCRIPT_PATH="${BASH_SOURCE[0]:-$0}"
SCRIPT_DIR="${SCRIPT_PATH%/*}"
TEST_FILE_GOOD="${SCRIPT_DIR}/../test/testdata/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431.xml"
TEST_FILE_AMMEND="${SCRIPT_DIR}/../test/testdata/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_ammend_utf.xml"
TEST_FILE_BAD="${SCRIPT_DIR}/../test/testdata/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml"
# CURRENT_DIR="$PWD"

xml_validate_xml_xsd_dir(){
  # SIGHUP SIGQUIT SIGTERM EXIT ERR
  ### BIND VARIABLES
  local TEST_DIR="${SCRIPT_DIR}/../test/testdata/"
  local SCHEMA_DIR="${TEST_DIR}/rundowns_schemas/"
  local SCHEMA_FILE="OM_LV7_schema.xsd"
  local XML_DIR_VALID="${TEST_DIR}/rundowns_additional/"
  local XML_DIR_INVALID="${TEST_DIR}/rundowns_invalid/"
  local xmllint_opts
  ### XMLLINT OPTS
  cd "$SCHEMA_DIR" || exit 1
  declare -a xmllint_opts=(
  '--quiet'
  '--nonet'
  --schema "$SCHEMA_FILE"
  )
  ### TEST VALID FILES 
  local valid_count=0
  local invalid_count=0
  #### find files in directory
  file=""
  declare -a xml_files_valid=(
    $( find "$XML_DIR_VALID" -type f -exec realpath {} \; | rg ".xml" | sort)
  )
  echo "Testing valid files: ${#xml_files_valid[@]}"
  for file in "${xml_files_valid[@]}"; do
    # echo "VALIDATING: $file"
    xmllint "${xmllint_opts[@]}" "$file" >/dev/null
    result="$?"
    if [[ $result != 0 ]] ; then
      echo Invalid file: "$file" >/dev/stderr
      ((invalid_count++))
      # return 1
    else
      ((valid_count++))
    fi
    echo "val: $valid_count, inval: $invalid_count"
  done

  ### TEST INVALID FILES
  #### find files in directory
  file=""
  declare -a xml_files_invalid=(
    $( find "$XML_DIR_INVALID" -type f -exec realpath {} \; | rg ".xml")
  )
  echo "Testing invalid files: ${#xml_files_invalid[@]}" >/dev/stderr
  for file in "${xml_files_invalid[@]}"; do
    # echo "$file"
    xmllint "${xmllint_opts[@]}" "$file" &>/dev/null
    result="$?"
    if [[ $result == 0 ]] ; then
      echo "File should not be validated: $file" > /dev/stderr
      # return 2
    fi
  done
}

get_encoding(){
  local filename="$1"
  file --mime-encoding --brief "$filename"
}

convert_to_utf8(){
  local filename="$1"
  iconv -f "$(get_encoding "$filename")" -t UTF-8 "$filename"
}

xml_validate(){
  local xmlfile="$1"
  xmllint --format "${xmlfile}" 
}

test_xml_validate(){
  local test_files
  declare -a test_files=(
    "$TEST_FILE_GOOD"
    "$TEST_FILE_BAD"
  )
  for file in "${test_files[@]}"; do
    xml_validate <(cat "$file") >/dev/null
    local result="$?"
    if [[ $result == 0 ]]; then
      echo "${file} is valid"
    else
      echo "${file} is invalid"
    fi
  done
}

xml_filter_out_empty_fields(){
  local xmlfile="$1"
  # xml_validate <(convert_to_utf8 "$xmlfile")
  convert_to_utf8 "$xmlfile" | grep -v "IsEmpty = \"yes\""
}

test_xml_filter2(){
  local xmlfile="$TEST_FILE_GOOD"

  convert_to_utf8 "$xmlfile" | grep --count -v "IsEmpty = \"yes\""
  convert_to_utf8 "$xmlfile" | awk '!/IsEmpty = "yes"/' | wc -l
  convert_to_utf8 "$xmlfile" | awk '!/IsEmpty = "yes"/'
# 10298
}

test_xml_filter_out_empty_fields(){
  local xmlfile="$TEST_FILE_GOOD"
  local count_before 
  count_before=$(convert_to_utf8 "$xmlfile" | wc -l)
  local count_after
  count_after=$(xml_filter_out_empty_fields "$xmlfile"  | wc -l)
  echo "Lines before: $count_before, after: $count_after"
}

archivate_file(){
  local srcfile="$1"
  local dstfile="$2"
  7z -mx=9 a "$dstfile" "$srcfile"
}

archivate_stdin(){
  local srcfile="$1"
  local dstfile="$2"
  local srcfile_name
  srcfile_name="$(basename "$srcfile")"
  7z -mx=9 a "$dstfile" "-si$srcfile_name" < /dev/stdin 
}

test_archivate(){
  local srcfile="$TEST_FILE_GOOD"
  filename="$(basename "$srcfile")"
  local dstfile="/dev/shm/${filename}.7z"
  archivate "$srcfile" "$dstfile" 
}

test_archivate_stdin(){
  local srcfile="$TEST_FILE_GOOD"
  filename="$(basename "$srcfile")"
  local dstfile="/dev/shm/${filename}.7z"
  local dstdir
  dstdir="$(dirname "$dstfile")"
  archivate_stdin "$srcfile" "$dstfile" <<< "$(xml_filter_out_empty_fields "$srcfile")"
  7z x "$dstfile" "-o${dstdir}"

}

test_conv(){
  local srcfile="$TEST_FILE_AMMEND"
  iconv -f "$(get_encoding "$srcfile")" -t UTF-8 "$srcfile" | xmllint --format - | xmllint --valid -
}

"$@"
