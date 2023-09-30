package main

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func (h handler) redirectTwitch(w http.ResponseWriter, r *http.Request) {
	conn, err := s.Upgrade(w, r)
	if err != nil {
		log.Printf("upgrading failed: %s", err)
		w.WriteHeader(500)
		return
	}
	// Handle the connection. Here goes the application logic.
}
