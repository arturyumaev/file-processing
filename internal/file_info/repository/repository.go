package repository

import (
	"context"

	"github.com/arturyumaev/file-processing/api/internal/file_info"
	"github.com/arturyumaev/file-processing/api/internal/file_info/queries"
	"github.com/arturyumaev/file-processing/api/models"
	"github.com/jackc/pgx/v5"
)

type repository struct {
	conn *pgx.Conn
}

func (r *repository) FindOne(ctx context.Context, name string) (*models.FileInfo, error) {
	fileInfo := &models.FileInfo{}

	err := r.conn.
		QueryRow(ctx, queries.SelectFileInfo, name).
		Scan(
			&fileInfo.Id,
			&fileInfo.Status,
			&fileInfo.TimeStamp,
			&fileInfo.FilenameHash,
		)
	if err == pgx.ErrNoRows {
		return nil, file_info.ErrNoSuchFile
	} else if err != nil {
		return nil, err
	}

	return fileInfo, nil
}

func (r *repository) FindAll(ctx context.Context) ([]*models.FileInfo, error) {
	all := []*models.FileInfo{}
	return all, nil
}

func New(conn *pgx.Conn) file_info.Repository {
	return &repository{conn}
}
