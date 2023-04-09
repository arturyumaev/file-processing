package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

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
	if os.Getenv("APPLICATION_MODE") == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	log := logger.Get()

	handlerTimeout, _ := strconv.Atoi(os.Getenv("APPLICATION_HANDLER_TIMEOUT"))
	log.Info().Msgf("default handler timeout is: %ds", handlerTimeout)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.RequestId())
	r.Use(middleware.Logger())
	r.Use(middleware.RequestTimeout(time.Duration(handlerTimeout) * time.Second))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ctx := context.Background()
	db, err := postgres.NewClient(ctx)
	if err != nil {
		panic(err)
	} else {
		log.Info().Msg("connected to postgres")
	}

	mux := http.NewServeMux()

	repository := fileInfoRepository.New(db)
	service := fileInfoService.New(repository)
	fileInfoHandler.RegisterHandlers(mux, service)

	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", os.Getenv("APPLICATION_PORT")),
		Handler:        mux,
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

	go func() {
		if err := a.server.ListenAndServe(); err != nil {
			a.log.Error().Msgf("failed to listen and serve: %+v", err)
		} else {
			a.log.Info().Msgf("server started at :%s", os.Getenv("APPLICATION_PORT"))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.server.Shutdown(ctx)
}
