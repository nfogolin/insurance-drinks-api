package instances

import (
	"github.com/insurance-drinks-api/src/api/core/entities"
	"github.com/insurance-drinks-api/src/api/core/entities/interfaces"
)

const (
	Whisky = "Whisky"
	Wine = "Wine"
	Water = "Water"
)

func CastDrinkInstances(obj entities.Drink) interfaces.Drink {
	var newObj interfaces.Drink

	switch obj.Type {
	case Whisky:
		newObj = &entities.Whisky{
			Drink: obj,
		}
	case Wine:
		newObj = &entities.Wine{
			Drink: obj,
		}
	case Water:
		newObj = &entities.Water{
			Drink: obj,
		}
	}

	return newObj
}
