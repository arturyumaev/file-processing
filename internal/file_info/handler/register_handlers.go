package handler

import "net/http"

func RegisterHandlers(mux *http.ServeMux, service Service) {
	h := New(service)
	mux.HandleFunc("/files", h.GetFileInfo)
}
