package entities

import "github.com/insurance-drinks-api/src/api/core/entities/dto"

type Water struct {
	dto.BaseDrink
}

func (t *Water) GetTotalWithTaxes() float64 {
	return t.Price * 1.21
}
