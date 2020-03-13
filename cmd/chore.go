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
	// "io"
	"bufio"
	"fmt"
	// "log"
	"os"
	"strconv"
	"strings"
	"encoding/gob"
	"encoding/json"
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
		type chore struct {
			Alias string `json:"alias"`
			Cmd   string `json:"cmd"`
		}
		// reader := bufio.NewReader(os.Stdin)
		// fmt.Println("Enter Command Alias (You'll use this to call the function later e.g. tidy dc for docker compose ): ")
		// text, _ := reader.ReadString('\n')
		// fmt.Println("Enter Command (Tidy will parse for variables and flags): ")
		// text_, _ := reader.ReadString('\n')
		text := "ba"
		text1 := "aws ssm get-parameter --name /test/just/bazoo --with-decryption "
		upObj := chore{text, text1}
		var flagCheck string = upObj.Cmd
		var cmdBroken = strings.Fields(flagCheck)
		for i := range cmdBroken {
			var flagBool = strings.HasPrefix(cmdBroken[i], "-")
			if flagBool == true {
				// actual command components
				// If the flag is not the last in the command, substitute
				// <| varN |> where N is the index in the command for later
				// parsing.
				if i+1 < len(cmdBroken) {
					cmdBroken[i+1] = "<| var" + strconv.Itoa(i+1) + " |>"
				}
				if i+1 > len(cmdBroken) {
					continue
				}
			}
		}
		var cmdFlagged [1]string
		type M map[string]interface{}
		for i := range cmdBroken {
			cmdFlagged[0] = cmdFlagged[0] + " " + cmdBroken[i]
		}
		homedir := os.Getenv(("HOME"))
		upFile := homedir+"/.tidy/up"
		encodeFile, err := os.OpenFile(upFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		UpToFile, err := json.Marshal(M{"Alias":text, "cmd":cmdFlagged[0]})
		check(err)
		// Since this is a binary format large parts of it will be unreadable
		encoder := gob.NewEncoder(encodeFile)
		// Write to the file
		if err := encoder.Encode(UpToFile); err != nil {
			panic(err)
		}	
		encodeFile.Close()
		// This needs to be moved up
		size, err := GetFileSize(upFile)
		check(err)
		// If there are chores already in the file, check for duplicates.
		if size > 0 {
			fmt.Println(size)
			auditCmd(string(UpToFile), upFile)
		} else {
			fmt.Println("No File Content")
		}
		// // Write comma seperated commands to up for later retrieval
		// f, err := os.OpenFile(homedir+"/.tidy/up", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		// check(err)
		// defer f.Close()
		// if _, err := f.WriteString(string(upMap) + "," + "\n"); err != nil {
		// 	log.Println(err)
		// }

	},
}

// Check file size of up to see if this is the first alias.
func GetFileSize(filepath string) (int64, error) {
    fi, err := os.Stat(filepath)
    if err != nil {
        return 0, err
    }
    // get the size
    return fi.Size(), nil
}


func auditCmd(upMap string, upFile string){
	inFile, err := os.Open(upFile)
	check(err)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		fmt.Printf("%T\n", scanner.Text())
		fmt.Println(scanner.Text()) // the line
	  }
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
