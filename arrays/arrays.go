package main

import "fmt"

func main() {

	// declaring arrays 
	//* var array1 [3] int
	//* array2 := [3] int {1, 2, 3}
	//* array3 := [...] int {1, 2, 3}
	//* array4 := [4] int{1, 2, 3}

	// declare an array and print values 
	var arr1 = [3] int {1,2,3}
	fmt.Println(arr1)

	arr2:= [5] int{1,2,3,4,5}
	fmt.Println(arr2)

	// declares an array of strings
	var cars = [5] string{"BMW", "Audi", "Nissan", "Ford", "Ferrari"}
	fmt.Println(cars)

	// Access Elements of an Array
	//* array indexes start at 0, 1, 2 etc.
	fmt.Println(cars[1])
	fmt.Println(cars[3])

	// declare an array and print through in for loop  
	myArray := [...] int {10, 20, 30, 40}

	// Change Elements of an Array
	prices := [3] int {10, 20, 30}
	prices[2] = 40 // change the 30 (index[2]) to 40 
	fmt.Println(prices)

	for i := 0; i < len(myArray); i ++ {
		item := myArray[i]
		fmt.Println(item)
	}
}