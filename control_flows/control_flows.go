package main

import (
  "fmt"
)

func count() int {
  return 8
}

func ifControl() {
  x := 6
  if x > 5 {
    fmt.Println("x is greater than 5")
  } else if x < 5 {
    fmt.Println("x is less than 5")
  } else {
    fmt.Println("x equals 5")
  }

  if y := 7; x > 5 && y > 5 {
    fmt.Println("x and y are greater than 5")
  }

  if z := count(); z > 5 {
    fmt.Println("z is greater than 5")
  }
}

func forLoop() {
  for i := 0; i < 10; i++ {
    fmt.Println(i)
  }

  sum := 1
  for sum < 1000 {
    sum += sum
  }
  fmt.Println("Sum =", sum)
}

func forEach() {
  var arr = [5]int{1, 2, 3, 4, 5}
  for index, value := range arr {
    fmt.Println(index, value)
  }

  fmt.Println()

  var slice = make([]int, 5)
  slice[4] = 4
  for index, value := range slice {
    fmt.Println(index, value)
  }

  ages := map[string]int{
    "John": 20,
    "Jane": 21,
    "Jack": 30,
  }
  for key, value := range ages {
    fmt.Println(key, value)
  }
}

func switchControl(c int) {
  switch c {
  case 0:
    fmt.Println("zero")
  case 2, 4, 6, 8:
    fmt.Println("even")
  case 1, 3, 5, 7:
    fmt.Println("odd")
  default:
    fmt.Println("unknown")
  }

  permissons := []string{}
  accountType := "anonymous"
  switch accountType {
  case "admin":
    permissons = append(permissons, "delete")
    fallthrough
  case "user":
    permissons = append(permissons, "update")
    fallthrough
  default:
    permissons = append(permissons, "read")
  }

  fmt.Println("Permissons :", permissons)
}

func main() {
  // ifControl()
  // forLoop()
  // switchControl(5)
  forEach()
}
