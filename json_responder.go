package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, playload interface{}) {
	data, err := json.Marshal(playload)
	if err != nil {
		log.Printf("failed to marshal json repsonse: %v", playload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
