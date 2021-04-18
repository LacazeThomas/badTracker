package models

type Winner struct {
	Player
}

type Loser struct {
	Player
}

type Player struct {
	ID          uint `gorm:"primary_key;AUTO_INCREMENT"`
	MatchID     uint
	Name        string
	FirstName   string
	Sex         string
	Nationality string
}
