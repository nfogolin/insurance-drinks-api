package dependencies

import (
	"github.com/insurance-drinks-api/src/api/config"
	"github.com/insurance-drinks-api/src/api/infrastructure/handlers"
	"github.com/insurance-drinks-api/src/api/infrastructure/handlers/interfaces"
	"github.com/insurance-drinks-api/src/api/infrastructure/repository"
	irepository "github.com/insurance-drinks-api/src/api/infrastructure/repository/interfaces"
)

type HandlerContainer struct {
	Ping interfaces.Handler
	GetDrinkById interfaces.Handler
	GetDrinks interfaces.Handler
	SaveDrink interfaces.Handler
}

var (
	drinkRepository irepository.Repository
	mongodbRepository irepository.Repository
)

func Initialize() *HandlerContainer {
	connection := config.CreateClient()
	mongodb := config.CreateMongoDBClient()

	drinkRepository = repository.MysqlRepository{
		DBClient:connection,
	}

	mongodbRepository = repository.MongoDbRepository{
		Client: mongodb,
	}

	handlerContainer := HandlerContainer{}

	handlerContainer.Ping = handlers.Ping{}

	handlerContainer.GetDrinkById = handlers.GetDrinkById{
		Repository:drinkRepository,
	}

	handlerContainer.GetDrinks = handlers.GetDrinks{
		Repository:mongodbRepository,
	}

	handlerContainer.SaveDrink = handlers.SaveDrink{
		Repository:mongodbRepository,
	}

	return &handlerContainer
}
