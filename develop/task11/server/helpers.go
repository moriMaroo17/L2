package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	calendar "github.com/moriMaroo17/l2/develop/task11/service"
)

type result struct {
	ID          int    `json:"id"`
	Date        string `json:"date"`
	Description string `json:"description"`
}
type formatedResult struct {
	Result result `json:"result"`
}

func validateUserID(w http.ResponseWriter, userID string) bool {
	if userID == "" {
		response, _ := json.Marshal(`{"error": "specify user id"}`)
		w.WriteHeader(400)
		w.Write(response)
		return false
	}
	return true
}

func validateEventID(w http.ResponseWriter, eventID string) (int, bool) {
	if eventID == "" {
		response, _ := json.Marshal(`{"error": "specify event id"}`)
		w.WriteHeader(400)
		w.Write(response)
		return -1, false
	}
	eventIDasInt, err := strconv.Atoi(eventID)
	if err != nil {
		response, _ := json.Marshal(`{"error": "event id must be an integer"}`)
		w.WriteHeader(400)
		w.Write(response)
		return -1, false
	}
	return eventIDasInt, true
}

func validateDate(w http.ResponseWriter, dateAsString string) (time.Time, bool) {
	date, err := time.Parse("2006-01-02", dateAsString)
	if err != nil {
		response, _ := json.Marshal(`{"error": "wrong date"}`)
		w.WriteHeader(400)
		w.Write(response)
		return time.Time{}, false
	}
	return date, true
}

func formatResult(events []calendar.Event) (forReturn []formatedResult) {
	for _, event := range events {
		forReturn = append(forReturn, formatedResult{
			result{
				event.ID,
				fmt.Sprintf("%d-%d-%d", event.Date.Year(), event.Date.Month(), event.Date.Day()),
				event.Description,
			},
		})
	}
	return
}
