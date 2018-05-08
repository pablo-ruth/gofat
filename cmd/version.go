package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionString = "undefined"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display gofat version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(versionString)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
