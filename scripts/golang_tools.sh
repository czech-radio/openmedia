#!/bin/bash -euC
# [[https://github.com/golang/tools][GitHub - golang/tools: [mirror] Go Tools]]

declare -a golang_tools=(
  ### 'asmfmt': This will format your assembler code in a similar way that gofmt formats your Go code.
  #'github.com/klauspost/asmfmt/cmd/asmfmt@latest'
  # 
  ### 'dlv': Delve is a debugger for the Go programming language., https://github.com/go-delve/delve
  'github.com/go-delve/delve/cmd/dlv@latest'
  #
  ### 'errcheck': errcheck is a program for checking for unchecked errors in Go code., https://github.com/kisielk/errcheck
  'github.com/kisielk/errcheck@latest'
  #
  ### 'fillstruct': fills a struct literal with default values., https://pkg.go.dev/github.com/davidrjenni/reftools/cmd/fillstruct#section-readme
  'github.com/davidrjenni/reftools/cmd/fillstruct@master'
  #
  ### 'godef': find symbol information in Go source, https://pkg.go.dev/github.com/rogpeppe/godef
  'github.com/rogpeppe/godef@latest'
  #
  ### 'goimports': updates your Go import lines, adding missing ones and removing unreferenced ones., https://pkg.go.dev/golang.org/x/tools/cmd/goimports  
  'golang.org/x/tools/cmd/goimports@master'
  #
  ### 'revive': ~6x faster, stricter, configurable, extensible, and beautiful drop-in replacement for golint,  https://github.com/mgechev/revive
  'github.com/mgechev/revive@latest'
  #
  ### 'gopls': official Go language server developed by the Go team. It provides IDE features to any LSP-compatible editor, https://pkg.go.dev/golang.org/x/tools/gopls
  'golang.org/x/tools/gopls@latest'
  #
  ### 'golangci-lint': is a fast Go linters runner. It runs linters in parallel, uses caching, supports yaml config, has integrations with all major IDE and has dozens of linters included., https://github.com/golangci/golangci-lint
  'github.com/golangci/golangci-lint/cmd/golangci-lint@latest'
  #
  ### 'staticcheck': Staticcheck is a state of the art linter for the Go programming language. Using static analysis, it finds bugs and performance issues, offers simplifications, and enforces style rules, https://staticcheck.dev/
  'honnef.co/go/tools/cmd/staticcheck@latest'
  #
  ### 'gomodifytags', Go tool to modify/update field tags in structs. gomodifytags makes it easy to update, add or delete the tags in a struct field. You can easily add new tags, update existing tags (such as appending a new key, i.e: db, xml, etc..) or remove existing tags. It also allows you to add and remove tag options. It's intended to be used by an editor, but also has modes to run it from the terminal. Read the usage section below for more information., https://github.com/fatih/gomodifytags
  'github.com/fatih/gomodifytags@latest'
  #
  ### 'gorename': The gorename command performs precise type-safe renaming of identifiers in Go source code., https://pkg.go.dev/golang.org/x/tools/cmd/gorename
  'golang.org/x/tools/cmd/gorename@master'
  #
  ### 'gotags': gotags is a ctags-compatible tag generator for Go., https://pkg.go.dev/github.com/lifeibo/gotags
  'github.com/jstemmer/gotags@master'
  #
  ### 'guru': a tool for answering questions about Go source code., https://pkg.go.dev/golang.org/x/tools/cmd/guru
  # 'golang.org/x/tools/cmd/guru@master'
  'golang.org/x/tools/cmd/guru@latest'
  #
  ### 'impl': impl generates method stubs for implementing an interface., https://pkg.go.dev/github.com/josharian/impl
  'github.com/josharian/impl@main'
  #
  ### 'keyify': Keyify turns unkeyed struct literals (T{1, 2, 3}) into keyed ones (T{A: 1, B: 2, C: 3}), https://pkg.go.dev/honnef.co/go/tools/cmd/keyify
  'honnef.co/go/tools/cmd/keyify@master'
  #
  ### 'motion': Motion is a tool that was designed to work with editors. It is providing contextual information for a given offset(option) from a file or directory of files. Editors can use these informations to implement navigation, text editing, etc... that are specific to a Go source code., https://github.com/fatih/motion
  'github.com/fatih/motion@latest'
  #
  ### 'iferr': Generate "if err != nil {" block, https://github.com/koron/iferr 
  'github.com/koron/iferr@master'
  #
  ### 'cobra': Cobra is a library for creating powerful modern CLI applications, https://github.com/spf13/cobra
  "github.com/spf13/cobra-cli@latest"
  #
  ### 'viper': Viper is a complete configuration solution for Go applications including 12-Factor apps. It is designed to work within an application, and can handle all types of configuration needs and formats., https://github.com/spf13/viper 
  "github.com/spf13/viper"
  # 'slog': Package slog provides structured logging, in which log records include a message, a severity level, and various other attributes expressed as key-value pairs., https://pkg.go.dev/golang.org/x/exp/slog
  #
  ### 'gotestsum': 'go test' runner with output optimized for humans, JUnit XML for CI integration, and a summary of the test results.  'https://github.com/gotestyourself/gotestsum'
  'gotest.tools/gotestsum@latest'
)

declare -a golang_packages=(
  ### 'memory'
  # "github.com/pbnjay/memory"
  #
  ### 'unicode': Package unicode provides Unicode encodings such as UTF-16., https://pkg.go.dev/golang.org/x/text/encoding/unicode 
	"golang.org/x/text/encoding/unicode"
  #
  ### 'charset': charset provides common text encodings for HTML documents
  "golang.org/x/net/html/charset"
  #
  ### 'charmap': https://pkg.go.dev/golang.org/x/text/encoding/charmap"
  "golang.org/x/text/encoding/charmap"
)

export GOPATH="${HOME}/go"
export GO111MODULE=on
export PATH="${GOPATH}/bin:$PATH"
for tool in "${golang_tools[@]}"; do
  go install -v -mod=mod "$tool"
done

for package in "${golang_packages[@]}"; do
  go install -v -mod=mod "$package"
done
