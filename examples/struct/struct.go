package main

import "fmt"

type Car struct {
	Make string
	Model string
	height int
	width int
	FrontWheel Wheel
	BackWheel Wheel
}

type Wheel struct {
	Radius int
	Material string
}

func main() {

	toyota := Car{}
	toyota.FrontWheel.Radius = 5
	fmt.Println(toyota.FrontWheel.Radius)
}