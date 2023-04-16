package main

import (
	"os"

	"github.com/arturyumaev/file-processing/docs"
	"github.com/arturyumaev/file-processing/internal/pkg/app"
	"github.com/arturyumaev/file-processing/pkg/logger"
)

// @title File processing API
// @version 1.0
// @description API for file processing

// @BasePath /
func main() {
	docs.SwaggerInfo.Host = "localhost:" + os.Getenv("APPLICATION_PORT")

	logger := logger.Get().Error()

	app := app.New()
	if err := app.Run(); err != nil {
		logger.Msg(err.Error())
	}
}
