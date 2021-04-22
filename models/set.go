package models

type Set struct {
	ID         uint `gorm:"primary_key"`
	MatchID    uint `json:"matchID" `
	ScoreTeam1 int  `json:"scoreTeam1" `
	ScoreTeam2 int  `json:"scoreTeam2" `
}
