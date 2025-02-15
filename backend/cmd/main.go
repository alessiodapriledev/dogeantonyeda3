package main

import (
	"carabinieri-backend/internal/config"
	"carabinieri-backend/internal/handlers"
	"carabinieri-backend/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDB()

	r := gin.Default()

	apiRoutes := r.Group("/api/v0")

	apiRoutes.Use(middlewares.AuthMiddleware())

	r.POST("/login", handlers.Login)

	handler := handlers.NewHandler(db)
	apiRoutes.GET("/alerts", handler.GetAlerts)
	apiRoutes.POST("/alerts", handler.AddAlert)
	apiRoutes.DELETE("/alerts/:id", handler.DeleteAlert)

	r.LoadHTMLGlob("templates/*")
	r.GET("/", middlewares.AuthMiddleware(), handlers.ServeHTML)

	apiRoutes.GET("/google-api-key", handlers.GetGoogleAPIKey)

	r.Run(":5000")
}
