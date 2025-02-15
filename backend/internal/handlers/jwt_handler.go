package handlers

import (
	"carabinieri-backend/internal/models"
	"carabinieri-backend/internal/variables"

	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24)
	claims := &models.Claims{
		Name: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(variables.SecretKey)
}
