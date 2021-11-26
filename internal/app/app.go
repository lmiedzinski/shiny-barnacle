package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/lmiedzinski/shiny-barnacle/config"
	"github.com/lmiedzinski/shiny-barnacle/internal/controller/http/apiv1"
	"github.com/lmiedzinski/shiny-barnacle/internal/usecase"
	"github.com/lmiedzinski/shiny-barnacle/internal/usecase/repository"
	"github.com/lmiedzinski/shiny-barnacle/pkg/httpserver"
	"github.com/lmiedzinski/shiny-barnacle/pkg/logger"
	"github.com/lmiedzinski/shiny-barnacle/pkg/postgres"

	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {
	logger := logger.New(cfg.LogLevel)
	logger.Info("Hello World!")

	// Repositories setup
	pg, err := postgres.New(logger, cfg.PostgresConnectionString)
	if err != nil {
		// logger.Fatal(err)
	}
	defer pg.Close()

	helloWorldUseCase := usecase.NewHelloWorldUseCase(repository.NewHelloWorldPostgresRepository(pg))

	// HTTP Server
	handler := gin.New()
	apiv1.NewRouter(handler, logger, helloWorldUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.PortNumber))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Info("signal: " + s.String())
	case err = <-httpServer.Notify():
		logger.Error(err)
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		logger.Error(err)
	}
}
