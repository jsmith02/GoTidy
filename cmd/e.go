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
	"os/exec"
	"strconv"
	"strings"

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
		edit := getRow(upFile)
		keyOrValue(edit, upFile)
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

// A menu to select which row you'd like to edit
func getRow(upFile string) string {
	var upTask string
	menu := wmenu.NewMenu("Which config do you want to edit?")
	menu.Action(func(opts []wmenu.Opt) error { upTask = opts[0].Text; return nil })
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
	return upTask
}

// Determines which part of a command we adjust, assembles new value.
// Returned as string to main body of command.
func keyOrValue(edit string, upFile string) {
	var changeReq string
	var oldUp chore
	eStr := []byte(edit)
	err := json.Unmarshal(eStr, &oldUp)
	check(err)
	fmt.Printf("Alias: %s", oldUp.Alias[0]+"\n")
	fmt.Printf("Command: %s", oldUp.Cmd[0]+"\n")
	fmt.Println("--------------------------------------------------")
	menu := wmenu.NewMenu("Which would you like to edit")
	menu.Action(func(opts []wmenu.Opt) error { changeReq = opts[0].Text; return nil })
	menu.Option("Alias", nil, true, nil)
	menu.Option("Command", nil, false, nil)
	err = menu.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("--------------------------------------------------")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Editing " + changeReq + "...")
	fmt.Println("Please update the " + changeReq + " to the new value: ")
	userEdit, _ := reader.ReadString('\n')
	fmt.Println("--------------------------------------------------")
	var index int
	index = findLine(oldUp.Alias[0], upFile)
	removeLine(upFile, index)
	writeReplace(changeReq, userEdit, oldUp.Alias[0], oldUp.Cmd[0])
}

func findLine(oldAlias string, upFile string) int {
	// Look for duplicate command aliases
	// Open up, look for keys
	file, err := os.Open(upFile)
	check(err)
	reader := bufio.NewReader(file)
	defer file.Close()
	var line string
	var cnt int
	cnt = 0
	for {
		line, err = reader.ReadString('\n')
		check(err)
		c := []byte(line)
		var oleAle chore
		err := json.Unmarshal(c, &oleAle)
		check(err)
		cnt = cnt + 1
		if fmt.Sprint(oleAle.Alias[0]) == oldAlias {
			return cnt
		}
	}
}

// Deletes old line.
func removeLine(upFile string, lineNumber int) {
	path, err := exec.LookPath("sed")
	check(err)
	cmd := exec.Command(path, "-i ", strconv.Itoa(lineNumber)+"d", upFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		os.Exit(0)
	}
}

// Writes Replacements to up
func writeReplace(changeReq string, userEdit string, oldAlias string, oldCmd string) {
	var newUp chore
	// if request was to change the alias, new chore
	if changeReq == "Alias" {
		newUp = chore{
			Alias: []string{strings.TrimSpace(userEdit)},
			Cmd:   []string{strings.TrimSpace(oldCmd)},
		}
		// Other way around down here.
	} else {
		newUp = chore{
			Alias: []string{strings.TrimSpace(oldAlias)},
			Cmd:   []string{strings.TrimSpace(userEdit)},
		}
	}
	var jsonData []byte
	jsonData, err := json.Marshal(newUp)
	check(err)
	var ToDo chore
	err = json.Unmarshal(jsonData, &ToDo)
	check(err)
	var flagCheck string = ToDo.Cmd[0]
	var cmdBroken = strings.Fields(flagCheck)
	// If we are editing the command, it has to be rebroken
	for i := range cmdBroken {
		var flagBool = strings.HasPrefix(cmdBroken[i], "-")
		if flagBool == true {
			// actual command components
			// If the flag is not the last in the command, substitute
			// <| varN |> where N is the index in the command for later
			// parsing.
			if i+1 < len(cmdBroken) {
				cmdBroken[i+1] = "|_var" + strconv.Itoa(i+1) + "_|"
			}
			if i+1 > len(cmdBroken) {
				continue
			}
		}
	}
	homedir := os.Getenv(("HOME"))
	upFile := homedir + "/.tidy/up"
	if changeReq == "Alias" {
		oldconv := []string{oldCmd}
		upList := chore{
			Alias: []string{strings.TrimSpace(userEdit)},
			Cmd:   []string{strings.Join(oldconv, " ")},
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
		fmt.Println("Alias configured for " + userEdit)
	} else if changeReq == "Command" {
		upList := chore{
			Alias: []string{strings.TrimSpace(oldAlias)},
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
		fmt.Println("Alias recconfigured for " + upList.Alias[0])
	}

}

type chore struct {
	Alias []string `json:"alias"`
	Cmd   []string `json:"cmd"`
}
