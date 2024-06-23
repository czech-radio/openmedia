#!/bin/bash -x

GenerateHelp(){
  go run main.go -H > HELP.md
}

GenerateUsage(){
  truncate -s 0 USAGE.md
  go test -v ./cmd/. -run Command_Root >> USAGE.md
  go test -v ./cmd/. -run Command_Archive >> USAGE.md
  sed -i "/^=== RUN.*$/d" USAGE.md
}

GenerateAll(){
  # GenerateHelp
  GenerateUsage
}

"$@"
