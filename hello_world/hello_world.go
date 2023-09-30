package main

import (
  "fmt"
  "rsc.io/quote" // an external package
)

func main() {
  
  fmt.Println("Hello World!")

  fmt.Println(quote.Go())
}