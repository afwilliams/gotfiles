package cmd

import (
	"fmt"
	"os"

	"github.com/afwilliams/gotfiles/internal/release"
	"github.com/spf13/cobra"
)

var releaseCmd = &cobra.Command{
	Use:     "release",
	Aliases: []string{"r"},
	Short:   "Analyze recent changes and produce a new version",
	Long: `The release command analyzes recent changes in your repository
and helps determine the appropriate version number for the next release.`,
	Run: func(cmd *cobra.Command, args []string) {
		version, err := release.Run()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Suggested version: %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(releaseCmd)
}
