// Copyright © 2020 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/dixonwille/wmenu"
	"github.com/spf13/cobra"
)

// eCmd represents the e command
var eCmd = &cobra.Command{
	Use:   "e",
	Short: "Edit rules through a CLI menu",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Function to edit configs
		homedir := os.Getenv(("HOME"))
		upFile := homedir + "/.tidy/up"
		menu := wmenu.NewMenu("Which config do you want to edit?")
		menu.Action(func(opts []wmenu.Opt) error { fmt.Printf(opts[0].Text + " is your favorite food."); return nil })
		f, err := os.Open(upFile)
		check(err)
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			menu.Option(line, nil, true, nil)
		}
		err = menu.Run()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(eCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// eCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// eCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
