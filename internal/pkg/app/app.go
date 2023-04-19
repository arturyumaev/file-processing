package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	_ "github.com/arturyumaev/file-processing/docs"
	fileInfoHandler "github.com/arturyumaev/file-processing/internal/file_info/handler"
	fileInfoRepository "github.com/arturyumaev/file-processing/internal/file_info/repository"
	fileInfoService "github.com/arturyumaev/file-processing/internal/file_info/service"
	"github.com/arturyumaev/file-processing/internal/middleware"
	"github.com/arturyumaev/file-processing/pkg/client/postgres"
	"github.com/arturyumaev/file-processing/pkg/logger"
)

type app struct {
	server *http.Server
	log    *zerolog.Logger
	db     *sqlx.DB
}

func New() *app {
	log := logger.Get()

	ctx := context.Background()
	db, err := postgres.NewClient(ctx)
	if err != nil {
		panic(err)
	} else {
		log.Info().Msg("connected to postgres")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	var rootHandler http.Handler
	rootHandler = middleware.Logger(mux)
	rootHandler = middleware.RequestId(rootHandler)
	rootHandler = middleware.PanicRecovery(rootHandler)

	repository := fileInfoRepository.New(db)
	service := fileInfoService.New(repository)
	fileInfoHandler.RegisterHandlers(mux, service)

	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", os.Getenv("APPLICATION_PORT")),
		Handler:        rootHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	app := &app{
		server,
		log,
		db,
	}

	return app
}

func (a *app) Run() error {
	defer a.db.Close()

	a.log.Info().Msgf("starting server at :%s", os.Getenv("APPLICATION_PORT"))
	go func() { // Runs in infinite cycle
		if err := a.server.ListenAndServe(); err != nil {
			a.log.Error().Msgf("failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.server.Shutdown(ctx)
}
