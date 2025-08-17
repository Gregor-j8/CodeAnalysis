package cmd

import (
	"fmt"

	worker "github.com/Gregor-j8/CodeAnalysis/Worker"
	"github.com/spf13/cobra"
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze [path]",
	Short: "Analyze a project or directory",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		worker.Scan(path)
		fmt.Println("Analyzing project at:", path)
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
}
