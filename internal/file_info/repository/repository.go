package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/arturyumaev/file-processing/internal/file_info"
	"github.com/arturyumaev/file-processing/internal/file_info/queries"
	"github.com/arturyumaev/file-processing/internal/file_info/service"
)

type repository struct {
	db *sqlx.DB
}

func (r *repository) FindOne(ctx context.Context, name string) (*file_info.FileInfo, error) {
	fileInfo := &file_info.FileInfo{}

	if name == "file3" {
		time.Sleep(5 * time.Second)
	}

	err := r.db.GetContext(ctx, fileInfo, r.db.Rebind(queries.SelectFileInfo), name)
	if err == sql.ErrNoRows {
		return nil, file_info.ErrNoSuchFile
	} else if err != nil {
		return nil, err
	}

	return fileInfo, nil
}

func New(db *sqlx.DB) service.Repository {
	return &repository{db}
}
