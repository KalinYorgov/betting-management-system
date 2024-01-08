package models

import "time"

// Entry represents an entry with all its fields
type Entry struct {
	ID           int       `json:"id"`
	Status       string    `json:"status"` // winner, loser, halfwin, halflose, void
	Counteragent string    `json:"counteragent"`
	Sport        string    `json:"sport"`       // Football, Basketball, Ice Hockey, etc.
	LiveStatus   string    `json:"live_status"` // Pregame, Live, PreLive
	Amount       float64   `json:"amount"`      // Euro
	Market       string    `json:"market"`      // Handicap -, Handicap +, Over, Under, etc.
	Selection    string    `json:"selection"`
	ODD          string    `json:"odd"`
	Notes        string    `json:"notes"`
	Tournament   string    `json:"tournament"`
	DateCreated  time.Time `json:"date_created"`
	DateFinish   time.Time `json:"date_finish"`
}
