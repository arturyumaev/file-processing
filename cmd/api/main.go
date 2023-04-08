package main

import (
	"github.com/arturyumaev/file-processing/internal/pkg/app"
	"github.com/arturyumaev/file-processing/pkg/logger"
)

// @title File processing API
// @version 1.0
// @description API for file processing

// @host localhost:8888
// @BasePath /
func main() {
	logger := logger.Get().Error()

	app := app.New()
	if err := app.Run(); err != nil {
		logger.Msg(err.Error())
	}
}
