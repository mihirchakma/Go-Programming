package main

import "fmt"

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

}
