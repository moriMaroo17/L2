package main

import (
	"log"
	"net/http"
	"time"
)

// LoggingHandler is a middleware that logs all requests to console
func LoggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cur := time.Now().UTC()

		h.ServeHTTP(w, r)

		log.Printf("method: %s  URI: %s  time: %s", r.Method, r.RequestURI, cur)
	})
}
