package main

import (
	"fmt"
)

func main() {

	// Maps are used to store data values in key:value pairs
	// an unordered and changeable collection, does not allow duplicates
	// Maps hold references to an underlying hash table 

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

	fmt.Println("Delete Age: ", details)



	//* Maps as a Reference 
	var m = map[string] string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	n := m

	fmt.Println(m)
	fmt.Println(n)

	n["year"] = "1970"
	fmt.Println("After change to n: ")

	fmt.Println(m)
	fmt.Println(n)

	//* Iterate Over Maps through range 
	for i, s := range m {
		fmt.Println(i, s)
	}


	//* ZTM 
	shoppingList := make(map[string]int)

	shoppingList["eggs"] = 11
	shoppingList["milk"] = 1
	shoppingList["bread"] += 1

	shoppingList["eggs"] += 1
	fmt.Println(shoppingList)

}