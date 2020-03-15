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
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var countCmd = &cobra.Command{
	Use:     "count",
	Aliases: []string{"#"},
	Short:   "The number of chores in your up.",
	Long:    `Iterates over file and indicates how many aliases are configured.`,
	Run: func(cmd *cobra.Command, args []string) {
		lineCounter()
	},
}

func init() {
	rootCmd.AddCommand(countCmd)
	countCmd.Flags().Int("number", 10, "A help for number")
}

func lineCounter() {
	homedir := os.Getenv(("HOME"))
	upFile := homedir + "/.tidy/up"
	counter := 0
	f, err := os.Open(upFile)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		counter++
		line = line[:0]
	}
	fmt.Println(strconv.Itoa(counter) + " configs stored in " + upFile)
}
