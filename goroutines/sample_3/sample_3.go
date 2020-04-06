package main

import (
  "fmt"
  "time"
)

func sum(arr []int) (s int) {
  for _, v := range arr {
    s += v
    time.Sleep(1 * time.Second)
  }

  return
}

func main() {
  arr1 := []int{1, 2, 3, 4, 5}
  arr2 := []int{6, 7, 8, 9}

  var x, y int

  go func() {
    x = sum(arr1)
    fmt.Printf("x = %d\n", x)
  }()

  go func() {
    y = sum(arr2)
    fmt.Printf("y = %d\n", y)
  }()

  fmt.Printf("x + y = %d\n", x+y)
}
