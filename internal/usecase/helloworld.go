package usecase

import (
	"context"

	"github.com/lmiedzinski/shiny-barnacle/internal/entity"
)

type HelloWorldUseCase struct {
	repository HelloWorldRepository
}

func NewHelloWorldUseCase(repository HelloWorldRepository) *HelloWorldUseCase {
	return &HelloWorldUseCase{repository: repository}
}

func (uc *HelloWorldUseCase) GetHelloWorld(ctx context.Context) ([]entity.HelloWorldDto, error) {
	helloWorldMessages, err := uc.repository.GetHelloWorldMessages(ctx)
	if err != nil {
		return nil, err
	}

	return helloWorldMessages, nil
}
