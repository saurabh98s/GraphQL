package api

import (
	"context"
	"graphql/graph/model"
	"graphql/repository"
)

// Coffee implements the functions required by the domain type
type Coffee interface {
	Add(ctx context.Context, input model.NewCoffee) (*model.Coffee, error)
}

type coffee struct {
	coffeeRepo repository.CoffeeRepo
}

//Add saves a new Data to the DB and subsequently converts the mongoDB model to GraphQL
