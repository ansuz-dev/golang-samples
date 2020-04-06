package main

import (
  "fmt"
  "sync"
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
  var wg sync.WaitGroup
  wg.Add(2)

  go func(wg *sync.WaitGroup) {
    defer wg.Done()

    x = sum(arr1)
    fmt.Printf("x = %d\n", x)
  }(&wg)

  go func(wg *sync.WaitGroup) {
    defer wg.Done()
    y = sum(arr2)
    fmt.Printf("y = %d\n", y)
  }(&wg)

  wg.Wait()

  fmt.Printf("x + y = %d\n", x+y)
}
