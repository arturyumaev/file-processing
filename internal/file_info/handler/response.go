package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *handler) WriteError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
}

func (h *handler) WriteSuccess(w http.ResponseWriter, entity interface{}) {
	bytes, err := json.Marshal(entity)
	if err != nil {
		h.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
