package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lacazethomas/badTracker/models"
)

// GET /matches
// Find all matches
func FindMatches(c *gin.Context) {
	var matches []models.Match
	DB.PreloadAll().Find(&matches)

	c.JSON(http.StatusOK, matches)
}

// GET /matches/:MatchId
// Find a matches
func FindMatch(c *gin.Context) {
	// Get model if exist
	var match models.Match
	if err := DB.PreloadAll().Where("id = ?", c.Param("MatchId")).First(&match).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, match)
}

// POST /matches
// Create new match
func CreateMatch(c *gin.Context) {
	// ValMatchIdate input
	var input models.Match
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create match
	DB.Create(&input)

	c.JSON(http.StatusOK, input)
}

// DELETE /match/:MatchId
// Delete a match
func DeleteMath(c *gin.Context) {
	// Get model if exist
	var match models.Match
	if err := DB.Where("id = ?", c.Param("MatchId")).First(&match).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	DB.Model(&match).Association("Location").Clear()
	DB.Model(&match).Association("Team1").Clear()
	DB.Model(&match).Association("Team2").Clear()
	DB.Model(&match).Association("Sets").Clear()
	DB.Delete(&match)

	c.JSON(http.StatusOK, true)
}
