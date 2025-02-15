package middlewares

import (
	"carabinieri-backend/internal/models"
	"carabinieri-backend/internal/variables"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the "Authorization" header
		tokenString := c.GetHeader("Authorization")
		fmt.Println("tokenString: ", tokenString)

		// Check if the token exists
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			c.Abort()
			return
		}

		// Remove the "Bearer " prefix if it exists
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Parse the token and validate it
		token, err := jwt.ParseWithClaims(tokenString,
			&models.Claims{},
			func(token *jwt.Token) (interface{}, error) {
				return variables.SecretKey, nil
			})

		// Handle parsing errors or invalid token
		if err != nil || !token.Valid {
			fmt.Println("Error parsing token:", err) // Print the error for debugging
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Token is valid, continue to the next handler
		c.Next()
	}
}
