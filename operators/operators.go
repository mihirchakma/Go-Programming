package main

import ("fmt")

func main() {
	
	// Arithmetic Operators 
	//* +, -, *, /, %, ++, --
	//* +=, -=, *=, /=, %= - Assignment Operators 
	var a = 10 + 20
	var b = 20 - 5
	var c = 30 * 3
	var d = 50 / 5
	var e = 45 % 6

	x := 10
	x++

	y := 10
	y--

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println()

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println()

	// Comparison Operators
	//* >, <, >=, <=, ==, !=
	var num1 = 5
	var num2 = 3
	fmt.Println(num1 > num2)
	fmt.Println(num1 < num2)
	fmt.Println(num1 >= num2)
	fmt.Println(num1 <= num2)
	fmt.Println(num1 == num2)
	fmt.Println(num1 != num2)
	fmt.Println()

	// Logical Operators
	//* && - AND = x < 5 &&  x < 10
	//* || - OR = x < 5 || x < 4
	//* ! - NOT = !(x < 5 && x < 10)
	var n1 = 5
	fmt.Println(n1 > 4 && n1 < 10)
	fmt.Println(n1 < 5 || n1 < 4)
	fmt.Println(!(n1 < 5 && n1 < 10))

	

	//* SEE MORE - Bitwise Operators
	
}