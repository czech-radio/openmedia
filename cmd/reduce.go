/*
Copyright © 2023 Czech Radio
*/
package cmd

import (
	"fmt"
	"github/czech-radio/openmedia-reduce/internal"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

// reduceCmd represents the reduce command
var reduceCmd = &cobra.Command{
	Use:     "reduce",
	Short:   "Reduce filesize of xml rundown files and create archive from them.",
	Long:    ``,
	Example: ``,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		slog.Info("reduce called", "args", args)
	},
	Run: Reduce,
}

func init() {
	rootCmd.AddCommand(reduceCmd)
	lf := reduceCmd.Flags()
	// lf.BoolP("load-env", "", false,
	// `Set commandline options from environment. Useful for example when running as systemd service or as docker container`)
	lf.StringP("input", "i", "",
		"input filename pattern")
	lf.StringP("output", "o", "",
		"output filename pattern")
	lf.BoolP("remove-orginals", "r", false,
		"Delete original files after processing if process finishes without any errors")
	lf.BoolP("invalid-file-continue", "", false,
		"Continue processing files in folder when invalid file found")
	lf.BoolP("output-csv", "", false,
		"convert xml rundown files also to csv")
	// lf.StringP("output-encoding", "E", "utf-8",
	// "Convert file to specified encoding")
	// lf.StringP("archive-compression", "", "default",
	// `Specify the algorithm which will be used for archive compression`)
	// lf.StringP("rename-invalid-files-pattern", "R", "",
	// `rename invalid files:
	// files which do not pass any xml-validate tests`)
	err := reduceCmd.MarkFlagRequired("input")
	if err != nil {
		slog.Error(err.Error())
	}
}

func Reduce(cmd *cobra.Command, args []string) {
	fl := cmd.Flags()
	input, err := fl.GetString("input")
	if err != nil {
		slog.Error(err.Error())
		return
	}
	output, err := fl.GetString("output")
	if err != nil {
		slog.Error(err.Error())
		return
	}
	dry_run, err := fl.GetBool("dry-run")
	if err != nil {
		slog.Error(err.Error())
		return
	}
	opts := internal.ProcessOptions{
		SourceDirectory:        input,
		DestinationDirectory:   output,
		InputEncoding:          "",
		OutputEncoding:         "",
		ValidateWithDefaultXSD: false,
		ValidateWithXSD:        "",
		ValidatePre:            false,
		ValidatePost:           false,
		CompressionType:        "zip",
		InvalidFileRename:      false,
		// InvalidFileContinue:    false,
		InvalidFileContinue: true,
	}
	if dry_run {
		tmpName := fmt.Sprintf("openmedia_archive_%d", os.Getpid())
		tmpPath := internal.DirectoryCreateTemporaryOrPanic(tmpName)
		opts.DestinationDirectory = tmpPath
	}
	process := internal.Process{Options: opts}
	//1. Check if destination/source directory is not empty
	//2. check if file exists there do no overwrite
	slog.Info("process", "options", opts)
	err = process.Folder()
	if err != nil {
		slog.Error(err.Error())
		return
	}
}
