package mappings

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/insurance-drinks-api/src/api/infrastructure/dependencies"
	"github.com/insurance-drinks-api/src/api/infrastructure/utils"
)

type Router struct {
	HandlerContainer dependencies.HandlerContainer
}

func (t *Router) ConfigureRouting() {
	router := gin.Default()

	router.GET(utils.URL_GET_PING, t.HandlerContainer.Ping.Ping)
	router.POST(utils.URL_POST_DRINKS, t.HandlerContainer.SaveDrink.SaveDrink)
	router.GET(utils.URL_GET_DRINKS, t.HandlerContainer.GetDrinks.GetDrinks)
	router.GET(utils.URL_GET_DRINK_BY_ID, t.HandlerContainer.GetDrinkById.GetDrinkById)

	if err := router.Run("localhost:8080"); err != nil {
		fmt.Sprintf("Error al inicializar API %s", err)
	}
}
