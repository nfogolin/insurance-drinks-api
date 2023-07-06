package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/insurance-drinks-api/src/api/core/entities/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDbRepository struct {
	Client *mongo.Client
}

func (r MongoDbRepository) GetByID(id int64) (response *dto.Drink, err error) {
	return response, nil
}

func (r MongoDbRepository) GetDrinks() (response []dto.Drink, err error) {
	collection := r.Client.Database("db-ida-api").Collection("drink")

	cur, err := collection.Find(context.TODO(), bson.D{{}})

	if err != nil {
		return nil, err
	}

	if err = cur.All(context.Background(), &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (r MongoDbRepository) SaveDrink(drink dto.Drink) (response *dto.Drink, err error) {
	collection := r.Client.Database("db-ida-api").Collection("drink")

	drink.Id = uuid.New().String()

	_, err = collection.InsertOne(context.TODO(), &drink)
	if err != nil {
		return nil, err
	}

	return &drink, nil
}
