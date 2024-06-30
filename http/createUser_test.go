package http

import (
	"net/http/httptest"
	"testing"

	mock_db "github.com/julianinsua/codis/db/mock"
	"github.com/stretchr/testify/require"
)

type createUserTestCase struct {
	name         string
	payload      signupRequest
	buildStubs   func(*mock_db.MockStore)
	checkResults func(t *testing.T, recorder *httptest.ResponseRecorder)
}

func createUserTest(t *testing.T) {
	// Your test code here
	require.False(t, false)
}
