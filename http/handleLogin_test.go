package http

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	mock_db "github.com/julianinsua/codis/db/mock"
	"github.com/julianinsua/codis/files"
	"github.com/julianinsua/codis/internal/database"
	"github.com/julianinsua/codis/parser"
	"github.com/julianinsua/codis/token"
	mock_token_mkr "github.com/julianinsua/codis/token/mock"
	"github.com/julianinsua/codis/util"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

type loginTestCase struct {
	name         string
	payload      loginRequest
	buildStubs   func(*mock_db.MockStore, *mock_token_mkr.MockMaker)
	checkResults func(t *testing.T, recorder *httptest.ResponseRecorder)
}

func TestHanldeLogin(t *testing.T) {
	usr, pass, err := database.RandomUser()
	session := database.Session{
		ID:           uuid.New(),
		UserID:       usr.ID,
		RefreshToken: "someRefreshToken",
		ClientAgent:  "",
		ClientIp:     "",
		IsBlocked:    false,
		ExpiresAt:    time.Now().Add(1 * time.Minute),
		CreatedAt:    time.Now(),
	}

	expiredSession := database.Session{
		ID:           uuid.New(),
		UserID:       usr.ID,
		RefreshToken: "someRefreshToken",
		ClientAgent:  "",
		ClientIp:     "",
		IsBlocked:    false,
		ExpiresAt:    time.Now().Add(-1 * time.Minute),
		CreatedAt:    time.Now(),
	}
	require.NoError(t, err)
	testCases := []loginTestCase{
		{
			name:    "OK no session",
			payload: loginRequest{Username: usr.Username, Password: pass},
			buildStubs: func(ms *mock_db.MockStore, mt *mock_token_mkr.MockMaker) {
				ms.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Times(1).Return(usr, nil)
				ms.EXPECT().GetSessions(gomock.Any(), gomock.Any()).Times(1).Return(database.Session{}, sql.ErrNoRows)
				mt.EXPECT().CreateToken(gomock.Any(), gomock.Any()).Times(2).Return("sometoken", &token.PASETOPayload{
					Payload: token.Payload{
						ID:       usr.ID,
						Username: usr.Username,
					},
					ExpiresAt: time.Now().Add(1 * time.Minute),
					IssuedAt:  time.Now(),
				}, nil)
				ms.EXPECT().CreateSession(gomock.Any(), gomock.Any()).Times(1).Return(session, nil)
			},
			checkResults: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:    "OK with session",
			payload: loginRequest{Username: usr.Username, Password: pass},
			buildStubs: func(ms *mock_db.MockStore, mt *mock_token_mkr.MockMaker) {
				ms.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Times(1).Return(usr, nil)
				ms.EXPECT().GetSessions(gomock.Any(), gomock.Any()).Times(1).Return(session, nil)
				mt.EXPECT().CreateToken(gomock.Any(), gomock.Any()).Times(1).Return("sometoken", &token.PASETOPayload{
					Payload: token.Payload{
						ID:       usr.ID,
						Username: usr.Username,
					},
					ExpiresAt: time.Now().Add(1 * time.Minute),
					IssuedAt:  time.Now(),
				}, nil)
			},
			checkResults: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:    "OK with expired session",
			payload: loginRequest{Username: usr.Username, Password: pass},
			buildStubs: func(ms *mock_db.MockStore, mt *mock_token_mkr.MockMaker) {
				ms.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Times(1).Return(usr, nil)
				ms.EXPECT().GetSessions(gomock.Any(), gomock.Any()).Times(1).Return(expiredSession, nil)
				ms.EXPECT().UpdateSession(gomock.Any(), gomock.Any()).Times(1)
				mt.EXPECT().CreateToken(gomock.Any(), gomock.Any()).Times(2).Return("sometoken", &token.PASETOPayload{
					Payload: token.Payload{
						ID:       usr.ID,
						Username: usr.Username,
					},
					ExpiresAt: time.Now().Add(1 * time.Minute),
					IssuedAt:  time.Now(),
				}, nil)
			},
			checkResults: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:    "Not found",
			payload: loginRequest{Username: usr.Username, Password: pass},
			buildStubs: func(ms *mock_db.MockStore, mt *mock_token_mkr.MockMaker) {
				ms.EXPECT().GetUserByEmail(gomock.Any(), gomock.Eq(usr.Username)).Times(1).Return(database.User{}, sql.ErrNoRows)
			},
			checkResults: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name:    "Wrong password",
			payload: loginRequest{Username: usr.Username, Password: util.RandomString(8)},
			buildStubs: func(ms *mock_db.MockStore, mt *mock_token_mkr.MockMaker) {
				ms.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Times(1).Return(usr, nil)
			},
			checkResults: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name:    "Error on get Session",
			payload: loginRequest{Username: usr.Username, Password: pass},
			buildStubs: func(ms *mock_db.MockStore, mt *mock_token_mkr.MockMaker) {
				ms.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Times(1).Return(usr, nil)
				ms.EXPECT().GetSessions(gomock.Any(), gomock.Any()).Times(1).Return(database.Session{}, sql.ErrConnDone)
				mt.EXPECT().CreateToken(gomock.Any(), gomock.Any()).Times(1).Return("sometoken", &token.PASETOPayload{
					Payload: token.Payload{
						ID:       usr.ID,
						Username: usr.Username,
					},
					ExpiresAt: time.Now().Add(1 * time.Minute),
					IssuedAt:  time.Now(),
				}, nil)
			},
			checkResults: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:    "Error on Create Token ",
			payload: loginRequest{Username: usr.Username, Password: pass},
			buildStubs: func(ms *mock_db.MockStore, mt *mock_token_mkr.MockMaker) {
				ms.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Times(1).Return(usr, nil)
				mt.EXPECT().CreateToken(gomock.Any(), gomock.Any()).Times(1).Return("", &token.PASETOPayload{}, fmt.Errorf("artificial error"))
			},
			checkResults: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mock_db.NewMockStore(ctrl)
			mkr := mock_token_mkr.NewMockMaker(ctrl)
			tc.buildStubs(store, mkr)

			// TODO: Create a NewHTTPTestServer to get a stubbed database and placeholders for the rest
			config, err := util.LoadConfig("..")
			require.NoError(t, err)

			parser := parser.NewMdParser()                                  // TODO: Create a test MdParser (stub) (not being used)
			fileManager := files.NewLocalFileManager(config.UploadFilePath) // TODO: Create a test file manager (stub) (not being used)

			server := NewServer(store, parser, config, mkr, fileManager)
			server.setRoutes()
			server.setAuthorizedRoutes()
			recorder := httptest.NewRecorder()
			url := "/api/login"

			data, err := json.Marshal(tc.payload)
			require.NoError(t, err)
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)

			tc.checkResults(t, recorder)
		})
	}
}

