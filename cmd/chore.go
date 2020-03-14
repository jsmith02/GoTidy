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
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	// "log"
	// "encoding/gob"

	"os"

	// "strconv"
	"strings"
	// "encoding/csv"
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

		// reader := bufio.NewReader(os.Stdin)
		// fmt.Println("Enter Command Alias (You'll use this to call the function later e.g. tidy dc for docker compose ): ")
		// text, _ := reader.ReadString('\n')
		// fmt.Println("Enter Command (Tidy will parse for variables and flags): ")
		// text_, _ := reader.ReadString('\n')
		ali := "ba"
		c := "aws ssm get-parameter --name /test/just/bazoo --with-decryptin "
		homedir := os.Getenv(("HOME"))
		upFile := homedir + "/.tidy/up"
		size, err := GetFileSize(upFile)
		check(err)
		if size != 0 {
			auditCmd(ali, upFile)
		} else {
			upObj := chore{
				Alias: []string{ali},
				Cmd:   []string{c},
			}
			var jsonData []byte
			jsonData, err := json.Marshal(upObj)
			check(err)
			var ToDo chore
			err = json.Unmarshal(jsonData, &ToDo)
			check(err)
			var flagCheck string = ToDo.Cmd[0]
			var cmdBroken = strings.Fields(flagCheck)
			for i := range cmdBroken {
				var flagBool = strings.HasPrefix(cmdBroken[i], "-")
				if flagBool == true {
					// actual command components
					// If the flag is not the last in the command, substitute
					// <| varN |> where N is the index in the command for later
					// parsing.
					if i+1 < len(cmdBroken) {
						cmdBroken[i+1] = "| var" + strconv.Itoa(i+1) + " |"
					}
					if i+1 > len(cmdBroken) {
						continue
					}
				}
			}
			upList := chore{
				Alias: []string{ali},
				Cmd:   []string{strings.Join(cmdBroken, " ")},
			}
			jsonData, err = json.Marshal(upList)
			check(err)
			var toList chore
			err = json.Unmarshal(jsonData, &toList)
			check(err)
			f, err := os.OpenFile(upFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			check(err)
			defer f.Close()
			if _, err = f.WriteString(string(jsonData)); err != nil {
				panic(err)
			}
			fmt.Println("Alias configured for " + upList.Alias[0] + ".")
		}
	},
}

/*
GetFileSize...Check file size of up to see if this is the first alias.
*/
func GetFileSize(filepath string) (int64, error) {
	fi, err := os.Stat(filepath)
	if err != nil {
		return 0, err
	}
	// get the size
	return fi.Size(), nil
}

func auditCmd(ali string, upFile string) {
	// Look for duplicate command aliases
	// Open up, look for keys
	file, err := os.Open(upFile)
	check(err)
	reader := bufio.NewReader(file)
	defer file.Close()
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != io.EOF {
			fmt.Printf(" > Failed!: %v\n", err)
		}
		c := []byte(line)
		var iot chore
		err := json.Unmarshal(c, &iot)
		check(err)
		if fmt.Sprint(iot.Alias[0]) == ali {
			fmt.Println("Woops, that alias is already used!")
			break
		}
	}

}

type chore struct {
	Alias []string `json:"alias"`
	Cmd   []string `json:"cmd"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
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
