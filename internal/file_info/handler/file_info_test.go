package handler

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/arturyumaev/file-processing/internal/file_info"
	mock_file_info "github.com/arturyumaev/file-processing/internal/file_info/mocks"
)

const (
	TEMP_FILE_NAME = "temp_file"
)

func generatePostRequestWithFile(t *testing.T, fieldname string) (*http.Request, error) {
	body := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(body)

	fileWriter, err := bodyWriter.CreateFormFile(fieldname, TEMP_FILE_NAME)
	if err != nil {
		return nil, err
	}

	emptyFileContent := strings.NewReader("")

	_, err = io.Copy(fileWriter, emptyFileContent)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	contentType := bodyWriter.FormDataContentType()
	err = bodyWriter.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, "/files", body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", contentType)

	return req, nil
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
			h := New(ms)

			response := httptest.NewRecorder()
			req := tc.imitateRequest(ms)
			h.GetFileInfo(response, req)

			assert.Equal(t, tc.expectedStatus, response.Code)
			assert.JSONEq(t, tc.expectedBody, response.Body.String())
		})
	}
}

func TestHandler_postFile(t *testing.T) {
	ctx := context.Background()

	t.Run("when method not allowed", func(t *testing.T) {
		expectedStatus := http.StatusMethodNotAllowed
		expectedBody := fmt.Sprintf(`{"error":"%s"}`, file_info.ErrMethodNotAllowed)

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		ms := mock_file_info.NewMockService(ctrl)

		h := New(ms)

		rw := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/files", nil)
		if err != nil {
			t.Error(err)
		}

		h.PostFile(rw, req)

		actualBody, err := io.ReadAll(rw.Body)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, expectedStatus, rw.Code)
		assert.Equal(t, expectedBody, string(actualBody))
	})

	t.Run("when error while uploading file", func(t *testing.T) {
		expectedStatus := http.StatusBadRequest
		expectedBody := fmt.Sprintf(`{"error":"%s"}`, file_info.ErrRetrievingFile)

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		ms := mock_file_info.NewMockService(ctrl)

		h := New(ms)

		rw := httptest.NewRecorder()
		req, err := generatePostRequestWithFile(t, "wrong_field_name")
		if err != nil {
			t.Log(err)
			t.Error(err)
		}

		h.PostFile(rw, req)

		actualBody, err := io.ReadAll(rw.Body)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, expectedStatus, rw.Code)
		assert.Equal(t, expectedBody, string(actualBody))
	})

	t.Run("when service responded ok", func(t *testing.T) {
		expectedStatus := http.StatusOK
		expectedBody := `{"id":1,"filename":"temp_file","status":"recieved","timestamp":""}`

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		ms := mock_file_info.NewMockService(ctrl)

		h := New(ms)

		rw := httptest.NewRecorder()

		fieldname := FORM_FIELD_FILE_NAME

		req, err := generatePostRequestWithFile(t, fieldname)
		if err != nil {
			t.Error(err)
		}

		file, fileHeader, _ := req.FormFile(fieldname)
		defer file.Close()

		fileInfo := &file_info.FileInfo{
			Id:       1,
			Filename: TEMP_FILE_NAME,
			Status:   file_info.StatusRecieved,
		}

		ms.EXPECT().UploadFile(ctx, file, fileHeader).Return(fileInfo, nil)

		h.PostFile(rw, req)

		actualBody, err := io.ReadAll(rw.Body)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, expectedStatus, rw.Code)
		assert.Equal(t, expectedBody, string(actualBody))
	})

	t.Run("when service call failed", func(t *testing.T) {
		expectedStatus := http.StatusInternalServerError
		expectedBody := `{"error":"unexpected error"}`

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		ms := mock_file_info.NewMockService(ctrl)

		h := New(ms)

		rw := httptest.NewRecorder()
		req, err := generatePostRequestWithFile(t, FORM_FIELD_FILE_NAME)
		if err != nil {
			t.Error(err)
		}

		file, fileHeader, _ := req.FormFile(FORM_FIELD_FILE_NAME)
		defer file.Close()

		ms.EXPECT().UploadFile(ctx, file, fileHeader).Return(nil, errors.New("unexpected error"))

		h.PostFile(rw, req)

		actualBody, err := io.ReadAll(rw.Body)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, expectedStatus, rw.Code)
		assert.Equal(t, expectedBody, string(actualBody))
	})
}
