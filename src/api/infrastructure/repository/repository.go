package repository

import (
	"errors"
	"github.com/insurance-drinks-api/src/api/core/entities/dto"
	"gorm.io/gorm"
)

type Repository struct {
	DBClient *gorm.DB
}

func (t Repository) GetByID(id int64) (response *dto.Drink, err error) {
	result := t.DBClient.Table("Drink").First(&response, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}

		return nil, result.Error
	}

	return response, nil
}

func (t Repository) GetDrinks() (response []dto.Drink, err error) {
	result := t.DBClient.Table("Drink").Scan(&response)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}

		return nil, result.Error
	}

	return response, nil
}

func (t Repository) SaveDrink(drink dto.Drink) (response *dto.Drink, err error) {
	result := t.DBClient.Table("Drink").Save(&drink)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}

		return nil, result.Error
	}

	return &drink, nil
}
