package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lacazethomas/badTracker/models"
)

// PATCH /match/:MatchId/team1
// Add Team1 to match
func ModifyTeam1(c *gin.Context) {
	// Get model if exist
	var match models.Match
	if err := DB.PreloadAll().Where("id = ?", c.Param("MatchId")).First(&match).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var player = []models.Player{}
	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Model(&match).Association("Team1").Append(&player)
	c.JSON(http.StatusOK, match)
}

// PATCH /match/:MatchId/team2
// Add Team2 to match
func ModifyTeam2(c *gin.Context) {
	// Get model if exist
	var match models.Match
	if err := DB.PreloadAll().Where("id = ?", c.Param("MatchId")).First(&match).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var player = []models.Player{}
	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Model(&match).Association("Team2").Append(&player)
	c.JSON(http.StatusOK, match)
}
