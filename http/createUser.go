package http

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/julianinsua/codis/internal/database"
)

type signupRequest struct {
	Username string `json:"username" validate:"required,min=3,alphanum"`
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required,min=8"`
}

type signupResponse struct {
	ID        uuid.UUID    `json:"id"`
	Username  string       `json:"username"`
	Email     string       `json:"email"`
	CreatedAt sql.NullTime `json:"createAt"`
}

func (srv Server) createUser(w http.ResponseWriter, r *http.Request) {
	body, err := getRequestBody[signupRequest](r)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to decode payload: %v", err), 400)
	}
	// Validate fields
	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(body)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("invalid request: %v", err))
		return
	}
	// check username and email
	list, err := srv.store.GetUsersByemailOrUsername(r.Context(), database.GetUsersByemailOrUsernameParams{
		Email:    body.Email,
		Username: body.Username,
	})

	if err != nil {
		if err != sql.ErrNoRows {
			respondWithError(w, 500, fmt.Sprintf("error while checking user: %v", err))
			return
		}
	}

	if len(list) != 0 {
		respondWithError(w, 422, "email or username already exists")
		return
	}
	// Create User
	usr, err := srv.store.CreateUser(r.Context(), database.CreateUserParams{
		Username: body.Username,
		Password: body.Password,
		Email:    body.Email,
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("error while creating user: %v", err))
		return
	}
	respondWithJson(w, 200, signupResponse{
		ID:        usr.ID,
		Username:  usr.Username,
		Email:     usr.Email,
		CreatedAt: usr.CreatedAt,
	})
}
