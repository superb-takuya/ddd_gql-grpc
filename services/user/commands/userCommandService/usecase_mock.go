// Code generated by MockGen. DO NOT EDIT.
// Source: commands/userCommandService/usecase.go

// Package usercommandservice is a generated GoMock package.
package usercommandservice

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUsecase is a mock of Usecase interface
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// create mocks base method
func (m *MockUsecase) create(req createRequest) (createResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "create", req)
	ret0, _ := ret[0].(createResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// create indicates an expected call of create
func (mr *MockUsecaseMockRecorder) create(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "create", reflect.TypeOf((*MockUsecase)(nil).create), req)
}

// update mocks base method
func (m *MockUsecase) update(req updateRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "update", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// update indicates an expected call of update
func (mr *MockUsecaseMockRecorder) update(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "update", reflect.TypeOf((*MockUsecase)(nil).update), req)
}

// delete mocks base method
func (m *MockUsecase) delete(req deleteRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "delete", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// delete indicates an expected call of delete
func (mr *MockUsecaseMockRecorder) delete(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "delete", reflect.TypeOf((*MockUsecase)(nil).delete), req)
}
