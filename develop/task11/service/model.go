package calendar

import "time"

// Event represents a new event in a calendar
type Event struct {
	ID          int       `json:"id"`
	UserID      string    `json:"user_id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description,omitempty"`
}
