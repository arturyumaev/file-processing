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

	lg := logger.Get().Error()

	application := app.New()
	if err := application.Run(); err != nil {
		lg.Msg(err.Error())
	}
}
