package main

import (
	"fmt" 
	"slices"
)

func main() {


	//* Syntax: slice_name := [] datatype {values}
	mySlice1 := [] int {10, 20, 30}
	fmt.Println(mySlice1)
	fmt.Println(len(mySlice1))
	fmt.Println(cap(mySlice1))

	mySlice2 := [] string{"Apple", "Orange", "Cherry", "Banana"}
	fmt.Println(mySlice2)
	fmt.Println(len(mySlice2))
	fmt.Println(cap(mySlice2))

	//* Create a Slice From an Array
	// Syntax:
	//* var myarray = [length]datatype{values} // An array
	array1 := [] int{1,2,3,5,10,30,50,90,110}

	//* myslice := myarray[start:end] // A slice made from the array
	mySlice3 := array1[1:6]
	fmt.Println(mySlice3)

	// Create a Slice With The make() Function
	//* Syntax: slice_name := make([]type, length, capacity)
	mySlice4 := make([]int, 5, 10)
	fmt.Println("Slices =",mySlice4)
	fmt.Println("Length =",len(mySlice4))
	fmt.Println("Capacity =",cap(mySlice4))

	fmt.Println("-----------------------------")


	
	//* Go by Example: Slices 

	var s []string
    fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    fmt.Println("len:", len(s))

	s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("append:", s)

	c := make([]string, len(s))
    copy(c, s)
    fmt.Println("copy:", c)

	l := s[2:5]
    fmt.Println("sl1:", l)

	l = s[:5]
    fmt.Println("sl2:", l)

	l = s[2:]
    fmt.Println("sl3:", l)

	t := []string {"g", "h", "i"}
    fmt.Println("declare:", t)

	t2 := []string{"g", "h", "i"}
    if slices.Equal(t, t2) {
        fmt.Println("t == t2")
    }

	twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
	fmt.Println()



	//* Slice Syntax Example

	numbers := [...] int {1,2,3,4}

	slice1 := numbers[:]

	slice2 := numbers[1:]
	slice3 := slice2[:1]

	slice4 := numbers[:2]

	slice5 := numbers[1:3]

	// slice6 := numbers[1:6] //* error of length array 

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice5)
	// fmt.Println(slice6)

	
	//* ZTM 
	route := []string {"Grocery", "Department Store", "Salon"}
	printSlice("Route 1", route)

	route = append(route, "Home")
	printSlice("Route 2", route)

	fmt.Println()
	fmt.Println(route[0], "Visited")
	fmt.Println(route[1], "Visited")
	
	route = route[2:]
	printSlice("Remaining locations", route)
}

func printSlice(title string, slice []string) {
	fmt.Println()
	fmt.Println("---", title, "---")
	for i := 0; i < len(slice); i++ {
		element := slice[i]
		fmt.Println(element)
	}
}