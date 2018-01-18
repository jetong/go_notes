package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	msg := "some msg to log"

	// default prints msg to console with timestamp
	log.Printf("%q", msg)
	// output: 2018/01/18 13:18:16 "some text to log"

	// print to file with timestamp
	log.SetOutput(file)
	log.Printf("%q", msg)

	// create new logger to customize the prefix to each log output
	// some common flag values:
	// 0 - nothing
	// 1 - year/month/day
	// 2 - time
	var logger = log.New(file, "_custom prefix_ ", 2)
	logger.Printf("%q", msg)
	// output: _custom prefix_ 13:25:30 "some text to log"
}
