package main

import (
  "fmt"
  "net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Index!")
}

func setCookie(w http.ResponseWriter, r *http.Request) {
  c1 := http.Cookie{
    Name:     "token",
    Value:    "my_token",
    Path:     "/",
    HttpOnly: true,
  }
  c2 := http.Cookie{
    Name:     "language",
    Value:    "vi",
    Path:     "/",
    HttpOnly: true,
    MaxAge:   10000,
  }

  w.Header().Set("Set-Cookie", c1.String())
  w.Header().Add("Set-Cookie", c2.String())

  // Or:
  http.SetCookie(w, &c1)
  http.SetCookie(w, &c2)

  fmt.Fprintf(w, "Set cookies successfully !")
}

func getCookie(w http.ResponseWriter, r *http.Request) {
  h := r.Header["Cookie"]
  fmt.Fprintln(w, h)

  // Or:
  c1, err := r.Cookie("token")
  if err != nil {
    fmt.Fprintln(w, "Cannot get the token cookie")
    return
  }
  cs := r.Cookies()
  fmt.Fprintln(w, c1.Value)
  fmt.Fprintln(w, cs)
}

func removeCookie(w http.ResponseWriter, r *http.Request) {
  c1 := http.Cookie{
    Name:     "token",
    Value:    "",
    Path:     "/",
    HttpOnly: true,
    MaxAge:   -1,
  }
  c2 := http.Cookie{
    Name:     "language",
    Value:    "",
    Path:     "/",
    HttpOnly: true,
    MaxAge:   -1,
  }

  w.Header().Set("Set-Cookie", c1.String())
  w.Header().Add("Set-Cookie", c2.String())

  // Or:
  http.SetCookie(w, &c1)
  http.SetCookie(w, &c2)

  fmt.Fprintln(w, "Remove cookies successfully !")
}

func main() {
  server := http.Server{
    Addr: "127.0.0.1:3000",
  }

  http.HandleFunc("/", index)
  http.HandleFunc("/cookies/set", setCookie)
  http.HandleFunc("/cookies/get", getCookie)
  http.HandleFunc("/cookies/remove", removeCookie)

  server.ListenAndServe()
}
