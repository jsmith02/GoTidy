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
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "delete",
	Aliases: []string{"del"},
	Short: "Delete value from up by specifying alias, can also be called with del",
	Long: `Specify an existing alias to delete from your up file. 
	Tidy will look for it in your up, if it can't find it, it will say so. 

	It will also print a confirmation of what it removed from your file. 
	
	jms@jms-desktop:~/go/src/GoTidy$ ./tidy up del gp
	Tidy couldn't find what you wanted in your up.

	jms@jms-desktop:~/go/src/GoTidy$ ./tidy del di
	Removed di from up.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		homedir := os.Getenv(("HOME"))
		upFile := homedir + "/.tidy/up"
		size, err := GetFileSize(upFile)
		check(err)
		if size == 0 {
			fmt.Println("Your up is empty, configure commands before trying to delete them.")
			os.Exit(0)
		} else {
			for i := range args {
				if args[i] != "" {
				index := findLine(args[i], upFile)
				removeLine(upFile, index)
				fmt.Println("Removed " + args[i] + " from up.")
			} else {
				fmt.Println("Please provide an alias to delete.")
				os.Exit(0)
			}
		}
	}
},
}

func init() {
	rootCmd.AddCommand(delCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
