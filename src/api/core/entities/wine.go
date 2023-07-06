package entities

import "github.com/insurance-drinks-api/src/api/core/entities/dto"

type Wine struct {
	dto.BaseDrink
}

func (t *Wine) GetTotalWithTaxes() float64 {
	return t.Price * 1.27
}
