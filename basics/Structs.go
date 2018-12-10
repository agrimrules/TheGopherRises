package main

import (
  "fmt"
)

type Square struct {
  length int
  Center struct {
    x int
    y int
  }
}

func (s *Square) Area() int {
  return s.length * s.length
}

func (s *Square) Move(dx int, dy int)  {
 s.Center.x += dx
 s.Center.y += dy
}

func main() {
  s := &Square{length: 4}
  s.Center.x = 1
  s.Center.y = 1
  s.Move(1,2)
  fmt.Printf("%+v\n",s)
  ar := s.Area()
  fmt.Printf("Area: %d",ar)
}
