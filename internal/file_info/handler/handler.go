package handler

import (
	"github.com/arturyumaev/file-processing/internal/file_info"
)

type handler struct {
	svc file_info.Service
}

func New(svc file_info.Service) file_info.Handler {
	return &handler{svc}
}
