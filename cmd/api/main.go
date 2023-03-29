package main

import (
	"flag"

	"github.com/arturyumaev/file-processing/config"
	"github.com/arturyumaev/file-processing/internal/pkg/app"
	"github.com/arturyumaev/file-processing/pkg/logger"
)

type Flags struct {
	ConfigPath string
}

func parseFlags() *Flags {
	var configPath string
	flag.StringVar(&configPath, "config", "./config/config.yaml", "path to config file")
	flag.Parse()

	flags := &Flags{
		ConfigPath: configPath,
	}

	return flags
}

// @title File processing API
// @version 1.0
// @description API for file processing

// @host localhost:8888
// @BasePath /
func main() {
	logger := logger.Get().Error()
	flags := parseFlags()

	cfg, err := config.Read(flags.ConfigPath)
	if err != nil {
		logger.Msg(err.Error())
	}

	app := app.New(cfg)
	if err = app.Run(); err != nil {
		logger.Msg(err.Error())
	}
}
