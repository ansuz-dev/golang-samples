package main

import (
  "fmt"
  "sync"
  "time"
)

var (
  counter = 0
  lock    sync.Mutex
)

func incr() {
  counter++
  fmt.Println(counter)
}

func incrWithMutex() {
  lock.Lock()
  defer lock.Unlock()

  counter++
  fmt.Println(counter)
}

func main() {
  for i := 0; i < 20; i++ {
    go incrWithMutex()
  }

  time.Sleep(1 * time.Second)
}
