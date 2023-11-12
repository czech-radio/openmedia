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
	Short:   "Delete empty fields in xml files",
	Long:    ``,
	Example: ``,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("reduce called", "args", cmd)
	},
}

func init() {
	rootCmd.AddCommand(reduceCmd)
	reduceCmd.Flags().StringP("input", "i", "", "input")
	reduceCmd.Flags().StringP("output", "o", "", "output")
	reduceCmd.MarkFlagRequired("input")
}
