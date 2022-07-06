package calendar

import "time"

// Event represents a calendar event
type Event struct {
	ID          int
	UserID      string
	Date        time.Time
	Description string
}
