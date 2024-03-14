package usecase

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/yourscarent/cars_manager/internal/db"
	"github.com/yourscarent/cars_manager/internal/usecase/models"
)

type Manager struct {
	repo db.Repository
}

type Params struct {
	Repo db.Repository
}

func NewManager(p Params) *Manager {
	return &Manager{
		repo: p.Repo,
	}
}

func (m Manager) CreateCar(ctx context.Context, car models.Car) error {
	car.Id = uuid.New().String()

	err := m.repo.CreateCar(ctx, car)
	if err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	return nil
}

func (m Manager) UpdateCar(ctx context.Context, updated models.Update) error {
	err := m.repo.UpdateCar(ctx, updated)
	if err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	return nil
}

func (m Manager) DeleteCar(ctx context.Context, carID string) error {
	err := m.repo.DeleteCar(ctx, carID)
	if err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	return nil
}
