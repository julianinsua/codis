package http

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	mock_db "github.com/julianinsua/codis/db/mock"
	"github.com/julianinsua/codis/files"
	"github.com/julianinsua/codis/internal/database"
	"github.com/julianinsua/codis/parser"
	mock_token_mkr "github.com/julianinsua/codis/token/mock"
	"github.com/julianinsua/codis/util"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

type createUserTestCase struct {
	name         string
	payload      signupRequest
	buildStubs   func(*mock_db.MockStore, *mock_token_mkr.MockMaker)
	checkResults func(t *testing.T, recorder *httptest.ResponseRecorder)
}

func TestCreateUser(t *testing.T) {
	suReq := signupRequest{
		Username: util.RandomUsername(),
		Email:    util.RandomEmail(),
		Password: util.RandomPassword(),
	}
	testCases := []createUserTestCase{
		{
			name:    "OK",
			payload: suReq,
			buildStubs: func(ms *mock_db.MockStore, mm *mock_token_mkr.MockMaker) {
				ms.EXPECT().GetUsersByemailOrUsername(gomock.Any(), gomock.Any()).Times(1).Return([]database.User{}, sql.ErrNoRows)
				ms.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Times(1).Return(database.User{
					ID:       uuid.New(),
					Username: suReq.Username,
					Password: suReq.Password,
					Email:    suReq.Email,
					CreatedAt: sql.NullTime{
						Time:  time.Now(),
						Valid: true,
					},
					UpdatedAt: sql.NullTime{
						Time:  time.Now(),
						Valid: true,
					},
				}, nil)
			},
			checkResults: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "invalid password",
			payload: signupRequest{
				Username: util.RandomUsername(),
				Email:    util.RandomEmail(),
				Password: util.RandomString(4),
			},
			buildStubs: func(ms *mock_db.MockStore, mm *mock_token_mkr.MockMaker) {

			},
			checkResults: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "invalid email",
			payload: signupRequest{
				Username: util.RandomUsername(),
				Email:    util.RandomString(6),
				Password: util.RandomPassword(),
			},
			buildStubs: func(ms *mock_db.MockStore, mm *mock_token_mkr.MockMaker) {

			},
			checkResults: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "invalid username",
			payload: signupRequest{
				Username: util.RandomString(2),
				Email:    util.RandomEmail(),
				Password: util.RandomPassword(),
			},
			buildStubs: func(ms *mock_db.MockStore, mm *mock_token_mkr.MockMaker) {

			},
			checkResults: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:    "error checking users",
			payload: suReq,
			buildStubs: func(ms *mock_db.MockStore, mm *mock_token_mkr.MockMaker) {
				ms.EXPECT().GetUsersByemailOrUsername(gomock.Any(), gomock.Any()).Times(1).Return([]database.User{}, sql.ErrConnDone)
			},
			checkResults: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:    "user already exists",
			payload: suReq,
			buildStubs: func(ms *mock_db.MockStore, mm *mock_token_mkr.MockMaker) {
				ms.EXPECT().GetUsersByemailOrUsername(gomock.Any(), gomock.Any()).Times(1).Return([]database.User{{
					ID:       uuid.New(),
					Username: suReq.Username,
					Password: suReq.Password,
					Email:    suReq.Email,
					CreatedAt: sql.NullTime{
						Time:  time.Now(),
						Valid: true,
					},
					UpdatedAt: sql.NullTime{
						Time:  time.Now(),
						Valid: true,
					},
				}}, nil)
			},
			checkResults: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
			},
		},
		{
			name:    "error creatin participant",
			payload: suReq,
			buildStubs: func(ms *mock_db.MockStore, mm *mock_token_mkr.MockMaker) {
				ms.EXPECT().GetUsersByemailOrUsername(gomock.Any(), gomock.Any()).Times(1).Return([]database.User{}, sql.ErrNoRows)
				ms.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Times(1).Return(database.User{}, sql.ErrConnDone)
			},
			checkResults: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		store := mock_db.NewMockStore(ctrl)
		mkr := mock_token_mkr.NewMockMaker(ctrl)
		tc.buildStubs(store, mkr)

		config, err := util.LoadConfig("..")
		require.NoError(t, err)

		parser := parser.NewMdParser()
		fileManager := files.NewLocalFileManager(config.UploadFilePath)
		server := NewServer(store, parser, config, mkr, fileManager)
		server.setRoutes()
		server.setAuthorizedRoutes()
		recorder := httptest.NewRecorder()
		url := "/api/signup"

		data, err := json.Marshal(tc.payload)
		require.NoError(t, err)
		request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
		require.NoError(t, err)
		server.router.ServeHTTP(recorder, request)
		tc.checkResults(t, recorder)
	}
}
