package responses

import "github.com/insurance-drinks-api/src/api/core/entities"

type GetDrinksTotalPricesResponse struct {
	Drink          []entities.Drink `json:"drinks"`
	TotalWithTaxes float64          `json:"total_with_taxes"`
}
