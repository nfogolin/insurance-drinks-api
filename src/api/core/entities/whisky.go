package entities

import (
	"github.com/insurance-drinks-api/src/api/core/entities/dto"
	"github.com/insurance-drinks-api/src/api/infrastructure/utils"
)

type Whisky struct {
	dto.BaseDrink
	dto.Aging
}

func (t *Whisky) GetTotalWithTaxes() float64 {
	return t.Price * 1.105 * float64(t.GetAging())
}

func (t *Whisky) GetAging() int32 {
	if utils.NullableInt32ToFloat64(t.Aging.Aging) > 0 {
		return *t.Aging.Aging
	}

	return 1
}
