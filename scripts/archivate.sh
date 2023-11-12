#!/bin/bash

get_encoding(){
  local filename="$1"
  file --mime-encoding --brief "$filename"
}

convert_to_utf8(){
  local filename="$1"
  iconv -f "$(get_encoding "$filename")" -t UTF-8 $filename
}

"$@"
