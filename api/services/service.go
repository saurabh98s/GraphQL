package service

import (
	"graphql/api"
	"graphql/instance"
	"graphql/repository"
)

// Services is the interface for enclosing all the services
type Services interface {
	Coffee() api.Coffee
}

type services struct {
	CoffeeService api.Coffee
}


func (svc *services)Coffee() api.Coffee {
	return svc.CoffeeService
}

// Init intialised the Services
func Init() Services  {
	return &services{
		CoffeeService: api.NewCoffee(repository.NewCoffeeRepo(instance.DB())),
	}
}