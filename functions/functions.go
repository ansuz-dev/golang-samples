package main

import (
  "errors"
  "fmt"
  "os"
)

// function without parameter
func sayHello() {
  fmt.Println("Hello")
}

// function has one parameter
func print(message string) {
  fmt.Println(message)
}

// function has 2 paramters and return a value
func sum(x int, y int) int {
  return x + y
}

/* or:
func sum(x, y int) int {
  return x + y
}
*/

// function has varying numbers of arguments (Variadic function)
func sumAll(initialValue int, params ...int) int {
  sum := initialValue

  for _, v := range params {
    sum += v
  }

  return sum
}

// function returns 2 values
func getBalance(account string) (int, error) {
  if account == "anonymous" {
    return -1, errors.New("Invalid account")
  }

  return 100, nil
}

// declare return variable in function prototype
func getBalance2(account string) (balance int, err error) {
  if account == "anonymous" {
    err = errors.New("Invalid account")
    return
  }

  balance = 100
  return
}

// recursive function
func incr(x int) int {
  if x >= 10 {
    return x
  }

  return incr(x + 1)
}

func double(x int) int {
  return x * 2
}

// anonymous function
func anon() func() int {
  x := 0
  return func() int {
    x++
    return x * x
  }
}

// deferred function call
func square(z *int) {
  *z = *z * *z
  fmt.Println("z in square :", *z)
}

func callDeferred(z *int) {
  *z = 2
  defer square(z)

  *z = *z + 1

  fmt.Println("z in callDeferred :", *z)
}

func doFile(filename string) (err error) {
  f, err := os.Open(filename)
  if err != nil {
    return
  }
  defer f.Close()
  // ...process f...

  return
}

func main() {
  balance, err := getBalance2("anonymous")
  if err == nil {
    fmt.Println("John's balance :", balance)
  } else {
    fmt.Println("Error :", err)
  }

  x := incr(4)
  fmt.Println("x :", x)

  y := 1
  var f func(int) int
  if y > 5 {
    f = incr
  } else {
    f = double
  }
  y = f(y)
  fmt.Println("y :", y)

  fa := anon()
  fmt.Println(fa())
  fmt.Println(fa())

  var z int
  callDeferred(&z)
  fmt.Println("z :", z)
}
