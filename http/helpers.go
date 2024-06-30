package http

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, e := json.Marshal(payload)
	if e != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	_, e = w.Write(data)
	if e != nil {
		log.Printf("Server unable to respond: %v", e)
	}

}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	// Internal server errors are logged
	if code > 499 {
		log.Println(msg)
	}

	type errResponse struct {
		Message string `json:"message"`
		Code    string `json:"code"`
	}

	respondWithJson(w, code, errResponse{Message: msg})
}

func getRequestBody[S any](r *http.Request) (S, error) {
	var body S
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return body, err
	}
	return body, nil
}
