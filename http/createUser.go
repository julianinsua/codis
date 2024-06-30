package http

import (
	"fmt"
	"net/http"
)

type signupRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (srv Server) createUser(w http.ResponseWriter, r *http.Request) {
	body, err := getRequestBody[signupRequest](r)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("unsable to decode payload: %v", err))
	}
	fmt.Println(body.Username)
	// Validate fields
	// check username and email
	// create user
}
