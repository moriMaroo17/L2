package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	calendar "github.com/moriMaroo17/l2/develop/task11/service"
)

var storage *calendar.Storage

type configuration struct {
	StorageFile string
	Port        string
}

var config = configuration{}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	body := r.Form

	userID := body.Get("user_id")
	if !validateUserID(w, userID) {
		return
	}

	eventIDasInt, ok := validateEventID(w, body.Get("event_id"))
	if !ok {
		return
	}

	e := calendar.Event{
		ID:     eventIDasInt,
		UserID: userID,
	}
	if err := storage.Delete(e); err != nil {
		response, _ := json.Marshal(fmt.Sprintf(`{"error": "%s"}`, err))
		w.WriteHeader(400)
		w.Write(response)
		return
	}
	response, _ := json.Marshal(fmt.Sprintf(`{"result": "event with id %d was successfully deleted"}`, eventIDasInt))
	w.Write(response)

}

func createHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	body := r.Form

	userID := body.Get("user_id")
	if !validateUserID(w, userID) {
		return
	}

	date, ok := validateDate(w, body.Get("date"))
	if !ok {
		return
	}

	e := calendar.Event{
		UserID:      userID,
		Date:        date,
		Description: body.Get("description"),
	}

	id := storage.Create(e)
	response, _ := json.Marshal(fmt.Sprintf(`{"result": "event with id %d was successfully created"}`, id))
	w.Write(response)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	body := r.Form

	userID := body.Get("user_id")
	if !validateUserID(w, userID) {
		return
	}

	date, ok := validateDate(w, body.Get("date"))
	if !ok {
		return
	}

	eventIDasInt, ok := validateEventID(w, body.Get("event_id"))
	if !ok {
		return
	}

	e := calendar.Event{
		ID:          eventIDasInt,
		UserID:      userID,
		Date:        date,
		Description: body.Get("description"),
	}

	if err := storage.Update(e); err != nil {
		response, _ := json.Marshal(fmt.Sprintf(`{"error": "%s"}`, err))
		w.WriteHeader(400)
		w.Write(response)
		return
	}
	response, _ := json.Marshal(fmt.Sprintf(`{"result": "event with id %d was successfully updated"}`, eventIDasInt))
	w.Write(response)
}

func getByDayHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	userID := params.Get("user_id")
	if !validateUserID(w, userID) {
		return
	}

	date, ok := validateDate(w, params.Get("date"))
	if !ok {
		return
	}

	result := storage.GetEventsByDay(userID, date)
	forReturn := formatResult(result)
	response, err := json.Marshal(forReturn)
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
	if !validateUserID(w, userID) {
		return
	}

	date, ok := validateDate(w, params.Get("date"))
	if !ok {
		return
	}

	result := storage.GetEventsByWeek(userID, date)
	forReturn := formatResult(result)
	response, err := json.Marshal(forReturn)
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
	if !validateUserID(w, userID) {
		return
	}

	date, ok := validateDate(w, params.Get("date"))
	if !ok {
		return
	}

	result := storage.GetEventsByMonth(userID, date)

	forReturn := formatResult(result)
	response, err := json.Marshal(forReturn)
	if err != nil {
		response, _ := json.Marshal(`{"error": "internal server error"}`)
		w.WriteHeader(500)
		w.Write(response)
		return
	}
	w.Write(response)
}

func loadConfigFromFile(filename string) error {
	byteArr, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(byteArr, &config); err != nil {
		return err
	}
	return err
}

func main() {
	if err := loadConfigFromFile(os.Args[1]); err != nil {
		log.Fatal(err)
	}
	storage = calendar.NewStorage(config.StorageFile)
	storage.Restore()

	r := http.NewServeMux()

	r.HandleFunc("/create_event", createHandler)
	r.HandleFunc("/update_event", updateHandler)
	r.HandleFunc("/delete_event", deleteHandler)

	r.HandleFunc("/events_for_day", getByDayHandler)
	r.HandleFunc("/events_for_week", getByWeekHandler)
	r.HandleFunc("/events_for_month", getByMonthHandler)

	m := LoggingHandler(r)
	http.ListenAndServe(fmt.Sprintf(":%s", config.Port), m)
}
