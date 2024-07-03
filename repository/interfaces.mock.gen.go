// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/SawitProRecruitment/UserService/repository (interfaces: RepositoryInterface)
//
// Generated by this command:
//
//	mockgen --build_flags=--mod=mod -destination=interfaces.mock.gen.go -package=repository . RepositoryInterface
//

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockRepositoryInterface is a mock of RepositoryInterface interface.
type MockRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryInterfaceMockRecorder
}

// MockRepositoryInterfaceMockRecorder is the mock recorder for MockRepositoryInterface.
type MockRepositoryInterfaceMockRecorder struct {
	mock *MockRepositoryInterface
}

// NewMockRepositoryInterface creates a new mock instance.
func NewMockRepositoryInterface(ctrl *gomock.Controller) *MockRepositoryInterface {
	mock := &MockRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryInterface) EXPECT() *MockRepositoryInterfaceMockRecorder {
	return m.recorder
}

// CreateEstate mocks base method.
func (m *MockRepositoryInterface) CreateEstate(arg0 context.Context, arg1, arg2 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEstate", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEstate indicates an expected call of CreateEstate.
func (mr *MockRepositoryInterfaceMockRecorder) CreateEstate(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEstate", reflect.TypeOf((*MockRepositoryInterface)(nil).CreateEstate), arg0, arg1, arg2)
}

// CreateTree mocks base method.
func (m *MockRepositoryInterface) CreateTree(arg0 context.Context, arg1 string, arg2 Tree) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTree", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTree indicates an expected call of CreateTree.
func (mr *MockRepositoryInterfaceMockRecorder) CreateTree(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTree", reflect.TypeOf((*MockRepositoryInterface)(nil).CreateTree), arg0, arg1, arg2)
}

// GetEstateByID mocks base method.
func (m *MockRepositoryInterface) GetEstateByID(arg0 context.Context, arg1 string) (Estate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEstateByID", arg0, arg1)
	ret0, _ := ret[0].(Estate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEstateByID indicates an expected call of GetEstateByID.
func (mr *MockRepositoryInterfaceMockRecorder) GetEstateByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEstateByID", reflect.TypeOf((*MockRepositoryInterface)(nil).GetEstateByID), arg0, arg1)
}

// GetTree mocks base method.
func (m *MockRepositoryInterface) GetTree(arg0 context.Context, arg1 string) ([]Tree, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTree", arg0, arg1)
	ret0, _ := ret[0].([]Tree)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTree indicates an expected call of GetTree.
func (mr *MockRepositoryInterfaceMockRecorder) GetTree(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTree", reflect.TypeOf((*MockRepositoryInterface)(nil).GetTree), arg0, arg1)
}
