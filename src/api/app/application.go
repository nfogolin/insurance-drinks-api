package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/insurance-drinks-api/src/api/config"
	"github.com/insurance-drinks-api/src/api/infrastructure/handlers"
	"github.com/insurance-drinks-api/src/api/infrastructure/repository"
	"github.com/insurance-drinks-api/src/api/infrastructure/utils"
)

func StartApi() {
	router := gin.Default()
	connection := config.CreateClient()

	repo := repository.Repository{
		DBClient:connection,
	}

	handler := handlers.Handler{
		Repository:repo,
	}

	router.GET(utils.URL_GET_PING, handler.Ping)
	router.POST(utils.URL_POST_DRINKS, handler.SaveDrink)
	router.GET(utils.URL_GET_DRINKS, handler.GetDrinks)
	router.GET(utils.URL_GET_DRINK_BY_ID, handler.GetDrinkById)

	if err := router.Run("localhost:8080"); err != nil {
		fmt.Sprintf("Error al inicializar API %s", err)
	}
}
