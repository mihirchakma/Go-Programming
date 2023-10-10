package main

import "fmt"

type Car struct {
	Make string
	Model string
	height int
	width int
	FrontWheel Wheel
	BackWheel Wheel

	/*Wheel struct { // anonymous structure
		Radius int
		Material string
	}*/
}

type Wheel struct {
	Radius int
	Material string
}

// Anonymous Structs
/*
myCar := struct {
	Make string
	Model string
} {
	Make: "Tesla"
	Model: "Model 3"
}
*/

func main() {

	toyota := Car{}
	toyota.Make = "Japan"
	toyota.Model = "Corolla"
	toyota.height = 50
	toyota.width = 40
	toyota.FrontWheel.Radius = 5
	fmt.Println(toyota.Make)
	fmt.Println(toyota.FrontWheel.Radius)

}