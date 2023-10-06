package main

import "fmt"

type Person struct {
	name   string
	age    int
	city   string
}

func main() {

	per1 := Person{"Mihir", 25, "London"}
	fmt.Println(per1.name)
	fmt.Println(per1.age)
	fmt.Println(per1.city)


	per2 := Person {name: "Lisa", age: 24, city: "Wales"}
	fmt.Println(per2.name, per2.age, per2.city)
}