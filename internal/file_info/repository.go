package file_info

import (
	"context"

	"github.com/arturyumaev/file-processing/api/models"
)

type Repository interface {
	FindOne(ctx context.Context, name string) (*models.FileInfo, error)
	FindAll(ctx context.Context) ([]*models.FileInfo, error)
}
