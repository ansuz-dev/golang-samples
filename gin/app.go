package main

import (
  "golang-samples/gin/routes"
  "golang-samples/gin/services"
)

func main() {
  err := services.Connect("dev:root@tcp(127.0.0.1:3306)/contacts?parseTime=true")
  if err != nil {
    panic(err)
  }
  defer services.DB.Close()

  g := routes.Create()
  g.Run("127.0.0.1:3000")
}
