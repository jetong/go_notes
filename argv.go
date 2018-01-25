// parsing command-line arguments

package main

import (
	"fmt"
	"os"
	"flag"
)

func main() {
	// contains current scriptname and command-line arguments
	fmt.Print("Args: ")
	for i:=0; i<len(os.Args); i++ {
		fmt.Printf("%s ", os.Args[i])
	}
	fmt.Println()

	// only contains command-line arguments 
	flag.Parse()
	fmt.Print("Args: ")
	for i:=0; len(flag.Arg(i)) != 0; i++ {
		fmt.Printf("%s ", flag.Arg(i))
	}
	fmt.Println()
} 
