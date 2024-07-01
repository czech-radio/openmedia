#!/bin/bash -ux

SCRIPT_PATH="${BASH_SOURCE[0]:-$0}"
SCRIPT_DIR="${SCRIPT_PATH%/*}"
PROJECT_DIR="${SCRIPT_DIR%/*}"

Go_test_all(){
  # go vet -v ./...
  go test -vet=all -v ./...
}

Go_build(){
  go build main.go
}

Git_push(){
  All
  read -pr 'Continue push to git? (y): ' cont
  if [[ $cont == 'y' ]]; then
    echo "yes"
  else
    exit 1
  fi
  git push origin
}


ChangeVersionInFile(){
  local fileName="$1"
  local string_before="$2"
  local string_after="$3"
  local new_version="$4"
  local file_path="${PROJECT_DIR}/${fileName}"
  # sed "s,\(${string_before}\).\+\(${string_after}\),\1${new_version}\2,g" "${file_path}"
  sed -i "s,\(${string_before}\).\+\(${string_after}\),\1${new_version}\2,g" "${file_path}"
}

ChangeVersionInFiles(){
  local new_version="$1"
  # README.md
  ChangeVersionInFile "README.md" "https://img.shields.io/badge/version-" "-blue.svg" "$new_version"
  # cmd/root.go
  local new_version="\"$1\""
  ChangeVersionInFile "cmd/root.go" "Version:\s\+" "\," "${new_version}"
}

Git_update_tag(){
  tag_name="${1}"
  ### check github action syntax
  if ! actionlint ; then
    return 1
  fi

  ### change version
  ChangeVersionInFiles tag_name

  ### push tags
  git add .
  git co -m "new version/tag: $tag_name" 
  git tag -d "$tag_name"
  git push --delete origin "$tag_name"
  sleep 1
  git tag -a "$tag_name" -m "$tag_name"
  git push origin --tags
}

"$@"
