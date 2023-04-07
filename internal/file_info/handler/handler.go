package handler

import (
	"context"

	"github.com/arturyumaev/file-processing/internal/file_info"
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=handler.go -destination=../mocks/service.go
type Service interface {
	GetFileInfo(ctx context.Context, name string) (*file_info.FileInfo, error)
}

type Handler interface {
	GetFileInfo(c *gin.Context)
}

type handler struct {
	svc Service
}

func New(svc Service) Handler {
	return &handler{svc}
}
