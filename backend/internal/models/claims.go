package models

import "github.com/dgrijalva/jwt-go"

// Claims struct for JWT token payload
type Claims struct {
	Name string `json:"username"`
	jwt.StandardClaims
}
