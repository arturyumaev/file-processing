package main

import (
	"flag"

	"github.com/arturyumaev/file-processing/api/config"
	"github.com/arturyumaev/file-processing/api/internal/logger"
	"github.com/arturyumaev/file-processing/api/internal/pkg/app"
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

func main() {
	logger := logger.GetLogger().Error()
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
