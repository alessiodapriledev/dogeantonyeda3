package handlers

import (
	"net/http"
	"strconv"

	"carabinieri-backend/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAlerts(c *gin.Context) {
	var alerts []models.Alert
	h.DB.Find(&alerts)
	c.JSON(http.StatusOK, alerts)
}

func (h *Handler) AddAlert(c *gin.Context) {
	var alert models.Alert
	if err := c.ShouldBindJSON(&alert); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	h.DB.Create(&alert)
	c.JSON(http.StatusCreated, alert)
}

func (h *Handler) DeleteAlert(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var alert models.Alert
	result := h.DB.First(&alert, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Alert not found"})
		return
	}
	h.DB.Delete(&alert)
	c.JSON(http.StatusOK, gin.H{"message": "Alert deleted"})
}
