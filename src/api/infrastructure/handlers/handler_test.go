package handlers

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/insurance-drinks-api/src/api/config"
	"github.com/insurance-drinks-api/src/api/infrastructure/repository"
	"github.com/insurance-drinks-api/src/api/infrastructure/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

func TestGetDrinkByIdOk(t *testing.T) {
	wantedId := "5"

	response := utils.GetTargetResponse()

	jsonResponse := `
	{
		"id": "5",
		"name": "Vino Cabernet",
		"type": "Vino",
		"price": 100
	}`

	c:= utils.GetMockedContext(http.MethodGet,
		fmt.Sprintf("%s", strings.Replace(utils.URL_GET_DRINK_BY_ID, ":drinkId",
			wantedId, 1)),
		bytes.NewBufferString(jsonResponse),
		response)

	c.Params = []gin.Param{
		{
			Key:   "drinkId",
			Value: wantedId,
		},
	}

	connection := config.CreateClient()

	repo := repository.Repository {
		DBClient:connection,
	}

	handler := Handler {
		Repository: repo,
	}

	handler.GetDrinkById(c)
	assert.EqualValues(t, http.StatusOK, response.Code)
}
