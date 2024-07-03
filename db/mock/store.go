// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/julianinsua/codis/internal/database (interfaces: Store)
//
// Generated by this command:
//
//	mockgen --build_flags=--mod=mod -destination db/mock/store.go -package mock_db github.com/julianinsua/codis/internal/database Store
//

// Package mock_db is a generated GoMock package.
package mock_db

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	uuid "github.com/google/uuid"
	database "github.com/julianinsua/codis/internal/database"
	gomock "go.uber.org/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreatePost mocks base method.
func (m *MockStore) CreatePost(arg0 context.Context, arg1 database.CreatePostParams) (database.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", arg0, arg1)
	ret0, _ := ret[0].(database.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockStoreMockRecorder) CreatePost(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockStore)(nil).CreatePost), arg0, arg1)
}

// CreatePostTag mocks base method.
func (m *MockStore) CreatePostTag(arg0 context.Context, arg1 database.CreatePostTagParams) (database.PostTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePostTag", arg0, arg1)
	ret0, _ := ret[0].(database.PostTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePostTag indicates an expected call of CreatePostTag.
func (mr *MockStoreMockRecorder) CreatePostTag(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePostTag", reflect.TypeOf((*MockStore)(nil).CreatePostTag), arg0, arg1)
}

// CreatePostTx mocks base method.
func (m *MockStore) CreatePostTx(arg0 context.Context, arg1 database.CreatePostTxParams) (database.CreatePostTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePostTx", arg0, arg1)
	ret0, _ := ret[0].(database.CreatePostTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePostTx indicates an expected call of CreatePostTx.
func (mr *MockStoreMockRecorder) CreatePostTx(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePostTx", reflect.TypeOf((*MockStore)(nil).CreatePostTx), arg0, arg1)
}

// CreateSession mocks base method.
func (m *MockStore) CreateSession(arg0 context.Context, arg1 database.CreateSessionParams) (database.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", arg0, arg1)
	ret0, _ := ret[0].(database.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockStoreMockRecorder) CreateSession(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockStore)(nil).CreateSession), arg0, arg1)
}

// CreateTag mocks base method.
func (m *MockStore) CreateTag(arg0 context.Context, arg1 database.CreateTagParams) (database.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTag", arg0, arg1)
	ret0, _ := ret[0].(database.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTag indicates an expected call of CreateTag.
func (mr *MockStoreMockRecorder) CreateTag(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTag", reflect.TypeOf((*MockStore)(nil).CreateTag), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 database.CreateUserParams) (database.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(database.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// GetPost mocks base method.
func (m *MockStore) GetPost(arg0 context.Context, arg1 uuid.UUID) (database.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPost", arg0, arg1)
	ret0, _ := ret[0].(database.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPost indicates an expected call of GetPost.
func (mr *MockStoreMockRecorder) GetPost(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPost", reflect.TypeOf((*MockStore)(nil).GetPost), arg0, arg1)
}

// GetPosts mocks base method.
func (m *MockStore) GetPosts(arg0 context.Context, arg1 sql.NullString) ([]database.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPosts", arg0, arg1)
	ret0, _ := ret[0].([]database.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPosts indicates an expected call of GetPosts.
func (mr *MockStoreMockRecorder) GetPosts(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPosts", reflect.TypeOf((*MockStore)(nil).GetPosts), arg0, arg1)
}

// GetPostsWithTags mocks base method.
func (m *MockStore) GetPostsWithTags(arg0 context.Context, arg1 uuid.UUID) ([]database.PostsView, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostsWithTags", arg0, arg1)
	ret0, _ := ret[0].([]database.PostsView)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostsWithTags indicates an expected call of GetPostsWithTags.
func (mr *MockStoreMockRecorder) GetPostsWithTags(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostsWithTags", reflect.TypeOf((*MockStore)(nil).GetPostsWithTags), arg0, arg1)
}

// GetSessions mocks base method.
func (m *MockStore) GetSessions(arg0 context.Context, arg1 uuid.UUID) (database.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessions", arg0, arg1)
	ret0, _ := ret[0].(database.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessions indicates an expected call of GetSessions.
func (mr *MockStoreMockRecorder) GetSessions(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessions", reflect.TypeOf((*MockStore)(nil).GetSessions), arg0, arg1)
}

// GetTagById mocks base method.
func (m *MockStore) GetTagById(arg0 context.Context, arg1 uuid.UUID) (database.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTagById", arg0, arg1)
	ret0, _ := ret[0].(database.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagById indicates an expected call of GetTagById.
func (mr *MockStoreMockRecorder) GetTagById(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagById", reflect.TypeOf((*MockStore)(nil).GetTagById), arg0, arg1)
}

// GetTagPosts mocks base method.
func (m *MockStore) GetTagPosts(arg0 context.Context, arg1 uuid.UUID) ([]database.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTagPosts", arg0, arg1)
	ret0, _ := ret[0].([]database.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagPosts indicates an expected call of GetTagPosts.
func (mr *MockStoreMockRecorder) GetTagPosts(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagPosts", reflect.TypeOf((*MockStore)(nil).GetTagPosts), arg0, arg1)
}

// GetUserByEmail mocks base method.
func (m *MockStore) GetUserByEmail(arg0 context.Context, arg1 string) (database.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", arg0, arg1)
	ret0, _ := ret[0].(database.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockStoreMockRecorder) GetUserByEmail(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockStore)(nil).GetUserByEmail), arg0, arg1)
}

// GetUserByID mocks base method.
func (m *MockStore) GetUserByID(arg0 context.Context, arg1 uuid.UUID) (database.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", arg0, arg1)
	ret0, _ := ret[0].(database.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockStoreMockRecorder) GetUserByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockStore)(nil).GetUserByID), arg0, arg1)
}

// GetUserTags mocks base method.
func (m *MockStore) GetUserTags(arg0 context.Context, arg1 uuid.UUID) ([]database.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserTags", arg0, arg1)
	ret0, _ := ret[0].([]database.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserTags indicates an expected call of GetUserTags.
func (mr *MockStoreMockRecorder) GetUserTags(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserTags", reflect.TypeOf((*MockStore)(nil).GetUserTags), arg0, arg1)
}

// GetUsersByemailOrUsername mocks base method.
func (m *MockStore) GetUsersByemailOrUsername(arg0 context.Context, arg1 database.GetUsersByemailOrUsernameParams) ([]database.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersByemailOrUsername", arg0, arg1)
	ret0, _ := ret[0].([]database.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersByemailOrUsername indicates an expected call of GetUsersByemailOrUsername.
func (mr *MockStoreMockRecorder) GetUsersByemailOrUsername(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersByemailOrUsername", reflect.TypeOf((*MockStore)(nil).GetUsersByemailOrUsername), arg0, arg1)
}

// UpdateSession mocks base method.
func (m *MockStore) UpdateSession(arg0 context.Context, arg1 database.UpdateSessionParams) (database.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSession", arg0, arg1)
	ret0, _ := ret[0].(database.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSession indicates an expected call of UpdateSession.
func (mr *MockStoreMockRecorder) UpdateSession(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSession", reflect.TypeOf((*MockStore)(nil).UpdateSession), arg0, arg1)
}

// UserTagExists mocks base method.
func (m *MockStore) UserTagExists(arg0 context.Context, arg1 database.UserTagExistsParams) (database.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserTagExists", arg0, arg1)
	ret0, _ := ret[0].(database.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserTagExists indicates an expected call of UserTagExists.
func (mr *MockStoreMockRecorder) UserTagExists(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserTagExists", reflect.TypeOf((*MockStore)(nil).UserTagExists), arg0, arg1)
}
