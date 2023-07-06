package handlers

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/insurance-drinks-api/src/api/core/entities/dto"
	"github.com/insurance-drinks-api/src/api/infrastructure/utils"
	"github.com/insurance-drinks-api/src/api/infrastructure/utils/instances"
	mock "github.com/insurance-drinks-api/src/api/mocks/repository"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

func TestGetDrinkByIdOk(t *testing.T) {
	var wantedId int64 = 5

	response := utils.GetTargetResponse()

	jsonResponse := `
	{
		"id": "5",
		"name": "Vino Cabernet",
		"type": "Vino",
		"price": 100
	}`

	drink := &dto.Drink{
		BaseDrink: dto.BaseDrink{
			Id: "_Id",
			Name: "Name",
			Type:instances.Whisky,
			Price:100,
		},
	}

	repo := new(mock.DrinkRepositoryMock)

	handler := GetDrinkById {
		Repository: repo,
	}

	repo.On("GetByID", wantedId).Return(drink, nil)

	c:= utils.GetMockedContext(http.MethodGet,
		fmt.Sprintf("%s", strings.Replace(utils.URL_GET_DRINK_BY_ID, ":drinkId",
			fmt.Sprintf("%d", wantedId), 1)),
		bytes.NewBufferString(jsonResponse),
		response)

	c.Params = []gin.Param{
		{
			Key:   "drinkId",
			Value: fmt.Sprintf("%d", wantedId),
		},
	}

	handler.Handle(c)
	assert.EqualValues(t, http.StatusOK, response.Code)
}
