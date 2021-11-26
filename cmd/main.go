package main

import (
	"github.com/lmiedzinski/shiny-barnacle/config"
	"github.com/lmiedzinski/shiny-barnacle/internal/app"
)

func main() {
	cfg := config.GetConfig()
	app.Run(cfg)
}
