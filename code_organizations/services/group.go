package services

import (
  "fmt"
  "golang-samples/code_organizations/models"
)

// a private function starts with lowercase
// and will not be exported
func testGroup() {
  fmt.Println("Test group function")
}

// CreateGroup creates a new group with given name
func CreateGroup(name string) (group *models.Group) {
  return &models.Group{
    Name: name,
  }
}
