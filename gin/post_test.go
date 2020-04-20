package main

import (
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestGetPosts(t *testing.T) {
  writer := httptest.NewRecorder()
  request, _ := http.NewRequest("GET", "/v1/post", nil)
  ts.ServeHTTP(writer, request)
  if writer.Code != 200 {
    t.Errorf("Response code should be Ok, was: %d", writer.Code)
  }
  if writer.HeaderMap.Get("Content-Type") != "application/json; charset=utf-8" {
    t.Errorf(
      "Content-Type should be application/json, was %s",
      writer.HeaderMap.Get("Content-Type"),
    )
  }
}
