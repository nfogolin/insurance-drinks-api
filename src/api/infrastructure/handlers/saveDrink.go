package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/insurance-drinks-api/src/api/core/entities/dto"
	"github.com/insurance-drinks-api/src/api/infrastructure/repository"
	"github.com/insurance-drinks-api/src/api/infrastructure/utils"
)

type SaveDrink struct {
	Repository repository.Repository
}

func (h SaveDrink) SaveDrink(c *gin.Context) {
	var drink dto.Drink

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
