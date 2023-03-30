package handler

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/arturyumaev/file-processing/internal/file_info"
	mock_file_info "github.com/arturyumaev/file-processing/internal/file_info/mocks"
	"github.com/arturyumaev/file-processing/models"
)

func getRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.UseRawPath = true
	return router
}

func TestHandler_getFileInfo(t *testing.T) {
	t.Run("when file with name 'file1' is available", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		fileInfoServiceMock := mock_file_info.NewMockService(ctrl)

		fileInfo := &models.FileInfo{
			FilenameHash: "826e8142e6baabe8af779f5f490cf5f5",
			Status:       "done",
		}
		fileInfoServiceMock.EXPECT().GetFileInfo(ctx, "file1").Return(fileInfo, nil)

		handler := New(fileInfoServiceMock)

		router := getRouter()
		router.GET("/files/:name", handler.GetFileInfo)
		response := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/files/file1", nil)

		router.ServeHTTP(response, req)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(
			t,
			`{"id":"00000000-0000-0000-0000-000000000000","filename_hash":"826e8142e6baabe8af779f5f490cf5f5","status":"done","timestamp":""}`,
			response.Body.String(),
		)
	})

	t.Run("when no such file", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		fileInfoServiceMock := mock_file_info.NewMockService(ctrl)

		noSuchFileError := file_info.ErrNoSuchFile
		fileInfoServiceMock.EXPECT().GetFileInfo(ctx, "file4").Return(nil, noSuchFileError)

		router := getRouter()
		RegisterHandlers(router, fileInfoServiceMock)
		response := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/files/file4", nil)

		router.ServeHTTP(response, req)
		assert.Equal(t, http.StatusBadRequest, response.Code)

		assert.JSONEq(
			t,
			fmt.Sprintf(`{"error":"%s"}`, noSuchFileError.Error()),
			response.Body.String(),
		)
	})
}
