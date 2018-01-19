package main

import "fmt"
import "math"

// ex1 string params
func greet(greeting string, names ...string) {
	for _, name := range names {
		fmt.Println(greeting + " " + name)
	}
}

// ex2 variable param type
func printValues(values ...interface{}) {
	for _, value := range values {
		switch value.(type) {
		case int:
			fmt.Println(value.(int))
		case string:
			fmt.Println(value.(string))
		case float64:
			fmt.Println(value.(float64))
		}
	}
}

func main() {
	greet("hello", "Bob", "Mary", "Joe")

	printValues(5, "10", math.Pi)
}
