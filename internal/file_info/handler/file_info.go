package handler

import (
	"context"
	"net/http"
	"regexp"

	"github.com/arturyumaev/file-processing/internal/file_info"
)

var (
	getFileRegexp = regexp.MustCompile(`^\/files\/(\w+)$`)
)

//go:generate mockgen -source=file_info.go -destination=../mocks/service.go
type Service interface {
	GetFileInfo(ctx context.Context, name string) (*file_info.FileInfo, error)
}

type Handler interface {
	GetFileInfo(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	svc Service
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
// @Failure      408  {object}  file_info.HttpResponseErr
// @Failure      500  {object}  file_info.HttpResponseErr
// @Router       /files/{name} [get]
func (h *handler) GetFileInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.WriteError(w, http.StatusMethodNotAllowed, file_info.ErrMethodNotAllowed)
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
		h.WriteError(w, http.StatusBadRequest, err)
		return
	}

	h.WriteSuccess(w, file)
}

func New(svc Service) Handler {
	return &handler{svc}
}
