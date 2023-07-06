package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/insurance-drinks-api/src/api/core/dto/requests"
	"github.com/insurance-drinks-api/src/api/infrastructure/repository/interfaces"
	"github.com/insurance-drinks-api/src/api/infrastructure/utils"
)

type GetDrinkById struct {
	Repository interfaces.Repository
}

func (h GetDrinkById) Handle(c *gin.Context) {
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
