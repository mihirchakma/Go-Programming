package main

import "fmt"

func main() {

	myArray := [...] int {10, 20, 30, 40}

	for i := 0; i < len(myArray); i ++ {
		item := myArray[i]
		fmt.Println(item)
	}
}