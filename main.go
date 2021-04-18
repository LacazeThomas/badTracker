package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lacazethomas/badTracker/config"
	"github.com/lacazethomas/badTracker/controllers"
)

func main() {
	r := gin.Default()

	// Connect to database
	controllers.SetupDatabase(config.GetAppConfig().GetDsn())

	// Routes
	r.GET("/matches", controllers.FindMatches)
	r.GET("/matches/:MatchId", controllers.FindMatch)
	r.POST("/matches", controllers.CreateMatch)
	r.DELETE("/matches/:MatchId", controllers.DeleteMath)
	r.PATCH("/matches/:MatchId/losers", controllers.ModifyLosers)
	r.PATCH("/matches/:MatchId/winners", controllers.ModifyWinners)
	r.PATCH("/matches/:MatchId/sets", controllers.AddSet)

	// Run the server
	r.Run()
}
