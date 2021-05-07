package models

type Match struct {
	ID       uint          `json:"ID" `
	Date     uint64        `json:"date" `
	Name     string        `json:"name" `
	Location MatchLocation `gorm:"foreignKey:MatchID" json:"location" `
	Team1    []Player1     `gorm:"foreignKey:MatchID" json:"team1" `
	Team2    []Player2     `gorm:"foreignKey:MatchID" json:"team2" `
	Sets     []Set         `gorm:"foreignKey:MatchID" json:"sets" `
}
