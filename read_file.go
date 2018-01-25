// reading files

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"bufio"
	"io"
)

func main() {
	// ioutil.Readfile() slurps entire file as a string and adds a newline at the very end of it
	fmt.Println("Using: ioutil.ReadFile()")
    var buffer []byte
    buffer, _ = ioutil.ReadFile("data.txt")
    fmt.Printf("<%s>", buffer)
	fmt.Println()
	fmt.Println()

	// bufio.Reader.ReadString() reads up to and including the specified delimiter
	fmt.Println("Using: bufio.Reader")
    file, _ := os.Open("data.txt")
	defer file.Close()

    reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
	    fmt.Printf("<%s>", line)
	}
	fmt.Println()
	fmt.Println()

    // bufio.Scanner.Scan() & Text() read line by line without including newlines
    fmt.Println("Using: bufio.Scanner")
	file, _ = os.Open("data.txt")
	defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Printf("<%s>", scanner.Text())
    }
    if scanner.Err() != nil {
        fmt.Println(scanner.Err())
    }
	fmt.Println()
}

/* Output:

Using: ioutil.ReadFile()
<line 1 of data
line 2 of data
line 3 of data
>

Using: bufio.Reader
<line 1 of data
><line 2 of data
><line 3 of data
>

Using: bufio.Scanner
<line 1 of data><line 2 of data><line 3 of data>

*/
