package main

import (
  "fmt"
)

func clone(x int) int {
  return x
}

// Declare variable using keyword `var`
func declare1() {
  // declare variable with data type
  var x int
  x = 1
  fmt.Println("x = ", x)

  // declare variable with initial value
  var y int = 2
  /*
     or: var y = 2
  */
  fmt.Println("y = ", y)

  var y2 int = clone(2)
  fmt.Println("y2 = ", y2)

  // declare many variables with data type at once
  var (
    a int
    b int
  )
  a, b = 3, 4
  fmt.Println("a = ", a, ", b = ", b)

  // declare variables having same data type
  var (
    c, d int
    s    string
  )
  c, d, s = 5, 6, "Hello"
  fmt.Println("c = ", c, ", d = ", d, ", s = ", s)

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
  fmt.Println("n = ", n, ", bo = ", bo, ", str = ", str)
}

// Declare variable using inferred typing
func declare2() {
  x := 1
  fmt.Println("x = ", x)

  a, b := 3, 4
  fmt.Println("a = ", a, ", b = ", b)

  c := clone(5)
  fmt.Println("c = ", c)

  c, d := 6, 7
  fmt.Println("c = ", c, ", d = ", d)

  print := fmt.Println
  print("Hello World")
}

func main() {
  // declare_1()

  declare2()
}
