package main

import "fmt"
import "reflect"

type Human struct {
  name string
}

type Dog struct {
  name string
  age int
}

type Object interface {
  greet()
}

// implement Object interface
func (h Human) greet() {
  fmt.Println("Greetings, my name is", h.name, "and I'm a", reflect.TypeOf(h), "object.")
}

// implement Object interface
func (d Dog) greet() {
  fmt.Println("Woorf!  Woorf!  I'm a", d.age, "year old", reflect.TypeOf(d), "object, that can talk!")
}

func main() {
  var Bob Object
  var Sparky Object
  Bob = Human{name:"Bob"}
  Sparky = Dog{name:"Sparky", age:3}
  Bob.greet()
  Sparky.greet()
}
