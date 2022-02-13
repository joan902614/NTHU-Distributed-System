// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/NTHU-LSALAB/NTHU-Distributed-System/modules/video/dao (interfaces: VideoDAO)

// Package daomock is a generated GoMock package.
package daomock

import (
	context "context"
	reflect "reflect"

	dao "github.com/NTHU-LSALAB/NTHU-Distributed-System/modules/video/dao"
	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// MockVideoDAO is a mock of VideoDAO interface.
type MockVideoDAO struct {
	ctrl     *gomock.Controller
	recorder *MockVideoDAOMockRecorder
}

// MockVideoDAOMockRecorder is the mock recorder for MockVideoDAO.
type MockVideoDAOMockRecorder struct {
	mock *MockVideoDAO
}

// NewMockVideoDAO creates a new mock instance.
func NewMockVideoDAO(ctrl *gomock.Controller) *MockVideoDAO {
	mock := &MockVideoDAO{ctrl: ctrl}
	mock.recorder = &MockVideoDAOMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVideoDAO) EXPECT() *MockVideoDAOMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockVideoDAO) Create(arg0 context.Context, arg1 *dao.Video) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockVideoDAOMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockVideoDAO)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockVideoDAO) Delete(arg0 context.Context, arg1 primitive.ObjectID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockVideoDAOMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVideoDAO)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockVideoDAO) Get(arg0 context.Context, arg1 primitive.ObjectID) (*dao.Video, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*dao.Video)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockVideoDAOMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockVideoDAO)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockVideoDAO) List(arg0 context.Context, arg1, arg2 int64) ([]*dao.Video, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*dao.Video)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockVideoDAOMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVideoDAO)(nil).List), arg0, arg1, arg2)
}

// Update mocks base method.
func (m *MockVideoDAO) Update(arg0 context.Context, arg1 *dao.Video) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockVideoDAOMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockVideoDAO)(nil).Update), arg0, arg1)
}
