package main

import "fmt"
import "time"

func main() {
	start := time.Now()
	fmt.Println(start.Unix())					// 1516219929
	fmt.Println(start.Format(time.UnixDate))	// Wed Jan 17 12:12:09 PST 2018
	//fmt.Println(start.Hour())

	time.Sleep(time.Second)
	duration := time.Since(start)
	fmt.Printf("Elasped time: %v\n", duration)	// Elasped time: 1.002504698s
}
