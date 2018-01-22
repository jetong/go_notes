package main

import "fmt"

type anyType interface{}
type sliceObjects []interface{}

func main() {
	var value anyType
	value = 3
	fmt.Println(value)
	value = "hello"
	fmt.Println(value)

	var mySlice sliceObjects
	mySlice = append(mySlice, 8)
	mySlice = append(mySlice, value, "goodbye", []int{1,2})
	fmt.Println(mySlice)
}
