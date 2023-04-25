package handler

import (
	"context"
	"errors"
	"mime/multipart"
	"net/http"
	"regexp"

	commonHandler "github.com/arturyumaev/file-processing/internal/common/handler"
	"github.com/arturyumaev/file-processing/internal/file_info"
)

var (
	getFileRegexp = regexp.MustCompile(`^\/files\/(\w+)$`)
)

const (
	MAX_FILE_SIZE_MB     = 10 << 20 // 10 Mb
	FORM_FIELD_FILE_NAME = "file"
)

//go:generate mockgen -source=file_info.go -destination=../mocks/service.go
type Service interface {
	GetFileInfo(ctx context.Context, name string) (*file_info.FileInfo, error)
	UploadFile(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) (*file_info.FileInfo, error)
}

type Handler interface {
	GetFileInfo(w http.ResponseWriter, r *http.Request)
	PostFile(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	svc Service
	*commonHandler.CommonHandler
}

// GetFileInfo godoc
// @Summary      Get meta information about a file
// @Description  get file info by its name
// @Tags         files
// @Produce      json
// @Param        name path string false "File name"
// @Success      200  {object}  file_info.FileInfo
// @Failure      400  {object}  file_info.HttpResponseErr
// @Failure      404  {object}  file_info.HttpResponseErr
// @Failure      500  {object}  file_info.HttpResponseErr
// @Router       /files/{name} [get]
func (h *handler) GetFileInfo(w http.ResponseWriter, r *http.Request) {
	if !h.IsMethodValid(w, r, http.MethodGet) {
		return
	}

	matches := getFileRegexp.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		h.WriteError(w, http.StatusBadRequest, file_info.ErrEmptyParameterName)
		return
	}
	filename := matches[1]

	file, err := h.svc.GetFileInfo(r.Context(), filename)
	if err != nil {
		if errors.Is(err, file_info.ErrNoSuchFile) {
			h.WriteError(w, http.StatusNotFound, file_info.ErrNoSuchFile)
		} else {
			h.WriteError(w, http.StatusInternalServerError, err)
		}

		return
	}

	h.WriteSuccess(w, file)
}

// PostFile godoc
// @Summary      Accepts a file and puts it in database
// @Description  Accepts a file and puts it in database
// @Tags         files
// @Produce      json
// @Success      200  {object}  file_info.FileInfo
// @Failure      400  {object}  file_info.HttpResponseErr
// @Failure      500  {object}  file_info.HttpResponseErr
// @Router       /files [post]
func (h *handler) PostFile(w http.ResponseWriter, r *http.Request) {
	if !h.IsMethodValid(w, r, http.MethodPost) {
		return
	}

	r.ParseMultipartForm(MAX_FILE_SIZE_MB)
	file, fileHeader, err := r.FormFile(FORM_FIELD_FILE_NAME)
	if err != nil {
		h.WriteError(w, http.StatusBadRequest, file_info.ErrRetrievingFile)
		return
	}
	defer file.Close()

	fileInfo, err := h.svc.UploadFile(r.Context(), file, fileHeader)
	if err != nil {
		h.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	h.WriteSuccess(w, fileInfo)
}

func New(svc Service) Handler {
	ch := commonHandler.New()
	return &handler{svc, ch}
}
