package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// data for buffer
	file, err := os.Open("buffer_data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println("-------------------------")
	// byte representation
	var b []byte
	b = []byte{'a', 'b', 'c'}	// equivalent to b = []byte("abc")
	fmt.Printf("%s\n", b)
	fmt.Printf("%v\n", b)

	fmt.Println("-------------------------")

	// bytes.Buffer for concatenating strings
	var buf bytes.Buffer  // buf := bytes.NewBufferString("")
	buf.Write([]byte("hello"))
	buf.Write([]byte(" world"))
	fmt.Println(buf.String())
	fmt.Println("Emptying buffer")
	buf.Reset()
	fmt.Println(buf.String())

	fmt.Println("-------------------------")

	// []byte buffer for reading entire file
	var fileBuffer []byte
	fileBuffer, _ = ioutil.ReadFile("buffer_data.txt")
	fmt.Printf("%s", fileBuffer)

	fmt.Println("-------------------------")

	// []byte buffer for reading 4 bytes at a time
	var buff = make([]byte, 4)
	for {
		if n, err := file.Read(buff); n > 0 {
			fmt.Printf("|%s|\n", buff[:n])
		} else if err != nil {   // EOF is detected as an "error", so check err for EOF
			if err == io.EOF {
				fmt.Println("Read all file content.")
				break
			} else {
				log.Fatal(err)
			}
		}
	}
	fmt.Println("-------------------------")
}
