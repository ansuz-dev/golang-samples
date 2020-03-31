package main

import (
  "fmt"
)

func main() {
  // declare variable with data type
  var x int
  x = 1

  // declare variable with initial value
  var y int = 2
  /*
     or: var y = 2
  */

  // declare many variables with data type at once
  var (
    a int
    b int
  )
  a, b = 3, 4

  // declare variables having same data type
  var (
    c, d int
    s    string
  )
  c, d, s = 5, 6, "Hello"

  // declare varibles with initial value at once
  var (
    n   int    = 7
    bo  bool   = true
    str string = "World"
  )
  /*
     or:
     var (
       n   = 7
       bo  = true
       str = "World"
     )
  */
  /*
     or: var n, bo, str = 7, true, "World"
  */

  // Print variables
  fmt.Println("x = ", x)
  fmt.Println("y = ", y)
  fmt.Println("a = ", a, ", b = ", b)
  fmt.Println("c = ", c, ", d = ", d, ", s = ", s)
  fmt.Println("n = ", n, ", bo = ", bo, ", str = ", str)
}
