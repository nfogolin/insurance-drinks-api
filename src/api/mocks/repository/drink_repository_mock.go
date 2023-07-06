package repository

import (
	"github.com/insurance-drinks-api/src/api/core/entities/dto"
	"github.com/stretchr/testify/mock"
)

type DrinkRepositoryMock struct {
	mock.Mock
}

func (mock *DrinkRepositoryMock) GetByID(id int64) (response *dto.Drink, err error) {
	responseArgs := mock.Called(id)
	response, _ = responseArgs.Get(0).(*dto.Drink)
	err, _ = responseArgs.Get(1).(error)

	return response, err
}

func (mock *DrinkRepositoryMock) GetDrinks() (response []dto.Drink, err error) {

	return response, nil
}

func (mock *DrinkRepositoryMock) SaveDrink(drink dto.Drink) (response *dto.Drink, err error) {
	responseArgs := mock.Called(drink)
	response, ok := responseArgs.Get(0).(*dto.Drink)
	if !ok {
		response = nil
	}
	err, _ = responseArgs.Get(1).(error)

	return response, nil
}
