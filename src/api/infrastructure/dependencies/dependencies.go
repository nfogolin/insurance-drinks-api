package dependencies

import (
	"github.com/insurance-drinks-api/src/api/config"
	"github.com/insurance-drinks-api/src/api/infrastructure/handlers"
	"github.com/insurance-drinks-api/src/api/infrastructure/repository"
)

type HandlerContainer struct {
	Ping handlers.Ping
	GetDrinkById handlers.GetDrinkById
	GetDrinks handlers.GetDrinks
	SaveDrink handlers.SaveDrink
}

var (
	drinkRepository repository.Repository
)

func Initialize() *HandlerContainer {
	connection := config.CreateClient()

	drinkRepository = repository.Repository{
		DBClient:connection,
	}

	handlerContainer := HandlerContainer{}

	handlerContainer.Ping = handlers.Ping{}

	handlerContainer.GetDrinkById = handlers.GetDrinkById{
		Repository:drinkRepository,
	}

	handlerContainer.GetDrinks = handlers.GetDrinks{
		Repository:drinkRepository,
	}

	handlerContainer.SaveDrink = handlers.SaveDrink{
		Repository:drinkRepository,
	}

	return &handlerContainer
}
