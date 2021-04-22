package models

type Match struct {
	ID       uint          `json:"ID" `
	Name     string        `json:"name" `
	Location MatchLocation `gorm:"foreignKey:MatchID" json:"location" `
	Winners  []Winner      `gorm:"foreignKey:MatchID" json:"winners" `
	Losers   []Loser       `gorm:"foreignKey:MatchID" json:"losers" `
	Sets     []Set         `gorm:"foreignKey:MatchID" json:"sets" `
}
