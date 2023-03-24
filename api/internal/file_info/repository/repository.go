package repository

import (
	"context"

	"github.com/arturyumaev/file-processing/api/internal/file_info"
	"github.com/arturyumaev/file-processing/api/models"
	"github.com/google/uuid"
)

type repository struct {
}

func (r *repository) GetFileInfo(ctx context.Context) (*models.FileInfo, error) {
	fileInfo := &models.FileInfo{
		Id:     uuid.New(),
		Hash:   "test hash",
		Status: models.FileInfoStatusInQueue,
	}

	return fileInfo, nil
}

func New() file_info.Repository {
	return &repository{}
}
