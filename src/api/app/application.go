package app

import (
	"github.com/insurance-drinks-api/src/api/infrastructure/dependencies"
	"github.com/insurance-drinks-api/src/api/infrastructure/mappings"
)

func StartApi() {
	handlers := dependencies.Initialize()
	router := mappings.Router{
		HandlerContainer : *handlers,
	}

	router.ConfigureRouting()
}
