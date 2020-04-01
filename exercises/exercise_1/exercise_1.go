package main

import (
  "fmt"
)

var names = []string{
  "Katrina", "Evan", "Neil",
  "Adam", "Martin", "Matt",
  "Emma", "Isabella", "Emily",
  "Madison", "Ava", "Olivia",
  "Sophia", "Abigail", "Elizabeth",
  "Chloe", "Samantha", "Addison",
  "Natalie", "Mia", "Alexis",
}

func main() {
  fmt.Println(names)

  groups := map[int][]string{}

  for _, name := range names {
    len := len(name)
    groups[len] = append(groups[len], name)
  }

  fmt.Println(groups)
}
