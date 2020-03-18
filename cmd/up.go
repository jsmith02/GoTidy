/*
Copyright © 2020 Joseph Smith smith.josephm@gmail.com

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
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Run your aliased command by running tidy up <cmd_here>",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal("Could not determine user home directory")
		}
		upFile := filepath.Join(homeDir, ".tidy", "up")
		size, err := GetFileSize(upFile)
		check(err)
		// If the up file contains data
		if size != 0 {
			cmdFrame := findCmd(args[0], upFile)
			var cmdFmt = strings.Fields(cmdFrame)
			var count int
			for i := range cmdFmt {
				matched, err := regexp.Match(`\|\_var\d*\_\|`, []byte(cmdFmt[i]))
				check(err)
				if matched == true {
					count = count + 1
					continue
				}
			}
			if count != len(args[1:]) {
				fmt.Println("Incorrect number of arguments passed for alias, \"" + args[0] + "\".")
				fmt.Printf("Arguments needed: %d\n", count)
				fmt.Printf("Arguments provided: %d\n", len(args[1:]))

			} else {
				taskAtHand := formatCmd(count, cmdFmt, args)
				exeCmd(taskAtHand)
			}

		} else {
			fmt.Println("No chores in up.")
		}
	},
}

func init() {
	rootCmd.AddCommand(upCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Executes command
func exeCmd(taskAtHand []string) {
	var command []string
	var binCmd string
	binCmd = taskAtHand[0]
	command = taskAtHand[1:]
	path, err := exec.LookPath(binCmd)
	fmt.Println(path)
	cmd := exec.Command(path, command...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	fmt.Println(err)
	if err != nil {
		os.Exit(0)
	}
}

// Formats chores with variables supplied in command
func formatCmd(count int, cmdFmt []string, args []string) []string {
	var trackFormat int
	var index [10]int
	for i := range cmdFmt {
		// track how many values to format left
		matched, err := regexp.Match(`\|\_(var\d*|var)\_\|`, []byte(cmdFmt[i]))
		check(err)
		if matched == true {
			index[i] = i
			trackFormat = trackFormat + 1
			cmdFmt[i] = args[trackFormat]
			continue
		}
	}
	return cmdFmt
}

func findCmd(args string, upFile string) string {
	// Look for duplicate command aliases
	// Open up, look for keys
	file, err := os.Open(upFile)
	check(err)
	reader := bufio.NewReader(file)
	defer file.Close()
	var line string
	for {
		line, err = reader.ReadString('\n')
		check(err)
		c := []byte(line)
		var iot chore
		err := json.Unmarshal(c, &iot)
		check(err)
		if fmt.Sprint(iot.Alias[0]) == args {
			return iot.Cmd[0]
		}
	}
}
