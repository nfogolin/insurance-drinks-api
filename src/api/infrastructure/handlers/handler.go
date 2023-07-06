package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/insurance-drinks-api/src/api/core/dto/requests"
	"github.com/insurance-drinks-api/src/api/core/dto/responses"
	"github.com/insurance-drinks-api/src/api/core/entities"
	"github.com/insurance-drinks-api/src/api/infrastructure/repository"
	"github.com/insurance-drinks-api/src/api/infrastructure/utils"
)

type Handler struct {
	Repository repository.Repository
}

func (h Handler) Ping(c *gin.Context) {
	utils.HandleResponseOk(c, "pong")
}

func (h Handler) GetDrinkById(c *gin.Context) {
	var getDrinkByIdRequest requests.GetDrinkByIdRequest

	if err := c.ShouldBindUri(&getDrinkByIdRequest); err != nil {
		utils.HandleResponseBadRequest(c, "drinkId is required")
		return
	}

	response, err := h.Repository.GetByID(getDrinkByIdRequest.Id)
	if err != nil {
		utils.HandleResponseNotFound(c, fmt.Sprintf("Not found drink id: %d", getDrinkByIdRequest.Id))
		return
	}

	utils.HandleResponseOk(c, response)
}

func (h Handler) GetDrinks(c *gin.Context) {
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

func (h Handler) SaveDrink(c *gin.Context) {
	var drink entities.Drink

	if err := c.ShouldBindJSON(&drink); err != nil {
		utils.HandleResponseBadRequest(c, err.Error())
		return
	}

	response, err := h.Repository.SaveDrink(drink)
	if err != nil {
		utils.HandleResponseNotFound(c, fmt.Sprintf(err.Error()))
		return
	}

	utils.HandleResponseOk(c, response)
}
