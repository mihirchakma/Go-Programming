package main

import "fmt"

// create a function that return 1
func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {

	day := 4

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Not a weekday!")
	}

	//* Multi-case switch
	number1 := 7

	switch number1 {
	case 1,3,5:
	 fmt.Println("Odd weekday")
	case 2,4:
	  fmt.Println("Even weekday")
	case 6,7:
	 fmt.Println("Weekend")
   default:
	 fmt.Println("Invalid day of day number.")
   }


   //* From ZTM course
   switch p := price(); {
   case p < 2:
		fmt.Println("Cheap item")
	case p < 10:
		fmt.Println("Moderately priced item")
	default:
		fmt.Println("Expensive item")
   }

   ticket := Economy

   switch ticket {
   case Economy:
		fmt.Println("Economy seating")
	case Business:
		fmt.Println("Business seating")
	case FirstClass:
		fmt.Println("FirstClass seating")
   }
   
}

