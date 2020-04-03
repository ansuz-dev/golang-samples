package services

import (
  "github.com/dgrijalva/jwt-go"
  "golang-samples/code_organizations/models"
  "time"
)

const (
  hmacSampleSecret = "tmP0O7/K6DbtIysuxnZEZ8tvDpbD9kgmc4KdDMFv/e98FkJXopQ6vxJ0LNedn8Cg"
)

// TokenClaims custom claims
type TokenClaims struct {
  jwt.StandardClaims
  Name string
}

// CreateToken creates the authorization token for a person
func CreateToken(p *models.Person) (tokenString string, err error) {
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims{
    Name: p.Name,
    StandardClaims: jwt.StandardClaims{
      NotBefore: time.Now().Unix(),
    },
  })

  tokenString, err = token.SignedString([]byte(hmacSampleSecret))

  return
}
