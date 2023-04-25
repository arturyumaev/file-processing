package service

import (
	"context"
	"mime/multipart"

	"github.com/arturyumaev/file-processing/internal/file_info"
	"github.com/arturyumaev/file-processing/internal/file_info/handler"
)

type Repository interface {
	FindOne(ctx context.Context, filename string) (*file_info.FileInfo, error)
	Create(ctx context.Context, filename string) error
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
	fileHeader *multipart.FileHeader,
) (*file_info.FileInfo, error) {
	filename := fileHeader.Filename

	if len(filename) > file_info.MAX_FILE_NAME_LENGTH {
		return nil, file_info.ErrFileNameLengthTooLong
	}

	if err := svc.r.Create(ctx, filename); err != nil {
		return nil, err
	}

	fileInfo, err := svc.GetFileInfo(ctx, filename)
	if err != nil {
		return nil, err
	}

	return fileInfo, nil
}

func New(r Repository) handler.Service {
	return &service{r}
}
