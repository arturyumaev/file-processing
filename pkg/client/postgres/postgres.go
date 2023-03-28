package postgres

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

var (
	host = os.Getenv("POSTGRES_HOST")
	db   = os.Getenv("POSTGRES_DB_NAME")
	port = os.Getenv("POSTGRES_PORT")
	user = os.Getenv("POSTGRES_USERNAME")
	pwd  = os.Getenv("POSTGRES_PASSWORD")
)

const (
	MAX_CONN_ATTEMPTS = 5
	CONN_DELAY        = 3 * time.Second
)

func NewClient(ctx context.Context) (conn *pgx.Conn, err error) {
	connString := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		user,
		pwd,
		host,
		port,
		db,
	)

	err = connectWithTries(func() error {
		conn, err = pgx.Connect(ctx, connString)
		if err != nil {
			return err
		}

		return nil
	}, MAX_CONN_ATTEMPTS, CONN_DELAY)

	return
}

func connectWithTries(fn func() error, attempts uint, delay time.Duration) (err error) {
	for attempts > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attempts--
			continue
		}
		return nil
	}
	return
}
