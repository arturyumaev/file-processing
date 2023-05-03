//go:build repository
// +build repository

package repository

import (
	"context"
	"errors"
	"log"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"

	"github.com/arturyumaev/file-processing/internal/file_info"
	"github.com/arturyumaev/file-processing/pkg/client/postgres"
)

var db *sqlx.DB

func TestMain(m *testing.M) {
	ctx := context.Background()
	conn, err := postgres.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	db = conn
	defer db.Close()

	code := m.Run()
	os.Exit(code)
}

func TestRepository_findOne(t *testing.T) {
	tt := []struct {
		name        string
		arg         string
		expected    *file_info.FileInfo
		expectedErr error
	}{
		{
			name: "when file found",
			arg:  "file1",
			expected: &file_info.FileInfo{
				Id:        4,
				Filename:  "file1",
				Status:    file_info.StatusDone,
				TimeStamp: "08.04.2023 18:20:00 GMT+00",
			},
		},
		{
			name:        "when file wasn't found",
			arg:         "wrong_file_name",
			expectedErr: file_info.ErrNoSuchFile,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			repo := New(db)

			actual, actualErr := repo.FindOne(ctx, tc.arg)
			if tc.expectedErr != nil {
				assert.Equal(t, tc.expectedErr, actualErr)
			}
			assert.EqualValues(t, tc.expected, actual)
		})
	}
}

func TestRepository_create(t *testing.T) {
	tt := []struct {
		name        string
		arg         string
		expectedErr error
	}{
		{
			name:        "success",
			arg:         "file10",
			expectedErr: nil,
		},
		{
			name:        "when file with name already exist",
			arg:         "file1",
			expectedErr: errors.New("duplicate key value"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			repo := New(db)

			actualErr := repo.Create(ctx, tc.arg)
			if tc.expectedErr == nil {
				assert.Nil(t, actualErr)
			} else {
				assert.ErrorContains(t, actualErr, tc.expectedErr.Error())
			}
		})
	}
}
