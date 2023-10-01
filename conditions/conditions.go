package main

import ("fmt")

func main() {
	
	//* if..else if..else 
	time := 10

	if time == 11 {
		fmt.Println("Good Morning")
	} else if time < 12 {
		fmt.Println("Good Afternoon")
	} else if time <= 9 {
		fmt.Println("Good Evening")
	} else {
		fmt.Println("Good Night")
	}

	//* Nested if 
	num := 20
	if num >= 10 {
		fmt.Println("Num is more than 10.")
		if num > 15 {
			fmt.Println("Num is also more than 15.")
		}
	} else {
		fmt.Println("Num is less than 10.")
	}

}