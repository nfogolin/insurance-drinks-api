package entities

import (
	"github.com/insurance-drinks-api/src/api/infrastructure/utils"
)

type Whisky struct {
	Drink
}

func (t *Whisky) GetTotalWithTaxes() float64 {
	return t.Price * 1.105
}

func (t *Whisky) GetAging() int32 {
	if utils.NullableInt32ToFloat64(t.Aging) > 0 {
		return *t.Aging
	}

	return 1
}
