package main

import (
  "testing"
)

func TestFib(t *testing.T) {
  n := Fib(10)
  if n != 55 {
    t.Error("The 10th Fibnonacy should equal 55")
  }
}

func TestFib2(t *testing.T) {
  n := Fib(20)
  if n != 6765 {
    t.Error("The 20th Fibnonacy should equal 6765")
  }
}

func TestFib3(t *testing.T) {
  n := Fib(30)
  if n != 832040 {
    t.Error("The 30th Fibnonacy should equal 832040")
  }
}

func TestFib4(t *testing.T) {
  n := Fib(40)
  if n != 102334155 {
    t.Error("The 40th Fibnonacy should equal 102334155")
  }
}

func TestFib5(t *testing.T) {
  t.Skip("Skipping long-running test")

  n := Fib(45)
  if n != 1134903170 {
    t.Error("The 45th Fibnonacy should equal 1134903170")
  }
}

// Benchmark Fib function
func BenchmarkFib(b *testing.B) {
  for i := 0; i < b.N; i++ {
    Fib(20)
  }
}
