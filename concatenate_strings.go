// concatenate strings via buffer

package main

import (
	"bytes"
	"fmt"
)

func main() {
	var str1, str2 string
	fmt.Println("Provide strings to concatenate")
	fmt.Print("string1: ")
	fmt.Scanln(&str1)
	fmt.Print("string2: ")
	fmt.Scanln(&str2)

	var buffer bytes.Buffer
	buffer.WriteString(str1)
	buffer.WriteString(str2)
	fmt.Printf("Result: %s\n", buffer.String())
}
