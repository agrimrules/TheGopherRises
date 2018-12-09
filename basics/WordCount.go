package main

import (
  "fmt"
  "strings"
)

func main()  {
  input := "Hello I should be able to count all the repating words words words"
  words := strings.Fields(input)
  wmap := map[string]int{}
  for _, k := range words {
    _, ok := wmap[k]
    if !ok {
      wmap[k]= 1
    } else {
      wmap[k]++
    }
  }

  for a,b :=  range wmap {
    fmt.Printf("%s -> %d \n",a, b)
  }
}
