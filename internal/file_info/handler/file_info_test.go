package handler

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/arturyumaev/file-processing/internal/file_info"
	mock_file_info "github.com/arturyumaev/file-processing/internal/file_info/mocks"
)

func getTestRouter() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}

func TestHandler_getFileInfo(t *testing.T) {
	tt := []struct {
		name           string
		expectedStatus int
		expectedBody   string
		imitateRequest func(*mock_file_info.MockService) *http.Request
	}{
		{
			name:           "when method is invalid",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   fmt.Sprintf(`{"error":"%s"}`, file_info.ErrMethodNotAllowed),
			imitateRequest: func(ms *mock_file_info.MockService) *http.Request {
				req, _ := http.NewRequest(http.MethodPost, "/files/file1", nil)
				return req
			},
		},
		{
			name:           "when wrong path",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   fmt.Sprintf(`{"error":"%s"}`, file_info.ErrEmptyParameterName),
			imitateRequest: func(ms *mock_file_info.MockService) *http.Request {
				req, _ := http.NewRequest(http.MethodGet, "/files/", nil)
				return req
			},
		},
		{
			name:           "when file with name 'file1' is available",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":0,"filename":"file1","status":"done","timestamp":""}`,
			imitateRequest: func(ms *mock_file_info.MockService) *http.Request {
				ctx := context.Background()

				fileInfo := &file_info.FileInfo{
					Filename: "file1",
					Status:   "done",
				}

				ms.EXPECT().GetFileInfo(ctx, "file1").Return(fileInfo, nil)
				req, _ := http.NewRequest(http.MethodGet, "/files/file1", nil)

				return req
			},
		},
		{
			name:           "when no such file",
			expectedStatus: http.StatusNotFound,
			expectedBody:   fmt.Sprintf(`{"error":"%s"}`, file_info.ErrNoSuchFile),
			imitateRequest: func(ms *mock_file_info.MockService) *http.Request {
				ctx := context.Background()
				ms.EXPECT().GetFileInfo(ctx, "file4").Return(nil, file_info.ErrNoSuchFile)
				req, _ := http.NewRequest(http.MethodGet, "/files/file4", nil)

				return req
			},
		},
		{
			name:           "when services return unknown error",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   fmt.Sprintf(`{"error":"%s"}`, errors.New("unknown sql error")),
			imitateRequest: func(ms *mock_file_info.MockService) *http.Request {
				ctx := context.Background()
				ms.EXPECT().GetFileInfo(ctx, "file4").Return(nil, errors.New("unknown sql error"))
				req, _ := http.NewRequest(http.MethodGet, "/files/file4", nil)

				return req
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			ms := mock_file_info.NewMockService(ctrl)
			router := getTestRouter()
			RegisterHandlers(router, ms)

			response := httptest.NewRecorder()
			req := tc.imitateRequest(ms)
			router.ServeHTTP(response, req)

			assert.Equal(t, tc.expectedStatus, response.Code)
			assert.JSONEq(t, tc.expectedBody, response.Body.String())
		})
	}
}

func TestHandler_postFile(t *testing.T) {
	tt := []struct {
		name           string
		expectedStatus int
		expectedBody   string
		imitateRequest func(*testing.T, *mock_file_info.MockService) (*http.Request, error)
	}{
		{
			name:           "when method is not allowed",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   fmt.Sprintf(`{"error":"%s"}`, file_info.ErrMethodNotAllowed),
			imitateRequest: func(t *testing.T, ms *mock_file_info.MockService) (*http.Request, error) {
				return http.NewRequest(http.MethodGet, "/files", nil)
			},
		},
		{
			name:           "when error while uploading file",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   fmt.Sprintf(`{"error":"%s"}`, file_info.ErrRetrievingFile),
			imitateRequest: func(t *testing.T, ms *mock_file_info.MockService) (*http.Request, error) {
				fileMock := &FormFileMock{fieldname: "wrong_field_name"}
				body, contentType := fileMock.Generate(t)
				req, err := http.NewRequest(http.MethodPost, "/files", body)
				req.Header.Add("Content-Type", contentType)

				return req, err
			},
		},
		{
			name:           "when service responded ok",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":1,"filename":"temp_file","status":"recieved","timestamp":""}`,
			imitateRequest: func(t *testing.T, ms *mock_file_info.MockService) (*http.Request, error) {
				ctx := context.Background()

				filename, fieldname := "temp_file", FORM_FIELD_FILE_NAME

				fileMock := &FormFileMock{filename: filename}
				body, contentType := fileMock.Generate(t)
				req, err := http.NewRequest(http.MethodPost, "/files", body)
				req.Header.Add("Content-Type", contentType)

				file, fileHeader, _ := req.FormFile(fieldname)
				defer file.Close()

				fileInfo := &file_info.FileInfo{
					Id:       1,
					Filename: filename,
					Status:   file_info.StatusRecieved,
				}

				ms.EXPECT().UploadFile(ctx, file, fileHeader).Return(fileInfo, nil)

				return req, err
			},
		},
		{
			name:           "when service call failed",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   fmt.Sprintf(`{"error":"%s"}`, errors.New("unexpected error")),
			imitateRequest: func(t *testing.T, ms *mock_file_info.MockService) (*http.Request, error) {
				ctx := context.Background()

				filename, fieldname := "temp_file", FORM_FIELD_FILE_NAME

				fileMock := &FormFileMock{filename: filename}
				body, contentType := fileMock.Generate(t)
				req, err := http.NewRequest(http.MethodPost, "/files", body)
				req.Header.Add("Content-Type", contentType)

				file, fileHeader, _ := req.FormFile(fieldname)
				defer file.Close()

				ms.EXPECT().UploadFile(ctx, file, fileHeader).Return(nil, errors.New("unexpected error"))

				return req, err
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			ms := mock_file_info.NewMockService(ctrl)

			router := getTestRouter()
			RegisterHandlers(router, ms)

			response := httptest.NewRecorder()
			req, err := tc.imitateRequest(t, ms)
			if err != nil {
				t.Fatal(err)
			}
			router.ServeHTTP(response, req)

			actualBody, _ := io.ReadAll(response.Body)

			assert.Equal(t, tc.expectedStatus, response.Code)
			assert.Equal(t, tc.expectedBody, string(actualBody))
		})
	}
}
