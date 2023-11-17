/*
Copyright © 2023 Czech Radio
*/
package cmd

import (
	"log/slog"

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

	lf.BoolP("load-env", "", false,
		`Set commandline options from environment.
Useful for example when running as systemd service or as docker container`)
	lf.StringP("input", "i", "",
		"input filename pattern")
	lf.StringP("output", "o", "",
		"output filename pattern")
	lf.StringP("output-encoding", "E", "utf-8",
		"Convert file to specified encoding")
	lf.BoolP("xml-validate-lines-pre", "", true,
		"Validate xml on line by line basis before processing.")
	lf.BoolP("xml-validate-file-pre", "", true,
		"Validate xml as whole file before processing.")
	lf.BoolP("xml-validate-lines-post", "", true,
		"Validate xml on line by line basis after processing.")
	lf.BoolP("xml-validate-file-post", "", true,
		"Validate xml as whole file after processing.")
	lf.StringP("archive-compression", "", "default",
		`Specify the algorithm which will be used for archive compression`)
	lf.StringP("rename-invalid-files-pattern", "R", "",
		`rename invalid files:
		files which do not pass any xml-validate tests`)
	reduceCmd.MarkFlagRequired("input")
}

func Reduce(cmd *cobra.Command, args []string) {
	load_env, err := cmd.Flags().GetBool("load-env")
	if err != nil {
		panic(err)
	}
	if load_env {
		slog.Info("loading environment variables")
	}
}
