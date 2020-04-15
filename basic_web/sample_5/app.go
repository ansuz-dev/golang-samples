package main

import (
  "fmt"
  "net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Index!")
}

func hello(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello!")
}

func world(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "World!")
}

func log(h http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Handler function called - " + r.RequestURI)

    h(w, r)
  }
}

func main() {
  server := http.Server{
    Addr: "127.0.0.1:3000",
  }

  http.HandleFunc("/", log(index))
  http.HandleFunc("/hello", log(hello))
  http.HandleFunc("/world", log(log(log(world))))

  server.ListenAndServe()
}
