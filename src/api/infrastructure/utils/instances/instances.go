package instances

import (
	"github.com/insurance-drinks-api/src/api/core/entities"
	"github.com/insurance-drinks-api/src/api/core/entities/dto"
	"github.com/insurance-drinks-api/src/api/core/entities/interfaces"
)

const (
	Whisky = "Whisky"
	Wine = "Wine"
	Water = "Water"
)

func CastDrinkInstances(obj dto.Drink) interfaces.Drink {
	var newObj interfaces.Drink

	switch obj.Type {
	case Whisky:
		newObj = &entities.Whisky{
			BaseDrink: obj.BaseDrink,
			Aging:obj.Aging,
		}
	case Wine:
		newObj = &entities.Wine{
			BaseDrink: obj.BaseDrink,
		}
	case Water:
		newObj = &entities.Water{
			BaseDrink: obj.BaseDrink,
		}
	}

	return newObj
}
