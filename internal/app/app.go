package app

import (
	"github.com/lmiedzinski/shiny-barnacle/config"
	"github.com/lmiedzinski/shiny-barnacle/pkg/logger"
)

func Run(cfg *config.Config) {
	logger := logger.New(cfg.LogLevel)
	logger.Info("Hello World!")
}
