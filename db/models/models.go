package models

import (
   "time"
   jwt "github.com/djrijalva/jwt-go"
   "github.com/ayush/secure-api/randomstrings"
)

const (
   RefreshTokenValidTime = time.Hour * 72
   AuthTokenValidTime = time.Minute * 15
)

type User struct {
   Username     string
   PasswordHash string
   Role         string
}

type TokenClaims struct {
   jwt.StandardClaims
   Role string `json:"role"`
   Csrf string `json:"csrf"`
}

func GenerateCsrfSecret() (string, error) {
   return randomstrings.GenerateRandomString(32)
}