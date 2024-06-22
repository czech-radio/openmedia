#!/bin/bash -x

GenerateHelp(){
  go run main.go -H > HELP.md
}

GenerateUsage(){
  go test -v ./cmd/. -run Command_Root > USAGE.md
  go test -v ./cmd/. -run Command_Archive >> USAGE.md
}

GenerateAll(){
  # GenerateHelp
  GenerateUsage
}

"$@"
