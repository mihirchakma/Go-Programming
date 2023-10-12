package main

import "fmt"

func main() {


	//* Syntax: slice_name := [] datatype {values}
	mySlice1 := [] int {10, 20, 30}
	fmt.Println(mySlice1)
	fmt.Println(len(mySlice1))
	fmt.Println(cap(mySlice1))

	mySlice2 := [] string{"Apple", "Orange", "Cherry", "Banana"}
	fmt.Println(mySlice2)
	fmt.Println(len(mySlice2))
	fmt.Println(cap(mySlice2))

	//* Create a Slice From an Array
	// Syntax:
	//* var myarray = [length]datatype{values} // An array
	array1 := [] int{1,2,3,5,10,30,50,90,110}
	
	//* myslice := myarray[start:end] // A slice made from the array
	mySlice3 := array1[1:6]
	fmt.Println(mySlice3)

}