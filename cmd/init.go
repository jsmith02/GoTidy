/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"i", "-i"},
	Short:   "Configures your local directories for app usage",
	Long: `Adds a directory called .tidy and a file to store aliases in called "up".
	It looks explicitly for the $HOME directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		tidypath := ".tidy"
		homedir := os.Getenv(("HOME"))
		up := "up"
		if _, err := os.Stat(homedir + "/" + tidypath); os.IsNotExist(err) {
			// .tidy does not already exist.
			fmt.Println("Creating your .tidy/up...")
			os.MkdirAll(homedir+"/"+tidypath, 0777)
			os.Create(homedir + "/" + tidypath + "/" + up)
			fmt.Println(".tidy/up configured.")

		} else {
			if _, err := os.Stat(homedir + "/" + tidypath + "/" + up); os.IsNotExist(err) {
				fmt.Println("Missing up... creating for you now...")
				os.Create(homedir + "/" + tidypath + "/" + up)
				fmt.Println(".tidy/up configured.")
			} else {
				fmt.Println(".tidy/up already configured.")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
