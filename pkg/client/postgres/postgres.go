package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	host   = os.Getenv("POSTGRES_HOST")
	dbname = os.Getenv("POSTGRES_DB")
	port   = os.Getenv("POSTGRES_PORT")
	user   = os.Getenv("POSTGRES_USER")
	pwd    = os.Getenv("POSTGRES_PASSWORD")
)

func NewClient(ctx context.Context) (db *sqlx.DB, err error) {
	connString := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		user,
		pwd,
		host,
		port,
		dbname,
	)

	return sqlx.ConnectContext(ctx, "postgres", connString)
}
