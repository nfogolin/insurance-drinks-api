package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/insurance-drinks-api/src/api/core/dto/responses"
	"github.com/insurance-drinks-api/src/api/core/entities/interfaces"
	repository "github.com/insurance-drinks-api/src/api/infrastructure/repository/interfaces"
	"github.com/insurance-drinks-api/src/api/infrastructure/utils"
	"github.com/insurance-drinks-api/src/api/infrastructure/utils/instances"
)

type GetDrinks struct {
	Repository repository.Repository
}

func (h GetDrinks) Handle(c *gin.Context) {
	var response responses.GetDrinksTotalPricesResponse

	drinks, err := h.Repository.GetDrinks()
	if err != nil {
		utils.HandleResponseNotFound(c, fmt.Sprintf("Not found drinks"))
		return
	}

	var castDrink interfaces.Drink

	for _, drink := range drinks {
		castDrink = instances.CastDrinkInstances(drink)

		response.Drink = append(response.Drink, castDrink)

		response.TotalWithTaxes = response.TotalWithTaxes +
			castDrink.GetTotalWithTaxes()
	}

	utils.HandleResponseOk(c, response)
}
