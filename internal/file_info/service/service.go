package service

import (
	"context"
	"mime/multipart"

	"github.com/arturyumaev/file-processing/internal/file_info"
	"github.com/arturyumaev/file-processing/internal/file_info/handler"
)

type Repository interface {
	FindOne(ctx context.Context, name string) (*file_info.FileInfo, error)
}

type service struct {
	r Repository
}

func (svc *service) GetFileInfo(ctx context.Context, name string) (*file_info.FileInfo, error) {
	return svc.r.FindOne(ctx, name)
}

func (svc *service) UploadFile(
	ctx context.Context,
	file multipart.File,
	handler *multipart.FileHeader,
) (*file_info.FileInfo, error) {
	return nil, nil
}

func New(r Repository) handler.Service {
	return &service{r}
}
