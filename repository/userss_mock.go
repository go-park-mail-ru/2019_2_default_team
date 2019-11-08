// Code generated by MockGen. DO NOT EDIT.
// Source: kino_backend/repository (interfaces: UsersRepository)

// Package repository is a generated GoMock package.
package repository

import (
	gomock "github.com/golang/mock/gomock"
	models "kino_backend/models"
	reflect "reflect"
)

// MockUsersRepository is a mock of UsersRepository interface
type MockUsersRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUsersRepositoryMockRecorder
}

// MockUsersRepositoryMockRecorder is the mock recorder for MockUsersRepository
type MockUsersRepositoryMockRecorder struct {
	mock *MockUsersRepository
}

// NewMockUsersRepository creates a new mock instance
func NewMockUsersRepository(ctrl *gomock.Controller) *MockUsersRepository {
	mock := &MockUsersRepository{ctrl: ctrl}
	mock.recorder = &MockUsersRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsersRepository) EXPECT() *MockUsersRepositoryMockRecorder {
	return m.recorder
}

// CheckExistenceOfEmail mocks base method
func (m *MockUsersRepository) CheckExistenceOfEmail(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckExistenceOfEmail", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckExistenceOfEmail indicates an expected call of CheckExistenceOfEmail
func (mr *MockUsersRepositoryMockRecorder) CheckExistenceOfEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckExistenceOfEmail", reflect.TypeOf((*MockUsersRepository)(nil).CheckExistenceOfEmail), arg0)
}

// CheckExistenceOfNickname mocks base method
func (m *MockUsersRepository) CheckExistenceOfNickname(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckExistenceOfNickname", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckExistenceOfNickname indicates an expected call of CheckExistenceOfNickname
func (mr *MockUsersRepositoryMockRecorder) CheckExistenceOfNickname(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckExistenceOfNickname", reflect.TypeOf((*MockUsersRepository)(nil).CheckExistenceOfNickname), arg0)
}

// CreateNewUser mocks base method
func (m *MockUsersRepository) CreateNewUser(arg0 *models.RegisterProfile) (models.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewUser", arg0)
	ret0, _ := ret[0].(models.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewUser indicates an expected call of CreateNewUser
func (mr *MockUsersRepositoryMockRecorder) CreateNewUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewUser", reflect.TypeOf((*MockUsersRepository)(nil).CreateNewUser), arg0)
}

// DeleteAvatar mocks base method
func (m *MockUsersRepository) DeleteAvatar(arg0 uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAvatar", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAvatar indicates an expected call of DeleteAvatar
func (mr *MockUsersRepositoryMockRecorder) DeleteAvatar(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAvatar", reflect.TypeOf((*MockUsersRepository)(nil).DeleteAvatar), arg0)
}

// GetCountOfUsers mocks base method
func (m *MockUsersRepository) GetCountOfUsers() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCountOfUsers")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCountOfUsers indicates an expected call of GetCountOfUsers
func (mr *MockUsersRepositoryMockRecorder) GetCountOfUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCountOfUsers", reflect.TypeOf((*MockUsersRepository)(nil).GetCountOfUsers))
}

// GetUserPassword mocks base method
func (m *MockUsersRepository) GetUserPassword(arg0 string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserPassword", arg0)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserPassword indicates an expected call of GetUserPassword
func (mr *MockUsersRepositoryMockRecorder) GetUserPassword(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserPassword", reflect.TypeOf((*MockUsersRepository)(nil).GetUserPassword), arg0)
}

// GetUserProfileByID mocks base method
func (m *MockUsersRepository) GetUserProfileByID(arg0 uint) (models.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserProfileByID", arg0)
	ret0, _ := ret[0].(models.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserProfileByID indicates an expected call of GetUserProfileByID
func (mr *MockUsersRepositoryMockRecorder) GetUserProfileByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserProfileByID", reflect.TypeOf((*MockUsersRepository)(nil).GetUserProfileByID), arg0)
}

// GetUserProfileByNickname mocks base method
func (m *MockUsersRepository) GetUserProfileByNickname(arg0 string) (models.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserProfileByNickname", arg0)
	ret0, _ := ret[0].(models.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserProfileByNickname indicates an expected call of GetUserProfileByNickname
func (mr *MockUsersRepositoryMockRecorder) GetUserProfileByNickname(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserProfileByNickname", reflect.TypeOf((*MockUsersRepository)(nil).GetUserProfileByNickname), arg0)
}

// UpdateUserByID mocks base method
func (m *MockUsersRepository) UpdateUserByID(arg0 uint, arg1 *models.RegisterProfile) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserByID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserByID indicates an expected call of UpdateUserByID
func (mr *MockUsersRepositoryMockRecorder) UpdateUserByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserByID", reflect.TypeOf((*MockUsersRepository)(nil).UpdateUserByID), arg0, arg1)
}

// UploadAvatar mocks base method
func (m *MockUsersRepository) UploadAvatar(arg0 uint, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadAvatar", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadAvatar indicates an expected call of UploadAvatar
func (mr *MockUsersRepositoryMockRecorder) UploadAvatar(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadAvatar", reflect.TypeOf((*MockUsersRepository)(nil).UploadAvatar), arg0, arg1)
}
