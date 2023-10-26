package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 10
}

func main() {

	// taking a normal variable 
	x := 45
	// declaration of pointer 
	var p *int

	// initialization of pointer 
	p = &x

	fmt.Println("Value stored in x : ", x)
	fmt.Println("Address of x : ", &x)
	fmt.Println("Value stored in variable p : ", p)

	// zero function 
	a := 5
	zero(&a)
	fmt.Println(a)
}