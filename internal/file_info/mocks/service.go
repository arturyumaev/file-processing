// Code generated by MockGen. DO NOT EDIT.
// Source: file_info.go

// Package mock_handler is a generated GoMock package.
package mock_handler

import (
	context "context"
	multipart "mime/multipart"
	http "net/http"
	reflect "reflect"

	file_info "github.com/arturyumaev/file-processing/internal/file_info"
	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// GetFileInfo mocks base method.
func (m *MockService) GetFileInfo(ctx context.Context, name string) (*file_info.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileInfo", ctx, name)
	ret0, _ := ret[0].(*file_info.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileInfo indicates an expected call of GetFileInfo.
func (mr *MockServiceMockRecorder) GetFileInfo(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileInfo", reflect.TypeOf((*MockService)(nil).GetFileInfo), ctx, name)
}

// UploadFile mocks base method.
func (m *MockService) UploadFile(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) (*file_info.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFile", ctx, file, fileHeader)
	ret0, _ := ret[0].(*file_info.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadFile indicates an expected call of UploadFile.
func (mr *MockServiceMockRecorder) UploadFile(ctx, file, fileHeader interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFile", reflect.TypeOf((*MockService)(nil).UploadFile), ctx, file, fileHeader)
}

// MockHandler is a mock of Handler interface.
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler.
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance.
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// GetFileInfo mocks base method.
func (m *MockHandler) GetFileInfo(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetFileInfo", w, r)
}

// GetFileInfo indicates an expected call of GetFileInfo.
func (mr *MockHandlerMockRecorder) GetFileInfo(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileInfo", reflect.TypeOf((*MockHandler)(nil).GetFileInfo), w, r)
}

// PostFile mocks base method.
func (m *MockHandler) PostFile(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PostFile", w, r)
}

// PostFile indicates an expected call of PostFile.
func (mr *MockHandlerMockRecorder) PostFile(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostFile", reflect.TypeOf((*MockHandler)(nil).PostFile), w, r)
}
