package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/arturyumaev/file-processing/api/config"
	fileInfoHandler "github.com/arturyumaev/file-processing/api/internal/file_info/handler"
	fileInfoRepository "github.com/arturyumaev/file-processing/api/internal/file_info/repository"
	fileInfoService "github.com/arturyumaev/file-processing/api/internal/file_info/service"
	"github.com/gin-gonic/gin"
)

type app struct {
	server *http.Server
}

func New(config *config.Config) (*app, error) {
	r := gin.Default()

	repository := fileInfoRepository.New()
	fileInfoService := fileInfoService.New(repository)
	fileInfoHandler.RegisterHandlers(r, fileInfoService)

	server := &http.Server{
		Addr:           ":" + config.Server.Port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	app := &app{server}

	return app, nil
}

func (a *app) Run() error {
	log.Println("server started")

	go func() {
		if err := a.server.ListenAndServe(); err != nil {
			log.Fatalf("failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.server.Shutdown(ctx)
}
