package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var developer string

var rootCmd = &cobra.Command{
	Use:   "tidy",
	Short: "Store your CLI command aliases locally and call them with tidy",
	Long:  `An attempt at a CLI aliaser and my first foray into GoLang`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
