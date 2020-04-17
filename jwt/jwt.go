package main

import (
  "fmt"
  "github.com/dgrijalva/jwt-go"
  "time"
)

const (
  hmacSecret = "iQasdq4MMdY2wxZCpAm1SxpbkGQopM4wx9QLgtVaHfjGCavuMLcuAZG6CvFxJaMd"
)

type TokenClaims struct {
  jwt.StandardClaims
  AccountID uint `json:"account_id"`
}

func createToken(accountId uint) (tokenStr string, err error) {
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims{
    AccountID: accountId,
    StandardClaims: jwt.StandardClaims{
      ExpiresAt: time.Now().Unix() + 1*60*60, // will be expired in 1 hour
      NotBefore: time.Now().Unix(),
    },
  })

  tokenStr, err = token.SignedString([]byte(hmacSecret))
  return
}

func verifyToken(tokenStr string) (claims *TokenClaims, err error) {
  token, err := jwt.ParseWithClaims(tokenStr, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, fmt.Errorf("Invalid token")
    }

    return []byte(hmacSecret), nil
  })

  if err != nil {
    return
  }

  if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
    return claims, nil
  } else {
    return nil, fmt.Errorf("Invalid token")
  }
}

func main() {
  token, err := createToken(1)
  if err != nil {
    panic(err)
  }
  fmt.Println("Token:", token)

  claims, err := verifyToken(token)
  if err != nil {
    panic(err)
  }
  fmt.Println("Account ID:", claims.AccountID)

  // verify expired token
  claims, err = verifyToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODcxMTI5OTQsIm5iZiI6MTU4NzExMjk5MywiYWNjb3VudF9pZCI6MX0.dDWisMMJ2ZkI60AkQGKILtGVMyowQxzBo88nC0tBHVc")
  if err != nil {
    panic(err)
  }
}
