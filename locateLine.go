// search a file for search_string and print line number where found

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	search_string := os.Args[1]
	file := os.Args[2]

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	line := 1
	found := false
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), search_string) {
			fmt.Printf("Found at line: %d\n", line)
			found = true
			break
		}
		line++
	}

	if !found {
		fmt.Println("Not found")
	}
}
