package main

import (
  "fmt"
  "net/http"
  "strconv"
)

func Fib(n int) int {
  if n < 2 {
    return n
  }

  return Fib(n-1) + Fib(n-2)
}

func fibonacci(w http.ResponseWriter, r *http.Request) {
  nStr, ok := r.URL.Query()["n"]
  if !ok || len(nStr[0]) < 1 {
    w.WriteHeader(400)
    fmt.Fprintf(w, "Missing n")
    return
  }

  n, err := strconv.Atoi(nStr[0])
  if err != nil {
    w.WriteHeader(400)
    fmt.Fprintf(w, "Invalid n")
    return
  }

  fibn := Fib(n)
  fmt.Fprintf(w, "%d", fibn)
}

func main() {
  server := http.Server{
    Addr: "127.0.0.1:3000",
  }

  http.HandleFunc("/fib", fibonacci)

  server.ListenAndServe()
}
