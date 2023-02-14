// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/adetunjii/netflakes/port (interfaces: MovieApi)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	model "github.com/adetunjii/netflakes/model"
	gomock "github.com/golang/mock/gomock"
)

// MockMovieApi is a mock of MovieApi interface.
type MockMovieApi struct {
	ctrl     *gomock.Controller
	recorder *MockMovieApiMockRecorder
}

// MockMovieApiMockRecorder is the mock recorder for MockMovieApi.
type MockMovieApiMockRecorder struct {
	mock *MockMovieApi
}

// NewMockMovieApi creates a new mock instance.
func NewMockMovieApi(ctrl *gomock.Controller) *MockMovieApi {
	mock := &MockMovieApi{ctrl: ctrl}
	mock.recorder = &MockMovieApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMovieApi) EXPECT() *MockMovieApiMockRecorder {
	return m.recorder
}

// FetchMovie mocks base method.
func (m *MockMovieApi) FetchMovie(arg0 context.Context, arg1 int64) (*model.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchMovie", arg0, arg1)
	ret0, _ := ret[0].(*model.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchMovie indicates an expected call of FetchMovie.
func (mr *MockMovieApiMockRecorder) FetchMovie(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchMovie", reflect.TypeOf((*MockMovieApi)(nil).FetchMovie), arg0, arg1)
}

// FetchMovieCharacters mocks base method.
func (m *MockMovieApi) FetchMovieCharacters(arg0 context.Context, arg1 int64) ([]model.Character, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchMovieCharacters", arg0, arg1)
	ret0, _ := ret[0].([]model.Character)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchMovieCharacters indicates an expected call of FetchMovieCharacters.
func (mr *MockMovieApiMockRecorder) FetchMovieCharacters(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchMovieCharacters", reflect.TypeOf((*MockMovieApi)(nil).FetchMovieCharacters), arg0, arg1)
}

// FetchMovies mocks base method.
func (m *MockMovieApi) FetchMovies(arg0 context.Context) ([]model.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchMovies", arg0)
	ret0, _ := ret[0].([]model.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchMovies indicates an expected call of FetchMovies.
func (mr *MockMovieApiMockRecorder) FetchMovies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchMovies", reflect.TypeOf((*MockMovieApi)(nil).FetchMovies), arg0)
}

// GetCharacter mocks base method.
func (m *MockMovieApi) GetCharacter(arg0 context.Context, arg1 string) (*model.Character, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharacter", arg0, arg1)
	ret0, _ := ret[0].(*model.Character)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCharacter indicates an expected call of GetCharacter.
func (mr *MockMovieApiMockRecorder) GetCharacter(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharacter", reflect.TypeOf((*MockMovieApi)(nil).GetCharacter), arg0, arg1)
}
