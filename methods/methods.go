package main

import (
  "fmt"
)

// Rectangle structure
type Rectangle struct {
  w int
  h int
}

func (s Rectangle) area() int {
  return s.w * s.h
}

func (s Rectangle) doubleWidth() {
  fmt.Printf("In doubleWidth, address of s : %p\n", &s)
  s.w = s.w * 2
}

func (s *Rectangle) doubleHeight() {
  fmt.Printf("In doubleHeight, address of s : %p\n", s)
  s.h = s.h * 2
}

func main() {
  s := Rectangle{
    w: 10,
    h: 5,
  }

  fmt.Printf("s : %#v\n", s)
  fmt.Println("Area of s :", s.area())

  fmt.Printf("Address of s : %p\n", &s)
  s.doubleWidth()
  fmt.Printf("After doubleWidth, s : %#v\n", s)

  s.doubleHeight()
  fmt.Printf("After doubleHeight, s : %#v\n", s)
  fmt.Println()

  var s1 Rectangle
  fmt.Printf("s1 : %#v\n", s1)
  fmt.Println("Area of s1 :", s1.area())
  fmt.Println()

  var s2 = &Rectangle{
    w: 6,
    h: 4,
  }
  fmt.Printf("s2 : %p\n", s2)
  fmt.Printf("Value of s2 : %#v\n", *s2)
  fmt.Println("Area of s2 :", s2.area())
  s2.doubleWidth()
  fmt.Printf("After doubleWidth, s2 : %#v\n", s2)
  s2.doubleHeight()
  fmt.Printf("After doubleHeight, s2 : %#v\n", s2)
}
