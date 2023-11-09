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
	Use:   "reduce",
	Short: "Delete empty fields in xml files",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("reduce called")
	},
}

func init() {
	rootCmd.AddCommand(reduceCmd)
}
