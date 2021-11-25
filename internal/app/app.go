package app

import (
	"github.com/lmiedzinski/shiny-barnacle/config"
	"github.com/lmiedzinski/shiny-barnacle/pkg/logger"
	"github.com/lmiedzinski/shiny-barnacle/pkg/postgres"
)

func Run(cfg *config.Config) {
	logger := logger.New(cfg.LogLevel)
	logger.Info("Hello World!")

	// Repositories setup
	pg, err := postgres.New(logger, cfg.PostgresConnectionString)
	if err != nil {
		logger.Fatal(err)
	}
	defer pg.Close()

}
