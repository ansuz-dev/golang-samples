package main

import (
  "fmt"
  "os"
)

/*
 * Create functions to:
 * - Create file `test.txt`
 * - Write "Hello world" to `test.txt`
 * - Read "world" from `test.txt`
 */

func main() {
  f, err := os.Open("test.txt")
  if err != nil {
    panic(err)
  }
  defer f.Close()

  buff := make([]byte, 5)
  _, err = f.Seek(6, 0)
  if err != nil {
    panic(err)
  }

  _, err = f.Read(buff)
  if err != nil {
    panic(err)
  }

  str := string(buff)
  fmt.Println(str)

}
