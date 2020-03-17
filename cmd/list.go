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
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list your commands from .tidy",
	Long:  `list your commands from .tidy`,
	Run: func(cmd *cobra.Command, args []string) {
		prntCmd()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func prntCmd() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Could not determine user home directory")
	}

	upFile := filepath.Join(homeDir, ".tidy", "up")
	file, err := os.Open(upFile)
	check(err)

	defer func() {
		//generally just a good idea to defer the file close immediately after you know it's open
		//this will run as soon as auditCmd returns

		err = file.Close()
		if err != nil {
			log.Print("Failed to close upfile\n", err)
		}
	}()

	reader := bufio.NewScanner(file)
	for reader.Scan() {
		c := []byte(reader.Text())
		var iot chore
		err = json.Unmarshal(c, &iot)
		check(err)
		fmt.Println(iot.Alias[0], ":", iot.Cmd[0])
	}
}
