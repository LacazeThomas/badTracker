package models

type Set struct {
	ID         uint `gorm:"primary_key"`
	MatchID    uint
	ScoreTeam1 int
	ScoreTeam2 int
}
