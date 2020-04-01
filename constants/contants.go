package main

import (
  "fmt"
)

func main() {
  const x = 4
  fmt.Printf("x = %d\n", x)

  const (
    a = 'a'
    b = 'b'
    c = 'c'
  )
  fmt.Printf("a = %v, b = %v, c = %v\n", a, b, c)

  const (
    A = "A"
    B = "B"
    C = "C"
  )
  fmt.Printf("A = %v, B = %v, C = %v\n", A, B, C)

}
