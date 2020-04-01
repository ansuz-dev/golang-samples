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
  fmt.Printf("capacity of arr2 = %d\n", cap(arr2))
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
  fmt.Printf("capacity of slice2 = %d\n", cap(slice2))
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

func sliceCapacity() {
  slice := make([]int, 5)
  fmt.Printf("length of slice = %d\n", len(slice))
  fmt.Printf("capacity of slice = %d\n", cap(slice))
  fmt.Printf("slice = %v\n", slice)
  fmt.Println()

  slice = append(slice, 1)
  fmt.Printf("length of slice = %d\n", len(slice))
  fmt.Printf("capacity of slice = %d\n", cap(slice))
  fmt.Println()

  for i := 0; i < 5; i++ {
    slice = append(slice, i)
  }
  fmt.Printf("length of slice = %d\n", len(slice))
  fmt.Printf("capacity of slice = %d\n", cap(slice))
  fmt.Println()

  for i := 0; i < 10; i++ {
    slice = append(slice, i)
  }
  fmt.Printf("length of slice = %d\n", len(slice))
  fmt.Printf("capacity of slice = %d\n", cap(slice))
  fmt.Println()
}

func sliceASlice() {
  // var array [5]int
  var slice = make([]int, 5)
  fmt.Printf("length of slice = %d\n", len(slice))
  fmt.Printf("capacity of slice = %d\n", cap(slice))
  fmt.Println(slice)
  fmt.Println()

  slice1 := slice[0:3]
  fmt.Printf("length of slice1 = %d\n", len(slice1))
  fmt.Printf("capacity of slice1 = %d\n", cap(slice1))
  slice1[0] = 8
  fmt.Println(slice)
  fmt.Println()

  slice2 := slice[1:3]
  fmt.Printf("length of slice2 = %d\n", len(slice2))
  fmt.Printf("capacity of slice2 = %d\n", cap(slice2))
  fmt.Println()
  slice2[0] = 8
  fmt.Println(slice)

  slice3 := slice1[2:3]
  slice3[0] = 9
  fmt.Println(slice)

  slice4 := make([]int, 3)
  copy(slice4, slice[0:3])
  slice4[0] = 5
  fmt.Println(slice4)
  fmt.Println(slice)
}

func mapType() {
  ages := map[string]int{
    "John": 20,
    "Jane": 21,
    "Jack": 30,
  }
  fmt.Printf("Ages: %v\n", ages)
  fmt.Printf("Length of ages: %v\n", len(ages))
  fmt.Println()

  fmt.Printf("Jane's age: %v\n", ages["Jane"])
  fmt.Println()

  ages["Jane"] = 22
  fmt.Printf("Now, Jane's age: %v\n", ages["Jane"])
  fmt.Println()

  delete(ages, "Jane")
  _, ok := ages["Jane"]
  if !ok {
    fmt.Printf("Jane is removed\n")
    fmt.Println()
  }

  addresses := make(map[string]string, 3)
  addresses["Jane"] = "Hà Nội"
  fmt.Printf("Length of addresses: %v\n", len(addresses))
}

func main() {
  // arrayType()
  // sliceType()
  // sliceCapacity()
  // mapType()
  sliceASlice()
}
