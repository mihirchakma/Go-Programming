package main
import ("fmt")

func main() {

	fmt.Print("Hello, World!\n") // default format

	fmt.Println("Hello, Go!") // a whitespace and a newline will be added

	var say = "World"
	fmt.Printf("Hello, %v!\n\n", say) // formatting verbs based on given format

	var saySomething string = "My name is PI"
	fmt.Printf("%x\n", saySomething)
	fmt.Printf("% x\n", saySomething)
	fmt.Printf("Value = %v, Type = %T\n", saySomething, saySomething)

	//* SEE MORE GO FORMATS - https://pkg.go.dev/fmt
	//* SEE MORE GO FORMATS - https://www.w3schools.com/go/go_output.php
}