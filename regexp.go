package main

import (
	"fmt"
	"regexp"
	"time"
)

func main() {
	str := "bin bun brain"
	re, err := regexp.Compile(`b(\w+n)`)  // or ("b(\\w+n)")
	if err != nil {
		fmt.Println("Error with regular expression.")
		return
	}
	// or re := regexp.MustCompile("b(\\w+n)"), which panics rather than return an error value

	// returns empty string or slice if no match found.  -1 to match as many as possible
	fmt.Printf("%q\n", re.FindString(str))					// "pin"
	fmt.Printf("%q\n", re.FindStringSubmatch(str))			// ["bin" "in"]
	fmt.Printf("%q\n", re.FindAllString(str, -1))			// ["bin" "bun" "brain"]
	fmt.Printf("%q\n", re.FindAllStringSubmatch(str, -1))	// [["bin" "in"] ["bun" "un"] ["brain" "rain"]]

	now := time.Now()
	fmt.Println(now)
	re2 := regexp.MustCompile(`\d+-\d+-\d+`)
	// MatchString() returns boolean
	if re2.MatchString(now.String()) {
		fmt.Printf("Match found: %q\n", re2.FindString(now.String()))
	}else{
		fmt.Printf("No match found\n")
	}

	// match any number of spaces
	re3 := regexp.MustCompile(" +")
	fmt.Printf("%q\n", re3.Split("hello   world    !!!", -1))			// ["hello" "world" "!!!"]
	fmt.Printf("%q\n", re3.ReplaceAllString("hello   world !", "-"))	// "hello-world-!"
}
