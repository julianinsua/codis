package http

import (
	"testing"

	mock_db "github.com/julianinsua/codis/db/mock"
	"github.com/julianinsua/codis/files"
	"github.com/julianinsua/codis/parser"
	mock_token_mkr "github.com/julianinsua/codis/token/mock"
	"github.com/julianinsua/codis/util"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

// TODO: use the function configuration approach from the book "100 go mistakes", so that if there's no config
// argument you get the normal functionality. Otherwise you can pass a mock instance.
func NewTestServer(t *testing.T, store *mock_db.MockStore, mkr *mock_token_mkr.MockMaker) (*Server, error) {
	config, err := util.LoadConfig("..")
	if err != nil {
		return nil, err
	}

	parser := parser.NewMdParser() // TODO: Create a test MdParser (stub) (not being used)

	fileManager := files.NewLocalFileManager(config.UploadFilePath) // TODO: Create a test file manager (stub) (not being used)

	server := NewServer(store, parser, config, mkr, fileManager)
	server.setRoutes()
	server.setAuthorizedRoutes()
	return server, nil
}

func serverTest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mock_db.NewMockStore(ctrl)
	mkr := mock_token_mkr.NewMockMaker(ctrl)
	srv, err := NewTestServer(t, store, mkr)
	require.NoError(t, err)
	require.NotEmpty(t, srv)
}
