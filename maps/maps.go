package main

import (
	"fmt"
)

func main() {

	//* Create Maps Using var and := 
	var car1 = map[string]string{"Brand": "BMW", "Model": "5", "Year": "1972"}
	fmt.Printf("%v\n", car1)

	car2 := map[string]string{"Brand":"Ford", "Model":"Mustang", "Year":"1964"}
	fmt.Println(car2)

	/*
	Syntax: 
	var a = make(map[KeyType]ValueType)
	b := make(map[KeyType]ValueType)
	*/
	//* Create Maps Using make()Function:
	details := make(map[string]string)
	details["Name"] = "Mihir"
	details["Age"] = "25"
	details["City"] = "Gampaha"

	fmt.Printf("%v\n", details)

	//* Create an Empty Map 
	// Syntax: var a map[KeyType]ValueType
	var a = make(map[string]string)
	var b map[string]string

	fmt.Println(a == nil)
	fmt.Println(b == nil)

	//* Access Map Elements
	// Syntax: map_name[key] = value
	fmt.Println(details["Name"])
	fmt.Println(details["City"])

	//* Update and Add Map Elements 
	details["Age"] = "26" // Updating an element 
	details["Study"] = "SLTC" // Adding an element 

	fmt.Println(details)

	//* Remove Element from Map 
	// Syntax: delete(map_name, key) 
	delete(details, "Age")

	fmt.Println("Age Deleted: ", details)

}