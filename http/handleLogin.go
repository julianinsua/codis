package http

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/julianinsua/codis/internal/database"
	"github.com/julianinsua/codis/token"
	"github.com/julianinsua/codis/util"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	Username         string    `json:"username"`
	Email            string    `json:"email"`
	AuthToken        string    `json:"authToken"`
	RefreshToken     string    `json:"refreshToken"`
	AuthExpiresAt    time.Time `json:"authExpiresAt"`
	RefreshExpiresAt time.Time `json:"refreshExpiresAt"`
}

func (srv *Server) handleLogin(w http.ResponseWriter, r *http.Request) {
	// get the username and password from the request
	body, err := getRequestBody[loginRequest](r)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("unable to decode payload: %v", err))
		return
	}

	// get user by email
	usr, err := srv.store.GetUserByEmail(r.Context(), body.Username)
	if err != nil {
		respondWithError(w, 401, "wrong email or password")
		return
	}

	//compare password sent with hash
	err = util.CheckPassword(body.Password, usr.Password)
	if err != nil {
		respondWithError(w, 401, "wrong email or password")
		return
	}

	// create token and session
	authToken, authPayload, err := srv.tokenMaker.CreateToken(usr.Username, 15*time.Minute)
	if err != nil {
		respondWithError(w, 500, err.Error())
		return
	}

	log.Println(usr.ID)
	refreshToken, expiration, err := ResolveSession(r.Context(), srv.store, srv.tokenMaker, usr.ID, usr.Username)
	if err != nil {
		respondWithError(w, 500, err.Error())
		return
	}
	// return token and refresh token
	respondWithJson(w, 200, loginResponse{
		Username:         usr.Username,
		Email:            usr.Email,
		AuthToken:        authToken,
		RefreshToken:     refreshToken,
		AuthExpiresAt:    authPayload.ExpiresAt,
		RefreshExpiresAt: expiration,
	})
}

func ResolveSession(ctx context.Context, store database.Store, tokenMaker token.Maker, usrID uuid.UUID, username string) (rfshToken string, expiration time.Time, err error) {
	var refreshPayload *token.PASETOPayload
	// check if there's a session
	session, err := store.GetSessions(ctx, usrID)
	if err != nil {
		if err != sql.ErrNoRows {
			return "", time.Time{}, fmt.Errorf("error while getting user session")
		}
		// There's no session, create it
		rfshToken, refreshPayload, err = tokenMaker.CreateToken(username, 24*time.Hour)
		if err != nil {
			return "", time.Time{}, fmt.Errorf("error while creating refresh token: %v", err)
		}

		// Create session for user
		params := database.CreateSessionParams{
			UserID:       usrID,
			RefreshToken: rfshToken,
			ClientAgent:  "",
			ClientIp:     "",
			IsBlocked:    false,
			ExpiresAt:    refreshPayload.ExpiresAt,
			CreatedAt:    time.Now(),
		}
		session, err = store.CreateSession(ctx, params)
		if err != nil {
			return "", time.Time{}, fmt.Errorf("error while creating session %v", err)
		}
		expiration = session.ExpiresAt
		return
	}
	// Session expired
	if session.ExpiresAt.Before(time.Now()) {
		// create token
		rfshToken, refreshPayload, err = tokenMaker.CreateToken(username, 24*time.Hour)
		if err != nil {
			return "", time.Time{}, fmt.Errorf("error while creating token")
		}
		// update session
		params := database.UpdateSessionParams{
			ID:           session.ID,
			RefreshToken: rfshToken,
			ClientAgent:  "",
			ClientIp:     "",
			ExpiresAt:    refreshPayload.ExpiresAt,
		}
		session, err = store.UpdateSession(ctx, params)
		if err != nil {
			return "", time.Time{}, fmt.Errorf("error while updating session")
		}
		expiration = session.ExpiresAt
	}
	rfshToken = session.RefreshToken
	return
}
