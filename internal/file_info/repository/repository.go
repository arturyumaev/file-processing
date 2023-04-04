package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/arturyumaev/file-processing/internal/file_info"
	"github.com/arturyumaev/file-processing/internal/file_info/queries"
	"github.com/arturyumaev/file-processing/models"
)

type repository struct {
	db *sqlx.DB
}

func (r *repository) FindOne(ctx context.Context, name string) (*models.FileInfo, error) {
	fileInfo := &models.FileInfo{}

	if name == "file3" {
		time.Sleep(5 * time.Second)
	}

	err := r.db.GetContext(ctx, fileInfo, queries.SelectFileInfo, name)
	if err == sql.ErrNoRows {
		return nil, file_info.ErrNoSuchFile
	} else if err != nil {
		return nil, err
	}

	return fileInfo, nil
}

func New(db *sqlx.DB) file_info.Repository {
	return &repository{db}
}
