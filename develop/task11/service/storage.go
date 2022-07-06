package calendar

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

// Storage is an implementation for storing a calendar events
type Storage struct {
	sync.RWMutex
	storage []Event
	file    string
}

//NewStorage creates a new storage (constructor)
func NewStorage(filename string) (*Storage, error) {
	return &Storage{storage: make([]Event, 0), file: filename}, nil
}

// Delete event by key
func (s *Storage) Delete(key int) error {
	s.Lock()
	defer s.Unlock()
	defer s.write()
	if key < len(s.storage) {
		return fmt.Errorf("no such Storage key: %d", key)
	}
	s.storage = append(s.storage[:key], s.storage[key+1:]...)
	return nil
}

// Create a new event
func (s *Storage) Create(m Event) {
	s.Lock()
	defer s.Unlock()
	defer s.write()
	m.ID = len(s.storage)
	s.storage = append(s.storage, m)
	fmt.Printf("%v\n", s.storage)
}

// Update event by key
func (s *Storage) Update(key int, m Event) error {
	s.Lock()
	defer s.Unlock()
	defer s.write()
	if key < len(s.storage) {
		return fmt.Errorf("no such Storage key: %d", key)
	}
	s.storage[key] = m
	return nil
}

func (s *Storage) write() error {
	byteArr, err := json.Marshal(s.storage)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(s.file, byteArr, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// Restore events from json file to storage
func (s *Storage) Restore() error {
	// byteArr := make([]byte, 0)
	byteArr, err := ioutil.ReadFile(s.file)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(byteArr, &s.storage); err != nil {
		return err
	}
	return nil
}

// GetEventsByDay returns events for given day
func (s *Storage) GetEventsByDay(userID string, date time.Time) []Event {
	s.RLock()
	defer s.RUnlock()

	result := make([]Event, 0)
	for _, event := range s.storage {
		if event.Date.Year() == date.Year() && event.Date.Month() == date.Month() && event.Date.Day() == date.Day() && event.UserID == userID {
			result = append(result, event)
		}
	}
	return result
}

// GetEventsByWeek returns events for given week
func (s *Storage) GetEventsByWeek(userID string, date time.Time) []Event {
	s.RLock()
	defer s.RUnlock()

	result := make([]Event, 0)
	for _, event := range s.storage {
		year1, week1 := event.Date.ISOWeek()
		year2, week2 := date.ISOWeek()
		if year1 == year2 && week1 == week2 && event.UserID == userID {
			result = append(result, event)
		}
	}
	return result
}

// GetEventsByMonth returns events for given month
func (s *Storage) GetEventsByMonth(userID string, date time.Time) []Event {
	s.RLock()
	defer s.RUnlock()

	result := make([]Event, 0)
	for _, event := range s.storage {
		if event.Date.Year() == date.Year() && event.Date.Month() == date.Month() && event.UserID == userID {
			result = append(result, event)
		}
	}
	return result
}
