package main

import ("fmt")

func add(a, b int) int {
	return a + b
}

func double(x int) int {
	return x + x
}

func greet()  {
	fmt.Println("Hello from greet function")
}

func main()  {
	greet()

	dozen := double(6)
	fmt.Println("A dozen is", dozen)

	bakerDozen := add(dozen, 1)
	fmt.Println("A baker's dozen is", bakerDozen)

	anotherBakerDozen := add(double(6), 1)
	fmt.Println("Another baker's dozen is", anotherBakerDozen)
}