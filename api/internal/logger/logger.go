package logger

import (
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var once sync.Once
var logger *zerolog.Logger

// singleton
func GetLogger() *zerolog.Logger {
	once.Do(func() {
		logger = &log.Logger
	})

	return logger
}
