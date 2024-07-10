package http

import (
	"net/http/httptest"
	"testing"

	mock_db "github.com/julianinsua/codis/db/mock"
	mock_token_mkr "github.com/julianinsua/codis/token/mock"
	"github.com/stretchr/testify/require"
)

type userContentTestCase struct {
	name         string
	Username     string
	buildStubs   func(ms mock_db.MockStore, mm mock_token_mkr.MockMaker)
	checkResults func(t testing.T, recorder httptest.ResponseRecorder)
}

func TestHandleUserContent(t *testing.T) {
	// Your test code here
	require.True(t, true)
}
