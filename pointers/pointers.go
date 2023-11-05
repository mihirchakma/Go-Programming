package main

import (
	"fmt"
)

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
	fmt.Println()


	//* pointers to basic types
	i, j := 45, 270

	ptr := &i // point to i
	fmt.Println(*ptr) // read i through the pointer
	*ptr = 10 // set i through the pointer
	fmt.Println(i) // see the new value of i

	p = &j // point to j
	*p = *p / 37 // divide j through the pointer
	fmt.Println(j)

}