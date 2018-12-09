package main

import (
  "fmt"
)

func main()  {
  nums := []int{16, 8, 42, 4, 23, 15}
  max := nums[0]
  for _, i := range nums[1:] {
    if i > max{
      max = i
    }
  }
  fmt.Printf("Max = %d ", max)
}
