package main

import (
  "fmt"
  "net/http"
)

func main() {
  fmt.Println("Server is running on localhost at port 3000 ...")

  // ListenAndServe takes 2 parameters:
  //  - The first parameter is the network address.
  //    The default is all network interfaces at port 80.
  //  - The second parameter is the requests handler.
  //    If it's nil, the default multiplexer, DefaultServeMux, is used.
  http.ListenAndServe("127.0.0.1:3000", nil)
}
