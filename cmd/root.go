package cmd

import (
	"fmt"
	"os"

	humanize "github.com/dustin/go-humanize"
	"github.com/pablo-ruth/gofat/profiler"
	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "gofat",
	Short: "Binary size profiler for golang ",
	RunE: func(cmd *cobra.Command, args []string) error {
		dependencies, err := profiler.Profile()
		if err != nil {
			return err
		}
		fmt.Println("")

		for _, dependency := range dependencies {
			fmt.Printf("%s %s\n", humanize.Bytes(dependency.Size), dependency.Name)
		}

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
