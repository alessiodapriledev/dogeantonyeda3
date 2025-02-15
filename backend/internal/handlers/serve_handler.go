package handlers

import (
	"net/http"

	"carabinieri-backend/internal/models"
	"carabinieri-backend/internal/variables"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if user.Username == "admin" && user.Password == "password" {
		token, _ := GenerateToken(user.Username)
		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}

func GetGoogleAPIKey(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"google_api_key": variables.GoogleAPIKey})
}

func ServeHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
