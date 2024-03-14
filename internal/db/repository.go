package db

import (
	"context"
	"github.com/yourscarent/cars_manager/internal/usecase/models"
)

type Repository interface {
	CreateCar(ctx context.Context, car models.Car) error
	DeleteCar(ctx context.Context, carID string) error
	UpdateCar(ctx context.Context, updated models.Update) error
}
