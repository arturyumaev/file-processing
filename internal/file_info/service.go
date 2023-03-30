package file_info

import (
	"context"

	"github.com/arturyumaev/file-processing/models"
)

//go:generate mockgen -source=service.go -destination=mocks/service.go

type Service interface {
	GetFileInfo(ctx context.Context, name string) (*models.FileInfo, error)
}
