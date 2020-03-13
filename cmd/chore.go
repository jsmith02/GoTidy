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
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// choreCmd represents the chore command
var choreCmd = &cobra.Command{
	Use:     "chore",
	Aliases: []string{"ch", "-c"},
	Short:   "Alias a command by adding a chore to your .tidy/up",
	Long: `This function prompts for the alias you'd like to use for a command, and the command
	itself to store in the .tidy/up directory`,
	Run: func(cmd *cobra.Command, args []string) {
		type chore struct {
			alias string
			cmd   string
		}
		type up struct {
			alias string
			up    []string
		}
		// reader := bufio.NewReader(os.Stdin)
		// fmt.Println("Enter Command Alias (You'll use this to call the function later e.g. tidy dc for docker compose ): ")
		// text, _ := reader.ReadString('\n')
		// fmt.Println("Enter Command (Tidy will parse for variables and flags): ")
		// text_, _ := reader.ReadString('\n')
		text := "ssmg"
		text1 := "aws ssm get-parameter --name /test/just/bazoo --with-decryption "
		upObj := chore{text, text1}
		var flagCheck string = upObj.cmd
		// var aliasCheck = upObj.alias
		var cmdBroken = strings.Fields(flagCheck)
		for i := range cmdBroken {
			var flagBool = strings.HasPrefix(cmdBroken[i], "-")
			if flagBool == true {
				// actual command components
				if i+1 < len(cmdBroken) {
					cmdBroken[i+1] = "<| var" + strconv.Itoa(i+1) + " |>"
				}
				if i+1 > len(cmdBroken) {
					continue
				}
			}
		}
		upToDo := up{text, cmdBroken}
		fmt.Println(upToDo)
		// homedir := os.Getenv(("HOME"))
		u, _ := json.Marshal(upToDo)
		fmt.Println(string(u))
		// file, err := os.OpenFile(homedir+"/.tidy/up", os.O_WRONLY|os.O_APPEND, 0644)
		// if err != nil {
		// 	log.Fatalf("failed opening file: %s", err)
		// }
		// defer file.Close()
		// len, err := file.WriteString(upFile)
		// if err != nil {
		// 	log.Fatalf("failed writing to file: %s", err)
		// }
		// fmt.Printf("\nLength: %d bytes", len)
		// fmt.Printf("\nFile Name: %s", file.Name())
	},
}

func init() {
	rootCmd.AddCommand(choreCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// choreCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// choreCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
