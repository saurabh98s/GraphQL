package api

import (
	"context"
	"graphql/db/models"
	"graphql/graph/model"
	"graphql/logger"
	"graphql/repository"
	"net/http"
)

// Coffee implements the functions required by the domain type
type Coffee interface {
	Add(ctx context.Context, input model.NewCoffee) (*model.Coffee, error)
}

type coffee struct {
	coffeeRepo repository.CoffeeRepo
}

//Add saves a new Data to the DB and subsequently converts the mongoDB model to GraphQL
func (coffee *coffee) Add(ctx context.Context, input model.NewCoffee) (*model.Coffee, error) {
	var sendResult *model.Coffee
	doc:=&models.Coffee{
		Name: &input.Name,
		Key: &input.Key,
	}
	result,err:=coffee.coffeeRepo.Save(ctx,doc)
	if err!=nil {
		logger.Log.WithError(err).Error(http.StatusInternalServerError)
	}
	sendResult=result.Serialize()
	return sendResult,nil
}
// NewCoffee creates a new copy of the domain type
func NewCoffee(coffeeRepo repository.CoffeeRepo) Coffee {
	return &coffee{
		coffeeRepo: coffeeRepo,
	}
}