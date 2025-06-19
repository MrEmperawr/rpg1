package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mremperor-atwork/rpg1/api1/internal/database"
)

// HealthCheck returns the health status of the API and database
func HealthCheck(c *gin.Context) {
	db := database.GetDB()

	// Check database connectivity
	var result int
	if err := db.Raw("SELECT 1").Scan(&result).Error; err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":   "unhealthy",
			"database": "disconnected",
			"error":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "healthy",
		"database": "connected",
		"message":  "API and database are running",
	})
}
