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
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// choreCmd represents the chore command
var choreCmd = &cobra.Command{
	Use:     "chore",
	Aliases: []string{"ch"},
	Short:   "Alias a command by adding a chore to your .tidy/up",
	Long: `This function prompts for the alias you'd like to use for a command, and the command
        itself to store in the .tidy/up directory`,
	Run: func(cmd *cobra.Command, args []string) {

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter Command Alias (You'll use this to call the function later e.g. tidy dc for docker compose ): ")
		ali, _ := reader.ReadString('\n')

		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal("Could not determine user home directory")
		}

		upFile := filepath.Join(homeDir, ".tidy", "up")
		size, err := GetFileSize(upFile)
		check(err)
		if size != 0 {
			auditCmd(ali, upFile)
		}
		fmt.Println("Enter command with flags e.g. (aws ssm get-parameter --name |_var_| --with-decryption): ")
		c, _ := reader.ReadString('\n')
		// ali := "gp"
		// c := "git push"
		writeToDo(ali, c, homeDir, upFile)
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
		isIdentic := strings.Compare(iot.Alias[0], ali)
		if isIdentic == 0 {
			fmt.Println("Woops, that alias is already used!")
			os.Exit(0)
		} else if err := reader.Err(); err != nil {
			log.Fatal("Fatal Error while iterating reader in auditCmd\n", err)
		} else {
			continue
		}
	}
}

// Writes aliases to up
func writeToDo(ali string, c string, homedir string, upFile string) {

	upObj := chore{
		Alias: []string{strings.TrimSpace(ali)},
		Cmd:   []string{c},
	}
	var jsonData []byte
	jsonData, err := json.Marshal(upObj)
	check(err)
	var ToDo chore
	err = json.Unmarshal(jsonData, &ToDo)
	check(err)
	flagCheck := ToDo.Cmd[0]
	var cmdBroken = strings.Fields(flagCheck)
	for i := range cmdBroken {
		var flagBool = strings.HasPrefix(cmdBroken[i], "-")
		if flagBool == true {
			// actual command components
			// If the flag is not the last in the command, substitute
			// <| varN |> where N is the index in the command for later
			// parsing.
			if i+1 < len(cmdBroken) {
				cmdBroken[i+1] = fmt.Sprintf("|_var%d_|", i+1)
			}
			if i+1 > len(cmdBroken) {
				continue
			}
		}
	}
	upList := chore{
		Alias: []string{strings.TrimSpace(ali)},
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
	if _, err = f.WriteString(string(jsonData) + "\n"); err != nil {
		panic(err)
	}
	log.Printf("Alias configured for %s\n", upList.Alias[0])
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
