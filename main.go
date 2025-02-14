package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Alert struct {
	ID      uint    `json:"id" gorm:"primaryKey"`
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
	Message string  `json:"message"`
}

var db *gorm.DB
var googleAPIKey string

func initDatabase() {
	var err error
	db, err = gorm.Open(sqlite.Open("alerts.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Alert{})
}

func getAlerts(c *gin.Context) {
	var alerts []Alert
	db.Find(&alerts)
	c.JSON(http.StatusOK, alerts)
}

func addAlert(c *gin.Context) {
	var alert Alert
	if err := c.ShouldBindJSON(&alert); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	db.Create(&alert)
	c.JSON(http.StatusCreated, alert)
}

func deleteAlert(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var alert Alert
	result := db.First(&alert, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Alert not found"})
		return
	}
	db.Delete(&alert)
	c.JSON(http.StatusOK, gin.H{"message": "Alert deleted"})
}

func getGoogleAPIKey(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"google_api_key": googleAPIKey})
}

func serveHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func main() {
	initDatabase()
	googleAPIKey = os.Getenv("GOOGLE_MAPS_API_KEY")
	if googleAPIKey == "" {
		fmt.Println("Warning: GOOGLE_MAPS_API_KEY is not set.")
	}

	r := gin.Default()

	// Load HTML templates from the "templates" directory
	r.LoadHTMLGlob("templates/*")

	r.GET("/", serveHTML) // Serve the HTML page when accessing the root
	r.GET("/alerts", getAlerts)
	r.POST("/alerts", addAlert)
	r.DELETE("/alerts/:id", deleteAlert)
	r.GET("/google-api-key", getGoogleAPIKey)

	r.Run(":5000")
}

/*
curl -X POST http://localhost:5000/alerts \
-H "Content-Type: application/json" \
-d '{"lat": 37.616669, "lng": 15.166667, "message": "Alert from Acireale, Catania"}'

*/
