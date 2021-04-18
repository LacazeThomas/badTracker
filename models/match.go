package models

type Match struct {
	ID       uint
	Name     string
	Location MatchLocation `gorm:"foreignKey:MatchID"`
	Winners  []Winner      `gorm:"foreignKey:MatchID"`
	Losers   []Loser       `gorm:"foreignKey:MatchID"`
	Sets     []Set         `gorm:"foreignKey:MatchID"`
}
