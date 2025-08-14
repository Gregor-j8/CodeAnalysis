package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "code-analyzer",
	Short: "Analyze code complexity and maintainability",
	Long:  "Scan a project and report high complexity functions, large files, and maintainability metrics.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
