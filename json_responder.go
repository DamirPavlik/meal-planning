package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, playload interface{}) {
	response := map[string]interface{}{
		"code": code,
		"data": playload,
	}

	data, err := json.Marshal(response)
	if err != nil {
		log.Printf("failed to marshal json repsonse: %v", playload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code < 499 {
		log.Println("res with 5xx error: ", msg)
	}
	response := map[string]interface{}{
		"code": code,
		"error": map[string]interface{}{
			"msg": msg,
		},
	}
	data, err := json.Marshal(response)
	if err != nil {
		log.Printf("error marshaling json: %v", err)
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
