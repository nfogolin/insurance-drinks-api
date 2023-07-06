package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/insurance-drinks-api/src/api/core/dto/responses"
	"github.com/insurance-drinks-api/src/api/core/entities"
	"github.com/insurance-drinks-api/src/api/infrastructure/repository"
	"github.com/insurance-drinks-api/src/api/infrastructure/utils"
)

type GetDrinks struct {
	Repository repository.Repository
}

func (h GetDrinks) GetDrinks(c *gin.Context) {
	var response responses.GetDrinksTotalPricesResponse

	drinks, err := h.Repository.GetDrinks()
	if err != nil {
		utils.HandleResponseNotFound(c, fmt.Sprintf("Not found drinks"))
		return
	}

	for _, drink := range drinks {
		response.Drink = append(response.Drink, drink)

		switch drink.Type {
		case entities.VINO:
			response.TotalWithTaxes = response.TotalWithTaxes + drink.Price * 1.21
		case entities.WHISKY:
			response.TotalWithTaxes = response.TotalWithTaxes +
				drink.Price * 1.105 * utils.NullableInt32ToFloat64(drink.Aging)
		case entities.WATER:
			response.TotalWithTaxes = response.TotalWithTaxes + drink.Price * 2.5
		}
	}

	utils.HandleResponseOk(c, response)
}
