package main

import (
  "fmt"
  "golang-samples/code_organizations/models"
  "golang-samples/code_organizations/services"
)

func main() {
  name := "My_Group"
  group := services.CreateGroup(name)
  fmt.Println("group =", group)

  john := models.Person{Name: "John", Age: 20, Address: "HN"}
  group.AddMember(&john)
  fmt.Println("group =", group)

  jane := &models.Person{Name: "Jane", Age: 20, Address: "HN"}
  group.AddMember(jane)
  fmt.Println("group =", group)

  // Create JWT token for john
  token, err := services.CreateToken(&john)
  if err != nil {
    fmt.Println("There's an error while creating JWT token")
  } else {
    fmt.Println("token: ", token)
  }
}
