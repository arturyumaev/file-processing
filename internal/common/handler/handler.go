package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arturyumaev/file-processing/internal/file_info"
)

type CommonHandler struct{}

func (h *CommonHandler) IsMethodValid(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		h.WriteError(w, http.StatusMethodNotAllowed, file_info.ErrMethodNotAllowed)
		return false
	}
	return true
}

func (h *CommonHandler) WriteError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
}

func (h *CommonHandler) WriteSuccess(w http.ResponseWriter, entity interface{}) {
	bytes, err := json.Marshal(entity)
	if err != nil {
		h.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func New() *CommonHandler {
	return &CommonHandler{}
}
