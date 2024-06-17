#!/bin/bash -x

GenerateUsage(){
  go test -v ./cmd/. -run CommandRoot > usage.txt
}