// Forging a bad request
// type errReader int
//
// func (errReader) Read(p []byte) (n int, err error) {
//     return 0, errors.New("test error")
// }
// request, err := http.NewRequest(http.MethodPost, url, errReader(0))
// testRequest := httptest.NewRequest(http.MethodPost, "/something", errReader(0))

type resolveSessionPayload struct {
	ID       uuid.UUID
	username string
}

type resolveSessionTestCase struct {
	name         string
	Payload      resolveSessionPayload
	buildStubs   func(*mock_db.MockStore, *mock_token_mkr.MockMaker)
	checkResults func(t *testing.T, token string, exp time.Time, err error)
}

func TestResolveSession(t *testing.T) {
	usr, _, err := database.RandomUser()
	require.NoError(t, err)
	expiredSession := database.Session{
		ID:           uuid.New(),
		UserID:       usr.ID,
		RefreshToken: "someRefreshToken",
		ClientAgent:  "",
		ClientIp:     "",
		IsBlocked:    false,
		ExpiresAt:    time.Now().Add(-1 * time.Minute),
		CreatedAt:    time.Now(),
	}

	testCases := []resolveSessionTestCase{
		{
			name: "CreateTokenError_1",
			Payload: resolveSessionPayload{
				ID:       usr.ID,
				username: usr.Username,
			},
			buildStubs: func(ms *mock_db.MockStore, mt *mock_token_mkr.MockMaker) {
				ms.EXPECT().GetSessions(gomock.Any(), gomock.Any()).Times(1).Return(database.Session{}, sql.ErrNoRows)
				mt.EXPECT().CreateToken(gomock.Any(), gomock.Any()).Times(1).Return("", &token.PASETOPayload{}, fmt.Errorf("artificial error"))
			},
			checkResults: func(t *testing.T, token string, exp time.Time, err error) {
				require.Error(t, err)
			},
		},
		{
			name: "CreateTokenError_2",
			Payload: resolveSessionPayload{
				ID:       usr.ID,
				username: usr.Username,
			},
			buildStubs: func(ms *mock_db.MockStore, mt *mock_token_mkr.MockMaker) {
				ms.EXPECT().GetSessions(gomock.Any(), gomock.Any()).Times(1).Return(expiredSession, nil)
				mt.EXPECT().CreateToken(gomock.Any(), gomock.Any()).Times(1).Return("", &token.PASETOPayload{}, fmt.Errorf("artificial error"))
			},
			checkResults: func(t *testing.T, token string, exp time.Time, err error) {
				require.Error(t, err)
			},
		},
		{
			name: "Create session error",
			Payload: resolveSessionPayload{
				ID:       usr.ID,
				username: usr.Username,
			},
			buildStubs: func(ms *mock_db.MockStore, mt *mock_token_mkr.MockMaker) {
				ms.EXPECT().GetSessions(gomock.Any(), gomock.Any()).Times(1).Return(database.Session{}, sql.ErrNoRows)
				mt.EXPECT().CreateToken(gomock.Any(), gomock.Any()).Times(1).Return("sometoken", &token.PASETOPayload{
					Payload: token.Payload{
						ID:       usr.ID,
						Username: usr.Username,
					},
					ExpiresAt: time.Now().Add(1 * time.Minute),
					IssuedAt:  time.Now(),
				}, nil)
				ms.EXPECT().CreateSession(gomock.Any(), gomock.Any()).Times(1).Return(database.Session{}, fmt.Errorf("artificial error"))
			},
			checkResults: func(t *testing.T, token string, exp time.Time, err error) {
				require.Error(t, err)
			},
		},
		{
			name: "UpdateSessionError",
			Payload: resolveSessionPayload{
				ID:       usr.ID,
				username: usr.Username,
			},
			buildStubs: func(ms *mock_db.MockStore, mt *mock_token_mkr.MockMaker) {
				ms.EXPECT().GetSessions(gomock.Any(), gomock.Any()).Times(1).Return(expiredSession, nil)
				mt.EXPECT().CreateToken(gomock.Any(), gomock.Any()).Times(1).Return("sometoken", &token.PASETOPayload{
					Payload: token.Payload{
						ID:       usr.ID,
						Username: usr.Username,
					},
					ExpiresAt: time.Now().Add(1 * time.Minute),
					IssuedAt:  time.Now(),
				}, nil)
				ms.EXPECT().UpdateSession(gomock.Any(), gomock.Any()).Times(1).Return(database.Session{}, fmt.Errorf("artificial error"))
			},
			checkResults: func(t *testing.T, token string, exp time.Time, err error) {
				require.Error(t, err)
			},
		},
	}

	for _, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		store := mock_db.NewMockStore(ctrl)
		mkr := mock_token_mkr.NewMockMaker(ctrl)
		tc.buildStubs(store, mkr)
		token, exp, err := ResolveSession(context.Background(), store, mkr, tc.Payload.ID, tc.Payload.username)
		tc.checkResults(t, token, exp, err)
	}
}
