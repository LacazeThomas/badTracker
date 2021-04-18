package models

type Set struct {
	ID      uint `gorm:"primary_key"`
	MatchID uint
	Team1   int
	Team2   int
}
