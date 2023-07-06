package responses

import (
	"github.com/insurance-drinks-api/src/api/core/entities/dto"
)

type SaveDrinkResponse struct {
	Drink *dto.Drink `json:"drink"`
}
