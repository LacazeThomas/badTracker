package models

type Player1 struct {
	ID          uint   `gorm:"primary_key;AUTO_INCREMENT"`
	MatchID     uint   `json:"matchID" `
	Name        string `json:"name" `
	FirstName   string `json:"firstName" `
	Sex         string `json:"sex" `
	Nationality string `json:"nationality" `
}

type Player2 struct {
	ID          uint   `gorm:"primary_key;AUTO_INCREMENT"`
	MatchID     uint   `json:"matchID" `
	Name        string `json:"name" `
	FirstName   string `json:"firstName" `
	Sex         string `json:"sex" `
	Nationality string `json:"nationality" `
}
