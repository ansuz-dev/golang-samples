package main

import (
  "fmt"
)

func main() {
  x := 1
  px := &x

  fmt.Println("x =", x)
  fmt.Println("px =", px)
  fmt.Println("Value of px =", *px)

  arr := [...]int{1, 2, 3, 4, 5}
  pa := &arr
  fmt.Println("arr =", arr)
  fmt.Println("pa =", pa)
  fmt.Printf("pa = %p\n", pa)

  pa[2] = 0
  fmt.Println("arr =", arr)

  var y int
  py := &y
  fmt.Println("y =", y)
  fmt.Println("py =", py)

  removeIndex(arr, 0)
  fmt.Println("arr =", arr)

  removeIndexP(&arr, 0)
  fmt.Println("arr =", arr)

  s := sum(arr)
  fmt.Println("s =", s)

  s = sumP(&arr)
  fmt.Println("s =", s)

  pu := new(int)
  *pu = 2
  fmt.Println("pu =", pu)
  fmt.Println("Value of pu =", *pu)
}

func removeIndex(arr [5]int, index int) {
  arr[index] = 0
}

func removeIndexP(arr *[5]int, index int) {
  arr[index] = 0
}

func sum(arr [5]int) (sum int) {
  for _, v := range arr {
    sum += v
  }

  return
}

func sumP(arr *[5]int) (sum int) {
  for _, v := range arr {
    sum += v
  }

  return
}
