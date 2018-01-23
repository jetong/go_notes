package main

import (
  "fmt"
  "github.com/jetong/go_pract/fraction"
)

func main() {
  num1 :=fraction.Fraction{4,-6}
  num2 :=fraction.Fraction{-5,10}
  num3 :=num1.Add(num2)

  fmt.Println(num1)		// -2/3
  fmt.Println(num2)		// -1/2
  fmt.Println(num3)		// -1 & 1/6
}
