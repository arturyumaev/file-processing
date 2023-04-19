package repository

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"

	"github.com/arturyumaev/file-processing/internal/file_info"
	"github.com/arturyumaev/file-processing/internal/file_info/queries"
)

func TestRepository_findOne(t *testing.T) {
	tt := []struct {
		name             string
		args             string
		expected         *file_info.FileInfo
		simulateSqlQuery func(sqlmock.Sqlmock)
		expectedErr      error
	}{
		{
			name: "when file is found",
			args: "file1",
			expected: &file_info.FileInfo{
				Id:        5,
				Filename:  "file1",
				Status:    file_info.StatusDone,
				TimeStamp: "08.04.2023 21:20:00 GMT+03",
			},
			simulateSqlQuery: func(mockSQL sqlmock.Sqlmock) {
				rows := sqlmock.
					NewRows([]string{"id", "filename", "status", "timestamp"}).
					AddRow(5, "file1", "done", "08.04.2023 21:20:00 GMT+03")

				mockSQL.
					ExpectQuery(regexp.QuoteMeta(queries.SelectFileInfo)).
					WithArgs("file1").
					WillReturnRows(rows)
			},
		},
		{
			name:        "when file wasn't found",
			args:        "wrong_file_name",
			expectedErr: file_info.ErrNoSuchFile,
			simulateSqlQuery: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(regexp.QuoteMeta(queries.SelectFileInfo)).
					WithArgs("wrong_file_name").
					WillReturnError(sql.ErrNoRows)
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			mockDB, mockSQL, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer mockDB.Close()

			db := sqlx.NewDb(mockDB, "sqlmock")
			mockRepository := New(db)

			tc.simulateSqlQuery(mockSQL)

			actual, actualErr := mockRepository.FindOne(ctx, tc.args)
			if tc.expectedErr != nil {
				assert.EqualError(t, actualErr, tc.expectedErr.Error())
			}
			assert.EqualValues(t, actual, tc.expected)
		})
	}
}
