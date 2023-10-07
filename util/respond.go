package util

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, e := json.Marshal(payload)
	if e != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)

}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	// Internal server errors are logged
	if code > 499 {
		log.Println(msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	RespondWithJson(w, code, errResponse{Error: msg})
}
