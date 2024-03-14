package utils

import (
	"github.com/yourscarent/cars_manager/internal/usecase/models"
	cm "github.com/yourscarent/cars_manager/pkg/proto/gen"
)

type Converter struct {
}

func NewConverter() Converter {
	return Converter{}
}

func (c Converter) NewCar(in *cm.Car) models.Car {
	return models.Car{
		Id:          in.Id,
		Brand:       in.Brand,
		Model:       in.Model,
		Type:        in.Type,
		Speed:       in.Speed,
		Seats:       in.Seats,
		Color:       in.Color,
		Description: in.Description,
	}
}

func (c Converter) NewUpdate(in *cm.Update) models.Update {
	return models.Update{
		Color:       in.Color,
		Description: in.Description,
	}
}
