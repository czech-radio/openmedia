package cmd

import "fmt"

type Config_archivate struct {
	SourceDirectory      string `cmd:"source_directory; i; ; directory to be processed"`
	DestinationDirectory string `cmd:"destination_directory; o; ; otput files"`
	InvalidFileContinue  bool   `cmd:"invalid_file_continue; f; 0; continue even though unprocesable file encountered"`
	InvalidFileRename    bool   `cmd:"invalid_file_rename; ifr; true; rename"`
}

func RunArchivate(cfg *Config_archivate) {
	fmt.Printf("%+v\n", cfg)
}
