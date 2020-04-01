package main

import (
  "fmt"
)

func wordCount(s string) map[string]int {
  // write code here
  bytes := []byte(s)
  words := map[string]int{}

  start := 0
  for index, b := range bytes {
    if b == ' ' {
      word := string(bytes[start:index])
      if len(word) > 0 {
        if _, ok := words[word]; ok {
          words[word]++
        } else {
          words[word] = 1
        }
      }

      start = index + 1
    }
  }

  // last word
  word := string(bytes[start:])
  if _, ok := words[word]; ok {
    words[word]++
  } else {
    words[word] = 1
  }

  return words
}

func main() {
  str := "lorem ipsum dolor sit amet dolor sit amet"
  result := wordCount(str)
  fmt.Println(result)
}
