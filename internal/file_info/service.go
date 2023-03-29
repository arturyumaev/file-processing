package file_info

import (
	"context"

	"github.com/arturyumaev/file-processing/models"
)

type Service interface {
	GetFileInfo(ctx context.Context, name string) (*models.FileInfo, error)
}
