package postgres

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/arturyumaev/file-processing/pkg/logger"
)

var (
	host   = os.Getenv("POSTGRES_HOST")
	dbname = os.Getenv("POSTGRES_DB_NAME")
	port   = os.Getenv("POSTGRES_PORT")
	user   = os.Getenv("POSTGRES_USERNAME")
	pwd    = os.Getenv("POSTGRES_PASSWORD")
)

const (
	MAX_CONN_ATTEMPTS = 5
	CONN_DELAY        = 3 * time.Second
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

	log := logger.Get()

	err = connectWithTries(func() error { // убрать
		log.Info().Msgf("trying to establish database connection..., %s", connString)
		db, err = sqlx.ConnectContext(ctx, "postgres", connString)
		if err != nil {
			log.Error().Msgf("error: %s", err.Error())
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
