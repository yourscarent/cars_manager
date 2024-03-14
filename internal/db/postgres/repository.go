package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/yourscarent/cars_manager/internal/db"
	"github.com/yourscarent/cars_manager/internal/usecase/models"
)

var _ db.Repository = &Postgres{}

func NewRepository(conn *sqlx.DB) *Postgres {
	return &Postgres{}
}

type Postgres struct {
}

func (p Postgres) CreateCar(ctx context.Context, car models.Car) error {
	//TODO implement me
	panic("implement me")
}

func (p Postgres) DeleteCar(ctx context.Context, carID string) error {
	//TODO implement me
	panic("implement me")
}

func (p Postgres) UpdateCar(ctx context.Context, updated models.Update) error {
	//TODO implement me
	panic("implement me")
}
