package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// will remove the desired line from the desired file, and writes the results into the provided output file
// ***THIS FOLLOWS SED -i nd RULES. THAT MEANS IT IS 1-INDEXED!!!!!!***
func RemoveLineFromFile(path string, n int, outputPath string) error {
	if _, err := os.Stat(path); err != nil {
		return err
	}

	f, err := os.Open(path)
	if err != nil {
		return err
	}

	content, err := ioutil.ReadAll(f)
	if err != nil {
		//Gotta close this here or else we're going to leak open files.
		//I know I know... We should defer it above, but i'm tryna do something here!
		defer f.Close()
		return err
	}

	f.Close() //we're not deferring because we might be writing into the file that we already had

	nlIdx := findNewLineIndicies(&content)
	if n > len(*nlIdx)+1 {
		return errors.New("line number out of range (not enough new lines in the file")
	}

	var start int
	var end int
	if n == 1 {
		start = 0
	} else {
		start = (*nlIdx)[n-2] + 1
	}

	if n == len(*nlIdx)+1 {
		end = len(content)
	} else {
		end = (*nlIdx)[n-1] + 1
	}

	//the indexing here is weird, newline 0 is the end of line 1, so we have to adjust down by two

	subLen := end - start
	bSize := len(content) - subLen

	outBuffer := make([]byte, bSize)

	copy(outBuffer[0:start], content[0:start])
	copy(outBuffer[start:], content[end:])

	if strings.EqualFold(path, outputPath) {
		err := os.Remove(path)
		if err != nil {
			return err
		}
		_, err = os.Create(path)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(path, outBuffer, 777)
		if err != nil {
			return err
		}
	} else {
		err := ioutil.WriteFile(outputPath, outBuffer, 777)
		if err != nil {
			return err
		}
	}
	return nil
}

func findNewLineIndicies(buffer *[]byte) *[]int {
	indicies := make([]int, 0, 50) //init with 50 indicies, decent starting size, it will get auto sized if we need more
	for i := range *buffer {
		c := (*buffer)[i]
		if c == '\n' {
			indicies = append(indicies, i)
		}
	}
	return &indicies
}
