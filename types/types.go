package main

import (
  "fmt"
  "unsafe"
)

func main() {
  // boolean
  var bo bool
  bo = true
  bo = false
  fmt.Printf("bo = %t\n", bo)
  fmt.Println()

  // string
  var str string
  str = "Hà Nội"
  fmt.Printf("str = %s\n", str)
  fmt.Printf("length in bytes of str = %d\n", len(str))
  fmt.Printf("length in bytes of H = %d\n", len("H"))
  fmt.Printf("length in bytes of à = %d\n", len("à"))
  fmt.Printf("length in bytes of ' ' = %d\n", len(" "))
  fmt.Printf("length in bytes of N = %d\n", len("N"))
  fmt.Printf("length in bytes of ộ = %d\n", len("ộ"))
  fmt.Printf("length in bytes of i = %d\n", len("i"))

  bytes := []byte(str)
  fmt.Println(bytes)

  runes := []rune(str)
  fmt.Println(runes)
  fmt.Println()

  // integer
  var (
    i   int
    i8  int8
    i16 int16
    i32 int32
    i64 int64
  )
  fmt.Printf("size in bytes of i = %d\n", unsafe.Sizeof(i))
  fmt.Printf("size in bytes of i8 = %d\n", unsafe.Sizeof(i8))
  fmt.Printf("size in bytes of i16 = %d\n", unsafe.Sizeof(i16))
  fmt.Printf("size in bytes of i32 = %d\n", unsafe.Sizeof(i32))
  fmt.Printf("size in bytes of i64 = %d\n", unsafe.Sizeof(i64))
  fmt.Println()

  // unsigned integer
  var (
    ui   int
    ui8  int8
    ui16 int16
    ui32 int32
    ui64 int64
  )
  fmt.Printf("size in bytes of ui = %d\n", unsafe.Sizeof(ui))
  fmt.Printf("size in bytes of ui8 = %d\n", unsafe.Sizeof(ui8))
  fmt.Printf("size in bytes of ui16 = %d\n", unsafe.Sizeof(ui16))
  fmt.Printf("size in bytes of ui32 = %d\n", unsafe.Sizeof(ui32))
  fmt.Printf("size in bytes of ui64 = %d\n", unsafe.Sizeof(ui64))
  fmt.Println()

  // float
  var (
    f32 float32
    f64 float64
  )
  fmt.Printf("size in bytes of f32 = %d\n", unsafe.Sizeof(f32))
  fmt.Printf("size in bytes of f64 = %d\n", unsafe.Sizeof(f64))
  fmt.Println()

  // complex : a + bi
  var (
    c64  complex64
    c128 complex128
  )
  fmt.Printf("size in bytes of c64 = %d\n", unsafe.Sizeof(c64))
  fmt.Printf("size in bytes of c128 = %d\n", unsafe.Sizeof(c128))
  fmt.Println()

  // byte and rune
  var (
    b byte
    r rune
  )
  fmt.Printf("size in bytes of b = %d\n", unsafe.Sizeof(b))
  fmt.Printf("size in bytes of r = %d\n", unsafe.Sizeof(r))
  fmt.Println()
}
