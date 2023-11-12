### Go test
Go_test_all(){
  # go vet -v ./...
  go test -vet=all -v ./...
}

### Go build
Go_build(){
  go build main.go
}


### GIT
Git_push(){
  All
  read -p 'Continue push to git? (y): ' cont
  if [[ $cont == 'y' ]]; then
    echo "yes"
  else
    exit 1
  fi
  git push origin
}

"$@"
