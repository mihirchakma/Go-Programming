package main

import "fmt"

func main() {


    // for loop
    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }
    
    // while loop
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    // break
    for {
        fmt.Println("loop")
        break
    }

    // continue
    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }

    // Basic for loop
    sum := 0
    fmt.Println("Sum is", sum)

    for i := 1; i <= 10; i++ {
        sum += i // sum = sum + 1
        fmt.Println("Increment. Sum is", sum)
    }

    fmt.Println("----------------------")

    // while loop in go
    for sum > 10 {
        sum -= 5
        fmt.Println("Decrement. Sum is", sum)
    }
}