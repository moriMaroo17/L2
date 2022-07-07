package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	calendar "github.com/moriMaroo17/l2/develop/task11/service"
)

var storage *calendar.Storage

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	body := r.Form

	userID := body.Get("user_id")
	if userID == "" {
		response, _ := json.Marshal(`{"error": "specify user id"}`)
		w.WriteHeader(400)
		w.Write(response)

		return
	}

	eventID := body.Get("event_id")
	if eventID == "" {
		response, _ := json.Marshal(`{"error": "specify event id"}`)
		w.WriteHeader(400)
		w.Write(response)
		return
	}
	eventIDasInt, err := strconv.Atoi(eventID)
	if err != nil {
		response, _ := json.Marshal(`{"error": "event id must be an integer"}`)
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	e := calendar.Event{
		ID:     eventIDasInt,
		UserID: userID,
	}
	if err = storage.Delete(e); err != nil {
		response, _ := json.Marshal(fmt.Sprintf(`{"error": "%s"}`, err))
		w.WriteHeader(400)
		w.Write(response)
		return
	}
	response, _ := json.Marshal(fmt.Sprintf(`{"result": "event with id %d was successfully deleted"}`, eventIDasInt))
	// w.WriteHeader(200)
	w.Write(response)

}

func createHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	body := r.Form

	userID := body.Get("user_id")
	if userID == "" {
		response, _ := json.Marshal(`{"error": "specify user id"}`)
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	date, err := time.Parse("2006-01-02", body.Get("date"))
	if err != nil {
		response, _ := json.Marshal(`{"error": "wrong date"}`)
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	e := calendar.Event{
		UserID:      userID,
		Date:        date,
		Description: body.Get("description"),
	}

	id := storage.Create(e)
	response, _ := json.Marshal(fmt.Sprintf(`{"result": "event with id %d was successfully created"}`, id))
	// w.WriteHeader(200)
	w.Write(response)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	body := r.Form

	userID := body.Get("user_id")
	if userID == "" {
		response, _ := json.Marshal(`{"error": "specify user id"}`)
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	date, err := time.Parse("2006-01-02", body.Get("date"))
	if err != nil {
		response, _ := json.Marshal(`{"error": "wrong date"}`)
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	eventID := body.Get("event_id")
	if eventID == "" {
		response, _ := json.Marshal(`{"error": "specify event id"}`)
		w.WriteHeader(400)
		w.Write(response)
		return
	}
	eventIDasInt, err := strconv.Atoi(eventID)
	if err != nil {
		response, _ := json.Marshal(`{"error": "event id must be an integer"}`)
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	e := calendar.Event{
		ID:          eventIDasInt,
		UserID:      userID,
		Date:        date,
		Description: body.Get("description"),
	}

	if err = storage.Update(e); err != nil {
		response, _ := json.Marshal(fmt.Sprintf(`{"error": "%s"}`, err))
		w.WriteHeader(400)
		w.Write(response)
		return
	}
	response, _ := json.Marshal(fmt.Sprintf(`{"result": "event with id %d was successfully updated"}`, eventIDasInt))
	// w.WriteHeader(200)
	w.Write(response)
}

func getByDayHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	userID := params.Get("user_id")
	if userID == "" {
		response, _ := json.Marshal(`{"error": "specify user id"}`)
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	date, err := time.Parse("2006-01-02", params.Get("date"))
	if err != nil {
		response, _ := json.Marshal(`{"error": "wrong date"}`)
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	result := storage.GetEventsByDay(userID, date)

	response, err := json.Marshal(fmt.Sprintf(`{"result": "%v"}`, result))
	if err != nil {
		response, _ := json.Marshal(`{"error": "internal server error"}`)
		w.WriteHeader(500)
		w.Write(response)
		return
	}
	w.Write(response)
}

func getByWeekHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	userID := params.Get("user_id")
	if userID == "" {
		response, _ := json.Marshal(`{"error": "specify user id"}`)
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	date, err := time.Parse("2006-01-02", params.Get("date"))
	if err != nil {
		response, _ := json.Marshal(`{"error": "wrong date"}`)
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	result := storage.GetEventsByWeek(userID, date)
	response, err := json.Marshal(fmt.Sprintf(`{"result": "%v"}`, result))
	if err != nil {
		response, _ := json.Marshal(`{"error": "internal server error"}`)
		w.WriteHeader(500)
		w.Write(response)
		return
	}
	w.Write(response)
}

func getByMonthHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	userID := params.Get("user_id")
	if userID == "" {
		response, _ := json.Marshal(`{"error": "specify user id"}`)
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	date, err := time.Parse("2006-01-02", params.Get("date"))
	if err != nil {
		response, _ := json.Marshal(`{"error": "wrong date"}`)
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	result := storage.GetEventsByMonth(userID, date)
	response, err := json.Marshal(fmt.Sprintf(`{"result": "%v"}`, result))
	if err != nil {
		response, _ := json.Marshal(`{"error": "internal server error"}`)
		w.WriteHeader(500)
		w.Write(response)
		return
	}
	w.Write(response)
}

func main() {
	storage = calendar.NewStorage("storage.json")
	storage.Restore()

	r := http.NewServeMux()

	r.HandleFunc("/create_event", createHandler)
	r.HandleFunc("/update_event", updateHandler)
	r.HandleFunc("/delete_event", deleteHandler)

	r.HandleFunc("/events_for_day", getByDayHandler)
	r.HandleFunc("/events_for_week", getByWeekHandler)
	r.HandleFunc("/events_for_month", getByMonthHandler)

	m := LoggingHandler(r)
	http.ListenAndServe(":8080", m)
}
