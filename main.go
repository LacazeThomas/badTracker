package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lacazethomas/badTracker/config"
	"github.com/lacazethomas/badTracker/controllers"
	"go.uber.org/zap"
)

func main() {
	var logger *zap.Logger
	var err error

	appConfig := config.GetAppConfig()

	// Set log level
	if appConfig.IsDev() {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		log.Println("Error to initialize logger")
	}

	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	r := gin.Default()

	// Connect to database
	controllers.SetupDatabase(appConfig.GetDsn())

	// Routes
	r.GET("/matches", controllers.FindMatches)
	r.GET("/matches/:MatchId", controllers.FindMatch)
	r.POST("/matches", controllers.CreateMatch)
	r.DELETE("/matches/:MatchId", controllers.DeleteMath)
	r.PATCH("/matches/:MatchId/losers", controllers.ModifyLosers)
	r.PATCH("/matches/:MatchId/winners", controllers.ModifyWinners)
	r.PATCH("/matches/:MatchId/sets", controllers.AddSet)

	// Run the server
	err = r.Run()
	if err != nil {
		log.Println("Unable to start web server")
	}
}
