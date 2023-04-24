package handler

import "net/http"

func RegisterHandlers(mux *http.ServeMux, service Service) {
	handlers := New(service)

	mux.HandleFunc("/files/", handlers.GetFileInfo)
	mux.HandleFunc("/files", handlers.PostFile)
}
