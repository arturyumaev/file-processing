package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/arturyumaev/file-processing/internal/file_info"
	"github.com/arturyumaev/file-processing/internal/file_info/queries"
	"github.com/arturyumaev/file-processing/models"
)

type repository struct {
	conn *pgx.Conn
}

func (r *repository) FindOne(ctx context.Context, name string) (*models.FileInfo, error) {
	fileInfo := &models.FileInfo{}

	if name == "file3" {
		time.Sleep(5 * time.Second)
	}

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

func New(conn *pgx.Conn) file_info.Repository {
	return &repository{conn}
}
