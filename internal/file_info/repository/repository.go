package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/arturyumaev/file-processing/internal/file_info"
	"github.com/arturyumaev/file-processing/internal/file_info/queries"
	"github.com/arturyumaev/file-processing/internal/file_info/service"
)

type repository struct {
	db *sqlx.DB
}

func (r *repository) FindOne(ctx context.Context, filename string) (*file_info.FileInfo, error) {
	fileInfo := &file_info.FileInfo{}

	if err := r.db.GetContext(ctx, fileInfo, r.db.Rebind(queries.SelectFileInfo), filename); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, file_info.ErrNoSuchFile
		}

		return nil, err
	}

	return fileInfo, nil
}

func (r *repository) Create(ctx context.Context, filename string) error {
	tx := r.db.MustBeginTx(ctx, nil)
	defer tx.Rollback()

	_, err := tx.ExecContext(ctx, r.db.Rebind(queries.InsertFile), filename)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func New(db *sqlx.DB) service.Repository {
	return &repository{db}
}
