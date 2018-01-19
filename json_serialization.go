// Serializing (marshaling/unmarshaling) JSON data.  (error checks omitted)

// func Marshal(v interface{}) ([]byte, error)
// func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)  //prefix for appending optional prefix to each line
// func Unmarshal(data []byte, v interface{}) error

package main

import (
	"encoding/json"
	"fmt"
)

// sample struct for marshaling
// capitalize first letters of identifiers to make them visible to external json pkg
type Person struct {
	Name string
	Id   int
}

type Data struct {
	Department string
	Persons    []Person
}

var Persons = []Person{Person{"Joe", 123}, Person{"Mary", 221}}
var s = Data{"Engineering", Persons}

// create jsonData for unmarshaling
var jsonData = []byte(`{"Department":"Engineering","Persons":[{"Name":"Joe","Id":123},{"Name":"Mary","Id":221}]}`)

// create map data
var DataMap = map[string]interface{}{
	"Department":"Engineering",
	"Persons":[]Person{Person{"Joe",123},Person{"Mary",221}},
}

// marshaling from struct object
func marshalStruct() {
	byteSlice, _ := json.Marshal(s)
	display("marshalStruct:", string(byteSlice))
}

// marshaling from struct object with formatting
func marshalIndentStruct() {
	byteSlice, _ := json.MarshalIndent(s, "", "\t")
	display("marshalIndentStruct:", string(byteSlice))
}

// unmarshaling json data to a struct
func unmarshalStruct() {
	_ = json.Unmarshal(jsonData, &s)
	display("unmarshalStruct:", s)

	fmt.Println("Department: " + s.Department)
	fmt.Printf("Names: %s, %s\n", s.Persons[0].Name, s.Persons[1].Name)
}

// unmarshaling json data to empty interface{}
func unmarshalInterface() {
	var i interface{}
	_ = json.Unmarshal(jsonData, &i)
	display("unmarshalInterface:", i)

	// type assertion to access underlyng concrete values of interface{}
	m := i.(map[string]interface{})
	fmt.Println(m["Department"])
	persons := m["Persons"].([]interface{})
//	fmt.Println(persons[0])
	person := persons[0].(map[string]interface{})
	fmt.Println(person["Name"])
}

// marshaling from map
func marshalMap() {
	byteSlice, _ := json.Marshal(DataMap)
	display("marshalMap:", string(byteSlice))
}

// unmarshaling to a map
func unmarshalMap() {
	var myMap map[string]interface{}
	_ = json.Unmarshal(jsonData, &myMap)
	display("unmarshalMap:", myMap)
}

func display(s string, i interface{}) {
	lineBreak()
	switch i.(type) {
		case string:
			fmt.Printf("%v\n%v\n", s, i)
		case map[string]interface{}:
			fmt.Printf("%v\n%v\n", s, i.(map[string]interface{}))
		case Data:
			fmt.Printf("%v\n%v\n", s, i)
		default:
			fmt.Println("No matching type in switch.")
	}
}

func lineBreak() {
	fmt.Println("-----------------------------------------------------")
}

func main() {
	// marshaling from struct into json format
	marshalStruct()
	marshalIndentStruct()

	// unmarshaling json data into a struct
	unmarshalStruct()

	// unmarshaling json data into empty interface
	unmarshalInterface()

	// marshaling from map to json format
	marshalMap()

	// unmarshaling to a map
	unmarshalMap()
}
