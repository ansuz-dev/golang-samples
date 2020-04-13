package main

import (
  "encoding/json"
  "fmt"
  "net/http"
)

type Person struct {
  // Field appears in JSON as key "name"
  Name string `json:"name"`
}

func handlePostHello(w http.ResponseWriter, r *http.Request) (err error) {
  len := r.ContentLength
  body := make([]byte, len)
  r.Body.Read(body)

  var person Person
  err = json.Unmarshal(body, &person)
  if err != nil {
    return
  }

  w.WriteHeader(200)
  w.Write([]byte("POST: /hello " + person.Name))

  // Reponse person as JSON data
  // output, err := json.Marshal(&person)
  // if err != nil {
  //   return
  // }
  // w.Header().Set("Content-Type", "application/json")
  // w.Write(output)

  return
}

func handlePutHello(w http.ResponseWriter, r *http.Request) (err error) {
  var person Person
  decoder := json.NewDecoder(r.Body)
  err = decoder.Decode(&person)
  if err != nil {
    return
  }

  w.WriteHeader(200)
  w.Write([]byte("PUT: /hello " + person.Name))

  // Response person as JSON data
  // encoder := json.NewEncoder(w)
  // err = encoder.Encode(&person)
  // if err != nil {
  //   fmt.Println("Error encoding JSON:", err)
  // }

  return
}

func hello(w http.ResponseWriter, r *http.Request) {

  switch r.Method {
  case "GET":
    fmt.Fprintf(w, "Handle GET /hello")
  case "POST":
    // curl -X POST http://127.0.0.1:3000/hello
    fmt.Fprintf(w, "Handle POST /hello")

    // curl -X POST -H "Content-Type: application/json" -d '{"name": "Aoshi"}' http://127.0.0.1:3000/hello
    // err := handlePostHello(w, r)
    // if err != nil {
    //   http.Error(w, err.Error(), http.StatusInternalServerError)
    //   return
    // }
  case "PUT":
    // curl -X PUT http://127.0.0.1:3000/hello
    fmt.Fprintf(w, "Handle PUT /hello")

    // curl -X PUT -H "Content-Type: application/json" -d '{"name": "Aoshi"}' http://127.0.0.1:3000/hello
    // err := handlePutHello(w, r)
    // if err != nil {
    //   http.Error(w, err.Error(), http.StatusInternalServerError)
    //   return
    // }
  case "DELETE":
    // curl -X DELETE http://127.0.0.1:3000/hello
    fmt.Fprintf(w, "Handle DELETE /hello")
  default:
    // curl -X PATCH http://127.0.0.1:3000/hello
    fmt.Fprintf(w, "Handle other methods on /hello")
  }

}

func main() {
  server := http.Server{
    Addr: "127.0.0.1:3000",
  }

  http.HandleFunc("/hello", hello)

  server.ListenAndServe()
}
