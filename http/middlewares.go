package http

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/julianinsua/codis/internal/database"
	"github.com/julianinsua/codis/util"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (srv *Server) authorizedHandler(handler authHandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// Grab token from the header
		token, err := util.GetToken(r.Header)
		if err != nil {
			// TODO: Respond with 401
			log.Printf("authorization middleware: %v", err)
			return
		}
		// Parse token
		payload, err := srv.tokenMaker.VerifyToken(token)
		if err != nil {
			respondWithError(w, 401, "Unauthorized")
			return
		}

		// Check expiration
		if payload.ExpiresAt.Before(time.Now()) {
			respondWithError(w, 401, "Unauthorized")
			return
		}
		// get user from DB using srv
		usr, err := srv.store.GetUserByID(r.Context(), payload.ID)
		if err != nil {
			respondWithError(w, 401, "Unauthorized")
			return
		}

		handler(w, r, usr)
	}
}

type userHandler func(http.ResponseWriter, *http.Request, database.User)

func (srv *Server) usernameHandler(handler userHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")
		usr, err := srv.store.GetUserByUsername(r.Context(), username)
		if err != nil {
			log.Printf("user not found")
			return
		}
		handler(w, r, usr)
	}
}
