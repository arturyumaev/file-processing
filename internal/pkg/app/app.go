package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/arturyumaev/file-processing/api/config"
	fileInfoHandler "github.com/arturyumaev/file-processing/api/internal/file_info/handler"
	fileInfoRepository "github.com/arturyumaev/file-processing/api/internal/file_info/repository"
	fileInfoService "github.com/arturyumaev/file-processing/api/internal/file_info/service"
	"github.com/arturyumaev/file-processing/api/internal/middleware"
	"github.com/arturyumaev/file-processing/api/pkg/client/postgres"
	"github.com/arturyumaev/file-processing/api/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type app struct {
	server *http.Server
	config *config.Config
	log    *zerolog.Logger
}

func New(config *config.Config) *app {
	log := logger.Get()

	if config.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.RequestId())
	r.Use(middleware.Logger())

	db, err := postgres.NewClient()
	if err != nil {
		panic(err)
	} else {
		log.Info().Msg("connected to postgres")
	}
	defer db.Close()

	repository := fileInfoRepository.New(db)
	service := fileInfoService.New(repository)
	fileInfoHandler.RegisterHandlers(r, service)

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.Server.Port),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	app := &app{
		server,
		config,
		log,
	}

	return app
}

func (a *app) Run() error {
	go func() {
		if err := a.server.ListenAndServe(); err != nil {
			a.log.Error().Msgf("failed to listen and serve: %+v", err)
		}
	}()
	a.log.Info().Msgf("server started at :%d", a.config.Server.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.server.Shutdown(ctx)
}
