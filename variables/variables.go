package main

import ("fmt")

func main() {
	
	var myName = "Mihir Chakma"
	fmt.Println("My name is", myName)

	// Variable Declaration With Initial Value 
	var name string = "Lisa"
	fmt.Println("Name =", name)

	userName := "Admin"
	fmt.Println("Username =", userName)

	var sum int
	fmt.Println("The sum is =", sum)

	// Multiple Variable Declaration
	part1, other := 1, 5
	fmt.Println("Part 1 =", part1, "other =", other)

	part2, other := 2, 0
	fmt.Println("Part 2 =", part2, "other =", other)

	sum = part1 + part2
	fmt.Println("Sum is =", sum)

	// Variable Declaration in a Block 
	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)
	fmt.Println("Lesson Name =", lessonName)
	fmt.Println("Lesson Type =", lessonType)

	word1, word2, _ := "Hello", "World", "!"
	fmt.Println(word1, word2)


	// Variables in Go
	var student1 string = "Mihir" // type is string
	var student2 = "Elsa" // type is inferred
	x := 10 // type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	var n1 int
	fmt.Println(n1)

	n1 = 15
	fmt.Println(n1)

}