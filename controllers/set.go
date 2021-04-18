package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lacazethomas/badTracker/models"
)

// PATCH /match/:MatchId/sets
// Add set to match
func AddSet(c *gin.Context) {
	// Get model if exist
	var match models.Match
	if err := DB.PreloadAll().Where("id = ?", c.Param("MatchId")).First(&match).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var addSet = models.Set{}
	if err := c.ShouldBindJSON(&addSet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Model(&match).Association("Sets").Append(&addSet)
	c.JSON(http.StatusOK, match)
}
