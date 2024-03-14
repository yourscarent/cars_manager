package utils

import (
	"errors"
	cm "github.com/yourscarent/cars_manager/pkg/proto/gen"
)

type Validator struct {
}

func NewValidator() Converter {
	return Converter{}
}

func (v Validator) ValidateCar(car *cm.Car) error {
	if car.Id == "" {
		return errors.New("id cannot be empty")
	}
	if len(car.Description) < 10 || len(car.Description) > 200 {
		return errors.New("description must have length from 10 to 200")
	}
	if car.Brand == "" {
		return errors.New("brand cannot be empty")
	}
	if car.Model == "" {
		return errors.New("model cannot be empty")
	}
	if car.Type == "" {
		return errors.New("type cannot be empty")
	}
	if car.Speed < 1 {
		return errors.New("speed must be a positive value")
	}
	if car.Seats < 1 {
		return errors.New("seats must be a positive value")
	}
	if car.Color == "" {
		return errors.New("color cannot be empty")
	}
	return nil
}

func (v Validator) ValidateId(id *cm.ID) error {
	if id.Id == "" {
		return errors.New("id cannot be empty")
	}
	return nil
}

func (v Validator) ValidateUpdate(update *cm.Update) error {
	if len(update.Description) < 10 || len(update.Description) > 200 {
		return errors.New("description must have length from 10 to 200")
	}
	if update.Color == "" {
		return errors.New("color cannot be empty")
	}
	return nil
}
