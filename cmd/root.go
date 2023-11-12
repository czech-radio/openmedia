/*
Copyright © 2023 Czech Radio

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"github/czech-radio/openmedia-reduce/internal"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "openmedia-reduce",
	Version: "v0.0.1",
	Short:   "Archivates rundown xml files",
	Long:    `Program operates on Rundown files. It strips down unnecessary or empty fields and produces light version of an original file. Then it can create packed archive to furher reduce size of files`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		verboseFlag, _ := cmd.Flags().GetString("verbose")
		internal.SetLogLevel(verboseFlag)
		slog.Debug("verbose level set", "verbose", verboseFlag)
	},
	// Bare application action:
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("root called")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Persistent flags (global for whole app/commands)
	rootCmd.PersistentFlags().BoolP("dry-run", "n", false, "perform a trial run with no changes")
	rootCmd.PersistentFlags().IntP("verbose", "v", 0, "Set verbosity level.")
	// rootCmd.PersistentFlags().IntVar

	// Config flags
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.openmedia-reduce.yaml)")

	// Local flags
	// rootCmd.Flags().BoolP("version", "V", false, "print version of program")
}
