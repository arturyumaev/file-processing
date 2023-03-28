package repository

import (
	"context"
	"database/sql"

	"github.com/arturyumaev/file-processing/api/internal/file_info"
	"github.com/arturyumaev/file-processing/api/models"
	"github.com/google/uuid"
)

type repository struct {
	db *sql.DB
}

func (r *repository) GetFileInfo(ctx context.Context, name string) (*models.FileInfo, error) {
	fileInfo := &models.FileInfo{
		Id:     uuid.New(),
		Hash:   "test hash",
		Status: models.FileInfoStatusInQueue,
	}

	return fileInfo, nil
}

func New(db *sql.DB) file_info.Repository {
	return &repository{db}
}
