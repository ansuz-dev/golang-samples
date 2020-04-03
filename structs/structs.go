package main

import (
  "fmt"
)

// Name structure
type Name struct {
  Name string
}

// Person structure
type Person struct {
  Name    string
  Age     int
  Address string
}

// Couple structure
type Couple struct {
  P1 Person
  P2 *Person
}

// Group structure
type Group struct {
  Persons []Person
}

func main() {
  // basic initialization
  john := Person{
    Name:    "John",
    Age:     20,
    Address: "HN",
  }
  fmt.Println("John =", john)

  // skip some fields
  jane := Person{
    Name: "Jane",
  }
  jane.Age = 21
  fmt.Println("Jane =", jane)

  // skip field names
  jack := &Person{"Jack", 30, "HN"}
  fmt.Println("Jack =", jack)
  fmt.Println()

  c := Couple{
    P1: Person{
      Name: "John",
      Age:  20,
    },
    P2: &Person{
      Name: "Jane",
      Age:  20,
    },
  }
  fmt.Println("c =", c)
  fmt.Println("c.P2 =", c.P2)
  fmt.Println()

  g := Group{
    Persons: []Person{
      Person{"John", 20, "HN"},
      Person{"Jane", 20, "HN"},
      Person{"Jack", 20, "HN"},
    },
  }
  fmt.Println("g =", g)
}
