// search file for line containing search_str and replace with replacement_line

package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	search_str := os.Args[1]
	replacement_line := os.Args[2]
	file := os.Args[3]

	f, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(f), "\n")

	for i, line := range lines {
		if strings.Contains(line, search_str) {
			lines[i] = replacement_line
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(file, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
