package file_info

import (
	"context"

	"github.com/arturyumaev/file-processing/api/models"
)

type Repository interface {
	GetFileInfo(ctx context.Context) (*models.FileInfo, error)
}
