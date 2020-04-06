package main

import (
  "fmt"
  "time"
)

func write(c chan int) {
  for i := 0; i < 5; i++ {
    fmt.Printf("Prepare to write to channel: %d\n", i)
    c <- i
    fmt.Printf("Wrote to channel: %d\n", i)
  }
}

func read(c chan int) {
  for i := 0; i < 4; i++ {
    v := <-c
    fmt.Printf("Read from channel: %d\n", v)
    time.Sleep(1 * time.Second)
  }
}

func readByRange(c chan int) {
  for i := range c {
    fmt.Printf("Read from channel: %d\n", i)
    time.Sleep(1 * time.Second)
  }

  fmt.Printf("Out of range !\n")
}

func main() {
  c := make(chan int)

  // start to write to channel
  go write(c)

  // read from channel after 2 seconds
  time.Sleep(2 * time.Second)
  go read(c)

  // wait for 6 seconds for everything's finished
  time.Sleep(6 * time.Second)
  fmt.Printf("Finished !\n")
}
