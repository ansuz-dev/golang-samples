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

func goSum() {
  arr := []int{1, 2, 3, 4, 5}

  var x int
  go func() {
    x = sum(arr)
    fmt.Printf("x = %d\n", x)
  }()

  fmt.Printf("Finished !\n")
}

// Wait for a goroutine finished by channel
func goSumAndWait() {
  arr := []int{1, 2, 3, 4, 5}

  var x int
  finished := make(chan bool)
  go func(finished chan bool) {
    x = sum(arr)
    fmt.Printf("x = %d\n", x)

    // write value to channel
    finished <- true
  }(finished)

  // read value from channel
  <-finished

  fmt.Printf("Finished !\n")
}

// Wait for a goroutine finished by WaitGroup
func goSumAndWait2() {
  arr := []int{1, 2, 3, 4, 5}

  var x int
  var wg sync.WaitGroup
  // add number of goroutine to wait for
  wg.Add(1)
  go func(wg *sync.WaitGroup) {
    // tell WaitGroup a goroutine finished
    defer wg.Done()

    x = sum(arr)
    fmt.Printf("x = %d\n", x)
  }(&wg)

  // wait for a group finished
  wg.Wait()

  fmt.Printf("Finished !\n")
}

func main() {
  // goSum()
  goSumAndWait()
  // goSumAndWait2()
}
