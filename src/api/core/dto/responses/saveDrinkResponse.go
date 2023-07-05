package responses

import "github.com/insurance-drinks-api/src/api/core/entities"

type SaveDrinkResponse struct {
	Drink *entities.Drink `json:"drink"`
}
