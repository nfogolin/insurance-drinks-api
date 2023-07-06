package interfaces

import "github.com/insurance-drinks-api/src/api/core/entities/dto"

type Repository interface {
	GetByID(int64) (*dto.Drink, error)
	GetDrinks() ([]dto.Drink, error)
	SaveDrink(dto.Drink) (*dto.Drink, error)
}
