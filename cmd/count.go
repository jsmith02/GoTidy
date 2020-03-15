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
