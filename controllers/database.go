package controllers

import (
	"github.com/lacazethomas/badTracker/models"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseDetails struct {
	*gorm.DB
}

var DB DatabaseDetails

func SetupDatabase(dsn string) {

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		zap.S().Panic(err)
	}

	err = database.AutoMigrate(&models.Match{}, &models.MatchLocation{}, &models.Set{}, &models.Winner{}, &models.Loser{})

	if err != nil {
		zap.S().Panic(err, "migration failed")
	}

	DB = DatabaseDetails{database}
}

func (DB DatabaseDetails) PreloadAll() *gorm.DB {
	return DB.Preload("Location").Preload("Winners").Preload("Losers").Preload("Sets")
}
