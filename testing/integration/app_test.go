package main

import (
  "net/http"
  "net/http/httptest"
  "testing"
  "strconv"
)

func TestHandleFib(t *testing.T) {
  mux := http.NewServeMux()
  mux.HandleFunc("/fib", fibonacci)

  writer := httptest.NewRecorder()
  request, _ := http.NewRequest("GET", "/fib?n=10", nil)
  mux.ServeHTTP(writer, request)
  if writer.Code != 200 {
    t.Errorf("Response code is %v", writer.Code)
  }
  fibn, err := strconv.Atoi(writer.Body.String())
  if err != nil {
    t.Error("Cannot convert fibonacci number")
  }
  if fibn != 55 {
    t.Error("Invalid fibonacci number")
  }
}
