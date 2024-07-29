#!/bin/bash -x

GenerateHelp(){
  go run main.go -H > ./docs/HELP.md
}

GenerateUsage(){
  #NOTE: do not forget update command list bellow
  truncate -s 0 ./docs/USAGE.md
  {
    go test -v ./cmd/. -run Command_Root;
    go test -v ./cmd/. -run Command_Archive;
    go test -v ./cmd/. -run Command_ExtractArchive;
    go test -v ./cmd/. -run Command_ExtractFile;
  } >> ./docs/USAGE.md
  # sed -i "/^=== RUN.*$/d" ./docs/USAGE.md
  # sed -i "/=== RUN.*$/d" ./docs/USAGE.md
  sed -i "s/\(.*\)=== RUN.*$/\1/g" ./docs/USAGE.md
}

GenerateAll(){
  GenerateHelp
  GenerateUsage
}

"$@"
