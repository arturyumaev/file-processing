package file_info

import (
	"context"

	"github.com/arturyumaev/file-processing/models"
)

type Repository interface {
	FindOne(ctx context.Context, name string) (*models.FileInfo, error)
}
