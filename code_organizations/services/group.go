package services

import (
  "golang-samples/code_organizations/models"
)

// CreateGroup creates a new group with given name
func CreateGroup(name string) (group *models.Group) {
  return &models.Group{
    Name: name,
  }
}
