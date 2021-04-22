package models

type MatchLocation struct {
	ID        uint    `json:"ID" `
	MatchID   uint    `json:"matchID" `
	Latitude  float64 `json:"latitude" `
	Longitude float64 `json:"longitude" `
}
