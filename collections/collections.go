package main

import (
  "fmt"
)

func arrayType() {
  var arr [2]string
  fmt.Printf("length of arr = %d\n", len(arr))
  fmt.Printf("arr = %v\n", arr)
  arr[0] = "Hà"
  arr[1] = "Nội"
  fmt.Printf("Now, arr = %v\n", arr)
  // arr[2] = "VN"
  fmt.Println()

  arr2 := [2]string{"Hello", "world"}
  fmt.Printf("length of arr2 = %d\n", len(arr2))
  fmt.Printf("arr2 = %v\n", arr2)
  // arr2[2] = "VN"
  // arr2 = append(arr2, "VN")
  fmt.Println()
}

func sliceType() {
  var slice1 []int
  fmt.Printf("length of slice1 = %d\n", len(slice1))
  fmt.Printf("capacity of slice1 = %d\n", cap(slice1))
  fmt.Println()

  slice2 := []int{1, 2}
  fmt.Printf("length of slice2 = %d\n", len(slice2))
  fmt.Printf("capacity of slice2 = %d\n", cap(slice2))
  slice2 = append(slice2, 3)
  fmt.Printf("slice2 = %v\n", slice2)
  fmt.Printf("Now, length of slice2 = %d\n", len(slice2))
  fmt.Println()

  slice3 := make([]int, 10)
  fmt.Printf("length of slice3 = %d\n", len(slice3))
  fmt.Printf("capacity of slice3 = %d\n", cap(slice3))
  fmt.Printf("slice3 = %v\n", slice3)
  slice3[7] = 8
  slice3 = append(slice3, 11)
  fmt.Printf("Now, length of slice3 = %d\n", len(slice3))
  fmt.Printf("Now, slice3 = %v\n", slice3)
  fmt.Println()

  slice4 := make([]int, 1, 10)
  fmt.Printf("length of slice4 = %d\n", len(slice4))
  fmt.Printf("capacity of slice4 = %d\n", cap(slice4))
  fmt.Printf("slice4 = %v\n", slice4)
  fmt.Println()
}

func main() {
  arrayType()
  sliceType()
}
