package handler

import (
	"errors"
	"net/http"
)

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
		h.WriteError(w, http.StatusMethodNotAllowed, errors.New("method not allowed"))
		return
	}

	filename := r.URL.Query().Get("name")
	if filename == "" {
		h.WriteError(w, http.StatusBadRequest, errors.New("empty parameter: name"))
		return
	}

	file, err := h.svc.GetFileInfo(r.Context(), filename)
	if err != nil {
		h.WriteError(w, http.StatusBadRequest, err)
		return
	}

	h.WriteSuccess(w, file)
}
