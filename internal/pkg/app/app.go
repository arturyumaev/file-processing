package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/arturyumaev/file-processing/config"
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
	config *config.Config
	log    *zerolog.Logger
	conn   *pgx.Conn
}

func New(config *config.Config) *app {
	log := logger.Get()

	if config.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	log.Info().Msgf("default handler timeout is: %ds", config.ApplicationHandlerTimeout)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.RequestId())
	r.Use(middleware.Logger())
	r.Use(middleware.RequestTimeout(time.Duration(config.ApplicationHandlerTimeout) * time.Second))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ctx := context.Background()
	conn, err := postgres.NewClient(ctx)
	if err != nil {
		panic(err)
	} else {
		log.Info().Msg("connected to postgres")
	}

	repository := fileInfoRepository.New(conn)
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
		conn,
	}

	return app
}

func (a *app) Run() error {
	defer a.conn.Close(context.Background())

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
