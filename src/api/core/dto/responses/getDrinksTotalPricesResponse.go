package responses

import (
	"github.com/insurance-drinks-api/src/api/core/entities/interfaces"
)

type GetDrinksTotalPricesResponse struct {
	Drink          []interfaces.Drink `json:"drinks"`
	TotalWithTaxes float64          `json:"total_with_taxes"`
}
