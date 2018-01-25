// reading user input

package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	// fmt package: Scanln()
	fmt.Print("Enter 3 Fahrenheit temperatures separated by spaces: ")
	var t1, t2, t3 float64
	num_args, _ := fmt.Scanln(&t1, &t2, &t3)
	fmt.Println("Entered:", t1, t2, t3)
	fmt.Println("num_args:", num_args)
	fmt.Println()

	// bufio package: ReadString()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter 3 Fahrenheit temperatures separated by spaces: ")
	temps, _ := reader.ReadString('\n')
	fmt.Printf("Entered: %v", temps)
	fmt.Println()

	// bufio package: Scan()
	scanner := bufio.NewScanner(os.Stdin)	
	fmt.Println("Looping through user input. ('q' to quit)")
	fmt.Print("Enter something: ")
	for scanner.Scan() {
		if scanner.Text() == "q" {
			fmt.Println("Exiting loop.")
			break
		}
		fmt.Printf("You entered: %s\n", scanner.Text()) 
		fmt.Print("Enter something: ")
	}
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}
}
