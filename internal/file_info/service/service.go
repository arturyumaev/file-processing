package service

import (
	"context"

	"github.com/arturyumaev/file-processing/api/internal/file_info"
	"github.com/arturyumaev/file-processing/api/models"
)

type service struct {
	r file_info.Repository
}

func (svc *service) GetFileInfo(ctx context.Context, name string) (*models.FileInfo, error) {
	return svc.r.FindOne(ctx, name)
}

func New(r file_info.Repository) file_info.Service {
	return &service{r}
}
