package repository

import (
	"context"

	"github.com/lmiedzinski/shiny-barnacle/internal/entity"
	"github.com/lmiedzinski/shiny-barnacle/pkg/postgres"
)

const _defaultEntityCap = 64

type HelloWorldRepo struct {
	*postgres.Postgres
}

func NewHelloWorldPostgresRepository(pg *postgres.Postgres) *HelloWorldRepo {
	return &HelloWorldRepo{pg}
}

func (r *HelloWorldRepo) GetHelloWorldMessages(ctx context.Context) ([]entity.HelloWorldDto, error) {
	sql, _, err := r.Builder.
		Select("message").
		From("hello_world_messages").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	entities := make([]entity.HelloWorldDto, 0, _defaultEntityCap)

	for rows.Next() {
		e := entity.HelloWorldDto{}

		err = rows.Scan(&e.Message)
		if err != nil {
			return nil, err
		}

		entities = append(entities, e)
	}

	return entities, nil
}
