package main
import ("fmt")

const PI float32 = 3.14 //* Unchangeable and READ ONLY 

func main() {

	fmt.Println(PI)

	const NAME = "Mihir Chakma"
	fmt.Println("My name is", NAME)

	// NAME = "Lisa" // error 
	// fmt.Println(NAME)

	//* Multiple Constants Declaration
	const (
		grade int = 12
		number = 235
		greet string = "Hi!"
	)
	fmt.Printf("Grade: %v\n",grade) // %v = to print value, %T = to print type
	fmt.Println("Number:",number)
	fmt.Println("Greeting:",greet)

}