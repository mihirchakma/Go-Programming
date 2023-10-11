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

	// Change Elements of an Array
	prices := [3] int {10, 20, 30}
	prices[2] = 40 // change the 30 (index[2]) to 40 
	fmt.Println(prices)


	//* Array Initialization
	nArr1 := [5]int{} //not initialized
	nArr2 := [5]int{1,2} //partially initialized
	nArr3 := [5]int{1,2,3,4,5} //fully initialized

	fmt.Println(nArr1, nArr2, nArr3)

	// declare an array and print through in for loop using len() function 
	myArray := [...] int {10, 20, 30, 40}

	for i := 0; i < len(myArray); i ++ {
		item := myArray[i]
		fmt.Println(item)
	}


	// ZTM 
	rooms := [...] Room {
		{name: "Office"},
		{name: "Warhouse"},
		{name: "Reception"},
		{name: "Ops"},
	}

	// calling function 
	checkCleanliness(rooms)

	fmt.Println("Performing cleaning...")

	rooms[1].cleaned = true
	rooms[2].cleaned = true
	rooms[3].cleaned = true

	checkCleanliness(rooms)

}

// ZTM 
type Room struct {
	name string
	cleaned bool
}

func checkCleanliness(rooms[4] Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.cleaned {
			fmt.Println(room.name, "is clean")
		} else {
			fmt.Println(room.name, "is dirty")
		}
	}
}