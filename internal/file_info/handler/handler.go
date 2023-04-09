package handler

import (
	"context"
	"net/http"

	"github.com/arturyumaev/file-processing/internal/file_info"
)

//go:generate mockgen -source=handler.go -destination=../mocks/service.go
type Service interface {
	GetFileInfo(ctx context.Context, name string) (*file_info.FileInfo, error)
}

type Handler interface {
	GetFileInfo(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	svc Service
}

func New(svc Service) Handler {
	return &handler{svc}
}
