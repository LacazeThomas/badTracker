package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lacazethomas/badTracker/models"
)

// PATCH /match/:MatchId/winners
// Add set to match
func ModifyWinners(c *gin.Context) {
	// Get model if exist
	var match models.Match
	if err := DB.PreloadAll().Where("id = ?", c.Param("MatchId")).First(&match).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var player = []models.Winner{}
	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Model(&match).Association("Winners").Append(&player)
	c.JSON(http.StatusOK, match)
}

// PATCH /match/:MatchId/losers
// Add set to match
func ModifyLosers(c *gin.Context) {
	// Get model if exist
	var match models.Match
	if err := DB.PreloadAll().Where("id = ?", c.Param("MatchId")).First(&match).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var player = []models.Loser{}
	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Model(&match).Association("Losers").Append(&player)
	c.JSON(http.StatusOK, match)
}
