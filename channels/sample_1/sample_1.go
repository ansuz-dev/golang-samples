package main

import (
  "fmt"
  "time"
)

func handleFile(filePath string, newPath chan string) {
  // Process file here
  // Simulate by time.Sleep
  time.Sleep(2 * time.Second)

  // write same path to the channel
  newPath <- filePath
}

func main() {
  filePath := "/path/to/file"

  var path = make(chan string)
  go handleFile(filePath, path)

  // get processed file path from channel
  newPath := <-path

  fmt.Printf("New path: %s\n", newPath)
}
