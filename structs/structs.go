package main

import "fmt"

// declare Person struct 
type Person struct {
	name   string
	age    int
	city   string
}

func main() {

	//* Access Person Struct Members
	per1 := Person{"Mihir", 25, "London"}
	fmt.Println("Name:", per1.name)
	fmt.Println("Age:", per1.age)
	fmt.Println("City:", per1.city)

	per2 := Person {name: "Lisa", age: 24, city: "Wales"}
	fmt.Println(per2.name, per2.age, per2.city)

	fmt.Println("--------------------")


	//* Create a passenger object
	ayaan := Passenger {"Ayaan", 1, false}
	fmt.Println(ayaan)

	var (
		alisha = Passenger {Name: "Alisha", TicketNumber: 2}
		hiya = Passenger {Name: "Hiya", TicketNumber: 3}
	)

	fmt.Println(alisha, hiya)

	var heer Passenger
	heer.Name = "Heer"
	heer.TicketNumber = 4
	fmt.Println(heer)

	alisha.Boarded = true
	hiya.Boarded = true
	if hiya.Boarded {
		fmt.Println("Hiya has boarded the bus")
	}
	if alisha.Boarded {
		fmt.Println(alisha.Name, "has boarded the bus")
	}

	heer.Boarded = true
	bus := Bus {heer}
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, "is in the front seat")

}

//* declare Passenger struct 
type Passenger struct {
	Name string
	TicketNumber int
	Boarded bool
}

// delcare Bus struct and struct representing the Passenger
type Bus struct {
	FrontSeat Passenger
}