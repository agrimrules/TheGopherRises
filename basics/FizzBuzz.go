package main

import (
  "fmt"
)

func main() {
  for i:= 0; i <= 20; i ++ {
    switch true {
    case i % 5 == 0 && i % 3 == 0:
      fmt.Println("FizzBuzz")
    case i % 3 == 0:
      fmt.Println("Fizz")
    case i % 5 == 0:
      fmt.Println("Buzz")
    default:
      fmt.Println(i)
    }
  }
}
