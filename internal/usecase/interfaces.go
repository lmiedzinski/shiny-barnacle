package usecase

import (
	"context"

	"github.com/lmiedzinski/shiny-barnacle/internal/entity"
)

type (
	HelloWorld interface {
		GetHelloWorld(ctx context.Context) ([]entity.HelloWorldDto, error)
	}

	HelloWorldRepository interface {
		GetHelloWorldMessages(ctx context.Context) ([]entity.HelloWorldDto, error)
	}
)
