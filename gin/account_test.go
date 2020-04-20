package main

import (
  "github.com/gin-gonic/gin"
  "golang-samples/gin/routes"
  "golang-samples/gin/services"
  "net/http"
  "net/http/httptest"
  "os"
  "strings"
  "testing"
)

var ts *gin.Engine

func TestMain(m *testing.M) {
  setUp()
  code := m.Run()
  tearDown()
  os.Exit(code)
}

func setUp() {
  _ = services.Connect("dev:root@tcp(127.0.0.1:3306)/contacts?parseTime=true")
  ts = routes.Create()
}

func tearDown() {
  defer services.DB.Close()
}

func TestGetAccount(t *testing.T) {
  writer := httptest.NewRecorder()
  request, _ := http.NewRequest("GET", "/v1/account/1", nil)
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

func TestUpdateAccount(t *testing.T) {
  json := strings.NewReader(`{"username":"Updated username"}`)

  writer := httptest.NewRecorder()
  request, _ := http.NewRequest("PUT", "/v1/account/1", json)
  ts.ServeHTTP(writer, request)
  if writer.Code != 200 {
    t.Errorf("Response code should be Ok, was: %d", writer.Code)
  }
}
